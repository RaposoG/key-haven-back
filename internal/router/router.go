package router

import (
	"key-haven-back/internal/handler"

	"github.com/gofiber/fiber/v3"
)

// RegisterRoutesFunc type of function to register routes
type RegisterRoutesFunc func(app *fiber.App, authHandler *handler.AuthHandler)
type RegisterSwaggerRoutesFunc func(app *fiber.App)

// RegisterSwaggerRoutesFunc type of function to register swagger routes
func RegisterRoutesFuncProvider() RegisterRoutesFunc {
	return RegisterRoutes
}

// RegisterSwaggerRoutesFuncProvider provider for Fx
func RegisterSwaggerRoutesFuncProvider() RegisterSwaggerRoutesFunc {

	return RegisterSwaggerRoutes
}
