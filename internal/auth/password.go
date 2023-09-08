package auth

import (
	"github.com/quabynah-bilson/quantia/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

// PasswordHelper implements the PasswordHelper interface
type PasswordHelper struct {
	auth.PasswordHelper
}

// NewPasswordHelper creates a new password helper
func NewPasswordHelper() auth.PasswordHelper {
	return &PasswordHelper{}
}

// HashPassword hashes the given password
func (p *PasswordHelper) HashPassword(password string) (string, error) {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// ComparePassword compares the given password with the hashed password
func (p *PasswordHelper) ComparePassword(hashedPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashedPassword)
	return bcrypt.CompareHashAndPassword(hw, pw)
}
