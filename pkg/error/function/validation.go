package function

import (
	resp "key-haven-back/internal/http/res"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var TypeValidatorError validator.ValidationErrors

func ResponseValidatorHandler(c fiber.Ctx, err error) error {
	httpError := fiber.ErrUnprocessableEntity
	return c.Status(httpError.Code).JSON(resp.HTTPResponseError{
		Massage: httpError.Message,
		Err:     err.Error(),
		Code:    httpError.Code,
		Causes:  nil,
	})
}
