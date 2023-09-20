package payment

type Repository interface {
	// Pay pays an amount to a given URL.
	Pay(amount float32, url string) (*Transaction, error)

	// Subscribe subscribes to a given webhook URL.
	Subscribe(url string, queue chan *WebhookPayload) error
}
