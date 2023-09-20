package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/quabynah-bilson/quantia/interfaces/http/models"
	"github.com/quabynah-bilson/quantia/pkg"
	"net/http"
)

// PaymentHandler is a struct that holds the dependencies for the payment handlers
type PaymentHandler struct {
	useCase *pkg.PaymentUseCase
}

// NewPaymentHandler is a function that creates a new payment handler
func NewPaymentHandler(useCase *pkg.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{useCase: useCase}
}

// PayHandler is a function that handles the payment of a user
func (h *PaymentHandler) PayHandler(c *gin.Context) {
	// parse the request body into the MakePaymentRequest struct.
	// if there is an error, return a 400 Bad Request error
	var paymentReq models.MakePaymentRequest
	if err := c.ShouldBindJSON(&paymentReq); err != nil {
		c.JSON(http.StatusBadRequest, &models.APIResponse{Error: &models.APIError{
			Message: err.Error(),
			Code:    http.StatusBadRequest}},
		)
		return
	}

	// call the use case to make the payment
	transaction, err := h.useCase.MakePayment(paymentReq.Amount, paymentReq.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.APIResponse{Error: &models.APIError{
			Message: err.Error(),
			Code:    http.StatusBadRequest}},
		)
		return
	}

	// return a 201 Created response
	c.JSON(http.StatusCreated, &models.APIResponse{
		Success: true,
		Message: "Successfully made payment",
		Data: &models.MakePaymentResponse{
			Transaction: transaction,
		},
	})
}
