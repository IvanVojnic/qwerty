package service

import (
	"EFpractic2/models"
	"context"
)

type ImgUpload interface {
	CreateImg(context.Context, *models.Image) error
	GetImages(context.Context) ([]models.Image, error)
}

// ImgUpSrv wrapper for ImgUpSr repo
type ImgUpSrv struct {
	repo ImgUpload
}

// NewImgUpSrv used to init BookAP
func NewImgUpSrv(repo ImgUpload) *ImgUpSrv {
	return &ImgUpSrv{repo: repo}
}

// CreateImg used to create img
func (s *ImgUpSrv) CreateImg(ctx context.Context, img *models.Image) error {
	return s.repo.CreateImg(ctx, img)
}

func (s *ImgUpSrv) GetImages(ctx context.Context) ([]models.Image, error) {
	return s.repo.GetImages(ctx)
}
