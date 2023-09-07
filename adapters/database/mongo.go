package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// MongoDatabase is the implementation of the AppDatabase interface for MongoDB.
type MongoDatabase struct {
	DB *mongo.Database
}

// NewMongoDatabase creates a new instance of MongoDatabase.
func NewMongoDatabase(connectionString, dbName string) *MongoDatabase {
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

	return &MongoDatabase{DB: db.Database(dbName)}
}
