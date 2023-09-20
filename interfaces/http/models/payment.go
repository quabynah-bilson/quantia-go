package models

import "github.com/quabynah-bilson/quantia/pkg/payment"

// MakePaymentRequest represents the JSON structure expected for payment requests.
type MakePaymentRequest struct {
	Amount float32 `json:"amount"`
	Url    string  `json:"url"`
}

// MakePaymentResponse represents the JSON structure returned for payment requests.
type MakePaymentResponse struct {
	Transaction *payment.Transaction `json:"transaction"`
}

// SubscribeToWebhookRequest represents the JSON structure expected for webhook subscription requests.
type SubscribeToWebhookRequest struct {
	Url string `json:"url"`
}
