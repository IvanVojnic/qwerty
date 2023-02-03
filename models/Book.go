// Package models model Book
package models

// Book is a book
type Book struct {
	BookID   int    `json:"id" db:"id"`
	BookName string `json:"name" db:"name"`
	BookYear int    `json:"year" db:"year"`
	BookNew  bool   `json:"new" db:"new"`
}
