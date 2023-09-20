package payment

import (
	"github.com/quabynah-bilson/quantia/pkg/payment"
)

// RepositoryConfiguration is a function that configures a repository
type RepositoryConfiguration func(*Repository) error

// Repository is the payment repository implementation
type Repository struct {
	DB payment.Database
	payment.Repository
}

// NewRepository creates a new payment repository
func NewRepository(configs ...RepositoryConfiguration) *Repository {
	r := &Repository{}

	for _, config := range configs {
		_ = config(r)
	}

	return r
}

// Pay pays an amount to a given URL.
func (r *Repository) Pay(amount float32, url string) (*payment.Transaction, error) {
	return r.DB.SendWebhook(amount, url)
}

// Subscribe subscribes to a given webhook URL.
func (r *Repository) Subscribe(url string, queue chan *payment.WebhookPayload) error {
	return r.DB.SubscribeToWebhook(url, queue)
}
