package account

import "errors"

var (
	// ErrInvalidPassword is returned when the password is invalid.
	ErrInvalidPassword = errors.New("invalid password")

	// ErrPasswordMismatch is returned when the password does not match.
	ErrPasswordMismatch = errors.New("password mismatch")
)

// PasswordHelper is the interface that wraps the basic password methods.
type PasswordHelper interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) error
}
