package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBClient is an interface that represents a MongoDB client
type MongoDBClient interface {
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
	Disconnect(ctx context.Context) error
}

// Ensure *mongo.Client implements MongoDBClient
var _ MongoDBClient = (*mongo.Client)(nil)
