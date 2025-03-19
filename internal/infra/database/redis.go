package database

import (
	"context"
	"fmt"
	"key-haven-back/config"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       0,
	})

	status := client.Ping(context.Background())
	if status.Err() != nil {
		log.Panicf("Failed to connect to Redis: %v", status.Err())
	}

	return client
}
