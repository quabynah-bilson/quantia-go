package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	accountAdapter "github.com/quabynah-bilson/quantia/adapters/account/datastore"
	paymentAdapter "github.com/quabynah-bilson/quantia/adapters/payment/datastore"
	tokenAdapter "github.com/quabynah-bilson/quantia/adapters/token/datastore"
	"github.com/quabynah-bilson/quantia/interfaces/http/routes"
	"github.com/quabynah-bilson/quantia/internal/account"
	"github.com/quabynah-bilson/quantia/internal/payment"
	"github.com/quabynah-bilson/quantia/internal/token"
	"github.com/quabynah-bilson/quantia/pkg"
	"log"
	"os"
)

// StartAuthServer is a function that starts the http server for the auth group using the gin framework
func StartAuthServer() {

	// create a gin router
	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	// create a group for the auth routes
	authRoutes := router.Group("/api/v1/auth")

	// register the auth routes
	routes.SetupAuthRoutes(authRoutes, setupAuth())

	// create a group for the payment routes
	paymentRoutes := router.Group("/api/v1/payments")

	// register the payment routes
	routes.SetupPaymentRoutes(paymentRoutes, setupPayment())

	// start the server
	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

// setupAuth is a function that sets up the auth use case
func setupAuth() *pkg.AuthUseCase {
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

	return authUseCase
}

// setupPayment is a function that sets up the payment use case
func setupPayment() *pkg.PaymentUseCase {
	// create a new payment repository (with a database configuration)
	paymentRepo := payment.NewRepository(
		paymentAdapter.WithRedisPaymentDatabase(os.Getenv("REDIS_URI")),
	)

	// create a new payment use case
	paymentUseCase := pkg.NewPaymentUseCase(paymentRepo)

	return paymentUseCase
}
