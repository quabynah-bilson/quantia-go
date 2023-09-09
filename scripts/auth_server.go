package scripts

import (
	accountAdapter "github.com/quabynah-bilson/quantia/adapters/account/datastore"
	tokenAdapter "github.com/quabynah-bilson/quantia/adapters/token/datastore"
	"github.com/quabynah-bilson/quantia/internal/account"
	"github.com/quabynah-bilson/quantia/internal/token"
	"github.com/quabynah-bilson/quantia/pkg"
	"os"
)

// StartAuthServer starts the auth server.
// It creates a new account repository, token repository and auth use case.
// It also uses Go's dependency injection to inject the account and token repositories into the auth use case.
func StartAuthServer() {
	// Create a new password helper utility
	pwHelper := account.NewBcryptPasswordHelper()

	// Use the password helper utility to create a new account repository (with a database configuration)
	accountRepo := account.NewRepository(
		accountAdapter.WithMongoAccountDatabase(os.Getenv("MONGO_URI"), pwHelper),
		//accountAdapter.WithPostgresAccountDatabase(os.Getenv("POSTGRES_URI"), pwHelper),
	)

	// Create a new token repository (with a database configuration)
	tokenRepo := token.NewRepository(
		tokenAdapter.WithRedisTokenDatabase(os.Getenv("REDIS_URI")),
	)

	// Create a new auth use case
	_ = pkg.NewAuthUseCase(accountRepo, tokenRepo)

	// @todo Start the auth HTTP server with the use case

}
