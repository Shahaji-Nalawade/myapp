package services

import (
	"github.com/Shahaji2016/myapp/internal/models"
	"github.com/Shahaji2016/myapp/internal/validations" // Import the validations package

	"github.com/go-playground/validator/v10"
)

type UserService struct {
	validator *validator.Validate
}

func NewUserService() *UserService {
	validate := validator.New()
	validations.RegisterCustomValidations(validate) // Register custom validations
	return &UserService{validator: validate}
}

// ValidateUser validates the user data.
func (s *UserService) ValidateUser(req models.UserRequest) error {
	if err := s.validator.Struct(req); err != nil {
		return err
	}

	// Additional validations can be added here
	return nil
}
