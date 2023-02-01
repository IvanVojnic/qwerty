// Package repository declare func for user
package repository

import (
	"context"
	"fmt"

	"EFpractic2/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// UserAuthPostgres has an internal db object
type UserAuthPostgres struct {
	db *pgxpool.Pool
}

// NewUserAuthPostgres used to init UsesAP
func NewUserAuthPostgres(db *pgxpool.Pool) *UserAuthPostgres {
	return &UserAuthPostgres{db: db}
}

// CreateAuthUser used to create user
func (r *UserAuthPostgres) CreateAuthUser(ctx context.Context, user *models.UserAuth) error {
	_, err := r.db.Exec(ctx, "insert into usersauth (id, name, age, regular, password) values($1, $2, $3, $4, $5)",
		user.UserID, user.UserName, user.UserAge, user.UserIsRegular, user.Password)
	if err != nil {
		return fmt.Errorf("error while user creating: %v", err)
	}
	return nil
}

// UpdateRefreshToken used to update rt
func (r *UserAuthPostgres) UpdateRefreshToken(ctx context.Context, rt string, id uuid.UUID) error {
	_, errInsert := r.db.Exec(ctx, "UPDATE usersauth SET refreshtoken = $1 WHERE id = $2", rt, id)
	if errInsert != nil {
		return fmt.Errorf("update user error %w", errInsert)
	}
	return nil
}

// GetUserByID used to get user by ID
func (r *UserAuthPostgres) GetUserByID(ctx context.Context, userID uuid.UUID) (models.UserAuth, error) {
	user := models.UserAuth{}
	err := r.db.QueryRow(ctx,
		"select usersauth.id, usersauth.name, usersauth.age, usersauth.regular, usersauth.password, usersauth.refreshtoken from usersauth where id=$1",
		userID).Scan(&user.UserID, &user.UserName, &user.UserAge, &user.UserIsRegular, &user.Password, &user.RefreshToken)
	if err != nil {
		return user, fmt.Errorf("get user error %w", err)
	}

	return user, nil
}

/*
create table usersauth(
    id uuid primary key not null,
    name varchar(255) not null,
    age int not null,
    regular bool not null,
    password varchar(255) not null,
    refreshtoken varchar(255)
);*/

// SignInUser used to sign in user
func (r *UserAuthPostgres) SignInUser(ctx context.Context, user *models.UserAuth) error {
	err := r.db.QueryRow(ctx,
		"select usersauth.id, usersauth.name, usersauth.age, usersauth.regular, usersauth.password, usersauth.refreshtoken from usersauth where name=$1",
		user.UserName).Scan(&user.UserID, &user.UserName, &user.UserAge, &user.UserIsRegular, &user.Password, &user.RefreshToken)
	if err != nil {
		return fmt.Errorf("error while getting user %w", err)
	}
	return nil
}
