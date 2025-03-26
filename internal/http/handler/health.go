package handler

import (
	_ "key-haven-back/docs"

	"github.com/gofiber/fiber/v3"
)

type ResponseSuccess struct {
	Message string `json:"message,omitempty"`
}

// Health godoc
// @Summary Check if the service is healthy
// @Description Check if the service is healthy
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} ResponseSuccess
// @Router / [get]
func Health(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(ResponseSuccess{
		Message: "Healthy",
	})
}
