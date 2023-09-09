package main

import (
	"github.com/joho/godotenv"
	accountAdapter "github.com/quabynah-bilson/quantia/adapters/account/datastore"
	tokenAdapter "github.com/quabynah-bilson/quantia/adapters/token/datastore"
	"github.com/quabynah-bilson/quantia/internal/account"
	"github.com/quabynah-bilson/quantia/internal/token"
	"github.com/quabynah-bilson/quantia/pkg"
	"log"
	"os"
)

// entry point of the application
func main() {
	// Load .env file
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// @TODO: Initialize the server
	pwHelper := account.NewBcryptPasswordHelper()
	accountRepo := account.NewRepository(
		accountAdapter.WithMongoAccountDatabase(os.Getenv("MONGO_URI"), pwHelper),
		//accountAdapter.WithPostgresAccountDatabase(os.Getenv("POSTGRES_URI"), pwHelper),
	)
	tokenRepo := token.NewRepository(
		tokenAdapter.WithRedisTokenDatabase(os.Getenv("REDIS_URI")),
	)
	usc := pkg.NewAuthUseCase(accountRepo, tokenRepo)
	authToken, err := usc.Login("quabynah@gmail.com", "password")
	if err != nil {
		log.Fatalf("error logging in: %v", err)
	}
	log.Printf("auth token: %v", *authToken)
}
