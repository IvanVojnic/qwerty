// Package repository init db
package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// StorageConfig is struct with postgres url
type StorageConfig struct {
	PostgresURL string `json:"pUrl"`
}

// NewPostgresDB func to init and connect to db
func NewPostgresDB() (pool *pgxpool.Pool, err error) {
	pool, err = pgxpool.New(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("invalid configuration data: %v", err)
	}
	if err = pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("database not responding: %v", err)
	}
	return pool, err
}

// ClosePool is a func to close connection to db
func ClosePool(myPool *pgxpool.Pool) {
	if myPool != nil {
		myPool.Close()
	}
}
