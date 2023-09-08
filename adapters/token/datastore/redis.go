package datastore

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/quabynah-bilson/quantia/adapters"
	"github.com/quabynah-bilson/quantia/adapters/token"
	"github.com/quabynah-bilson/quantia/pkg/auth"
	"log"
	"time"
)

// RedisTokenDatabase is the implementation of the TokenDatabase interface for Redis.
type RedisTokenDatabase struct {
	token.Database
	client *redis.Client
}

// WithRedisTokenDatabase configures the database to use Redis.
func WithRedisTokenDatabase(connectionString string) adapters.DatabaseConfig {
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

	return func(database *adapters.DatabaseAdapter) error {
		database.TokenDB = &RedisTokenDatabase{client: client}
		return nil
	}
}

// CreateToken generates a token for the given username.
func (db *RedisTokenDatabase) CreateToken(id string) (string, error) {
	// set context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// create a new token for the given id
	session := auth.Session{
		ID:        uuid.NewString(),
		AccountID: id,
		Token:     uuid.NewString(), // @todo generate a random token
	}

	if err := db.client.HSet(ctx, id, fromSession(&session)).Err(); err != nil {
		log.Printf("error creating token: %v", err)
		return "", token.ErrTokenNotCreated
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
		return token.ErrInvalidToken
	}

	// check if the token is valid
	if storedToken != authToken {
		return token.ErrInvalidToken
	}

	// @todo -> check if the token has expired

	return nil
}

// DeleteToken invalidates the given token.
func (db *RedisTokenDatabase) DeleteToken(accountID string) error {
	// set context with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// delete the token
	if err := db.client.HDel(ctx, accountID, "token", "account_id", "id").Err(); err != nil {
		log.Printf("error deleting token: %v", err)
		return token.ErrCannotDeleteToken
	}

	return nil
}

// fromSession converts the given session to a JSON string.
func fromSession(session *auth.Session) map[string]interface{} {
	return map[string]interface{}{
		"id":         session.ID,
		"account_id": session.AccountID,
		"token":      session.Token,
	}
}
