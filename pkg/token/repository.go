package token

// Repository is the interface that wraps the basic token methods.
type Repository interface {
	// GenerateToken generates a token for the given username.
	GenerateToken(username string) (string, error)

	// ValidateToken validates the given token.
	ValidateToken(token string) error

	// InvalidateToken invalidates the given token.
	InvalidateToken(token string) error
}
