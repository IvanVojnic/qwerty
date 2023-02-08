package models

import "github.com/google/uuid"

// BookImage is an image
type BookImage struct {
	ImageID uuid.UUID `json:"imgId" bson:"_ImgId" db:"imgId"`
	BookID  uuid.UUID `json:"bookId" bson:"_bookId" db:"bookId"`
}
