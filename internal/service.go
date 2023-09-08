package internal

import (
	"github.com/quabynah-bilson/quantia/pkg/account"
	"github.com/quabynah-bilson/quantia/pkg/token"
)

// ServiceConfiguration is a function that configures a service
type ServiceConfiguration func(*Service) error

// Service is the account service
type Service struct {
	AccountRepo account.Repository
	TokenRepo   token.Repository
}

// NewService creates a new account service
func NewService(configs ...ServiceConfiguration) (*Service, error) {
	s := &Service{}

	for _, config := range configs {
		if err := config(s); err != nil {
			return nil, err
		}
	}

	return s, nil
}

// WithAuthRepo sets the account repository
func WithAuthRepo(repo account.Repository) ServiceConfiguration {
	return func(s *Service) error {
		s.AccountRepo = repo
		return nil
	}
}

// WithTokenRepo sets the token repository
func WithTokenRepo(repo token.Repository) ServiceConfiguration {
	return func(s *Service) error {
		s.TokenRepo = repo
		return nil
	}
}
