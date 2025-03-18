package http

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

func StartServer(lc fx.Lifecycle, app *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {

				if err := app.Listen(":8081"); err != nil {
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

func NewServer() (*fiber.App, error) {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return app, nil
}
