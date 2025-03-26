package router

import (
	handler2 "key-haven-back/internal/http/handler"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App, authHandler *handler2.AuthHandler, passwordHandler *handler2.PasswordHandler) {
	auth := app.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Post("/logout", authHandler.Logout)

	password := app.Group("/password")
	password.Post("/", passwordHandler.Register)
}
