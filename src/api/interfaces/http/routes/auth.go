package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/quabynah-bilson/quantia/interfaces/http/handlers"
	"github.com/quabynah-bilson/quantia/pkg"
)

// SetupAuthRoutes is a function that registers all the routes for the auth group
// It uses Go's dependency injection to inject the auth use case into the handlers
func SetupAuthRoutes(route *gin.RouterGroup, useCase *pkg.AuthUseCase) {
	// create a new auth handler
	authHandler := handlers.NewAuthHandler(useCase)

	// register the auth routes
	route.POST("/register", authHandler.RegisterHandler)
	route.POST("/login", authHandler.LoginHandler)
	route.POST("/logout", authHandler.LogoutHandler)
}
