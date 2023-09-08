package mocks

import (
	"github.com/quabynah-bilson/quantia/pkg/token"
)

const (
	SuggestedToken   = "token123"
	InvalidatedToken = "token456"
)

// MockTokenizerHelper is the token helper implementation for testing.
type MockTokenizerHelper struct {
	tokens map[string]bool
}

// NewMockTokenizerHelper creates a new mock token helper.
func NewMockTokenizerHelper() *MockTokenizerHelper {
	return &MockTokenizerHelper{tokens: make(map[string]bool)}
}

// GenerateToken generates a token for the given claim.
func (m *MockTokenizerHelper) GenerateToken(claim string) (string, error) {
	if claim != "" {
		m.tokens[SuggestedToken] = true
		return SuggestedToken, nil
	}
	return "", token.ErrInvalidClaim
}

// ValidateToken validates the given token.
func (m *MockTokenizerHelper) ValidateToken(suggestedToken string) error {
	if !m.tokens[suggestedToken] {
		return token.ErrInvalidToken
	}
	return nil
}

// InvalidateToken invalidates the given token.
func (m *MockTokenizerHelper) InvalidateToken(suggestedToken string) error {
	if !m.tokens[suggestedToken] {
		return token.ErrTokenRevoked
	}

	delete(m.tokens, suggestedToken)
	return nil
}
