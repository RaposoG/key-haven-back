package handler

import (
	"errors"
	"key-haven-back/internal/model"
	"key-haven-back/internal/repository"
	"key-haven-back/internal/service"
	apierror "key-haven-back/pkg/error"
	"key-haven-back/pkg/validator"
	"time"

	"github.com/gofiber/fiber/v3"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error *apierror.APIError `json:"error"`
}

// SuccessResponse represents a generic success response
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// AuthHandler handles authentication requests
type AuthHandler struct {
	authService service.AuthService
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// handleError centralizes error handling
func handleError(c fiber.Ctx, status int, apiError *apierror.APIError) error {
	return c.Status(status).JSON(ErrorResponse{
		Error: apiError,
	})
}

// handleValidationError handles validation errors
func handleValidationError(c fiber.Ctx, statusCode int, validationErrors []apierror.ValidationError) error {
	return handleError(c, statusCode,
		apierror.NewAPIError(apierror.ErrorUnprocessableEntity, "Validation failed").
			WithValidationErrors(validationErrors))
}

// setAuthCookie sets the authentication cookie
func setAuthCookie(c fiber.Ctx, token string, duration time.Duration) {
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

// Register godoc
// @Summary Register a new user
// @Description Creates a new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body apierror.CreateUserRequest true "User registration data"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse "Invalid request body or validation errors"
// @Failure 409 {object} ErrorResponse "Email already in use"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /auth/register [post]
func (h *AuthHandler) Register(c fiber.Ctx) error {
	var req model.CreateUserRequest
	if err := c.Bind().Body(&req); err != nil {
		return handleError(c, fiber.StatusUnprocessableEntity,
			apierror.NewAPIError(apierror.ErrorValidation, "Invalid request body"))
	}

	// Validate the request
	if valid, resErrors := validator.Validate(&req); !valid {
		return handleValidationError(c, fiber.StatusUnprocessableEntity, resErrors)
	}

	user, err := h.authService.Register(c.Context(), &req)
	if err != nil {
		if errors.Is(err, repository.ErrEmailAlreadyUsed) {
			return handleError(c, fiber.StatusConflict,
				apierror.NewAPIError(apierror.ErrorDuplicateEntry, "Email already used"))
		}
		return handleError(c, fiber.StatusInternalServerError,
			apierror.NewAPIError(apierror.ErrorInternal, "Failed to register user"))
	}

	return c.Status(fiber.StatusCreated).JSON(SuccessResponse{Data: user})
}

// Login godoc
// @Summary User login
// @Description Authenticates a user and returns a token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body apierror.LoginRequest true "User login data"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 401 {object} ErrorResponse "Invalid email or password"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req model.LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		return handleError(c, fiber.StatusBadRequest,
			apierror.NewAPIError(apierror.ErrorValidation, "Invalid request body"))
	}

	response, err := h.authService.Login(c.Context(), &req)
	if err != nil {
		if errors.Is(err, repository.ErrInvalidCredentials) {
			return handleError(c, fiber.StatusUnauthorized,
				apierror.NewAPIError(apierror.ErrorUnauthorized, "Invalid login"))
		}
		return handleError(c, fiber.StatusInternalServerError,
			apierror.NewAPIError(apierror.ErrorInternal, "Failed to login"))
	}

	setAuthCookie(c, response.Token, 24*time.Hour)
	return c.Status(fiber.StatusOK).JSON(SuccessResponse{Data: response})
}

// Logout godoc
// @Summary User logout
// @Description Logs out the user by clearing the authentication cookie
// @Tags Auth
// @Produce json
// @Success 200 {object} SuccessResponse
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Security ApiKeyAuth
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c fiber.Ctx) error {
	setAuthCookie(c, "", -time.Hour) // Expire the cookie
	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Data: "Logged out successfully",
	})
}
