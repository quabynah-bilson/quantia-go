package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/quabynah-bilson/quantia/interfaces/http/handlers"
	"github.com/quabynah-bilson/quantia/pkg"
)

// SetupPaymentRoutes is a function that sets up the payment routes
func SetupPaymentRoutes(router *gin.RouterGroup, paymentUseCase *pkg.PaymentUseCase) {
	// create a new payment handler
	pay := handlers.NewPaymentHandler(paymentUseCase)

	// set up the routes
	router.POST("/pay", pay.PayHandler)
}
