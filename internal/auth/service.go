package auth

import "github.com/quabynah-bilson/quantia/pkg/auth"

// ServiceConfiguration is a function that configures a service
type ServiceConfiguration func(*Service) error

// Service is the account service
type Service struct {
	AuthRepo  auth.Repository
	TokenRepo auth.TokenRepository
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
func WithAuthRepo(repo auth.Repository) ServiceConfiguration {
	return func(s *Service) error {
		s.AuthRepo = repo
		return nil
	}
}

// WithTokenRepo sets the token repository
func WithTokenRepo(repo auth.TokenRepository) ServiceConfiguration {
	return func(s *Service) error {
		s.TokenRepo = repo
		return nil
	}
}
