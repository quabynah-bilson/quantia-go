package unit_test

import (
	"errors"
	"github.com/quabynah-bilson/quantia/pkg/account"
	"github.com/quabynah-bilson/quantia/tests/auth/mocks"
	"testing"
)

// TestMockPasswordHelper_HashPassword tests the hash password method of the mock password helper.
func TestMockPasswordHelper_HashPassword(t *testing.T) {
	type testCase struct {
		name          string
		password      string
		expectedHash  string
		expectedError error
	}

	testCases := []testCase{
		{
			name:         "hash a non-empty password",
			password:     "password123",
			expectedHash: "hashedpassword123",
		},
		{
			name:          "try to hash an empty password",
			password:      "",
			expectedHash:  "",
			expectedError: account.ErrInvalidPassword,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			helper := mocks.MockPasswordHelper{}
			hash, err := helper.HashPassword(tc.password)
			if hash != tc.expectedHash {
				t.Errorf("Expected hash: %v, got: %v", tc.expectedHash, hash)
			}
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

// TestMockPasswordHelper_ComparePassword tests the compare password method of the mock password helper.
func TestMockPasswordHelper_ComparePassword(t *testing.T) {
	type testCase struct {
		name           string
		hashedPassword string
		password       string
		expectedError  error
	}

	testCases := []testCase{
		{
			name:           "compare correct password",
			hashedPassword: "hashedpassword123",
			password:       "password123",
		},
		{
			name:           "compare incorrect password",
			hashedPassword: "wrongpassword",
			password:       "password123",
			expectedError:  account.ErrPasswordMismatch,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			helper := mocks.MockPasswordHelper{}
			err := helper.ComparePassword(tc.hashedPassword, tc.password)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}
