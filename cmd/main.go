package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/quabynah-bilson/quantia/adapters/database"
	"log"
	"os"
)

// entry point of the application
func main() {
	// Load .env file
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fmt.Println("Config loaded successfully")

	// @todo start server and initialize all components
	db := database.NewRedisDatabase(os.Getenv("REDIS_URI"))
	if db == nil {
		log.Fatalf("Error connecting to database")
	}

	fmt.Println("Database connected successfully")
}
