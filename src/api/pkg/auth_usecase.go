package pkg

import (
	"errors"
	"github.com/quabynah-bilson/quantia/pkg/account"
	"github.com/quabynah-bilson/quantia/pkg/token"
	"log"
	"regexp"
)

var (
	// ErrInvalidUsername is returned when the username is invalid.
	ErrInvalidUsername = errors.New("invalid username. username must be a valid email address or a valid phone number")

	// ErrInvalidPassword is returned when the password is invalid.
	ErrInvalidPassword = errors.New("invalid password. password must be at least 8 characters long")

	// ErrInvalidToken is returned when the token is invalid.
	ErrInvalidToken = errors.New("invalid token. token must be a valid JWT token")
)

// AuthUseCase is the auth use case. It contains the necessary repositories to perform auth operations.
type AuthUseCase struct {
	accountRepo account.Repository
	tokenRepo   token.Repository
}

// NewAuthUseCase creates a new account use case.
func NewAuthUseCase(authRepo account.Repository, tokenRepo token.Repository) *AuthUseCase {
	return &AuthUseCase{
		accountRepo: authRepo,
		tokenRepo:   tokenRepo,
	}
}

// Register registers a new user.
func (uc *AuthUseCase) Register(username string, password string) (*string, error) {
	if err := validateUsername(username); err != nil {
		log.Printf("error validating username: %v", err)
		return nil, err
	}

	if err := validatePassword(password); err != nil {
		log.Printf("error validating password: %v", err)
		return nil, err
	}

	userAccount, err := uc.accountRepo.Register(username, password)
	if err != nil {
		log.Printf("error registering user: %v", err)
		return nil, err
	}

	generatedToken, err := uc.tokenRepo.GenerateToken(userAccount.ID)
	if err != nil {
		log.Printf("error generating token: %v", err)
		return nil, err
	}

	return &generatedToken, nil
}

// Login logs in a user.
func (uc *AuthUseCase) Login(username string, password string) (*string, error) {
	if err := validateUsername(username); err != nil {
		log.Printf("error validating username: %v", err)
		return nil, err
	}

	if err := validatePassword(password); err != nil {
		log.Printf("error validating password: %v", err)
		return nil, err
	}

	userAccount, err := uc.accountRepo.Login(username, password)
	if err != nil {
		log.Printf("error logging in user: %v", err)
		return nil, err
	}

	generatedToken, err := uc.tokenRepo.GenerateToken(userAccount.ID)
	if err != nil {
		log.Printf("error generating token: %v", err)
		return nil, err
	}

	return &generatedToken, nil
}

// Logout logs out a user.
func (uc *AuthUseCase) Logout(rawToken, accountID string) error {
	if err := uc.tokenRepo.InvalidateToken(rawToken, accountID); err != nil {
		log.Printf("error invalidating token: %v", err)
		return ErrInvalidToken
	}

	return nil
}

// ValidateToken validates the given token.
func (uc *AuthUseCase) ValidateToken(rawToken, accountID string) error {
	if err := uc.tokenRepo.ValidateToken(rawToken, accountID); err != nil {
		log.Printf("error validating token: %v", err)
		return ErrInvalidToken
	}

	return nil
}

// validateUsername validates the given username.
func validateUsername(username string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	phoneRegex := regexp.MustCompile(`^\+[0-9]{11,}$`)

	if !emailRegex.MatchString(username) && !phoneRegex.MatchString(username) {
		return ErrInvalidUsername
	}

	return nil
}

// validatePassword validates the given password.
func validatePassword(password string) error {
	if len(password) < 8 {
		return ErrInvalidPassword
	}

	return nil
}
