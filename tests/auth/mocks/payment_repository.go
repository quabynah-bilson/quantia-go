package mocks

import "github.com/quabynah-bilson/quantia/pkg/payment"

type MockPaymentRepository struct {
	PayFn       func(amount float32, url string) (*payment.Transaction, error)
	SubscribeFn func(url string, queue chan *payment.WebhookPayload) error
}

func (m *MockPaymentRepository) Pay(amount float32, url string) (*payment.Transaction, error) {
	return m.PayFn(amount, url)
}

func (m *MockPaymentRepository) Subscribe(url string, queue chan *payment.WebhookPayload) error {
	return m.SubscribeFn(url, queue)
}
