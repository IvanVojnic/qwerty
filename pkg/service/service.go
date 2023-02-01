package service

import (
	"EFpractic2/models"
	"EFpractic2/pkg/repository"
	"context"
)

type BookAct interface {
	CreateBook(context.Context, models.Book) error
	UpdateBook(context.Context, models.Book) error
	GetBook(context.Context, int) (models.Book, error)
	DeleteBook(context.Context, int) error
	GetAllBooks(context.Context) ([]models.Book, error)
}

type Authorization interface {
	CreateUserVerified(context.Context, models.UserAuth, string) error
	GetUserVerified(context.Context, interface{}) (models.UserAuth, error)
	SignInUser(context.Context, *models.UserAuth) (bool, error)
}

type Service struct {
	BookAct
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		BookAct:       NewBookActSrv(repos.BookAct),
		Authorization: NewAuthService(repos.Authorization),
	}
}
