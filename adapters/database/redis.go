package database

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/quabynah-bilson/quantia/adapters"
	"log"
	"time"
)

// RedisDatabase is the implementation of the AppDatabase interface for Redis.
type RedisDatabase struct {
	Client *redis.Client
}

// WithRedisDatabase configures the database to use Redis.
func WithRedisDatabase(connectionString string) adapters.AppDatabaseConfig {
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

	return func(database *adapters.Database) error {
		database.DB = &RedisDatabase{Client: client}
		return nil
	}
}

// Get retrieves a record from the database.
func (r *RedisDatabase) Get(key string, data interface{}) error {
	// get data from redis
	dataString, err := r.Client.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}

	// convert data to json
	if err = json.Unmarshal([]byte(dataString), &data); err != nil {
		return err
	}
	return nil
}

func (r *RedisDatabase) Set(key string, value interface{}) (*string, error) {
	// convert data to json
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	// save data to redis
	data := string(jsonBytes)
	if err := r.Client.Set(context.Background(), key, data, 0).Err(); err != nil {
		return nil, err
	}

	// return the data
	return &data, nil
}

func (r *RedisDatabase) Delete(key string) error {
	if err := r.Client.Del(context.Background(), key).Err(); err != nil {
		return err
	}
	return nil
}
