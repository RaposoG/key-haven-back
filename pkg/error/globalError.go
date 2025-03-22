// Package apierror provides standardized error handling for the API
package apierror

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

// Error code constants organized by HTTP status code
const (
	// --- 400 Bad Request ---
	ErrorValidation       = "VALIDATION_ERROR"
	ErrorInvalidJSON      = "INVALID_JSON"
	ErrorMissingField     = "MISSING_FIELD"
	ErrorValidationField  = "VALIDATION_FIELD"
	ErrorInvalidEmail     = "INVALID_EMAIL"
	ErrorInvalidPassword  = "INVALID_PASSWORD"
	ErrorFieldTooShort    = "FIELD_TOO_SHORT"
	ErrorFieldTooLong     = "FIELD_TOO_LONG"
	ErrorInvalidEnumValue = "INVALID_ENUM_VALUE"
	ErrorTooManyFields    = "TOO_MANY_FIELDS"

	// --- 401 Unauthorized ---
	ErrorUnauthorized       = "UNAUTHORIZED"
	ErrorMissingToken       = "MISSING_TOKEN"
	ErrorInvalidToken       = "INVALID_TOKEN"
	ErrorExpiredToken       = "EXPIRED_TOKEN"
	ErrorInvalidCredentials = "INVALID_CREDENTIALS"

	// --- 403 Forbidden ---
	ErrorForbidden               = "FORBIDDEN"
	ErrorInsufficientPermissions = "INSUFFICIENT_PERMISSIONS"
	ErrorAccountSuspended        = "ACCOUNT_SUSPENDED"
	ErrorActionNotAllowed        = "ACTION_NOT_ALLOWED"

	// --- 404 Not Found ---
	ErrorNotFound      = "NOT_FOUND"
	ErrorUserNotFound  = "USER_NOT_FOUND"
	ErrorEmailNotFound = "EMAIL_NOT_FOUND"
	ErrorPageNotFound  = "PAGE_NOT_FOUND"

	// --- 409 Conflict ---
	ErrorDuplicateEntry         = "DUPLICATE_ENTRY"
	ErrorEmailAlreadyRegistered = "EMAIL_ALREADY_REGISTERED"
	ErrorUsernameTaken          = "USERNAME_TAKEN"
	ErrorConflictingUpdate      = "CONFLICTING_UPDATE"

	// --- 422 Unprocessable Entity ---
	ErrorUnprocessableEntity   = "UNPROCESSABLE_ENTITY"
	ErrorInvalidPayload        = "INVALID_PAYLOAD"
	ErrorBusinessRuleViolation = "BUSINESS_RULE_VIOLATION"
	ErrorInconsistentData      = "INCONSISTENT_DATA"

	// --- 500 Internal Server Error ---
	ErrorInternal     = "INTERNAL_ERROR"
	ErrorDatabase     = "DATABASE_ERROR"
	ErrorCache        = "CACHE_ERROR"
	ErrorUnknown      = "UNKNOWN_ERROR"
	ErrorEncryption   = "ENCRYPTION_ERROR"
	ErrorThirdParty   = "THIRD_PARTY_ERROR"
	ErrorEmailService = "EMAIL_SERVICE_FAILURE"
)

// ValidationError represents a validation error for a specific field
type ValidationError struct {
	Code    string `json:"code"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

// APIError represents an API error response
type APIError struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Details string            `json:"details,omitempty"`
	Errors  []ValidationError `json:"errors,omitempty"`
}

// NewValidationError creates a new validation error for a specific field
func NewValidationError(code, field, message string) ValidationError {
	return ValidationError{
		Code:    code,
		Field:   field,
		Message: message,
	}
}

// NewAPIError creates a new API error with the specified code and message
func NewAPIError(code, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}

// WithDetails adds details to the API error and returns the modified error
func (e *APIError) WithDetails(details string) *APIError {
	e.Details = details
	return e
}

// WithValidationErrors adds validation errors to the API error and returns the modified error
func (e *APIError) WithValidationErrors(errors []ValidationError) *APIError {
	e.Errors = errors
	return e
}

var validationErrorType validator.ValidationErrors

// Handler is the central error handler for API errors
// It routes errors to appropriate handlers based on their type
func Handler(c fiber.Ctx, err error) error {
	switch {
	case errors.As(err, &validationErrorType):
		return HandleGenericError(c, fiber.StatusBadRequest, ErrorValidation, err)
	default:
		return HandleGenericError(c, fiber.StatusInternalServerError, ErrorInternal, err)
	}
}

// HandleGenericError creates a standardized error response with the specified status code and error code
// It handles special cases like validation errors automatically
func HandleGenericError(c fiber.Ctx, statusCode int, errorCode string, err error) error {
	// Get appropriate HTTP error message based on status code
	var message string
	switch statusCode {
	case fiber.StatusBadRequest:
		message = "Bad Request"
	case fiber.StatusUnauthorized:
		message = "Unauthorized"
	case fiber.StatusForbidden:
		message = "Forbidden"
	case fiber.StatusNotFound:
		message = "Not Found"
	case fiber.StatusMethodNotAllowed:
		message = "Method Not Allowed"
	case fiber.StatusUnprocessableEntity:
		message = "Unprocessable Entity"
	case fiber.StatusInternalServerError:
		message = "Internal Server Error"
	case fiber.StatusServiceUnavailable:
		message = "Service Unavailable"
	default:
		message = "Error occurred"
	}

	apiError := NewAPIError(errorCode, message).
		WithDetails(err.Error())

	// Handle validation errors if present
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {

		var errors []ValidationError

		for _, e := range validationErrors {
			validationError := NewValidationError(
				ErrorValidationField,
				e.Field(),
				"Field validation failed: "+e.Tag(),
			)
			errors = append(errors, validationError)
		}
		apiError.WithValidationErrors(errors)
	}

	return c.Status(statusCode).JSON(apiError)
}

//// The following handlers are deprecated and kept for backward compatibility
//// Use HandleGenericError instead
//
//// handleInternalServerError creates a standardized 500 Internal Server Error response
//func handleInternalServerError(c fiber.Ctx, err error) error {
//  return HandleGenericError(c, fiber.StatusInternalServerError, ErrorInternal, err)
//}

// handleValidationError creates a standardized 422 Unprocessable Entity response
// with detailed validation errors
//func handleValidationError(c fiber.Ctx, err error) error {
//	return HandleGenericError(c, fiber.StatusUnprocessableEntity, ErrorUnprocessableEntity, err)
//}
