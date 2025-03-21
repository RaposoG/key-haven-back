package router

import (
	"key-haven-back/internal/handler"

	"github.com/gofiber/fiber/v3"
)

// Tipos de funções para injeção no Fx
type RegisterRoutesFunc func(app *fiber.App, authHandler *handler.AuthHandler)
type RegisterSwaggerRoutesFunc func(app *fiber.App)

// Providers para o Fx
func RegisterRoutesFuncProvider() RegisterRoutesFunc {
	return RegisterRoutes
}

func RegisterSwaggerRoutesFuncProvider() RegisterSwaggerRoutesFunc {

	return RegisterSwaggerRoutes
}
