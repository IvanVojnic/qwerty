package repository

import (
	"EFpractic2/models"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserAct interface {
	CreateUser(context.Context, models.User) error
	UpdateUser(context.Context, models.User) error
	GetUser(context.Context, int) (models.User, error)
	DeleteUser(context.Context, int) error
	GetAllUsers(context.Context) ([]models.User, error)
}

type AuthUser interface {
}

type Repository struct {
	UserAct
	AuthUser
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		UserAct:  NewUserActPostgres(db),
		AuthUser: NewUserAuthPostgres(db),
	}
}
