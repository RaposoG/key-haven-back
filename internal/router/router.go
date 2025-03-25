package router

import (
	"key-haven-back/internal/handler"

	"github.com/gofiber/fiber/v3"
)

type RegisterRoutesFunc func(app *fiber.App, authHandler *handler.AuthHandler, passwordHandler *handler.PasswordHandler)
type RegisterSwaggerRoutesFunc func(app *fiber.App)

func RegisterRoutesFuncProvider() RegisterRoutesFunc {
	return RegisterRoutes
}

func RegisterSwaggerRoutesFuncProvider() RegisterSwaggerRoutesFunc {
	return RegisterSwaggerRoutes
}
