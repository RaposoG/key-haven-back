package main

import (
	"key-haven-back/config"
	"key-haven-back/internal/infra/database"
	"key-haven-back/internal/repository"
	"key-haven-back/internal/service"
	"log"

	_ "key-haven-back/docs"

	"github.com/joho/godotenv"
	"go.uber.org/fx"

	httpapi "key-haven-back/internal/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file")
	}

	app := fx.New(
		// Core modules
		config.Module,
		database.Module,

		// Repository modules
		repository.Module,

		// Service modules
		service.Module,

		// HTTP API modules
		httpapi.Module,
	)

	app.Run()

}
