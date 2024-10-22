package models

type UserRequest struct {
	Name   string `json:"name" validate:"required"`
	PAN    string `json:"pan" validate:"required,pan"`       // Custom PAN validation
	Mobile string `json:"mobile" validate:"required,len=10"` // Ensure it's exactly 10 digits
	Email  string `json:"email" validate:"required,email"`   // Standard email validation
}
