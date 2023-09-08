package token

import (
	"github.com/quabynah-bilson/quantia/pkg/token"
)

// RepositoryConfiguration is a function that configures a repository
type RepositoryConfiguration func(*Repository) error

// Repository is the token repository implementation
type Repository struct {
	token.Repository
	generator token.TokenizerHelper
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
func (r *Repository) GenerateToken(claim string) (string, error) {
	return r.generator.GenerateToken(claim)
}

// ValidateToken validates the given token.
func (r *Repository) ValidateToken(token string) error {
	return r.generator.ValidateToken(token)
}

// InvalidateToken invalidates the given token.
func (r *Repository) InvalidateToken(token string) error {
	return r.generator.InvalidateToken(token)
}
