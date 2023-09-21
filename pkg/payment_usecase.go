package pkg

import (
	"errors"
	"github.com/quabynah-bilson/quantia/pkg/payment"
	"log"
	"regexp"
)

var (
	// ErrInvalidAmount is the error returned when an amount is invalid.
	ErrInvalidAmount = errors.New("invalid amount. Please check and try again")

	// ErrInvalidURL is the error returned when a URL is invalid.
	ErrInvalidURL = errors.New("invalid URL. Please check and try again")
)

// PaymentUseCase is the payment use case. It contains the necessary repositories to perform payment operations.
type PaymentUseCase struct {
	paymentRepo payment.Repository
}

// NewPaymentUseCase creates a new payment use case.
func NewPaymentUseCase(paymentRepo payment.Repository) *PaymentUseCase {
	return &PaymentUseCase{
		paymentRepo: paymentRepo,
	}
}

// MakePayment makes a payment.
func (uc *PaymentUseCase) MakePayment(amount float32, url string) (*payment.Transaction, error) {
	if err := validateAmount(amount); err != nil {
		log.Printf("error validating amount: %v", err)
		return nil, err
	}

	return uc.paymentRepo.Pay(amount, url)
}

// Subscribe subscribes to a webhook.
func (uc *PaymentUseCase) Subscribe(url string, queue chan *payment.WebhookPayload) error {
	if err := validateURL(url); err != nil {
		log.Printf("error validating URL: %v", err)
		return err
	}

	return uc.paymentRepo.Subscribe(url, queue)
}

// validateAmount validates an amount.
func validateAmount(amount float32) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	return nil
}

// validateURL validates a URL.
func validateURL(url string) error {
	urlPattern := `^(http|https)://[^\s/$.?#].[^\s]*$`
	urlMatcher := regexp.MustCompile(urlPattern)
	if matched := urlMatcher.MatchString(url); !matched {
		return ErrInvalidURL
	}

	return nil
}
