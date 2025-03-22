package res

import (
	"key-haven-back/pkg/error"
	"time"

	"github.com/gofiber/fiber/v3"
)

type ErrorResponse struct {
	Error *apierror.APIError `json:"error"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

func HandleError(c fiber.Ctx, status int, apiError *apierror.APIError) error {
	return c.Status(status).JSON(ErrorResponse{
		Error: apiError,
	})
}

func HandleValidationError(c fiber.Ctx, statusCode int, validationErrors []apierror.ValidationError) error {
	return HandleError(c, statusCode,
		apierror.NewAPIError(apierror.ErrorValidation, "Validation failed").
			WithValidationErrors(validationErrors))
}

func SetAuthCookie(c fiber.Ctx, token string, duration time.Duration) {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(duration),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})
}
