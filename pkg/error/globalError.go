package handler

import (
	"errors"
	"key-haven-back/pkg/error/function"

	"github.com/gofiber/fiber/v3"
)

func GlobalErrorHandler(c fiber.Ctx, err error) error {
	switch {
	case errors.As(err, &function.TypeValidatorError):
		return function.ResponseValidatorHandler(c, err)
	default:
		return function.ResponseInternalServerErrorHandler(c, err)
	}
}
