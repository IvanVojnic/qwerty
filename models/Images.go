package models

import "github.com/google/uuid"

// Image is an image
type Image struct {
	ImageID    uuid.UUID `json:"id" bson:"_id" db:"id"`
	ImageRoute string    `json:"name" bson:"name" db:"name"`
}
