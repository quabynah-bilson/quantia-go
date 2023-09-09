package datastore

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	internal "github.com/quabynah-bilson/quantia/internal/token"
	pkg "github.com/quabynah-bilson/quantia/pkg/token"
	"log"
	"time"
)

// RedisTokenDatabase is the implementation of the TokenDatabase interface for Redis.
type RedisTokenDatabase struct {
	client    *redis.Client
	generator pkg.TokenizerHelper
	pkg.Database
}

// WithRedisTokenDatabase creates a new RedisTokenDatabase.
func WithRedisTokenDatabase(connectionString string) internal.RepositoryConfiguration {
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
		r.DB = &RedisTokenDatabase{
			client:    client,
			generator: internal.NewPasetoTokenizerHelper(),
		}

		return nil
	}
}

// CreateToken generates a token for the given username.
func (db *RedisTokenDatabase) CreateToken(id string) (string, error) {
	// set context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// generate a new token
	generatedToken, err := db.generator.GenerateToken(id)
	if err != nil {
		return "", err
	}

	// create a new token for the given id
	session := pkg.Session{
		ID:        uuid.NewString(),
		AccountID: id,
		Token:     generatedToken,
	}

	if err := db.client.HSet(ctx, id, fromSession(&session)).Err(); err != nil {
		log.Printf("error creating token: %v", err)
		return "", pkg.ErrTokenNotCreated
	}

	return session.Token, nil
}

// ValidateToken validates the given token.
func (db *RedisTokenDatabase) ValidateToken(authToken, accountID string) error {
	// set context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	storedToken, err := db.client.HGet(ctx, accountID, "token").Result()
	if err != nil {
		log.Printf("error validating token: %v", err)
		return pkg.ErrInvalidToken
	}

	// check if the token is valid
	if storedToken != authToken {
		return pkg.ErrInvalidToken
	}

	return db.generator.ValidateToken(storedToken)
}

// DeleteToken invalidates the given token.
func (db *RedisTokenDatabase) DeleteToken(accountID string) error {
	// set context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// check if the token exists
	if exists, err := db.client.HExists(ctx, accountID, "token").Result(); !exists || err != nil {
		log.Printf("error checking if token exists: %v", err)
		return pkg.ErrInvalidToken
	}

	// delete the token
	if deleteCount, err := db.client.HDel(ctx, accountID, "token", "account_id", "id").Result(); deleteCount == 0 || err != nil {
		log.Printf("error deleting token: %v", err)
		return pkg.ErrCannotDeleteToken
	}

	return nil
}

// fromSession converts the given session to a JSON string.
func fromSession(session *pkg.Session) map[string]interface{} {
	return map[string]interface{}{
		"id":         session.ID,
		"account_id": session.AccountID,
		"token":      session.Token,
	}
}
