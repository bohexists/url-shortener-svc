package validator

import (
	"github.com/go-playground/validator/v10"
)

// CustomValidator структура для валидации
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate метод для валидации входных данных
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

// NewValidator создает новый экземпляр валидатора
func NewValidator() *CustomValidator {
	return &CustomValidator{Validator: validator.New()}
}
