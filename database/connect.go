package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"backend-go/config"
)

var DB *mongo.Database

func ConnectDB() (client *mongo.Client, db *mongo.Database) {
	// Get MongoDB URI from .env using the Config() function
	mongoURI := config.Config("MONGO_URI")

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get your database
	dbName := config.Config("MONGO_DB_NAME")
	DB = client.Database(dbName)

	return client, DB
}
