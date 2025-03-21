package http

import (
	"context"
	"fmt"
	"key-haven-back/config"
	"key-haven-back/internal/handler"
	"key-haven-back/internal/router"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"go.uber.org/fx"

	docsPkg "key-haven-back/pkg/docs"
)

func StartServer(lc fx.Lifecycle, app *fiber.App, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := app.Listen(fmt.Sprintf(":%s", cfg.Port)); err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}

func NewServer(
	cfg *config.Config,
	authHandler *handler.AuthHandler,
	registerRoutes router.RegisterRoutesFunc,
	registerSwagger router.RegisterSwaggerRoutesFunc,
	registerDocs docsPkg.RegisterDocsRouterFunc,
) *fiber.App {
	app := fiber.New()

	// Configure middlewares
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(recover.New())

	// Health check route
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Key Haven API is running!")
	})

	// Register API routes
	registerRoutes(app, authHandler)
	registerSwagger(app)
	registerDocs(app)

	return app
}
