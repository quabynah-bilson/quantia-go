package main

import (
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/quabynah-bilson/quantia/internal"
	"github.com/quabynah-bilson/quantia/internal/token"
	"log"
)

// entry point of the application
func main() {
	// Load .env file
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// @todo start server and initialize all components
	//pwHelper := account.NewPasswordHelper()
	//_ = adapters.NewAdapter(
	//	ads.WithPostgresAccountDatabase(os.Getenv("POSTGRES_URI"), pwHelper),
	//	tds.WithRedisTokenDatabase(os.Getenv("REDIS_URI")),
	//)

	// create service (to be used in the route handler)
	svc, err := internal.NewService(
		internal.WithTokenRepo(token.NewRepository(
			token.WithPasetoTokenizerHelper(),
		)),
	)
	if err != nil {
		log.Fatalf("error creating service: %v", err)
	}

	// generate token
	generateToken, err := svc.TokenRepo.GenerateToken(uuid.NewString())
	if err != nil {
		log.Fatalf("error generating token: %v", err)
	}
	log.Printf("token: %s", generateToken)

	//account, err := adpt.AccountDB.GetAccountByUsernameAndPassword("quabynah@gmail.com", "password")
	//if err != nil {
	//	log.Fatalf("error creating account: %v", err)
	//}

	//token, err := adpt.TokenDB.CreateToken(account.ID)
	//if err != nil {
	//	log.Fatalf("error creating token: %v", err)
	//}

	//log.Printf("account: %+v & token: %s", account, token)
}
