// Package service for user
package service

import (
	"context"
	"crypto/sha256"
	"fmt"

	"EFpractic2/models"
	"github.com/google/uuid"
)

const salt = "s53d42fg98gh7j6kkbver"

// Authorization interface consists of methos to communicate with user repo
type Authorization interface {
	CreateAuthUser(context.Context, *models.UserAuth) error
	GetUserByID(context.Context, uuid.UUID) (models.UserAuth, error)
	UpdateRefreshToken(context.Context, string, uuid.UUID) error
	SignInUser(context.Context, *models.UserAuth) error
}

// AuthService is wrapper for user repo
type AuthService struct {
	repo Authorization
}

// NewAuthService used to init AS
func NewAuthService(repo Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUserVerified used to
func (s *AuthService) CreateUserVerified(ctx context.Context, user *models.UserAuth, rt string) error {
	user.Password = generatePasswordHash(user.Password)
	err := s.repo.CreateAuthUser(ctx, user)
	if err != nil {
		return fmt.Errorf("error create auth user %w", err)
	}
	errRtInsert := s.repo.UpdateRefreshToken(ctx, rt, user.UserID)
	if errRtInsert != nil {
		return fmt.Errorf("error update rt token %w", errRtInsert)
	}
	return err
}

// GetUserVerified used to get user
func (s *AuthService) GetUserVerified(ctx context.Context, id uuid.UUID) (models.UserAuth, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	return user, err
}

// SignInUser used to sign in user
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

// generatePasswordHash used to generate hash password
func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
