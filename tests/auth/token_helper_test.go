package auth_test

import (
	"errors"
	"github.com/quabynah-bilson/quantia/pkg/token"
	"github.com/quabynah-bilson/quantia/tests/auth/mocks"
	"testing"
)

func TestMockTokenizerHelper_GenerateToken(t *testing.T) {
	type testCase struct {
		name          string
		claim         string
		expectedToken string
		expectedError error
	}

	testCases := []testCase{
		{
			name:          "generate a token for non-empty claim",
			claim:         "claim123",
			expectedToken: mocks.SuggestedToken,
		},
		{
			name:          "try to generate a token for an empty claim",
			claim:         "",
			expectedError: token.ErrInvalidClaim,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			helper := mocks.NewMockTokenizerHelper()
			generatedToken, err := helper.GenerateToken(tc.claim)

			if generatedToken != tc.expectedToken {
				t.Errorf("Expected token: %s, got: %s", tc.expectedToken, generatedToken)
			}

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestMockTokenizerHelper_ValidateToken(t *testing.T) {
	type testCase struct {
		name           string
		suggestedToken string
		expectedError  error
	}

	testCases := []testCase{
		{
			name:           "validate generated token",
			suggestedToken: mocks.SuggestedToken,
		},
		{
			name:           "validate non-generated token",
			suggestedToken: "nonGeneratedToken",
			expectedError:  token.ErrInvalidToken,
		},
		{
			name:           "validate invalidated token",
			suggestedToken: mocks.InvalidatedToken,
			expectedError:  token.ErrInvalidToken,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			helper := mocks.NewMockTokenizerHelper()
			if _, err := helper.GenerateToken("claim123"); err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			validateErr := helper.ValidateToken(tc.suggestedToken)

			if !errors.Is(validateErr, tc.expectedError) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, validateErr)
			}
		})
	}
}
