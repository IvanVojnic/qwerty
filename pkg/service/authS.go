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
	id, err := s.repo.CreateAuthUser(ctx, &user)
	if err != nil {
		return "", "", fmt.Errorf("Error create auth user %w", err)
	}
	at, errAT := utils.GenerateToken(id, false)
	if errAT != nil {
		log.WithFields(log.Fields{
			"ERROR":        errAT,
			"access token": at,
		}).Info("Error while generating access token")
	}
	rt, errRT := utils.GenerateToken(id, true)
	if errRT != nil {
		log.WithFields(log.Fields{
			"ERROR":        errRT,
			"access token": rt,
		}).Info("Error while generating access token")
	}
	errRtInsert := s.repo.UpdateRefreshToken(ctx, rt, id)
	if errRtInsert != nil {

	}
	return rt, at, err
}

func (s *AuthService) GetUserVerified(ctx context.Context, id int /*at string, rt string*/) (models.UserAuth, error) {
	/*userIdByAT, err := utils.ParseToken(at)
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
	}*/
	user, err := s.repo.GetUserById(ctx, id)
	return user, err
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
