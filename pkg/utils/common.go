package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

const SigningKey = "barband"

const TokenRTDuration = 1 * time.Hour
const TokenATDuretion = 100 * time.Minute

type tokenClaims struct {
	jwt.StandardClaims
	UserId uuid.UUID `json:"user_id"`
}

func ParseToken(tokenToParse string) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenToParse, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(SigningKey), nil
	})
	if err != nil {
		return false, err
	}
	_, ok := token.Claims.(*tokenClaims)
	if !ok {
		return false, errors.New("token claims are not type of tokenClaims")
	}
	return false, nil
}

func GenerateToken(id uuid.UUID, tokenDuration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})
	return token.SignedString([]byte(SigningKey))
}

func IsAuthorized(requestToken string) (bool, error) {
	_, err := ParseToken(requestToken)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(requestToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SigningKey), nil
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if ok && !token.Valid {
		return uuid.UUID{}, fmt.Errorf("Invalid Token")
	}

	return claims.UserId, nil
}

func IsTokenExpired(requestToken string) bool {
	token, err := jwt.ParseWithClaims(requestToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(SigningKey), nil
	})
	v, _ := err.(*jwt.ValidationError)
	tokenExpired := false

	if tk := token.Claims.(jwt.StandardClaims); v.Errors == jwt.ValidationErrorExpired && tk.VerifyExpiresAt(time.Now().Unix(), tokenExpired) {
		return true
	}
	return tokenExpired
}
