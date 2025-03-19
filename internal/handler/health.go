package handler

import (
	_ "key-haven-back/docs"

	"github.com/gofiber/fiber/v3"
)

type ResponseSuccess struct {
	Message string `json:"message,omitempty"`
}

// Exemple Doc
// @Summary      Get Hello
// @Description  Simple get hello word
// @Tags         Hello
// @Produce      json
// @Success      200				 {object}  ResponseSuccess
// @Router       / [get]
func Health(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(ResponseSuccess{
		Message: "Healthy",
	})
}
