package internal

import (
	"github.com/Shahaji2016/myapp/internal/handlers"
	"github.com/Shahaji2016/myapp/internal/middleware"
	"github.com/Shahaji2016/myapp/internal/services"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Create a new UserService
	userService := services.NewUserService()

	// Register middleware
	r.Use(middleware.LatencyLogger)

	// Register routes with UserHandler using UserService
	userHandler := handlers.NewUserHandler(userService)
	r.HandleFunc("/user", userHandler.CreateUser).Methods("POST")

	return r
}
