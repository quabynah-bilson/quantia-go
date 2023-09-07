package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/quabynah-bilson/quantia/adapters"
	"github.com/quabynah-bilson/quantia/adapters/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
)

type Account struct {
	Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username string             `json:"username,omitempty"`
}

// entry point of the application
func main() {
	// Load .env file
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// @todo start server and initialize all components
	db := adapters.NewDatabase(
		//database.WithMongoDatabase(os.Getenv("MONGO_URI"), "accounts"),
		database.WithPostgresDatabase(os.Getenv("POSTGRES_URI")),
		//database.WithRedisDatabase(os.Getenv("REDIS_URI")),
	)

	if _, err := db.DB.Set("quabynah1809", Account{Username: "Quabynah Junior"}); err != nil {
		log.Fatalf("Error saving data to database: %v", err)
	}

	var data Account
	if err := db.DB.Get("quabynah1809", &data); err != nil {
		log.Fatalf("Error getting data from database: %v", err)
	}

	fmt.Println(data)
	fmt.Println("\nDatabase connected successfully")
}
