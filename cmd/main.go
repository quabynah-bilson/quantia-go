package main

import (
	"github.com/joho/godotenv"
	"github.com/quabynah-bilson/quantia/adapters"
	ads "github.com/quabynah-bilson/quantia/adapters/account/datastore"
	tds "github.com/quabynah-bilson/quantia/adapters/token/datastore"
	"github.com/quabynah-bilson/quantia/internal/auth"
	"log"
	"os"
)

// entry point of the application
func main() {
	// Load .env file
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// @todo start server and initialize all components
	pwHelper := auth.NewPasswordHelper()
	adpt := adapters.NewAdapter(
		ads.WithPostgresAccountDatabase(os.Getenv("POSTGRES_URI"), pwHelper),
		tds.WithRedisTokenDatabase(os.Getenv("REDIS_URI")),
	)

	account, err := adpt.AccountDB.GetAccountByUsernameAndPassword("quabynah@gmail.com", "password")
	if err != nil {
		log.Fatalf("error creating account: %v", err)
	}

	token, err := adpt.TokenDB.CreateToken(account.ID)
	if err != nil {
		log.Fatalf("error creating token: %v", err)
	}

	log.Printf("account: %+v & token: %s", account, token)
}
