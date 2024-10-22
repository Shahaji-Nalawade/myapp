package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ValidatePAN is a custom validation function to check PAN format.
func ValidatePAN(fl validator.FieldLevel) bool {
	panRegex := regexp.MustCompile(`^[A-Z]{5}[0-9]{4}[A-Z]{1}$`)
	return panRegex.MatchString(fl.Field().String())
}

// RegisterCustomValidations registers custom validation functions.
func RegisterCustomValidations(validate *validator.Validate) {
	validate.RegisterValidation("pan", ValidatePAN)
}
