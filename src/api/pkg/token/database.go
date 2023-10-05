package token

import "errors"

var (
	// ErrInvalidClaim is the error returned when an invalid claim is provided.
	ErrInvalidClaim = errors.New("invalid claim")

	// ErrInvalidToken is the error returned when an invalid token is provided.
	ErrInvalidToken = errors.New("invalid token")

	// ErrTokenExpired is the error returned when an expired token is provided.
	ErrTokenExpired = errors.New("token expired")

	// ErrTokenNotCreated is returned when the token could not be created.
	ErrTokenNotCreated = errors.New("token not created")

	// ErrCannotDeleteToken is returned when the token could not be deleted.
	ErrCannotDeleteToken = errors.New("cannot delete token")
)

// Database is the interface that wraps the basic token database operations.
type Database interface {
	// CreateToken generates a token for the given username.
	CreateToken(id string) (string, error)

	// ValidateToken validates the given token for the given account ID.
	ValidateToken(authToken, accountID string) error

	// DeleteToken invalidates the given account ID.
	DeleteToken(accountID string) error
}
