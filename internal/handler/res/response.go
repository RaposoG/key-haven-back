package res

import "github.com/gofiber/fiber/v3"

type HttpResponse struct {
	Ctx fiber.Ctx
}

type Message struct {
	Message string `json:"message"`
}

func (r *HttpResponse) Send(status int, message string) error {
	return r.Ctx.Status(status).JSON(Message{Message: message})
}
