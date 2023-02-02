// Package repository init db
package repository

import (
	"EFpractic2/pkg/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPostgresDB func to init and connect to db
func NewPostgresDB(cfg *config.Config) (pool *pgxpool.Pool, err error) {
	pool, err = pgxpool.New(context.Background(), cfg.PostgresURL)
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
