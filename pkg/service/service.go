package service

import (
	"EFpractic2/models"
	"EFpractic2/pkg/repository"
	"context"
)

type UserAct interface {
	CreateUser(context.Context, models.User) error
	UpdateUser(context.Context, models.User) error
	GetUser(context.Context, int) (models.User, error)
	DeleteUser(context.Context, int) error
	GetAllUsers(context.Context) ([]models.User, error)
}

type Authorization interface {
	CreateUserVerified(context.Context, models.UserAuth) (string, string, error)
	GetUserVerified(context.Context, string, string) (models.UserAuth, string, error)
}

type Service struct {
	UserAct
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserAct:       NewUserActSrv(repos.UserAct),
		Authorization: NewAuthService(repos.Authorization),
	}
}
