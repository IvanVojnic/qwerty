// Package repository file to test book act
package repository

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"EFpractic2/models"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
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
		err := repos.CreateBook(ctx, &b)
		require.NoError(t, err, "create error")
		repos.db.Exec(ctx, "delete from books where name=$1", b.BookName)
	}
}

// TestBookActPostgres_GetBook used to get book
func TestBookActPostgres_GetBook(t *testing.T) {
	ctx := context.Background()
	repos := NewBookActPostgres(db)
	b := testValidData[0]
	_, errDel := repos.db.Exec(ctx, "delete from books where name=$1", b.BookName)
	if errDel != nil {
		log.Fatalf("Could not purge resource: %s", errDel)
	}
	errCreate := repos.CreateBook(ctx, &b)
	if errCreate != nil {
		log.Fatalf("Could not purge resource: %s", errCreate)
	}
	bookFromDB, errGet := repos.GetBookByName(ctx, "book")
	require.Equal(t, bookFromDB.BookName, b.BookName)
	require.Equal(t, bookFromDB.BookYear, b.BookYear)
	require.Equal(t, bookFromDB.BookNew, b.BookNew)
	require.NoError(t, errGet, "get book by ID")

	_, errDel = repos.db.Exec(ctx, "delete from books where name=$1", b.BookName)
	if errDel != nil {
		log.Fatalf("Could not purge resource: %s", errDel)
	}
}

// Test_UpdateBook used to update book
func Test_UpdateBook(t *testing.T) {
	ctx := context.Background()
	repos := NewBookActPostgres(db)
	book1 := testValidData[0]
	book2 := testValidData[1]
	var err error
	_, err = repos.db.Exec(ctx, "delete from books where name=$1", book1.BookName)
	if err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	err = repos.CreateBook(ctx, &book1)
	if err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	id, errGetId := repos.GetBookId(ctx, book1.BookName)
	if errGetId != nil {
		log.Fatalf("Could not purge resource: %s", errGetId)
	}
	book2.BookID = id
	err = repos.UpdateBook(ctx, book2)
	updatedBook, err := repos.GetBookByName(ctx, "book")
	require.NotEqual(t, updatedBook.BookName, book1.BookName)
	require.NotEqual(t, updatedBook.BookYear, book1.BookYear)
	require.NotEqual(t, updatedBook.BookNew, book1.BookNew)
	require.NoError(t, err, "create error")
	repos.db.Exec(ctx, "delete from books where name=$1", book1.BookName)
}

// Test_DeleteBook used to delete book
func Test_DeleteBook(t *testing.T) {
	ctx := context.Background()
	repos := NewBookActPostgres(db)
	b := testValidData[0]
	var err error
	err = repos.CreateBook(ctx, &b)
	if err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	err = repos.DeleteBook(ctx, "book")
	if err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	book, errGet := repos.GetBookByName(ctx, "book")
	require.Errorf(t, errGet, "create error, book: %s", book)
}
