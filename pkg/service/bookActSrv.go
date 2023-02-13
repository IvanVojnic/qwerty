// Package service for book
package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"

	"EFpractic2/models"
)

// BookAct interface consists of methos to communicate with boockAct repo
type BookAct interface {
	GetBookId(ctx context.Context, bookName string) (uuid.UUID, error)
	CreateBook(context.Context, *models.Book) error
	UpdateBook(context.Context, models.Book) error
	DeleteBook(context.Context, string) error
	GetAllBooks(context.Context) ([]models.Book, error)
	GetBookByName(context.Context, string) (models.Book, error)
}

type BookCache interface {
	GetBook(ctx context.Context, bookName string) (models.Book, error)
	CacheBook(ctx context.Context, book *models.Book) error
}

// BookActSrv wrapper for bookAP repo
type BookActSrv struct {
	repo  BookAct
	cache BookCache
}

// NewBookActSrv used to init BookAP
func NewBookActSrv(repo BookAct, cache BookCache) *BookActSrv {
	return &BookActSrv{repo: repo, cache: cache}
}

// CreateBook used to create book
func (s *BookActSrv) CreateBook(ctx context.Context, book models.Book) error {
	err := s.cache.CacheBook(ctx, &book)
	if err != nil {
		return fmt.Errorf("error while caching, %s", err)
	}
	return s.repo.CreateBook(ctx, &book)
}

func (s *BookActSrv) GetBookId(ctx context.Context, bookName string) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

// UpdateBook used update book
func (s *BookActSrv) UpdateBook(ctx context.Context, book models.Book) error {
	return s.repo.UpdateBook(ctx, book)
}

// DeleteBook used delete book
func (s *BookActSrv) DeleteBook(ctx context.Context, bookName string) error {
	return s.repo.DeleteBook(ctx, bookName)
}

// GetAllBooks used get all books
func (s *BookActSrv) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	return s.repo.GetAllBooks(ctx)
}

func (s *BookActSrv) GetBookByName(ctx context.Context, bookName string) (models.Book, error) {
	book, err := s.cache.GetBook(ctx, bookName)
	if err != nil {
		return s.repo.GetBookByName(ctx, bookName)
	}
	return book, nil
}
