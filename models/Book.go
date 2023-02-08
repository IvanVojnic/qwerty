// Package models model Book
package models

import "github.com/google/uuid"

// Book is a book
type Book struct {
	BookID   uuid.UUID `json:"id" bson:"_id" db:"id"`
	BookName string    `json:"name" bson:"name" db:"name"`
	BookYear int       `json:"year" bson:"year" db:"year"`
	BookNew  bool      `json:"new" bson:"new" db:"new"`
}
