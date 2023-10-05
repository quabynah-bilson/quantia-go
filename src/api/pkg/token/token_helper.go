package token

// TokenizerHelper is the interface that wraps the basic token methods.
type TokenizerHelper interface {
	// GenerateToken generates a token for the given ID
	GenerateToken(claim string) (string, error)

	// ValidateToken validates the given token.
	ValidateToken(rawToken string) error
}
