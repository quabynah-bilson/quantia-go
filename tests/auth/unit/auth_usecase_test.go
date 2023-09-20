package unit_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/quabynah-bilson/quantia/pkg"
	"github.com/quabynah-bilson/quantia/pkg/account"
	"github.com/quabynah-bilson/quantia/tests/auth/mocks"
	"testing"
)

// testCase is a struct that represents a test case.
type testCase struct {
	name              string
	username          string
	password          string
	expectedToken     string
	expectedAccountID string
	expectedErr       error
}

// getTestToken returns a test token.
func getTestToken() string {
	uuidToken, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	token := uuidToken.String()
	return token
}

// TestAuthUseCase_RegisterUser tests the register user method of the auth use case.
func TestAuthUseCase_RegisterUser(t *testing.T) {
	testCases := []testCase{
		{
			name:        "invalid username",
			username:    "user@quantia",
			password:    mocks.ValidPassword,
			expectedErr: pkg.ErrInvalidUsername,
		},
		{
			name:        "invalid password",
			username:    mocks.NewCustomerUsername,
			password:    "pass",
			expectedErr: pkg.ErrInvalidPassword,
		},
		{
			name:        "empty username",
			username:    "",
			password:    mocks.ValidPassword,
			expectedErr: pkg.ErrInvalidUsername,
		},
		{
			name:        "empty password",
			username:    mocks.NewCustomerUsername,
			password:    "",
			expectedErr: pkg.ErrInvalidPassword,
		},
		{
			name:        "user already exists",
			username:    mocks.ExistingCustomerUsername,
			password:    mocks.ValidPassword,
			expectedErr: mocks.ErrAlreadyExists,
		},
		{
			name:              "valid registration",
			username:          mocks.NewCustomerUsername,
			password:          mocks.ValidPassword,
			expectedToken:     getTestToken(),
			expectedAccountID: uuid.NewString(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			authRepo := &mocks.MockAccountRepository{
				RegisterFn: func(username, password string) (*account.Account, error) {
					var err error

					// simulate a user already exists
					if username == mocks.ExistingCustomerUsername {
						err = mocks.ErrAlreadyExists
					}

					// simulate a successful registration
					return &account.Account{
						ID: tc.expectedAccountID,
					}, err
				},
			}

			tokenRepo := &mocks.MockTokenRepository{
				GenerateTokenFn: func(claim string) (string, error) {
					return getTestToken(), nil
				},
			}

			uc := pkg.NewAuthUseCase(authRepo, tokenRepo)
			token, err := uc.Register(tc.username, tc.password)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}

			if err == nil && *token != tc.expectedToken {
				t.Errorf("expected token %v, got %v", tc.expectedToken, token)
			}
		})
	}
}

// TestAuthUseCase_LoginUser tests the login user method of the auth use case.
func TestAuthUseCase_LoginUser(t *testing.T) {
	testCases := []testCase{
		{
			name:        "invalid username",
			username:    "user@quantia",
			password:    mocks.ValidPassword,
			expectedErr: pkg.ErrInvalidUsername,
		},
		{
			name:        "invalid password",
			username:    mocks.NewCustomerUsername,
			password:    "pass",
			expectedErr: pkg.ErrInvalidPassword,
		},
		{
			name:        "empty username",
			username:    "",
			password:    mocks.ValidPassword,
			expectedErr: pkg.ErrInvalidUsername,
		},
		{
			name:        "empty password",
			username:    mocks.NewCustomerUsername,
			password:    "",
			expectedErr: pkg.ErrInvalidPassword,
		},
		{
			name:        "user not found",
			username:    mocks.NewCustomerUsername,
			password:    mocks.ValidPassword,
			expectedErr: mocks.ErrUserNotFound,
		},
		{
			name:        "wrong password",
			username:    mocks.ExistingCustomerUsername,
			password:    "password",
			expectedErr: mocks.ErrAuthenticationFailed,
		},
		{
			name:          "valid login",
			username:      mocks.ExistingCustomerUsername,
			password:      "password@1234",
			expectedToken: getTestToken(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			authRepo := &mocks.MockAccountRepository{
				LoginFn: func(username, password string) (*account.Account, error) {

					// simulate a successful login
					return &account.Account{
						ID: tc.expectedAccountID,
					}, nil
				},
			}

			tokenRepo := &mocks.MockTokenRepository{
				GenerateTokenFn: func(claim string) (string, error) {
					return getTestToken(), nil
				},
			}

			uc := pkg.NewAuthUseCase(authRepo, tokenRepo)
			token, err := uc.Login(tc.username, tc.password)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}

			if err == nil && *token != tc.expectedToken {
				t.Errorf("expected token %v, got %v", tc.expectedToken, token)
			}
		})
	}
}

// TestAuthUseCase_LogoutUser tests the logout user method of the auth use case.
func TestAuthUseCase_LogoutUser(t *testing.T) {
	testCases := []testCase{
		{
			name:              "valid token",
			expectedToken:     getTestToken(),
			expectedAccountID: uuid.NewString(),
			expectedErr:       nil,
		},
		{
			name:              "invalid token",
			expectedToken:     "",
			expectedAccountID: uuid.NewString(),
			expectedErr:       pkg.ErrInvalidToken,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tokenRepo := &mocks.MockTokenRepository{
				InvalidateTokenFn: func(rawToken, accountID string) error {
					var err error
					if rawToken == "" {
						err = pkg.ErrInvalidToken
					}
					return err
				},
			}

			uc := pkg.NewAuthUseCase(nil, tokenRepo)
			err := uc.Logout(tc.expectedToken, tc.expectedAccountID)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

// TestAuthUseCase_ValidateToken tests the validate token method of the auth use case.
func TestAuthUseCase_ValidateToken(t *testing.T) {
	testCases := []testCase{
		{
			name:              "valid token",
			expectedToken:     getTestToken(),
			expectedAccountID: uuid.NewString(),
			expectedErr:       nil,
		},
		{
			name:        "invalid token",
			expectedErr: pkg.ErrInvalidToken,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tokenRepo := &mocks.MockTokenRepository{
				ValidateTokenFn: func(rawToken, accountID string) error {
					var err error
					if rawToken == "" {
						err = pkg.ErrInvalidToken
					}
					return err
				},
			}

			uc := pkg.NewAuthUseCase(nil, tokenRepo)
			err := uc.ValidateToken(tc.expectedToken, tc.expectedAccountID)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
