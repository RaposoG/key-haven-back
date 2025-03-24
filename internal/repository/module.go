package repository

import (
	"key-haven-back/internal/infra/database"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"repository",
	fx.Provide(
		func(client database.MongoDBClient) *mongo.Collection {
			return client.Database("key-haven").Collection("users")
		},
		NewUserRepository,
		NewPasswordRepository,
	),
)
