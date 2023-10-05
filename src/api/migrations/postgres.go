package migrations

import (
	"context"
	"github.com/jackc/pgx/v5"
	"time"
)

// PerformMigrations performs all the migrations.
func PerformMigrations(conn *pgx.Conn, errChan chan error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	// create the database
	_, _ = conn.Exec(ctx, "CREATE DATABASE IF NOT EXISTS quantia")
	_, _ = conn.Exec(ctx, "USE quantia")

	// create the migrations table
	_, _ = conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS migrations (id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL, created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)")

	// create the accounts table
	_, _ = conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS accounts (id UUID PRIMARY KEY, username VARCHAR(255) NOT NULL, password VARCHAR(255) NOT NULL, created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)")

	// create the sessions table
	_, _ = conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS sessions (id UUID PRIMARY KEY, account_id UUID NOT NULL, token VARCHAR(255) NOT NULL, created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)")

	// alter the accounts table to add a unique constraint on the username column
	_, _ = conn.Exec(ctx, "ALTER TABLE accounts ADD CONSTRAINT unique_username UNIQUE (username)")

	errChan <- nil
}
