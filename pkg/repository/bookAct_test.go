// Package repository file to test book act
package repository

import (
	"context"
	"os"
	"testing"

	"EFpractic2/models"
	"EFpractic2/pkg/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

var Pool *pgxpool.Pool

var testValidData = []models.Book{
	{
		BookName: `algebra`,
		BookYear: 1950,
		BookNew:  false,
	},
	{
		BookName: `geometric`,
		BookYear: 2023,
		BookNew:  true,
	},
}

var testNotValidData = []models.Book{
	{
		BookName: `false`,
		BookYear: -5,
		BookNew:  true,
	},
	{
		BookName: `2023`,
		BookYear: 1000000000,
		BookNew:  false,
	},
}

// db object used fo testing
var db *pgxpool.Pool

// TestMain used to test main func
func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			"pUrl=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
			"DB=postgres",
			"PASSWORD=postgres",
			"USER=postgres",
			"PORT=5432",
		},
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	cfg, errConfig := config.NewConfig()
	if errConfig != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if errConnWrap := pool.Retry(func() error {
		var errConn error
		db, errConn = pgxpool.New(context.Background(), cfg.PostgresURL)
		if errConn != nil {
			return errConn
		}
		return db.Ping(context.Background())
	}); errConnWrap != nil {
		log.Fatalf("Could not connect to database: %s", errConnWrap)
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

// Test_CreateBook used to create book
func Test_CreateBook(t *testing.T) {
	ctx := context.Background()
	db := NewBookActPostgres(Pool)
	for _, u := range testValidData {
		_, err := db.CreateBook(ctx, u)
		require.NoError(t, err, "create error")
	}

}
