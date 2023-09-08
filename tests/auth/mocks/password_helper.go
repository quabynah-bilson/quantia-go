package mocks

import (
	"github.com/quabynah-bilson/quantia/pkg/account"
)

// MockPasswordHelper is a mock password helper
type MockPasswordHelper struct{}

// HashPassword hashes the given password
func (*MockPasswordHelper) HashPassword(password string) (string, error) {
	if password != "" {
		return "hashedpassword123", nil
	}
	return "", account.ErrInvalidPassword
}

// ComparePassword compares the given password with the hashed password
func (*MockPasswordHelper) ComparePassword(hashedPassword string, password string) error {
	if hashedPassword != "hashedpassword123" || password == "" {
		return account.ErrPasswordMismatch
	}
	return nil
}
