package token

import "github.com/quabynah-bilson/quantia/pkg/token"

// PasetoTokenizerHelper is the token helper implementation for Paseto.
type PasetoTokenizerHelper struct {
	token.TokenizerHelper
}

// NewPasetoTokenizerHelper creates a new Paseto token helper.
func NewPasetoTokenizerHelper() token.TokenizerHelper {
	return &PasetoTokenizerHelper{}
}

// GenerateToken generates a token for the given ID.
func (p *PasetoTokenizerHelper) GenerateToken(ID string) (string, error) {
	// @todo -> generate token here
	return "", nil
}
