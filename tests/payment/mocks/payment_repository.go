package mocks

import "github.com/quabynah-bilson/quantia/pkg/payment"

// MockPaymentRepository is a mock of the payment repository
type MockPaymentRepository struct {
	PayFn       func(amount float32, url string) (*payment.Transaction, error)
	SubscribeFn func(url string, queue chan *payment.WebhookPayload) error
}

// Pay calls the PayFn
func (m *MockPaymentRepository) Pay(amount float32, url string) (*payment.Transaction, error) {
	return m.PayFn(amount, url)
}

// Subscribe calls the SubscribeFn
func (m *MockPaymentRepository) Subscribe(url string, queue chan *payment.WebhookPayload) error {
	return m.SubscribeFn(url, queue)
}
