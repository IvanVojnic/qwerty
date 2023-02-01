package models

import "github.com/google/uuid"

type Book struct {
	BookID   uuid.UUID `json:"id" db:"id"`
	BookName string    `json:"name" db:"name"`
	BookYear int       `json:"year" db:"year"`
	BookNew  bool      `json:"new" db:"new"`
}
