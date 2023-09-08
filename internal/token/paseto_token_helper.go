package token

import (
	"github.com/google/uuid"
	"github.com/quabynah-bilson/quantia/pkg/token"
)

// PasetoTokenizerHelper is the token helper implementation for Paseto.
type PasetoTokenizerHelper struct {
	token.TokenizerHelper
}

// WithPasetoTokenizerHelper sets the token helper implementation to Paseto.
func WithPasetoTokenizerHelper() RepositoryConfiguration {
	return func(r *Repository) error {
		r.generator = &PasetoTokenizerHelper{}
		return nil
	}
}

// GenerateToken generates a token for the given claim.
func (p *PasetoTokenizerHelper) GenerateToken(claim string) (string, error) {
	// @todo -> generate token here
	return uuid.NewString(), nil
}

// ValidateToken validates the given token.
func (p *PasetoTokenizerHelper) ValidateToken(token string) error {
	// @todo -> validate token here
	return nil
}

// InvalidateToken invalidates the given token.
func (p *PasetoTokenizerHelper) InvalidateToken(token string) error {
	// @todo -> invalidate token here
	return nil
}
