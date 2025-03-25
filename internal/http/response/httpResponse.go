package response

import "github.com/gofiber/fiber/v3"

type HttpResponse struct {
	Ctx fiber.Ctx
}

func (hr *HttpResponse) Message(status int, message string) error {
	var data = fiber.Map{"message": message}
	return hr.Ctx.Status(status).JSON(data)
}

func (hr *HttpResponse) Ok(data any) error {
	return hr.Ctx.Status(fiber.StatusOK).JSON(data)
}

func (hr *HttpResponse) Created(data any) error {
	return hr.Ctx.Status(fiber.StatusCreated).JSON(data)
}

func (hr *HttpResponse) UnprocessableEntity(data any) error {
	return hr.Ctx.Status(fiber.StatusUnprocessableEntity).JSON(data)
}

func (hr *HttpResponse) InternalServerError() error {
	return hr.Message(fiber.StatusInternalServerError, "Internal Server Error")
}
