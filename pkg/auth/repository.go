package auth

// AccountRepository is the interface that wraps the basic account methods.
type AccountRepository interface {
	// Register registers a new user.
	Register(username string, password string) error

	// Login logs in a user.
	Login(username string, password string) error
}

// TokenRepository is the interface that wraps the basic token methods.
type TokenRepository interface {
	// GenerateToken generates a token for the given username.
	GenerateToken(username string) (string, error)

	// ValidateToken validates the given token.
	ValidateToken(token string) error

	// InvalidateToken invalidates the given token.
	InvalidateToken(token string) error
}
