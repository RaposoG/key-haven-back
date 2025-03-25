package function

import (
	"key-haven-back/internal/http/response"
	"key-haven-back/pkg/validator/langs"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type ErrorResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ResponseValidatorHandler(c fiber.Ctx, err error) error {
	var res = response.HTTPResponse{Ctx: c}
	var acceptLanguage = c.Get("accept-language", "en")
	var trans = langs.SetTranslate(acceptLanguage)

	var causes []Causes
	for _, e := range err.(validator.ValidationErrors) {
		causes = append(causes, Causes{
			Field:   e.Field(),
			Message: e.Translate(trans),
		})
	}

	var data ErrorResponse
	data.Code = 422
	data.Message = "Unprocessable Entity"
	data.Causes = causes

	return res.UnprocessableEntity(data)
}
