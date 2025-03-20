package docs

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path"

	"github.com/gofiber/fiber/v3"
)

type Provider struct {
	URL  string
	Name string
}

func RegisterDocsRouter(app *fiber.App) {
	app.Get("/public/openapi.json", func(ctx fiber.Ctx) error {
		dir, err := os.Getwd()
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Fail to get current directory: "+err.Error())
		}
		filename := path.Join(dir, "pkg", "docs", "swagger.json")
		return ctx.SendFile(filename)
	})

	app.Get("/docs", func(ctx fiber.Ctx) error {
		dir, err := os.Getwd()
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Fail to get current directory: "+err.Error())
		}

		name := ctx.Query("name", "scalar")
		filename := path.Join(dir, "pkg", "docs", fmt.Sprintf("%s.html", name))
		tmpl, err := template.ParseFiles(filename)
		if err != nil {
			fmt.Println("ParseFiles error")
			return fiber.NewError(fiber.StatusInternalServerError, "Fail to generate docs: "+err.Error())
		}

		provider := &Provider{
			URL:  "http://localhost:8080/public/openapi.json",
			Name: name,
		}

		var bufferHtml bytes.Buffer
		if err := tmpl.Execute(&bufferHtml, provider); err != nil {
			fmt.Println("Execute Replace Text")
			return fiber.NewError(fiber.StatusInternalServerError, "Fail to generate docs: "+err.Error())
		}

		return ctx.Type("html").SendString(bufferHtml.String())
	})
}
