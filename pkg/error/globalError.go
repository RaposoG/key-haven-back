package error

import (
	"key-haven-back/pkg/error/function"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func GlobalErrorHandler(c fiber.Ctx, err error) error {
	switch err.(type) {
	case validator.ValidationErrors:
		return function.ResponseValidatorHandler(c, err)
	default:
		return function.ResponseInternalServerErrorHandler(c, err)
	}
}
