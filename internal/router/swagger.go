package router

import (
	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gofiber/fiber/v3"
)

func RegisterSwaggerRoutes(app *fiber.App) {

	app.Get("/reference", func(c fiber.Ctx) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
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
}
