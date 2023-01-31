package service

import (
	"EFpractic2/models"
	"EFpractic2/pkg/repository"
	"EFpractic2/pkg/utils"
	"context"
	"crypto/sha1"
	"fmt"
	log "github.com/sirupsen/logrus"
)

const salt = "s53d42fg98gh7j6kkbver"

type AuthService struct {
	repo repository.Authorization
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
	at, errGT := utils.GenerateToken(id, false)
	if errGT != nil {
		log.WithFields(log.Fields{
			"ERROR":        err,
			"access token": at,
		}).Info("Error while generating access token")
	}
	return rt, at, err
}

func (s *AuthService) GetUserVerified(ctx context.Context, at string, rt string) (models.UserAuth, string, error) {
	userIdByAT, err := utils.ParseToken(at)
	var user models.UserAuth
	if err != nil {
		log.WithFields(log.Fields{
			"ERROR":        err,
			"access token": at,
		}).Info("Error while verified access token")
		user, err = s.repo.GetUserWithRefreshToken(ctx, rt)
		if err != nil {
			return user, "", err
		}
		newAt, errAt := utils.GenerateToken(user.UserId, false)
		if err != nil {
			log.WithFields(log.Fields{
				"ERROR":        errAt,
				"access token": newAt,
			}).Info("Error while generating access token")
		}
		return user, at, err
	}
	user, err = s.repo.GetUserById(ctx, userIdByAT)
	return user, at, err
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
