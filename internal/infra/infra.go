package infra

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Infra struct {
	MongoDB *mongo.Client
}

func NewInfra(mongodb *mongo.Client) *Infra {
	return &Infra{
		MongoDB: mongodb,
	}
}
