// Package service init service
package service

import (
	"context"

	"EFpractic2/models"
	"EFpractic2/pkg/repository"

	"github.com/google/uuid"
)

// BookAct interface consists of methos to communicate with boockAct repo
type BookAct interface {
	CreateBook(context.Context, models.Book) error
	UpdateBook(context.Context, models.Book) error
	GetBook(context.Context, int) (models.Book, error)
	DeleteBook(context.Context, int) error
	GetAllBooks(context.Context) ([]models.Book, error)
}

// Authorization interface consists of methos to communicate with user repo
type Authorization interface {
	CreateAuthUser(context.Context, *models.UserAuth) error
	GetUserByID(context.Context, uuid.UUID) (models.UserAuth, error)
	UpdateRefreshToken(context.Context, string, uuid.UUID) error
	SignInUser(context.Context, *models.UserAuth) error
}

// Service is a wrapper for to repos
type Service struct {
	BookAct       *BookActSrv
	Authorization *AuthService
}

// NewService used to init Service
func NewService(repos *repository.Repository) *Service {
	return &Service{
		BookAct:       NewBookActSrv(*repos.BookAct),
		Authorization: NewAuthService(*repos.Authorization),
	}
}
