package auth

import (
	"errors"
	"github.com/google/uuid"
	pkg "github.com/quabynah-bilson/quantia/pkg/auth"
	"github.com/quabynah-bilson/quantia/tests/auth/mocks"
	"testing"
)

type testCase struct {
	name          string
	username      string
	password      string
	expectedToken string
	expectedErr   error
}

func getTestToken() string {
	uuidToken, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	token := uuidToken.String()
	return token
}

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
			name:          "valid registration",
			username:      mocks.NewCustomerUsername,
			password:      mocks.ValidPassword,
			expectedToken: getTestToken(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			authRepo := &mocks.MockAccountRepository{
				RegisterFn: func(username, password string) error {
					var err error

					// simulate a user already exists
					if username == "bilson@quantia.com" {
						err = mocks.ErrAlreadyExists
					}

					// simulate a successful registration
					return err
				},
			}

			tokenRepo := &mocks.MockTokenRepository{
				GenerateTokenFn: func(username string) (string, error) {
					return getTestToken(), nil
				},
			}

			uc := pkg.NewUseCase(authRepo, tokenRepo)
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
			name:        "invalid password",
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
				LoginFn: func(username, password string) error {
					// simulate a successful login
					return nil
				},
			}

			tokenRepo := &mocks.MockTokenRepository{
				GenerateTokenFn: func(username string) (string, error) {
					return getTestToken(), nil
				},
			}

			uc := pkg.NewUseCase(authRepo, tokenRepo)
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

func TestAuthUseCase_LogoutUser(t *testing.T) {
	testCases := []testCase{
		{
			name:          "valid token",
			expectedToken: getTestToken(),
			expectedErr:   nil,
		},
		{
			name:          "invalid token",
			expectedToken: "",
			expectedErr:   pkg.ErrInvalidToken,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tokenRepo := &mocks.MockTokenRepository{
				InvalidateTokenFn: func(token string) error {
					var err error
					if token == "" {
						err = pkg.ErrInvalidToken
					}
					return err
				},
			}

			uc := pkg.NewUseCase(nil, tokenRepo)
			err := uc.Logout(tc.expectedToken)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
