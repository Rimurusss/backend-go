package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Rimurusss/backend-goconfig"
)

var DB *mongo.Database

func ConnectDB() {
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

	fmt.Println("Connected to MongoDB!")

	// Get your database
	dbName := config.Config("MONGO_DB_NAME")
	DB = client.Database(dbName)
}
