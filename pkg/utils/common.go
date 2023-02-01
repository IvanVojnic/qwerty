// Package utils tokens
package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// SigningKey is a secret key for tokens
const SigningKey = "barband"

// TokenRTDuration is a duration of rt life
const TokenRTDuration = 1 * time.Hour

// TokenATDuretion is a duration of at life
const TokenATDuretion = 100 * time.Minute

type tokenClaims struct {
	jwt.StandardClaims
	UserID uuid.UUID `json:"user_id"`
}

// ParseToken used to parse tokens with claims
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

// GenerateToken used to generate tokens with id
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

// IsAuthorized used to check is user authorized with the tokens
func IsAuthorized(requestToken string) (bool, error) {
	_, err := ParseToken(requestToken)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ExtractIDFromToken used to get id from the token
func ExtractIDFromToken(requestToken string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(requestToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SigningKey), nil
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if ok && !token.Valid {
		return uuid.UUID{}, fmt.Errorf("invalid Token")
	}

	return claims.UserID, nil
}

// IsTokenExpired used to check is token expired
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
