package unit

import (
	"errors"
	"github.com/quabynah-bilson/quantia/pkg"
	"github.com/quabynah-bilson/quantia/pkg/payment"
	"github.com/quabynah-bilson/quantia/tests/auth/mocks"
	"log"
	"testing"
)

// testCase is a struct that represents a test case.
type testCase struct {
	name                  string
	amount                float32
	url                   string
	expectedTransactionID string
	expectedErr           error
}

// TestPaymentUseCase_Pay tests the pay method of the payment use case.
func TestPaymentUseCase_Pay(t *testing.T) {
	testCases := []testCase{
		{
			name:                  "invalid amount",
			amount:                -1,
			url:                   "https://quantia.com",
			expectedTransactionID: "",
			expectedErr:           pkg.ErrInvalidAmount,
		},
		{
			name:                  "invalid URL",
			amount:                100,
			url:                   "quantia.com",
			expectedTransactionID: "",
			expectedErr:           pkg.ErrInvalidURL,
		},
		{
			name:                  "valid payment",
			amount:                100,
			url:                   "https://quantia.com",
			expectedTransactionID: "123e4567-e89b-12d3-a456-426614174000",
			expectedErr:           nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			paymentRepo := &mocks.MockPaymentRepository{
				PayFn: func(amount float32, url string) (*payment.Transaction, error) {
					return &payment.Transaction{
						ID:     "123e4567-e89b-12d3-a456-426614174000",
						Amount: amount,
						Status: payment.TransactionStatusSuccess,
					}, nil
				},
			}

			paymentUseCase := pkg.NewPaymentUseCase(paymentRepo)

			// Act
			transaction, err := paymentUseCase.MakePayment(tc.amount, tc.url)

			// Assert
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error: %v, got: %v", tc.expectedErr, err)
			}

			if transaction != nil && transaction.ID != tc.expectedTransactionID {
				t.Errorf("expected transaction ID: %s, got: %s", tc.expectedTransactionID, transaction.ID)
			}
		})
	}
}

// TestPaymentUseCase_SubscribeToWebhook tests the subscribe to webhook method of the payment use case.
func TestPaymentUseCase_SubscribeToWebhook(t *testing.T) {
	testCases := []testCase{
		{
			name:        "invalid URL",
			url:         "quantia.com",
			expectedErr: pkg.ErrInvalidURL,
		},
		{
			name:        "valid URL",
			url:         "https://quantia.com",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			paymentRepo := &mocks.MockPaymentRepository{
				SubscribeFn: func(url string, queue chan *payment.WebhookPayload) error {
					return nil
				},
			}

			paymentUseCase := pkg.NewPaymentUseCase(paymentRepo)

			// Act
			queueChan := make(chan *payment.WebhookPayload)
			err := paymentUseCase.SubscribeToWebhook(tc.url, queueChan)

			// Listen to the queue channel
			go func() {
				for {
					log.Printf("webhook channel received: %v", <-queueChan)
				}
			}()

			// Assert
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}
