package function

import (
  "github.com/go-playground/locales/pt_BR"
  ut "github.com/go-playground/universal-translator"
  "github.com/go-playground/validator/v10"
  pt_translations "github.com/go-playground/validator/v10/translations/pt_BR"
  "github.com/gofiber/fiber/v3"
)

var (
  transl   ut.Translator
  validate *validator.Validate
)

func init() {
  ptBR := pt_BR.New()
  uni := ut.New(ptBR, ptBR)
  transl, _ = uni.GetTranslator("pt_BR")

  validate = validator.New()
  err := pt_translations.RegisterDefaultTranslations(validate, transl)
  if err != nil {
    return
  }
}

var TypeValidatorError validator.ValidationErrors

func ResponseValidatorHandler(c fiber.Ctx, err error) error {
  httpError := fiber.ErrUnprocessableEntity

  var causes []Causes

  for _, err := range err.(validator.ValidationErrors) {
    causes = append(causes, Causes{
      Code:    "validation_error",
      Field:   err.Field(),
      Message: err.Translate(transl),
    })
  }

  return c.Status(httpError.Code).JSON(APIError{
    Code:    httpError.Message,
    Message: httpError.Message,
    Errors:  causes,
    Details: "",
  })

}

type Causes struct {
  Code    string `json:"code"`
  Field   string `json:"field"`
  Message string `json:"message"`
}

type APIError struct {
  Code    string   `json:"code"`
  Message string   `json:"message"`
  Details string   `json:"details,omitempty"`
  Errors  []Causes `json:"errors,omitempty"`
}
