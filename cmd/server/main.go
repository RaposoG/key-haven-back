package main

import (
	"key-haven-back/config"
	error_handler "key-haven-back/pkg/error"
	"key-haven-back/pkg/validator"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &config.Config{}
	config.LoadConfig(cfg)

	fiberConfig := fiber.Config{
		StructValidator: validator.NewStructValidator(),
		ErrorHandler:    error_handler.GlobaltErrorHandler,
	}

	app := fiber.New(fiberConfig)

	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(recoverer.New())
}
