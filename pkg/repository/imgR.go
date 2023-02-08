package repository

import (
	"EFpractic2/models"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ImgUpPostgres is a wrapper to db object
type ImgPostgres struct {
	db *pgxpool.Pool
}

// NewImgUpPostgres used to init BookAP
func NewImgPostgres(db *pgxpool.Pool) *ImgPostgres {
	return &ImgPostgres{db: db}
}

// CreateBook used to create book
func (r *ImgPostgres) CreateImg(ctx context.Context, img *models.Image) error {
	ID := uuid.New()
	img.ImageID = ID
	_, err := r.db.Exec(ctx, "insert into images (route, id) values($1, $2)",
		&img.ImageRoute, &img.ImageID)
	if err != nil {
		return fmt.Errorf("error while img creating: %v", err)
	}
	return nil
}
