package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	BookAct       *BookActPostgres
	Authorization *UserAuthPostgres
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		BookAct:       NewBookActPostgres(db),
		Authorization: NewUserAuthPostgres(db),
	}
}
