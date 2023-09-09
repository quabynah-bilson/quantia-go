package token

import (
	"github.com/quabynah-bilson/quantia/pkg/token"
)

// RepositoryConfiguration is a function that configures a repository
type RepositoryConfiguration func(*Repository) error

// Repository is the token repository implementation
type Repository struct {
	DB token.Database
	token.Repository
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
	return r.DB.CreateToken(claim)
}

// ValidateToken validates the given token.
func (r *Repository) ValidateToken(rawToken, accountID string) error {
	return r.DB.ValidateToken(rawToken, accountID)
}

// InvalidateToken invalidates the given account's ID
func (r *Repository) InvalidateToken(_, accountID string) error {
	return r.DB.DeleteToken(accountID)
}
