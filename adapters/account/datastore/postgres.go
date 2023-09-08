package datastore

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/quabynah-bilson/quantia/adapters"
	"github.com/quabynah-bilson/quantia/adapters/account"
	"github.com/quabynah-bilson/quantia/migrations"
	"github.com/quabynah-bilson/quantia/pkg/auth"
	"log"
	"time"
)

// AccountPostgresDatabase is the struct that wraps the basic account database operations for PostgreSQL.
type AccountPostgresDatabase struct {
	account.Database
	conn *pgx.Conn
}

// WithPostgresAccountDatabase is the function that returns a DatabaseConfig function that sets the account database to PostgreSQL.
func WithPostgresAccountDatabase(connectionString string) adapters.DatabaseConfig {
	// set a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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

	// perform migrations
	errChan := make(chan error)
	go migrations.PerformMigrations(conn, errChan)
	if err = <-errChan; err != nil {
		log.Printf("error performing migrations: %v", err)
		return nil
	}

	return func(a *adapters.DatabaseAdapter) error {
		a.AccountDB = &AccountPostgresDatabase{conn: conn}
		return nil
	}
}

// CreateAccount creates a new account.
func (d *AccountPostgresDatabase) CreateAccount(username, password string) (*auth.Account, error) {
	// set a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// check if the account already exists by username
	var userAccount auth.Account
	if err := d.conn.QueryRow(ctx, "SELECT id, username, password FROM accounts WHERE username = $1", username).Scan(&userAccount.ID, &userAccount.Username, &userAccount.Password); err == nil {
		return nil, account.ErrAccountAlreadyExists
	}

	// @todo -> hash the password

	// create a new account
	id := uuid.New()
	tag, err := d.conn.Exec(ctx, "INSERT INTO accounts (id, username, password) VALUES ($1, $2, $3)", id, username, password)
	if err != nil {
		log.Printf("error creating account: %v", err)
		return nil, account.ErrAccountNotCreated
	}

	// check if the account was created
	if tag.RowsAffected() == 0 {
		return nil, account.ErrAccountNotCreated
	}

	// get the account
	getAccountResult, err := d.GetAccount(id.String())
	if err != nil {
		log.Printf("error getting account: %v", err)
		return nil, account.ErrAccountNotFound
	}

	return getAccountResult, nil
}

// GetAccount gets an account by ID.
func (d *AccountPostgresDatabase) GetAccount(id string) (*auth.Account, error) {
	// set a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// parse the ID
	parsedID, err := parseID(id)
	if err != nil {
		return nil, err
	}

	// get the account
	var userAccount auth.Account
	if err := d.conn.QueryRow(ctx, "SELECT id, username, password FROM accounts WHERE id = $1", parsedID).Scan(&userAccount.ID, &userAccount.Username, &userAccount.Password); err != nil {
		log.Printf("error getting account: %v", err)
		return nil, account.ErrAccountNotFound
	}

	return &userAccount, nil
}

// GetAccountByUsernameAndPassword gets an account by username and password.
func (d *AccountPostgresDatabase) GetAccountByUsernameAndPassword(username, password string) (*auth.Account, error) {
	// set a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// get the account
	var userAccount auth.Account
	// @todo -> 1. remove password check from query
	if err := d.conn.QueryRow(ctx, "SELECT id, username, password FROM accounts WHERE username = $1 AND password = $2", username, password).Scan(&userAccount.ID, &userAccount.Username, &userAccount.Password); err != nil {
		log.Printf("error getting account: %v", err)
		return nil, account.ErrAccountNotFound
	}

	// @todo -> 2. compare the password with the hashed password

	return &userAccount, nil
}

// DeleteAccount deletes an account by ID.
func (d *AccountPostgresDatabase) DeleteAccount(id string) error {
	// set a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// parse the ID
	parsedID, err := parseID(id)
	if err != nil {
		return err
	}

	// delete the account
	tag, err := d.conn.Exec(ctx, "DELETE FROM accounts WHERE id = $1", parsedID)
	if err != nil {
		log.Printf("error deleting account: %v", err)
		return account.ErrAccountNotDeleted
	}

	// check if the account was deleted
	if tag.RowsAffected() == 0 {
		return account.ErrAccountNotDeleted
	}

	return nil
}

func parseID(id string) (uuid.UUID, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return [16]byte{}, account.ErrInvalidID
	}

	return parsed, nil
}
