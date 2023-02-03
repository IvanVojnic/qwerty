// Package repository declare func for book
package repository

import (
	"EFpractic2/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// BookActPostgres is a wrapper to db object
type BookActPostgres struct {
	db *pgxpool.Pool
}

// NewBookActPostgres used to init BookAP
func NewBookActPostgres(db *pgxpool.Pool) *BookActPostgres {
	return &BookActPostgres{db: db}
}

// CreateBook used to create book
func (r *BookActPostgres) CreateBook(ctx context.Context, book *models.Book) error {
	_, err := r.db.Exec(ctx, "insert into books (name, year, new) values($1, $2, $3)",
		&book.BookName, &book.BookYear, &book.BookNew)
	if err != nil {
		return fmt.Errorf("error while book creating: %v", err)
	}
	return nil
}

func (r *BookActPostgres) GetBookId(ctx context.Context, bookName string) (int, error) {
	var idBook int
	err := r.db.QueryRow(ctx, "select books.id from books where books.name=$1", bookName).Scan(&idBook)
	if err != nil {
		return 0, fmt.Errorf("error: cannot get id, %w", err)
	}
	return idBook, nil
}

// UpdateBook used to update book
func (r *BookActPostgres) UpdateBook(ctx context.Context, book models.Book) error {
	_, err := r.db.Exec(ctx,
		`UPDATE books 
			SET "name" = $1, "year" = $2, "new" =$3 
			WHERE id = $4`,
		book.BookName, book.BookYear, book.BookNew, book.BookID)
	if err != nil {
		return fmt.Errorf("update book error %w", err)
	}
	return nil
}

// GetBook used to get book
func (r *BookActPostgres) GetBook(ctx context.Context, bookID int) (models.Book, error) {
	book := models.Book{}
	err := r.db.QueryRow(ctx,
		`select books.id, books.name, books.year, books.new 
			 from books 
			 where books.id=$1`,
		bookID).Scan(&book.BookID, &book.BookName, &book.BookYear, &book.BookNew)
	if err != nil {
		return book, fmt.Errorf("get book error %w", err)
	}
	return book, nil
}

// DeleteBook used to delete book
func (r *BookActPostgres) DeleteBook(ctx context.Context, bookID int) error {
	_, err := r.db.Exec(ctx, "delete from books where id=$1", bookID)
	if err != nil {
		return fmt.Errorf("delete book error %w", err)
	}
	return nil
}

// GetAllBooks used to get all books
func (r *BookActPostgres) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	books := make([]models.Book, 0)
	rows, err := r.db.Query(ctx, "select books.id, books.name, books.year, books.new, from books")
	if err != nil {
		return books, fmt.Errorf("get all books sql script error %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var book models.Book
		errScan := rows.Scan(&book.BookID, &book.BookName, &book.BookYear, &book.BookNew)
		if errScan != nil {
			return books, fmt.Errorf("get all books scan rows error %w", errScan)
		}
		books = append(books, book)
	}
	return books, err
}
