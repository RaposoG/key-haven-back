package handler

import (
	_ "key-haven-back/docs"

	"github.com/gofiber/fiber/v3"
)

type ResponseSuccess struct {
	Message string `json:"message,omitempty"`
}

func Health(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(ResponseSuccess{
		Message: "Healthy",
	})
}
