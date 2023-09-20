package payment

// Database is the interface that wraps the basic payment database operations.
type Database interface {
	// SendWebhook sends a webhook to a URL
	SendWebhook(amount float32, url string) (*Transaction, error)

	// SubscribeToWebhook subscribes to a webhook
	SubscribeToWebhook(url string, queue chan *WebhookPayload) error
}
