package repository

import (
	"EFpractic2/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

type BookActPostgres struct {
	db *pgxpool.Pool
}

func NewBookActPostgres(db *pgxpool.Pool) *BookActPostgres {
	return &BookActPostgres{db: db}
}

func (r *BookActPostgres) CreateBook(ctx context.Context, book models.Book) error {
	_, err := r.db.Exec(ctx, "insert into books (name, age, regular, password) values($1, $2, $3, $4)",
		book.BookName, book.BookYear, book.BookNew)
	if err != nil {
		return fmt.Errorf("error while book creating: %v", err)
	}
	return nil
}

func (r *BookActPostgres) UpdateBook(ctx context.Context, book models.Book) error {
	_, err := r.db.Exec(ctx, "UPDATE books SET name = $1, age = $2, regular =$3 WHERE id = $4", book.BookName, book.BookYear, book.BookNew, book.BookID)
	if err != nil {
		return fmt.Errorf("update book error %w", err)
	}
	return nil
}

func (r *BookActPostgres) GetBook(ctx context.Context, bookId int) (models.Book, error) {
	book := models.Book{}
	err := r.db.QueryRow(ctx, "select books.id, books.name, books.year, book.new, from books where id=$1", bookId).Scan(
		&book.BookID, &book.BookName, &book.BookYear, &book.BookNew)
	if err != nil {
		return book, fmt.Errorf("get book error %w", err)
	}
	return book, nil
}

func (r *BookActPostgres) DeleteBook(ctx context.Context, bookId int) error {
	_, err := r.db.Exec(ctx, "delete from books where id=$1", bookId)
	if err != nil {
		return fmt.Errorf("delete book error %w", err)
	}
	return nil
}

func (r *BookActPostgres) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	books := make([]models.Book, 0)
	rows, err := r.db.Query(ctx, "select books.id, books.name, books.age, books.regular, from books")
	if err != nil {
		log.WithFields(log.Fields{
			"Error get all book": err,
			"rows":               rows,
		}).Info("SQL QUERY")
	}
	defer rows.Close()
	for rows.Next() {
		var book models.Book
		errScan := rows.Scan(&book.BookID, &book.BookName, &book.BookYear, &book.BookNew)
		if errScan != nil {
			log.WithFields(log.Fields{
				"Error while scan current row to get book model": err,
				"book": book,
			}).Info("SCAN ERROR. GET ALL BOOKS")
		}
		books = append(books, book)
	}
	if errRows := rows.Err(); errRows != nil {
		fmt.Errorf("get all books error %w", errRows)
	}
	return books, err
}
