package repository

import (
	"EFpractic2/models"
	"context"
	"fmt"
	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserAuthPostgres struct {
	db *pgxpool.Pool
}

func NewUserAuthPostgres(db *pgxpool.Pool) *UserAuthPostgres {
	return &UserAuthPostgres{db: db}
}

func (r *UserAuthPostgres) CreateAuthUser(ctx context.Context, user *models.UserAuth) error {
	_, err := r.db.Exec(ctx, "insert into usersauth (id, name, age, regular, password) values($1, $2, $3, $4, $5)", user.UserId, user.UserName, user.UserAge, user.UserIsRegular, user.Password) //nolint:lll
	if err != nil {
		return fmt.Errorf("error while user creating: %v", err)
	}
	return nil
}

func (r *UserAuthPostgres) UpdateRefreshToken(ctx context.Context, rt string, id uuid.UUID) error {
	_, errInsert := r.db.Exec(ctx, "UPDATE usersauth SET refreshtoken = $1 WHERE id = $2", rt, id)
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

func (r *UserAuthPostgres) SignInUser(ctx context.Context, user *models.UserAuth) error {
	err := r.db.QueryRow(ctx, "select usersauth.id, usersauth.name, usersauth.age, usersauth.regular, usersauth.password, usersauth.refreshtoken from usersauth where name=$1", user.UserName).Scan(&user.UserId, &user.UserName, &user.UserAge, &user.UserIsRegular, &user.Password, &user.RefreshToken)
	if err != nil {
		return fmt.Errorf("error while getting user %w", err)
	}
	return nil
}

/*
create table usersauth(
    id uuid primary key not null,
    name varchar(255) not null,
    age int not null,
    regular bool not null,
    password varchar(255) not null,
    refreshtoken varchar(255)
);
*/

/*func (r *UserAuthPostgres) GetUserWithRefreshToken(ctx context.Context, rt string) (models.UserAuth, error) {
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
}*/
