package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Shahaji2016/myapp/internal/models"
	"github.com/Shahaji2016/myapp/internal/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.UserRequest

	// Decode the JSON request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON provided", http.StatusBadRequest)
		return
	}

	// Use UserService to validate the user request
	if err := h.userService.ValidateUser(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User data is valid"})
}
