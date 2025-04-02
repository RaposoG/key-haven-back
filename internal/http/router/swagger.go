package router

import (
	scalar "github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gofiber/fiber/v3"
)

// RegisterSwaggerRoutes registers the Swagger documentation routes
func RegisterSwaggerRoutes(app *fiber.App) {
	// Swagger JSON
	app.Get("/swagger.json", func(c fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})

	app.Get("/reference", func(c fiber.Ctx) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Key Haven API Reference",
			},
			DarkMode: true,
		})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate API reference: "+err.Error())
		}
		return c.Type("html").SendString(htmlContent)
	})
}
