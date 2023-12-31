package main

import (
	"github.com/joho/godotenv"
	"github.com/quabynah-bilson/quantia/interfaces/http"
	"github.com/quabynah-bilson/quantia/interfaces/webhook"
	"log"
)

// entry point of the application
func main() {
	// Load .env file
	if err := godotenv.Load("src/api/configs/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Start the auth server
	go http.StartAuthServer()

	// Start the webhook worker
	webhook.StartWebhookWorker()
}
