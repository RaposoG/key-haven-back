package router

import (
	"key-haven-back/internal/handler"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App, authHandler *handler.AuthHandler) {
	auth := app.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Post("/logout", authHandler.Logout)
}
