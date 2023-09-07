package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

// RedisDatabase is the implementation of the AppDatabase interface for Redis.
type RedisDatabase struct {
	Client *redis.Client
}

// NewRedisDatabase creates a new instance of RedisDatabase.
func NewRedisDatabase(connectionString string) *RedisDatabase {
	// set a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

	return &RedisDatabase{Client: client}
}
