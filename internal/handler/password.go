package handler

import (
	"key-haven-back/internal/handler/res"
	"key-haven-back/internal/service"

	"github.com/gofiber/fiber/v3"
)

type PasswordHandler struct {
	passwordService service.PasswordService
}

func NewPasswordHandler(passwordService service.PasswordService) *PasswordHandler {
	return &PasswordHandler{
		passwordService: passwordService,
	}
}

func (h *PasswordHandler) Register(ctx fiber.Ctx) error {
	response := res.HttpResponse{Ctx: ctx}
	return response.Send(200, "PasswordHandler.Register")
}
