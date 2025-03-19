package main

import (
	"key-haven-back/config"
	"key-haven-back/internal/handler"
	error_handler "key-haven-back/pkg/error"
	"key-haven-back/pkg/validator"
	"log"

	_ "key-haven-back/docs"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/joho/godotenv"
)

// @title Key Haven API
// @version 1.0
// @description This is the API for Key Haven
// @host localhost:8080
// @BasePath /
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

	app.Get("/health", handler.Health)

	app.Get("/reference", func(c fiber.Ctx) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			// SpecURL: "https://generator3.swagger.io/openapi.json",// allow external URL or local path file
			SpecURL: "./docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "API Reference",
			},
			DarkMode: true,
		})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Fail to generate swagger API: "+err.Error())
		}
		return c.Type("html").SendString(htmlContent)
	})

	log.Fatal(app.Listen(":8080"))
}
