package account

import (
	"errors"
	"github.com/quabynah-bilson/quantia/pkg/account"
)

var (
	// ErrAccountNotFound is the error returned when an account is not found.
	ErrAccountNotFound = errors.New("account not found")

	// ErrAccountAlreadyExists is the error returned when an account already exists.
	ErrAccountAlreadyExists = errors.New("account already exists")

	// ErrAccountNotCreated is the error returned when an account is not created.
	ErrAccountNotCreated = errors.New("account not created. Please try again")

	// ErrAccountNotDeleted is the error returned when an account is not deleted.
	ErrAccountNotDeleted = errors.New("account not deleted. Please try again")

	// ErrInvalidID is the error returned when an ID is invalid.
	ErrInvalidID = errors.New("invalid ID. Please check and try again")

	// ErrInvalidCredentials is the error returned when credentials are invalid.
	ErrInvalidCredentials = errors.New("invalid credentials. Please check and try again")
)

// Database is the interface that wraps the basic account database operations.
type Database interface {
	// GetAccount gets an account by ID
	GetAccount(id string) (*account.Account, error)

	// GetAccountByUsernameAndPassword gets an account by username and password
	GetAccountByUsernameAndPassword(username, password string) (*account.Account, error)

	// CreateAccount creates a new account
	CreateAccount(username, password string) (*account.Account, error)

	// DeleteAccount deletes an account by ID
	DeleteAccount(id string) error
}
