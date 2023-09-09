package mocks

// MockTokenRepository is a mock of the token repository.
type MockTokenRepository struct {
	GenerateTokenFn   func(claim string) (string, error)
	ValidateTokenFn   func(rawToken, accountID string) error
	InvalidateTokenFn func(rawToken, accountID string) error
}

// GenerateToken mocks the generate token method.
func (m *MockTokenRepository) GenerateToken(claim string) (string, error) {
	return m.GenerateTokenFn(claim)
}

// ValidateToken mocks the validate token method.
func (m *MockTokenRepository) ValidateToken(rawToken, accountID string) error {
	return m.ValidateTokenFn(rawToken, accountID)
}

// InvalidateToken mocks the invalidate token method.
func (m *MockTokenRepository) InvalidateToken(rawToken, accountID string) error {
	return m.InvalidateTokenFn(rawToken, accountID)
}
