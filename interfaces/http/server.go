package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	accountAdapter "github.com/quabynah-bilson/quantia/adapters/account/datastore"
	tokenAdapter "github.com/quabynah-bilson/quantia/adapters/token/datastore"
	"github.com/quabynah-bilson/quantia/interfaces/http/routes"
	"github.com/quabynah-bilson/quantia/internal/account"
	"github.com/quabynah-bilson/quantia/internal/token"
	"github.com/quabynah-bilson/quantia/pkg"
	"log"
	"os"
)

// StartAuthServer is a function that starts the http server for the auth group using the gin framework
func StartAuthServer() {
	// create a new password helper utility
	pwHelper := account.NewBcryptPasswordHelper()

	// use the password helper utility to create a new account repository (with a database configuration)
	accountRepo := account.NewRepository(
		//accountAdapter.WithMongoAccountDatabase(os.Getenv("MONGO_URI"), pwHelper),
		accountAdapter.WithPostgresAccountDatabase(os.Getenv("POSTGRES_URI"), pwHelper),
	)

	// create a new token repository (with a database configuration)
	tokenRepo := token.NewRepository(
		tokenAdapter.WithRedisTokenDatabase(os.Getenv("REDIS_URI")),
	)

	// create a new auth use case
	authUseCase := pkg.NewAuthUseCase(accountRepo, tokenRepo)

	// create a gin router
	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	// create a group for the auth routes
	authRoutes := router.Group("/api/v1/auth")

	// register the auth routes
	routes.SetupAuthRoutes(authRoutes, authUseCase)

	// start the server
	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
