package utils

import (
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

// Init validator
func (cv *CustomValidator) Init() {
	cv.Validator.RegisterValidation("year", func(fl validator.FieldLevel) bool {
		var (
			lessZero       = false
			biggerThisYear = false
		)
		year := fl.Field().Int()
		if year < 0 {
			lessZero = true
		} else if year > 2023 {
			biggerThisYear = true
		}
		return lessZero && biggerThisYear
	})
}

// Validate Data
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
