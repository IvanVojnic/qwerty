package service

import (
	"EFpractic2/models"
	"EFpractic2/pkg/repository"
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"

	//"github.com/golang-jwt/jwt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const salt = "s53d42fg98gh7j6kkbver"
const signingKey = "barband"
const tokenTTL = 12 * time.Hour

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not type of tokenClaims")
	}
	return claims.UserId, nil
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUserVerified(ctx context.Context, user models.UserAuth) (string, string, error) {
	user.Password = generatePasswordHash(user.Password)
	rt, id, err := s.repo.CreateAuthUser(ctx, &user)
	if err != nil {
		return "", "", fmt.Errorf("Error create auth user %w", err)
	}
	at, errGT := s.GenerateToken(id)
	if errGT != nil {
		log.WithFields(log.Fields{
			"ERROR":        err,
			"access token": at,
		}).Info("Error while generating access token")
	}
	//return s.repo.CreateAuthUser(ctx, &user)
	return rt, at, err
}

func (s *AuthService) GetUserVerified(ctx context.Context, at string, rt string) (models.UserAuth, error) {
	userIdByAT, err := s.ParseToken(at)
	//var user models.UserAuth
	if err != nil {
		return models.UserAuth{}, err
	}
	user, err := s.repo.GetUserId(ctx, userIdByAT)
	/*if userIdByAT == userIdByRT {
		return s.repo.GetAuthUser(ctx, userIdByRT)
	}*/
	//return user.Id, nil
	return user, nil
}

func (s *AuthService) GenerateToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})
	return token.SignedString([]byte(signingKey))
	//return "qwerty", nil
}

/*
accessClaims := &CustomClaims{
		user.Login,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenStr, err = accessToken.SignedString(u.jwtKey)
	if err != nil {
		return "", "", fmt.Errorf("userService - CreateJWT - SignedString: %w", err)
	}

	refreshClaims := &CustomClaims{
		user.Login,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 10)),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenStr, err = refreshToken.SignedString(u.jwtKey)
	if err != nil {
		return "", "", fmt.Errorf("userService - CreateJWT - SignedString: %w", err)
	}

	err = u.rps.RefreshUser(ctx, user.Login, refreshTokenStr)
	if err != nil {
		return "", "", fmt.Errorf("userService - CreateJWT - RefreshUser: %w", err)
	}
	return
*/

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
