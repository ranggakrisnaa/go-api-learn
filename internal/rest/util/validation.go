package util

import (
	"gopkg.in/go-playground/validator.v9"
)

func NewValidator() *validator.Validate {
	return validator.New()
}

func ValidateStruct(m interface{}) (bool, error) {
	validate := NewValidator()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
