package database

import (
	"go.uber.org/fx"
)

// MongoDBClient is an alias for mongo.Client

var Module = fx.Module(
	"database",
	fx.Provide(
		NewMongoDBClient,
	),
)
