package main

import (
	"github.com/joho/godotenv"
	"github.com/quabynah-bilson/quantia/interfaces/http"
	"log"
)

// entry point of the application
func main() {
	// Load .env file
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Start the auth server
	http.StartAuthServer()
}
