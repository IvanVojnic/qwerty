package repository

import (
	"EFpractic2/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

type UserActPostgres struct {
	db *pgxpool.Pool
}

func NewUserActPostgres(db *pgxpool.Pool) *UserActPostgres {
	return &UserActPostgres{db: db}
}

func (r *UserActPostgres) CreateUser(ctx context.Context, user models.User) error {
	_, err := r.db.Exec(ctx, "insert into users (name, age, regular, password) values($1, $2, $3, $4)",
		user.UserName, user.UserAge, user.UserIsRegular, user.Password)
	if err != nil {
		return fmt.Errorf("error while user creating: %v", err)
	}
	return nil
}

func (r *UserActPostgres) UpdateUser(ctx context.Context, user models.User) error {
	_, err := r.db.Exec(ctx, "UPDATE users SET name = $1, age = $2, regular =$3 WHERE id = $4", user.UserName, user.UserAge, user.UserIsRegular, user.UserId)
	if err != nil {
		return fmt.Errorf("update user error %w", err)
	}
	return nil
}

func (r *UserActPostgres) GetUser(ctx context.Context, userId int) (models.User, error) {
	user := models.User{}
	err := r.db.QueryRow(ctx, "select users.id, users.name, users.age, users.regular, users.password from users where id=$1", userId).Scan(
		&user.UserId, &user.UserName, &user.UserAge, &user.UserIsRegular, &user.Password)
	if err != nil {
		return user, fmt.Errorf("get user error %w", err)
	}
	return user, nil
}

func (r *UserActPostgres) DeleteUser(ctx context.Context, userId int) error {
	_, err := r.db.Exec(ctx, "delete from users where id=$1", userId)
	if err != nil {
		return fmt.Errorf("delete user error %w", err)
	}
	return nil
}

func (r *UserActPostgres) GetAllUsers(ctx context.Context) ([]models.User, error) {
	users := make([]models.User, 0)
	rows, err := r.db.Query(ctx, "select users.id, users.name, users.age, users.regular, users.password from users")
	if err != nil {
		log.WithFields(log.Fields{
			"Error get all user": err,
			"rows":               rows,
		}).Info("SQL QUERY")
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		errScan := rows.Scan(&user.UserId, &user.UserName, &user.UserAge, &user.UserIsRegular, &user.Password)
		if errScan != nil {
			log.WithFields(log.Fields{
				"Error while scan current row to get user model": err,
				"user": user,
			}).Info("SCAN ERROR. GET ALL USERS")
		}
		users = append(users, user)
	}
	if errRows := rows.Err(); errRows != nil {
		fmt.Errorf("get all users error %w", errRows)
	}
	return users, err
}
