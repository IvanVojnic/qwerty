package repository

import "github.com/jackc/pgx/v5/pgxpool"

type UserAuthPostgres struct {
	db *pgxpool.Pool
}

func NewUserAuthPostgres(db *pgxpool.Pool) *UserActPostgres {
	return &UserActPostgres{db: db}
}
