package database

import (
	"context"
	"key-haven-back/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoDBClient(cfg *config.Config) MongoDBClient {
	// Use the value directly from the config
	mongodbURL := cfg.MongodbURL

	log.Printf("Connecting to MongoDB with URL: %s", mongodbURL)

	// Create a context with timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Configure the client options
	clientOptions := options.Client().
		ApplyURI(mongodbURL).
		SetDirect(true)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Panicf("Error connecting to MongoDB: %v", err)
	}

	// Ping the MongoDB server to verify connection
	pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pingCancel()

	if err := client.Ping(pingCtx, readpref.Primary()); err != nil {
		log.Panicf("Error connecting to MongoDB: %v", err)
	}

	log.Println("Successfully connected to MongoDB")
	return client
}
