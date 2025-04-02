package middleware

import (
	"errors"
	"key-haven-back/pkg/secret"

	"github.com/gofiber/fiber/v3"
)

func IsAuthenticatedHandler(c fiber.Ctx) error {
	tokenJwt := c.Cookies("token")
	if tokenJwt == "" {
		// Also check for Authorization header
		authHeader := c.Get("Authorization")
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenJwt = authHeader[7:]
		}

		if tokenJwt == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
	}

	// Validate token
	claims, err := secret.ValidateToken(tokenJwt)
	if err != nil {
		if errors.Is(err, secret.ErrTokenExpired) {
			return c.Status(401).JSON(fiber.Map{
				"message": "Token expired",
			})
		}
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	// Store claims in context for use in route handlers
	c.Locals("user_id", claims.UserID)
	c.Locals("email", claims.Email)

	return c.Next()
}
