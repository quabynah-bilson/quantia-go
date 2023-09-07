package database

import (
	"context"
	"log"
	"time"

	"github.com/quabynah-bilson/quantia/adapters"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const databaseName = "quantia"

// MongoDatabase is the implementation of the AppDatabase interface for MongoDB.
type MongoDatabase struct {
	Collection *mongo.Collection
}

// WithMongoDatabase configures the database to use MongoDB.
func WithMongoDatabase(connectionString, collectionName string) adapters.AppDatabaseConfig {
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

	return func(database *adapters.Database) error {
		database.DB = &MongoDatabase{Collection: db.Database(databaseName).Collection(collectionName)}
		return nil
	}
}

// Get retrieves a record from the database.
func (m *MongoDatabase) Get(key string, data interface{}) error {
	oid, err := primitive.ObjectIDFromHex(key)
	if err != nil {
		return err
	}

	// set a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var result bson.D
	if err = m.Collection.FindOne(ctx, bson.D{{Key: "_id", Value: oid}}).Decode(&result); err != nil {
		return err
	}

	// convert the result to the provided data type
	if marshal, err := bson.Marshal(result); err != nil {
		return err
	} else if err = bson.Unmarshal(marshal, data); err != nil {
		return err
	}

	return nil
}

// Set saves a record into the database.
func (m *MongoDatabase) Set(_ string, value interface{}) (*string, error) {
	// set a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := m.Collection.InsertOne(ctx, &value)
	if err != nil {
		return nil, err
	}

	oid := result.InsertedID.(primitive.ObjectID).Hex()
	return &oid, nil
}

// Delete removes a record from the database.
func (m *MongoDatabase) Delete(key string) error {
	// set a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(key)
	if err != nil {
		return err
	}

	if _, err = m.Collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: oid}}); err != nil {
		return err
	}

	log.Printf("Deleted a single document: %s", key)
	return nil
}
