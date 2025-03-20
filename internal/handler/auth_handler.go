package handler

import (
	"key-haven-back/internal/model"
	"key-haven-back/internal/repository"
	"key-haven-back/internal/service"
	"time"

	"github.com/gofiber/fiber/v3"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// MessageResponse represents a simple message response
type MessageResponse struct {
	Message string `json:"message"`
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

// Register godoc
// @Summary Register a new user
// @Description Creates a new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body model.CreateUserRequest true "User registration data"
// @Success 201 {object} model.User
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 409 {object} ErrorResponse "Email already in use"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /auth/register [post]
func (h *AuthHandler) Register(c fiber.Ctx) error {
	var req model.CreateUserRequest
	if err := c.Bind().Body(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	user, err := h.authService.Register(c.Context(), &req)
	if err != nil {
		if err == repository.ErrEmailAlreadyUsed {
			return fiber.NewError(fiber.StatusConflict, "Email already in use")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to register user")
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// Login godoc
// @Summary User login
// @Description Authenticates a user and returns a token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body model.LoginRequest true "User login data"
// @Success 200 {object} model.LoginResponse
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 401 {object} ErrorResponse "Invalid email or password"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req model.LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	response, err := h.authService.Login(c.Context(), &req)
	if err != nil {
		if err == repository.ErrInvalidCredentials {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to process login")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    response.Token,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.Status(fiber.StatusOK).JSON(response)
}

// Logout godoc
// @Summary User logout
// @Description Logs out the user by clearing the authentication cookie
// @Tags Auth
// @Produce json
// @Success 200 {object} MessageResponse
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Security ApiKeyAuth
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour), // Expired
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.Status(fiber.StatusOK).JSON(MessageResponse{
		Message: "Logged out successfully",
	})
}
