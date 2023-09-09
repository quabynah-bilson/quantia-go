package token

import "errors"

var (
	// ErrTokenNotCreated is returned when the token could not be created.
	ErrTokenNotCreated = errors.New("token not created")

	// ErrInvalidToken is returned when the token is invalid.
	ErrInvalidToken = errors.New("invalid token")

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
