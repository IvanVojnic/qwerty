// Package models UserAuth model
package models

import "github.com/google/uuid"

// UserAuth is a UserAuth
type UserAuth struct {
	UserID        uuid.UUID `json:"id" db:"id"`
	UserName      string    `json:"name" db:"name"`
	UserAge       int       `json:"age" db:"age"`
	UserIsRegular bool      `json:"isRegular" db:"regular"`
	Password      string    `json:"password" db:"password"`
	RefreshToken  string    `json:"refreshToken" db:"refreshToken"`
}
