package service

import (
	"EFpractic2/models"
	"EFpractic2/pkg/repository"
	"context"
	"github.com/google/uuid"
)

type BookAct interface {
	CreateBook(context.Context, models.Book) error
	UpdateBook(context.Context, models.Book) error
	GetBook(context.Context, int) (models.Book, error)
	DeleteBook(context.Context, int) error
	GetAllBooks(context.Context) ([]models.Book, error)
}

type Authorization interface {
	CreateAuthUser(context.Context, *models.UserAuth) error
	GetUserById(context.Context, uuid.UUID) (models.UserAuth, error)
	UpdateRefreshToken(context.Context, string, uuid.UUID) error
	SignInUser(context.Context, *models.UserAuth) error
}

type Service struct {
	BookAct       *BookActSrv
	Authorization *AuthService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		BookAct:       NewBookActSrv(*repos.BookAct),
		Authorization: NewAuthService(*repos.Authorization),
	}
}
