package util

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	Validator *validator.Validate
}

func (cu *CustomValidator) Validate(i interface{}) error {
	return cu.Validator.Struct(i)
}