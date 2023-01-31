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

type Authorization interface {
	CreateAuthUser(context.Context, *models.UserAuth) (int, error)
	GetAuthUser(context.Context, int) (models.UserAuth, error)
	GetUserById(ctx context.Context, userId int) (models.UserAuth, error)
	GetUserWithRefreshToken(ctx context.Context, rt string) (models.UserAuth, error)
	UpdateRefreshToken(ctx context.Context, rt string, id int) error
}

type Repository struct {
	UserAct
	Authorization
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		UserAct:       NewUserActPostgres(db),
		Authorization: NewUserAuthPostgres(db),
	}
}
