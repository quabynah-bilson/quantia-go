package token

// Repository is the interface that wraps the basic token methods.
type Repository interface {
	// GenerateToken generates a token for the given claim.
	GenerateToken(claim string) (string, error)

	// ValidateToken validates the given token.
	ValidateToken(rawToken, accountID string) error

	// InvalidateToken invalidates the given token.
	InvalidateToken(rawToken, accountID string) error
}
