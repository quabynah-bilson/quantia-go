package adapters

import (
	"github.com/quabynah-bilson/quantia/adapters/account"
	"github.com/quabynah-bilson/quantia/adapters/token"
	"log"
)

// DatabaseConfig is the type of the function that configures the account database.
type DatabaseConfig func(*DatabaseAdapter) error

// DatabaseAdapter is the type that provides access to the various database operations.
type DatabaseAdapter struct {
	AccountDB account.Database
	TokenDB   token.Database
}

// NewAdapter creates a new database adapter instance based on the provided configs.
func NewAdapter(configs ...DatabaseConfig) *DatabaseAdapter {
	database := &DatabaseAdapter{}
	for _, config := range configs {
		if err := config(database); err != nil {
			log.Printf("error configuring database: %v", err)
			return nil
		}
	}
	return database
}
