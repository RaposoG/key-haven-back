package function

import (
	"key-haven-back/internal/http/response"

	"github.com/gofiber/fiber/v3"
)

func ResponseInternalServerErrorHandler(c fiber.Ctx, err error) error {
	var res = response.HTTPResponse{Ctx: c}
	// TODO: Implementar log de erro
	return res.InternalServerError()
}
