package main

import (
	"context"
	"key-haven-back/config"
	"key-haven-back/internal/handler"
	"key-haven-back/internal/infra/database"
	"key-haven-back/internal/repository"
	"key-haven-back/internal/router"
	"key-haven-back/internal/service"
	"key-haven-back/pkg/docs"
	error_handler "key-haven-back/pkg/error"
	"key-haven-back/pkg/validator"
	"log"

	_ "key-haven-back/docs"

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
		log.Println("Warning: Error loading .env file")
	}

	cfg := &config.Config{}
	config.LoadConfig(cfg)

	// Initialize database clients
	mongoClient := database.NewMongoDBClient(cfg)

	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			log.Printf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	// Get the users collection from MongoDB
	usersCollection := mongoClient.Database("key-haven").Collection("users")

	// Initialize repositories
	userRepo := repository.NewMongoUserRepository(usersCollection)

	// Initialize services
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userService)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService)

	fiberConfig := fiber.Config{
		StructValidator: validator.NewStructValidator(),
		ErrorHandler:    error_handler.GlobaltErrorHandler,
	}

	app := fiber.New(fiberConfig)
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(recoverer.New())

	app.Get("/health", handler.Health)

	// Authentication routes
	router.RegisterRoutes(app, authHandler)
	router.RegisterSwaggerRoutes(app)
	docs.RegisterDocsRouter(app)

	// Start the server
	log.Fatal(app.Listen(":8080"))
}
