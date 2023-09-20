package payment

// TransactionStatus is the type that represents a transaction status
type TransactionStatus string

const (
	// TransactionStatusPending is the status of a pending transaction
	TransactionStatusPending TransactionStatus = "pending"

	// TransactionStatusFailed is the status of a failed transaction
	TransactionStatusFailed TransactionStatus = "failed"

	// TransactionStatusSuccess is the status of a successful transaction
	TransactionStatusSuccess TransactionStatus = "success"
)

// Transaction is the entity that represents a payment transaction
type Transaction struct {
	ID     string            `json:"id"`
	Amount float32           `json:"amount"`
	Status TransactionStatus `json:"status"`
}

// WebhookPayload is the entity that represents a webhook payload
type WebhookPayload struct {
	ID     string            `json:"id"`
	Status TransactionStatus `json:"status"`
	Url    string            `json:"url"`
	Data   struct {
		TransactionID string `json:"transaction_id"`
		Date          string `json:"created_at"`
	} `json:"data"`
}
