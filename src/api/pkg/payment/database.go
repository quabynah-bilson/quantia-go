package payment

import "errors"

var (
	// ErrFailedToMarshalTransaction is the error returned when a transaction fails to marshal
	ErrFailedToMarshalTransaction = errors.New("invalid transaction details. Please check and try again")

	// ErrFailedToSubscribeToWebhook is the error returned when a webhook subscription fails
	ErrFailedToSubscribeToWebhook = errors.New("failed to subscribe to webhook. Please check and try again")
)

// Database is the interface that wraps the basic payment database operations.
type Database interface {
	// SendWebhook sends a webhook to a URL
	SendWebhook(amount float32, url string) (*Transaction, error)

	// SubscribeToWebhook subscribes to a webhook
	SubscribeToWebhook(url string, queue chan *WebhookPayload) error
}
