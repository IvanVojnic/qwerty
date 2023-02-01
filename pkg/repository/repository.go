// Package repository is a repository
package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository struct that has two fields
type Repository struct {
	BookAct       *BookActPostgres
	Authorization *UserAuthPostgres
}

// NewRepository to init repo
func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		BookAct:       NewBookActPostgres(db),
		Authorization: NewUserAuthPostgres(db),
	}
}
