package token

import "errors"

var (
	// ErrInvalidClaim is the error returned when an invalid claim is provided.
	ErrInvalidClaim = errors.New("invalid claim")

	// ErrInvalidToken is the error returned when an invalid token is provided.
	ErrInvalidToken = errors.New("invalid token")

	// ErrTokenExpired is the error returned when an expired token is provided.
	ErrTokenExpired = errors.New("token expired")
)

// TokenizerHelper is the interface that wraps the basic token methods.
type TokenizerHelper interface {
	// GenerateToken generates a token for the given ID
	GenerateToken(claim string) (string, error)

	// ValidateToken validates the given token.
	ValidateToken(rawToken string) error
}
