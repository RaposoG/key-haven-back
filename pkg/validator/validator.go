package validator

import (
	"errors"
	"key-haven-back/pkg/error"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// getJSONFieldName extracts the field name from the JSON tag of a struct field
func getJSONFieldName(s interface{}, fieldName string) string {
	structField, ok := reflect.TypeOf(s).Elem().FieldByName(fieldName)
	if !ok {
		return strings.ToLower(fieldName)
	}

	jsonTag := structField.Tag.Get("json")
	if jsonTag == "" || jsonTag == "-" {
		return strings.ToLower(fieldName)
	}

	return strings.Split(jsonTag, ",")[0]
}

// generateErrorMessage creates a user-friendly error message and code for a validation error
func generateErrorMessage(tag string, param string) (string, string) {
	switch tag {
	case "required":
		return "This field is required", apierror.ErrorMissingField
	case "email":
		return "Invalid email format", apierror.ErrorInvalidEmail
	case "min":
		return "Field is below minimum length it must be at least " + param + " characters long", apierror.ErrorFieldTooShort
	case "max":
		return "Must be at most " + param + " characters long", apierror.ErrorFieldTooLong
	default:
		return "Failed validation on tag: " + tag, ""
	}
}

// Validate validates a struct and returns validation errors
func Validate(s interface{}) (bool, []apierror.ValidationError) {
	err := validate.Struct(s)
	if err == nil {
		return true, nil
	}

	var validationErrors validator.ValidationErrors
	ok := errors.As(err, &validationErrors)
	if !ok {
		// If error is not of expected type, return a generic error
		return false, []apierror.ValidationError{{
			Code:    "validation_error",
			Field:   "general",
			Message: "Invalid input",
		}}
	}

	var e []apierror.ValidationError
	for _, err := range validationErrors {
		fieldName := err.Field()
		jsonField := getJSONFieldName(s, fieldName)
		message, code := generateErrorMessage(err.Tag(), err.Param())

		e = append(e, apierror.ValidationError{
			Code:    code,
			Field:   jsonField,
			Message: message,
		})
	}

	return false, e
}
