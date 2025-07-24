package internal_errors

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}
	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]

	field := strings.ToLower(validationError.StructField())

	switch validationError.Tag() {
	case "required":
		return errors.New(fmt.Sprintf("%s is required", field))
	case "max":
		return errors.New(fmt.Sprintf("%s must be less than %s", field, validationError.Param()))
	case "min":
		return errors.New(fmt.Sprintf("%s must be greater than %s", field, validationError.Param()))
	case "email":
		return errors.New(fmt.Sprintf("%s is invalid", field))
	}
	return nil
}
