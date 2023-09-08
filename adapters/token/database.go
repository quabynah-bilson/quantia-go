package token

import "errors"

var (
	// ErrTokenNotCreated is returned when the token could not be created.
	ErrTokenNotCreated = errors.New("token not created")

	// ErrParsingSession is returned when the session could not be parsed.
	ErrParsingSession = errors.New("error parsing session")

	// ErrInvalidToken is returned when the token is invalid.
	//ErrInvalidToken = errors.New("invalid token")

	// ErrTokenExpired is returned when the token has expired.
	ErrTokenExpired = errors.New("token expired")
)

// Database is the interface that wraps the basic token database operations.
type Database interface {
	// CreateToken generates a token for the given username.
	CreateToken(id string) (string, error)

	// ValidateToken validates the given token.
	ValidateToken(authToken string) error

	// DeleteToken invalidates the given token.
	DeleteToken(authToken string) error
}
