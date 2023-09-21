package datastore

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	internal "github.com/quabynah-bilson/quantia/internal/payment"
	pkg "github.com/quabynah-bilson/quantia/pkg/payment"
	"log"
	"time"
)

// RedisPaymentDatabase is the implementation of the PaymentDatabase interface for Redis.
type RedisPaymentDatabase struct {
	client *redis.Client
	pkg.Database
}

// WithRedisPaymentDatabase creates a new RedisPaymentDatabase.
func WithRedisPaymentDatabase(connectionString string) internal.RepositoryConfiguration {
	// set a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// connect to the database
	client := redis.NewClient(&redis.Options{
		Addr: connectionString,
		DB:   0,
	})

	// ping the database to check if the connection is working
	if err := client.Ping(ctx).Err(); err != nil {
		log.Printf("error pinging Redis: %v", err)
		return nil
	}

	return func(r *internal.Repository) error {
		r.DB = &RedisPaymentDatabase{
			client: client,
		}

		return nil
	}
}

// SendWebhook sends a webhook to the given URL.
func (db *RedisPaymentDatabase) SendWebhook(amount float32, url string) (*pkg.Transaction, error) {
	// set context in background
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create a new transaction
	transaction := &pkg.Transaction{
		ID:     uuid.New().String(),
		Amount: amount,
		Status: pkg.TransactionStatusPending,
	}

	// marshal the transaction
	transactionJSON, err := marshalToJson(transaction)
	if err != nil {
		return nil, err
	}

	// send the transaction to the given URL
	if err = db.client.Publish(ctx, url, transactionJSON).Err(); err != nil {
		return nil, err
	}

	return transaction, nil
}

// SubscribeToWebhook subscribes to the given webhook URL.
func (db *RedisPaymentDatabase) SubscribeToWebhook(url string, queue chan *pkg.WebhookPayload) error {
	// set context in background
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// subscribe to the webhook
	pubSub := db.client.Subscribe(ctx, url)

	defer func(pubSub *redis.PubSub) {
		// close the channel
		_ = pubSub.Close()
	}(pubSub)

	// listen for messages
	for {
		msg, err := pubSub.ReceiveMessage(ctx)
		if err != nil {
			log.Printf("error receiving message from webhook: %v", err)
			return pkg.ErrFailedToSubscribeToWebhook
		}
		// unmarshal the message
		var payload *pkg.WebhookPayload
		if err := json.Unmarshal([]byte(msg.Payload), &payload); err != nil {
			log.Printf("error unmarshalling webhook payload: %v", err)
			continue
		}

		// send the payload to the queue
		queue <- payload
	}
}

// marshalToJson converts the given transaction to JSON.
func marshalToJson(transaction *pkg.Transaction) (interface{}, error) {
	// convert the transaction to JSON
	transactionJSON, err := json.Marshal(transaction)
	if err != nil {
		return nil, pkg.ErrFailedToMarshalTransaction
	}

	return transactionJSON, nil
}
