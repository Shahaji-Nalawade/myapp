package services

import (
	"testing"

	"github.com/Shahaji2016/myapp/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestValidateUser(t *testing.T) {
	userService := NewUserService()

	validUser := models.UserRequest{
		Name:   "John Doe",
		PAN:    "ABCDE1234F",
		Mobile: "9876543210",
		Email:  "john.doe@example.com",
	}

	invalidPAN := models.UserRequest{
		Name:   "John Doe",
		PAN:    "1234ABCDE",
		Mobile: "9876543210",
		Email:  "john.doe@example.com",
	}

	invalidMobile := models.UserRequest{
		Name:   "John Doe",
		PAN:    "ABCDE1234F",
		Mobile: "98765",
		Email:  "john.doe@example.com",
	}

	invalidEmail := models.UserRequest{
		Name:   "John Doe",
		PAN:    "ABCDE1234F",
		Mobile: "9876543210",
		Email:  "invalid-email",
	}

	t.Run("Valid user data", func(t *testing.T) {
		err := userService.ValidateUser(validUser)
		assert.NoError(t, err)
	})

	t.Run("Invalid PAN format", func(t *testing.T) {
		err := userService.ValidateUser(invalidPAN)
		assert.Error(t, err)
	})

	t.Run("Invalid mobile number", func(t *testing.T) {
		err := userService.ValidateUser(invalidMobile)
		assert.Error(t, err)
	})

	t.Run("Invalid email address", func(t *testing.T) {
		err := userService.ValidateUser(invalidEmail)
		assert.Error(t, err)
	})
}
