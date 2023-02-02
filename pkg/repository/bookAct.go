// Package repository declare func for book
package repository

import (
	"context"
	"fmt"

	"EFpractic2/models"

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
func (r *BookActPostgres) CreateBook(ctx context.Context, book models.Book) (error, error) {
	_, err := r.db.Exec(ctx, "insert into books (name, age, regular, password) values($1, $2, $3, $4)",
		book.BookName, book.BookYear, book.BookNew)
	if err != nil {
		return fmt.Errorf("error while book creating: %v", err), nil
	}
	return nil, nil
}

// UpdateBook used to update book
func (r *BookActPostgres) UpdateBook(ctx context.Context, book models.Book) error {
	_, err := r.db.Exec(ctx, "UPDATE books SET name = $1, age = $2, regular =$3 WHERE id = $4", book.BookName, book.BookYear, book.BookNew, book.BookID)
	if err != nil {
		return fmt.Errorf("update book error %w", err)
	}
	return nil
}

// GetBook used to get book
func (r *BookActPostgres) GetBook(ctx context.Context, bookID int) (models.Book, error) {
	book := models.Book{}
	err := r.db.QueryRow(ctx, "select books.id, books.name, books.year, book.new, from books where id=$1", bookID).Scan(
		&book.BookID, &book.BookName, &book.BookYear, &book.BookNew)
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
	rows, err := r.db.Query(ctx, "select books.id, books.name, books.age, books.regular, from books")
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
