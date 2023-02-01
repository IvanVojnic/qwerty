package service

import (
	"EFpractic2/models"
	"EFpractic2/pkg/repository"
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/google/uuid"
)

const salt = "s53d42fg98gh7j6kkbver"

type AuthService struct {
	repo repository.UserAuthPostgres
}

func NewAuthService(repo repository.UserAuthPostgres) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUserVerified(ctx context.Context, user models.UserAuth, rt string) error {
	user.Password = generatePasswordHash(user.Password)
	err := s.repo.CreateAuthUser(ctx, &user)
	if err != nil {
		return fmt.Errorf("Error create auth user %w", err)
	}
	errRtInsert := s.repo.UpdateRefreshToken(ctx, rt, user.UserId)
	if errRtInsert != nil {
		return fmt.Errorf("Error update rt token %w", errRtInsert)
	}
	return err
}

func (s *AuthService) GetUserVerified(ctx context.Context, id uuid.UUID) (models.UserAuth, error) {
	user, err := s.repo.GetUserById(ctx, id)
	return user, err
}

func (s *AuthService) SignInUser(ctx context.Context, user *models.UserAuth) (bool, error) {
	hashedPass := generatePasswordHash(user.Password)
	err := s.repo.SignInUser(ctx, user)
	if err != nil {
		return false, fmt.Errorf("error while sign in query %w", err)
	}
	if user.Password == hashedPass {
		return true, nil
	}
	return false, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
