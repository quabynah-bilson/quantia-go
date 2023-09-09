package account

import (
	"github.com/quabynah-bilson/quantia/pkg/account"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// PasswordHelper implements the PasswordHelper interface
type PasswordHelper struct {
	account.PasswordHelper
}

// NewBcryptPasswordHelper creates a new password helper that uses bcrypt
func NewBcryptPasswordHelper() account.PasswordHelper {
	return &PasswordHelper{}
}

// HashPassword hashes the given password
func (p *PasswordHelper) HashPassword(password string) (string, error) {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return "", account.ErrInvalidPassword
	}
	return string(result), nil
}

// ComparePassword compares the given password with the hashed password
func (p *PasswordHelper) ComparePassword(hashedPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashedPassword)
	if err := bcrypt.CompareHashAndPassword(hw, pw); err != nil {
		log.Printf("password mismatch: %v", err)
		return account.ErrPasswordMismatch
	}
	return nil
}
