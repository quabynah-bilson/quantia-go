package token

// TokenizerHelper is the interface that wraps the basic token methods.
type TokenizerHelper interface {
	// GenerateToken generates a token for the given ID
	GenerateToken(ID string) (string, error)
}
