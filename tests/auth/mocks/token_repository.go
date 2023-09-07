package mocks

// MockTokenRepository is a mock of the token repository.
type MockTokenRepository struct {
	GenerateTokenFn   func(username string) (string, error)
	ValidateTokenFn   func(token string) error
	InvalidateTokenFn func(token string) error
}

// GenerateToken mocks the generate token method.
func (m *MockTokenRepository) GenerateToken(username string) (string, error) {
	return m.GenerateTokenFn(username)
}

// ValidateToken mocks the validate token method.
func (m *MockTokenRepository) ValidateToken(token string) error {
	return m.ValidateTokenFn(token)
}

// InvalidateToken mocks the invalidate token method.
func (m *MockTokenRepository) InvalidateToken(token string) error {
	return m.InvalidateTokenFn(token)
}
