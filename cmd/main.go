package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

// entry point of the application
func main() {
	// Load .env file
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fmt.Println("Config loaded successfully")

	// @todo start server and initialize all components

}
