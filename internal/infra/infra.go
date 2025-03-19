package infra

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Infra struct {
	Redis   *redis.Client
	MongoDB *mongo.Client
}

func NewInfra(redis *redis.Client, mongodb *mongo.Client) *Infra {
	return &Infra{
		Redis:   redis,
		MongoDB: mongodb,
	}
}
