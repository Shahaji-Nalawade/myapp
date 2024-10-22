package main

import (
	"log"
	"net/http"

	"github.com/Shahaji2016/myapp/internal"
)

func main() {
	// Set up the router
	r := internal.SetupRouter()

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
