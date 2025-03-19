package middleware

import (
	"github.com/gofiber/fiber/v3"
)

func IsAuthenticatedHandler(c fiber.Ctx) error {
	tokenJwt := c.Cookies("token")
	if tokenJwt == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
