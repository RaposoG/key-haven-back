package router

import (
	handler2 "key-haven-back/internal/http/handler"

	"github.com/gofiber/fiber/v3"
)

type RegisterRoutesFunc func(app *fiber.App, authHandler *handler2.AuthHandler, passwordHandler *handler2.PasswordHandler)
type RegisterSwaggerRoutesFunc func(app *fiber.App)

func RegisterRoutesFuncProvider() RegisterRoutesFunc {
	return RegisterRoutes
}

func RegisterSwaggerRoutesFuncProvider() RegisterSwaggerRoutesFunc {
	return RegisterSwaggerRoutes
}
