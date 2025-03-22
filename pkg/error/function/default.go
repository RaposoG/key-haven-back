package function

import (
	resp "key-haven-back/internal/http/res"

	"github.com/gofiber/fiber/v3"
)

func ResponseInternalServerErrorHandler(c fiber.Ctx, err error) error {
	httpError := fiber.ErrInternalServerError
	return c.Status(httpError.Code).JSON(resp.HTTPResponseError{
		Message: httpError.Message,
		Err:     err.Error(),
		Code:    httpError.Code,
		Causes:  nil,
	})
}
