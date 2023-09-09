package mocks

import (
	"errors"
	"github.com/quabynah-bilson/quantia/pkg/account"
)

var (
	//ExistingCustomerUsername is a user that already exists.
	ExistingCustomerUsername = "bilson@quantia.com"

	// NewCustomerUsername is a user that does not exist.
	NewCustomerUsername = "user@quantia.com"

	// ValidPassword is a valid password.
	ValidPassword = "password@1234"

	// ErrAlreadyExists is returned when the user already exists.
	ErrAlreadyExists = errors.New("user already exists")

	// ErrUserNotFound is returned when the user is not found.
	ErrUserNotFound = errors.New("user not found")

	// ErrAuthenticationFailed is returned when the authentication fails.
	ErrAuthenticationFailed = errors.New("authentication failed. invalid username or password")
)

// MockAccountRepository is a mock of the account repository.
type MockAccountRepository struct {
	LoginFn    func(username, password string) (*account.Account, error)
	RegisterFn func(username, password string) (*account.Account, error)
}

// Login mocks the login method.
func (m *MockAccountRepository) Login(username, password string) (*account.Account, error) {
	if username == NewCustomerUsername {
		return nil, ErrUserNotFound
	}

	if password != ValidPassword {
		return nil, ErrAuthenticationFailed
	}

	return m.LoginFn(username, password)
}

// Register mocks the register method.
func (m *MockAccountRepository) Register(username, password string) (*account.Account, error) {
	if username == ExistingCustomerUsername {
		return nil, ErrAlreadyExists
	}

	return m.RegisterFn(username, password)
}
