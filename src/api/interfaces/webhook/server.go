package webhook

import (
	"github.com/quabynah-bilson/quantia/adapters/payment/datastore"
	"github.com/quabynah-bilson/quantia/internal/payment"
	paymentPkg "github.com/quabynah-bilson/quantia/pkg/payment"
	"log"
	"os"
)

// StartWebhookWorker starts the webhook worker (to process webhooks)
func StartWebhookWorker() {
	// create a new payment repository (with a database configuration)
	paymentRepo := payment.NewRepository(datastore.WithRedisPaymentDatabase(os.Getenv("REDIS_URI")))

	// queue for webhooks (buffer 100 webhooks (to avoid blocking the main thread))
	webhookQueue := make(chan *paymentPkg.WebhookPayload, 100)

	// process webhooks
	go payment.ProcessWebhooks(paymentRepo, webhookQueue)

	// subscribe to the payment channel
	if err := paymentRepo.Subscribe(os.Getenv("WEBHOOK_ADDRESS"), webhookQueue); err != nil {
		log.Fatalf("failed to subscribe to payment channel: %v", err)
	}

	select {}
}
