package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
)

// PostgresDatabase is the implementation of the AppDatabase interface for PostgreSQL.
type PostgresDatabase struct {
	Conn *pgx.Conn
}

// NewPostgresDatabase creates a new instance of PostgresDatabase.
func NewPostgresDatabase(connectionString string) *PostgresDatabase {
	// set a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect to the database
	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		log.Printf("error connecting to database: %v", err)
		return nil
	}

	// ping the database to ensure that the connection is alive
	if err := conn.Ping(ctx); err != nil {
		log.Printf("error pinging database: %v", err)
		return nil
	}

	return &PostgresDatabase{Conn: conn}
}
