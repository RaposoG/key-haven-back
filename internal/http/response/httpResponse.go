package response

import "github.com/gofiber/fiber/v3"

type HTTPResponse struct {
	Ctx fiber.Ctx
}

func (hr *HTTPResponse) Message(status int, message string) error {
	var data = fiber.Map{"message": message}
	return hr.Ctx.Status(status).JSON(data)
}

func (hr *HTTPResponse) Ok(data any) error {
	return hr.Ctx.Status(fiber.StatusOK).JSON(data)
}

func (hr *HTTPResponse) Created(data any) error {
	return hr.Ctx.Status(fiber.StatusCreated).JSON(data)
}

func (hr *HTTPResponse) UnprocessableEntity(data any) error {
	return hr.Ctx.Status(fiber.StatusUnprocessableEntity).JSON(data)
}

func (hr *HTTPResponse) InternalServerError() error {
	return hr.Message(fiber.StatusInternalServerError, "Internal Server Error")
}
