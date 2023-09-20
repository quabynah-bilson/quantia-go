package payment

import (
	pkg "github.com/quabynah-bilson/quantia/pkg/payment"
	"log"
	"time"
)

// ProcessWebhooks is a worker that processes webhooks
func ProcessWebhooks(repo *Repository, webhookQueue chan *pkg.WebhookPayload) {
	for payload := range webhookQueue {
		go func(p *pkg.WebhookPayload) {
			backoffTime, maxBackoffTime := time.Second, time.Minute
			retries, maxRetries := 0, 5
			for {
				// process the webhook payload
				if _, err := repo.Pay(p.Amount, p.Url); err == nil {
					break // success
				}

				// compare the retries to the max retries
				retries++
				if retries >= maxRetries {
					log.Printf("max retries reached for transaction %s", p.ID)
					break // max retries reached
				}

				// retry after backoff time
				log.Printf("retrying transaction %s in %s", p.ID, backoffTime)
				time.Sleep(backoffTime)

				// double the backoff time for the next iteration, capped at the max backoff time
				backoffTime *= 2
				if backoffTime > maxBackoffTime {
					backoffTime = maxBackoffTime
				}
			}
		}(payload)
	}
}
