package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SigningKey = "barband"
const TokenTTL = 12 * time.Hour

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(SigningKey), nil
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

func GenerateToken(id int, rt bool) (string, error) {
	var tokenLife int64
	if rt {
		tokenLife = time.Now().Add(1 * time.Minute).Unix()
	} else {
		tokenLife = time.Now().Add(1 * time.Hour).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: tokenLife,
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})
	return token.SignedString([]byte(SigningKey))
	//return "qwerty", nil
}
