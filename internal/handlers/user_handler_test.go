package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/Shahaji2016/myapp/internal/models"
	"github.com/Shahaji2016/myapp/internal/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	userService := services.NewUserService()
	userHandler := NewUserHandler(userService)

	tests := []struct {
		name               string
		payload            models.UserRequest
		expectedStatusCode int
		expectedMessage    string
	}{
		{
			name: "Valid User",
			payload: models.UserRequest{
				Name:   "John Doe",
				PAN:    "ABCDE1234F",
				Mobile: "9876543210",
				Email:  "john@example.com",
			},
			expectedStatusCode: http.StatusOK,
			expectedMessage:    "User data is valid",
		},
		{
			name: "Invalid PAN",
			payload: models.UserRequest{
				Name:   "Jane Doe",
				PAN:    "INVALIDPAN",
				Mobile: "9876543210",
				Email:  "jane@example.com",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedMessage:    "Key: 'UserRequest.PAN' Error:Field validation for 'PAN' failed on the 'pan' tag",
		},
		{
			name: "Invalid Mobile",
			payload: models.UserRequest{
				Name:   "Alice",
				PAN:    "ABCDE1234F",
				Mobile: "123",
				Email:  "alice@example.com",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedMessage:    "Key: 'UserRequest.Mobile' Error:Field validation for 'Mobile' failed on the 'len' tag",
		},
		{
			name: "Invalid Email",
			payload: models.UserRequest{
				Name:   "Bob",
				PAN:    "ABCDE1234F",
				Mobile: "9876543210",
				Email:  "not-an-email",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedMessage:    "Key: 'UserRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(payload))
			w := httptest.NewRecorder()

			userHandler.CreateUser(w, req)

			if w.Code != tt.expectedStatusCode {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatusCode, w.Code)
			}

			var response map[string]string
			json.Unmarshal(w.Body.Bytes(), &response)

			if msg, ok := response["message"]; ok && msg != tt.expectedMessage {
				t.Errorf("Expected message '%s', got '%s'", tt.expectedMessage, msg)
			}
		})
	}
}
