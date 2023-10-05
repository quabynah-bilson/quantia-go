package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/quabynah-bilson/quantia/interfaces/http/models"
	"github.com/quabynah-bilson/quantia/pkg"
	"net/http"
)

// AuthHandler is a struct that holds the dependencies for the auth handlers
// It uses Go's dependency injection to inject the auth use case into the handlers
type AuthHandler struct {
	useCase *pkg.AuthUseCase
}

// NewAuthHandler is a function that creates a new auth handler
func NewAuthHandler(useCase *pkg.AuthUseCase) *AuthHandler {
	return &AuthHandler{useCase: useCase}
}

// RegisterHandler is a function that handles the registration of a user
func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	// parse the request body into the RegistrationRequest struct.
	// if there is an error, return a 400 Bad Request error
	var regReq models.AuthenticationRequest
	if err := c.ShouldBindJSON(&regReq); err != nil {
		c.JSON(http.StatusBadRequest, &models.APIResponse{Error: &models.APIError{
			Message: err.Error(),
			Code:    http.StatusBadRequest}},
		)
		return
	}

	// call the use case to register the user
	authToken, err := h.useCase.Register(regReq.Username, regReq.Password)
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
		Message: "Successfully registered",
		Data: &models.AuthenticationResponse{
			AccessToken: *authToken,
		},
	})
}

// LoginHandler is a function that handles the login of a user
func (h *AuthHandler) LoginHandler(c *gin.Context) {
	// parse the request body into the AuthenticationRequest struct.
	// if there is an error, return a 400 Bad Request error
	var authReq models.AuthenticationRequest
	if err := c.ShouldBindJSON(&authReq); err != nil {
		c.JSON(http.StatusBadRequest, &models.APIResponse{Error: &models.APIError{
			Message: err.Error(),
			Code:    http.StatusBadRequest}},
		)
		return
	}

	// call the use case to authenticate the user
	authToken, err := h.useCase.Login(authReq.Username, authReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &models.APIResponse{Error: &models.APIError{
			Message: err.Error(),
			Code:    http.StatusUnauthorized}},
		)
		return
	}

	// return the auth token to the user
	c.JSON(http.StatusOK, &models.APIResponse{
		Success: true,
		Data: &models.AuthenticationResponse{
			AccessToken: *authToken,
		},
		Message: "Successfully logged in",
	})
}

// LogoutHandler is a function that handles the logout of a user
func (h *AuthHandler) LogoutHandler(c *gin.Context) {
	// parse the request bearer token
	// if there is an error, return a 401 Unauthorized error
	authToken := c.GetHeader("Authorization")
	if len(authToken) == 0 {
		c.JSON(http.StatusUnauthorized, &models.APIResponse{Error: &models.APIError{
			Message: "No authorization token provided",
			Code:    http.StatusUnauthorized}},
		)
		return
	}

	// remove the "Bearer " prefix from the token
	authToken = authToken[7:]

	// parse the request body into the LogoutRequest struct.
	// if there is an error, return a 400 Bad Request error
	var logoutReq models.LogoutRequest
	if err := c.ShouldBindJSON(&logoutReq); err != nil {
		c.JSON(http.StatusBadRequest, &models.APIResponse{Error: &models.APIError{
			Message: err.Error(),
			Code:    http.StatusBadRequest}},
		)
		return
	}

	// call the use case to log out the user
	if err := h.useCase.Logout(authToken, logoutReq.AccountID); err != nil {
		c.JSON(http.StatusUnauthorized, &models.APIResponse{Error: &models.APIError{
			Message: err.Error(),
			Code:    http.StatusUnauthorized}},
		)
		return
	}

	// return a 200 OK response
	c.JSON(http.StatusOK, &models.APIResponse{
		Success: true,
		Message: "Successfully logged out",
	})
}
