package datastore

import (
	"context"
	"encoding/json"
	"errors"
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

	// check if the username exists
	sessionJson, err := db.client.Get(ctx, id).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Printf("error checking if username exists: %v", err)
		return "", token.ErrTokenNotCreated
	}

	// check if the session exists
	if sessionJson != "" {
		// @todo -> check if the token has expired

		// return the token if the session exists
		var session auth.Session
		if err = toSession(sessionJson, &session); err != nil {
			return "", err
		}

		return session.Token, nil
	}

	// create a new token for the given id
	session := auth.Session{
		ID:        uuid.NewString(),
		AccountID: id,
		Token:     uuid.NewString(), // @todo generate a random token
	}

	// convert the session to a JSON string
	jsonString, err := fromSession(&session)
	if err != nil {
		return "", err
	}

	if err := db.client.Set(ctx, id, jsonString, 0).Err(); err != nil {
		log.Printf("error creating token: %v", err)
		return "", token.ErrTokenNotCreated
	}

	return session.Token, nil
}

// ValidateToken validates the given token.
func (db *RedisTokenDatabase) ValidateToken(authToken string) error {
	// @todo -> check if the token has expired

	// @todo -> implement this method

	return nil
}

// DeleteToken invalidates the given token.
func (db *RedisTokenDatabase) DeleteToken(authToken string) error {
	// @todo -> check if the token is valid

	// @todo -> implement this method

	return nil
}

// toSession converts the given JSON string to a session.
func toSession(jsonString string, session *auth.Session) error {
	if err := json.Unmarshal([]byte(jsonString), &session); err != nil {
		log.Printf("error unmarshalling session: %v", err)
		return token.ErrParsingSession
	}

	return nil
}

// fromSession converts the given session to a JSON string.
func fromSession(session *auth.Session) (string, error) {
	jsonString, err := json.Marshal(session)
	if err != nil {
		log.Printf("error marshalling session: %v", err)
		return "", token.ErrParsingSession
	}

	return string(jsonString), nil
}
