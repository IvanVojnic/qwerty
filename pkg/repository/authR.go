package repository

import (
	"EFpractic2/models"
	"EFpractic2/pkg/utils"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserAuthPostgres struct {
	db *pgxpool.Pool
}

func NewUserAuthPostgres(db *pgxpool.Pool) *UserAuthPostgres {
	return &UserAuthPostgres{db: db}
}

func (r *UserAuthPostgres) CreateAuthUser(ctx context.Context, user *models.UserAuth) error {
	sqlRow := `insert into usersauth (id, name, age, regular, password) values($1, $2, $3, $4, $5)`
	err := r.db.QueryRow(ctx, sqlRow, user.UserId, user.UserName, user.UserAge, user.UserIsRegular, user.Password)
	if err != nil {
		return fmt.Errorf("error while user creating: %v", err)
	}
	return nil
}

func (r *UserAuthPostgres) UpdateRefreshToken(ctx context.Context, rt string, id interface{}) error {
	_, errInsert := r.db.Exec(ctx, "UPDATE usersauth SET refreshtoken = $1 WHERE id = $2", rt, id.(int))
	if errInsert != nil {
		return fmt.Errorf("update user error %w", errInsert)
	}
	return nil
}

func (r *UserAuthPostgres) GetUserById(ctx context.Context, userId interface{}) (models.UserAuth, error) {
	user := models.UserAuth{}
	err := r.db.QueryRow(ctx, "select usersauth.id, usersauth.name, usersauth.age, usersauth.regular, usersauth.password, usersauth.refreshtoken from usersauth where id=$1", userId).Scan(
		&user.UserId, &user.UserName, &user.UserAge, &user.UserIsRegular, &user.Password, &user.RefreshToken)
	if err != nil {
		return user, fmt.Errorf("get user error %w", err)
	}

	return user, nil
}

func (r *UserAuthPostgres) GetUserWithRefreshToken(ctx context.Context, rt string) (models.UserAuth, error) {
	user := models.UserAuth{}
	userId, errRT := utils.ParseToken(rt)
	if errRT != nil {
		return models.UserAuth{}, fmt.Errorf("parse rt error %w", errRT)
	}
	err := r.db.QueryRow(ctx, "select usersauth.id, usersauth.name, usersauth.age, usersauth.regular, usersauth.password, usersauth.refreshtoken from usersauth where id=$1", userId).Scan(
		&user.UserId, &user.UserName, &user.UserAge, &user.UserIsRegular, &user.Password, &user.RefreshToken)
	if err != nil {
		return models.UserAuth{}, fmt.Errorf("get user error %w", err)
	}
	return user, nil
}
