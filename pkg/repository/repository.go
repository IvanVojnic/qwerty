package repository

import (
	"EFpractic2/models"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
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
	GetUserById(context.Context, interface{}) (models.UserAuth, error)
	UpdateRefreshToken(context.Context, string, uuid.UUID) error
	SignInUser(context.Context, *models.UserAuth) error
}

type Repository struct {
	BookAct
	Authorization
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		BookAct:       NewBookActPostgres(db),
		Authorization: NewUserAuthPostgres(db),
	}
}
