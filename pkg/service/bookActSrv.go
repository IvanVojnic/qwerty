// Package service for book
package service

import (
	"context"

	"EFpractic2/models"
	"EFpractic2/pkg/repository"
)

// BookActSrv wrapper for bookAP repo
type BookActSrv struct {
	repo repository.BookActPostgres
}

// NewBookActSrv used to init BookAP
func NewBookActSrv(repo repository.BookActPostgres) *BookActSrv {
	return &BookActSrv{repo: repo}
}

// CreateBook used to create book
func (s *BookActSrv) CreateBook(ctx context.Context, book models.Book) error {
	return s.repo.CreateBook(ctx, book)
}

// UpdateBook used update book
func (s *BookActSrv) UpdateBook(ctx context.Context, book models.Book) error {
	return s.repo.UpdateBook(ctx, book)
}

// GetBook used get book
func (s *BookActSrv) GetBook(ctx context.Context, bookID int) (models.Book, error) {
	return s.repo.GetBook(ctx, bookID)
}

// DeleteBook used delete book
func (s *BookActSrv) DeleteBook(ctx context.Context, bookID int) error {
	return s.repo.DeleteBook(ctx, bookID)
}

// GetAllBooks used get all books
func (s *BookActSrv) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	return s.repo.GetAllBooks(ctx)
}
