package token

import (
	"github.com/quabynah-bilson/quantia/pkg/token"
)

// RepositoryConfiguration is a function that configures a repository
type RepositoryConfiguration func(*Repository) error

// Repository is the token repository implementation
type Repository struct {
	token.Repository
	// @todo -> add token generation helper here
}

// NewRepository creates a new token repository
func NewRepository(configs ...RepositoryConfiguration) *Repository {
	r := &Repository{}

	for _, config := range configs {
		_ = config(r)
	}

	return r
}

// GenerateToken generates a token for the given username.
func (r *Repository) GenerateToken(username string) (string, error) {
	// @todo -> generate token here
	return "", nil
}

// ValidateToken validates the given token.
func (r *Repository) ValidateToken(token string) error {
	// @todo -> validate token here
	return nil
}

// InvalidateToken invalidates the given token.
func (r *Repository) InvalidateToken(token string) error {
	// @todo -> invalidate token here
	return nil
}
