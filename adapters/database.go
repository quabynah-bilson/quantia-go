package adapters

import "log"

// AppDatabaseConfig is the type of the function that configures the database.
type AppDatabaseConfig func(*Database) error

// AppDatabase is the interface that wraps the basic database operations.
type AppDatabase interface {

	// Get retrieves a record from the database.
	Get(key string, data interface{}) error

	// Set saves a record into the database.
	Set(key string, value interface{}) (*string, error)

	// Delete removes a record from the database.
	Delete(key string) error
}

// Database is the struct that holds the database instance (MongoDB, Redis, etc.).
// DB must implement the AppDatabase interface.
type Database struct {
	DB AppDatabase
}

// NewDatabase creates a new database instance based on the provided configs.
func NewDatabase(configs ...AppDatabaseConfig) *Database {
	database := &Database{}
	for _, config := range configs {
		if err := config(database); err != nil {
			log.Printf("error configuring database: %v", err)
			return nil
		}
	}
	log.Printf("database configured successfully -> %T", database.DB)
	return database
}
