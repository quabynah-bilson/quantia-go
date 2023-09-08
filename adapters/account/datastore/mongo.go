package datastore

import (
	"context"
	"errors"
	"github.com/quabynah-bilson/quantia/adapters"
	"github.com/quabynah-bilson/quantia/adapters/account"
	pkgAccount "github.com/quabynah-bilson/quantia/pkg/account"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	databaseName   = "quantia"
	collectionName = "accounts"
)

// MongoAccountDatabase is the struct that wraps the basic account database operations for MongoDB.
type MongoAccountDatabase struct {
	collection *mongo.Collection
	pwHelper   pkgAccount.PasswordHelper
	account.Database
}

// WithMongoAccountDatabase is the function that returns a DatabaseConfig function that sets the account database to MongoDB.
func WithMongoAccountDatabase(connectionString string, pwHelper pkgAccount.PasswordHelper) adapters.DatabaseConfig {
	// set a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect to the database
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Printf("error connecting to MongoDB: %v", err)
		return nil
	}

	// ping the database to check if the connection is working
	if err = db.Ping(ctx, nil); err != nil {
		log.Printf("error pinging MongoDB: %v", err)
		return nil
	}

	return func(a *adapters.DatabaseAdapter) error {
		a.AccountDB = &MongoAccountDatabase{
			collection: db.Database(databaseName).Collection(collectionName),
			pwHelper:   pwHelper,
		}
		return nil
	}
}

// GetAccount gets an account by ID.
func (db *MongoAccountDatabase) GetAccount(id string) (*pkgAccount.Account, error) {
	// set a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// find the account
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, account.ErrInvalidID
	}
	var acc pkgAccount.Account
	if err = db.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&acc); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, account.ErrAccountNotFound
		}
		return nil, err
	}

	return &acc, nil
}

// GetAccountByUsernameAndPassword gets an account by username and password
func (db *MongoAccountDatabase) GetAccountByUsernameAndPassword(username, password string) (*pkgAccount.Account, error) {
	// set a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// find the account
	var acc pkgAccount.Account
	if err := db.collection.FindOne(ctx, bson.M{"username": username}).Decode(&acc); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, account.ErrAccountNotFound
		}
		return nil, err
	}

	// compare the passwords
	if err := db.pwHelper.ComparePassword(acc.Password, password); err != nil {
		return nil, account.ErrInvalidCredentials
	}

	return &acc, nil
}

// CreateAccount creates a new account.
func (db *MongoAccountDatabase) CreateAccount(username, password string) (*pkgAccount.Account, error) {
	// set a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// hash the password before saving it to the database
	hashedPassword, err := db.pwHelper.HashPassword(password)
	if err != nil {
		return nil, account.ErrAccountNotCreated
	}

	// check if the account already exists
	if count, err := db.collection.CountDocuments(ctx, bson.M{"username": username}); err != nil {
		return nil, err
	} else if count > 0 {
		return nil, account.ErrAccountAlreadyExists
	}

	// create the account
	userAccount := &pkgAccount.Account{
		ID:       primitive.NewObjectID().Hex(),
		Username: username,
		Password: hashedPassword,
	}
	if _, err = db.collection.InsertOne(ctx, userAccount); err != nil {
		return nil, err
	}

	return userAccount, nil
}

// DeleteAccount deletes an account by ID.
func (db *MongoAccountDatabase) DeleteAccount(id string) error {
	// set a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// delete the account
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return account.ErrInvalidID
	}

	if result, err := db.collection.DeleteOne(ctx, bson.M{"_id": oid}); err != nil {
		return err
	} else if result.DeletedCount == 0 {
		return account.ErrAccountNotDeleted
	}

	return nil
}
