package service

import (
	"EFpractic2/models"
	"EFpractic2/pkg/repository"
	"context"
)

type BookActSrv struct {
	repo repository.BookAct
}

func NewBookActSrv(repo repository.BookAct) *BookActSrv {
	return &BookActSrv{repo: repo}
}

func (s *BookActSrv) CreateBook(ctx context.Context, book models.Book) error {
	return s.repo.CreateBook(ctx, book)
}

func (s *BookActSrv) UpdateBook(ctx context.Context, book models.Book) error {
	return s.repo.UpdateBook(ctx, book)
}

func (s *BookActSrv) GetBook(ctx context.Context, bookId int) (models.Book, error) {
	return s.repo.GetBook(ctx, bookId)
}

func (s *BookActSrv) DeleteBook(ctx context.Context, bookId int) error {
	return s.repo.DeleteBook(ctx, bookId)
}

func (s *BookActSrv) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	return s.repo.GetAllBooks(ctx)
}
