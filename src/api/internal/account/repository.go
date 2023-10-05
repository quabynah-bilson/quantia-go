package account

import (
	"github.com/quabynah-bilson/quantia/pkg/account"
)

// RepositoryConfiguration is a function that configures a repository
type RepositoryConfiguration func(*Repository) error

// Repository is the account repository implementation
type Repository struct {
	DB account.Database
	account.Repository
}

// NewRepository creates a new account repository
func NewRepository(configs ...RepositoryConfiguration) *Repository {
	r := &Repository{}

	for _, config := range configs {
		_ = config(r)
	}

	return r
}

// Register registers a new user.
func (r *Repository) Register(username string, password string) (*account.Account, error) {
	return r.DB.CreateAccount(username, password)
}

// Login logs in a user.
func (r *Repository) Login(username string, password string) (*account.Account, error) {
	return r.DB.GetAccountByUsernameAndPassword(username, password)
}
