package database

import (
	"context"
	"key-haven-back/config"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBClient(cfg *config.Config) *mongo.Client {
	mongodbUrl := os.Getenv(cfg.MongodbUrl)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUrl))
	if err != nil {
		log.Panicf("Error connecting to MongoDB: %v", err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Panicf("Error connecting to MongoDB: %v", err)
	}
	return client
}
