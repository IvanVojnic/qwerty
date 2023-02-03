// Package repository file to test book act
package repository

import (
	"EFpractic2/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"os"
	"os/exec"
	"testing"
)

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

// Test_Main used to test main func
func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("unix:///home/ivanvoynich/.docker/desktop/docker.sock")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			"POSTGRES_DB=postgres",
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=postgres",
		},
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseURL := fmt.Sprintf("postgresql://postgres:postgres@%s/postgres", hostAndPort)

	if errConnWrap := pool.Retry(func() error {
		var errConn error
		db, errConn = pgxpool.New(context.Background(), databaseURL)
		if errConn != nil {
			return errConn
		}
		return db.Ping(context.Background())
	}); errConnWrap != nil {
		log.Fatalf("Could not connect to database: %s", errConnWrap)
	}

	/* resource2, errFlyway := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "flyway/flyway",
		Cmd: []string{
			"-url=jdbc:postgresql://postgres:5432/postgres",
			"-user=postgres",
			"-password=postgres",
			"-locations=filesystem:/flyway/sql",
			"-connectRetries=10",
			"migrate",
		},
	})
	if errFlyway != nil {
		log.Fatalf("Could not start resource: %s", errFlyway)
	} */

	cmd := exec.Command("flyway", "-user=postgres", "-password=postgres",
		"-locations=filesystem:../../migrations/sql",
		fmt.Sprintf("-url=jdbc:postgresql://%v/postgres", hostAndPort), "migrate")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("can't run flyway: %s", err)
	}
	code := m.Run()

	if errR1 := pool.Purge(resource); errR1 != nil {
		log.Fatalf("Could not purge resource: %s", errR1)
	}

	os.Exit(code)
}

// Test_CreateBook used to create book
func Test_CreateBook(t *testing.T) {
	ctx := context.Background()
	repos := NewBookActPostgres(db)
	for _, b := range testValidData {
		err := repos.CreateBook(ctx, b)
		require.NoError(t, err, "create error")
		repos.db.Exec(ctx, "delete from books where name=$1", b.BookName)
	}
	/*	for _, b := range testNotValidData {
			err := repos.CreateBook(ctx, b)
			require.Error(t, err, "create error")
		}
		for _, b := range testValidData {
			err := repos.CreateBook(ctx, b)
			require.NoError(t, err, "create error")

			err = repos.CreateBook(ctx, b)
			require.Error(t, err, "create error")
		}*/
}

func TestBookActPostgres_GetBook(t *testing.T) {
	ctx := context.Background()
	repos := NewBookActPostgres(db)

	var bookFromDB models.Book
	for _, b := range testValidData {
		_, err := repos.db.Exec(ctx, "delete from books where name=$1", b.BookName)

		bookFromDB, err = repos.GetBook(ctx, b.BookID)
		require.Equal(t, bookFromDB.BookName, b.BookName)
		require.Equal(t, bookFromDB.BookYear, b.BookYear)
		require.Equal(t, bookFromDB.BookNew, b.BookNew)
		require.NoError(t, err, "get book by ID")

		_, err = repos.db.Exec(ctx, "delete from books where name=$1", b.BookName)
	}
}
