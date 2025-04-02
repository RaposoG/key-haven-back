package unittests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"key-haven-back/internal/domain/user"
	"key-haven-back/internal/http/handler"
	"key-haven-back/internal/repository"
	"key-haven-back/internal/service/dto"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuthService is a mock implementation of service.AuthService
type MockAuthService struct {
	mock.Mock
}

// ErrorResponse represents an error response structure
type ErrorResponse struct {
	Message string `json:"message"`
}

// SuccessResponse represents a success response structure
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

func (m *MockAuthService) Register(ctx context.Context, req *dto.CreateUserRequest) (*user.User, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockAuthService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.LoginResponse), args.Error(1)
}

func setupApp() *fiber.App {
	app := fiber.New()
	return app
}

func TestAuthHandler_Register(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Setup
		mockService := new(MockAuthService)
		handler := handler.NewAuthHandler(mockService)
		app := setupApp()

		user := &user.User{
			ID:    "user123",
			Email: "test@example.com",
			Name:  "Test",
		}

		reqBody := dto.CreateUserRequest{
			Email:    "test@example.com",
			Password: "password123",
			Name:     "Test",
		}
		jsonBody, _ := json.Marshal(reqBody)

		mockService.On("Register", mock.Anything, mock.MatchedBy(func(req *dto.CreateUserRequest) bool {
			return req.Email == reqBody.Email
		})).Return(user, nil)

		app.Post("/auth/register", handler.Register)

		// Execute
		req := httptest.NewRequest("POST", "/auth/register", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		// Assert
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		var response SuccessResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		responseUser, ok := response.Data.(map[string]interface{})
		assert.True(t, ok)
		assert.Equal(t, user.Email, responseUser["email"])

		mockService.AssertExpectations(t)
	})

	t.Run("Invalid Request Body", func(t *testing.T) {
		// Setup
		mockService := new(MockAuthService)
		handler := handler.NewAuthHandler(mockService)
		app := setupApp()

		// Invalid JSON
		invalidJSON := []byte(`{"email": "invalid`)

		app.Post("/auth/register", handler.Register)

		// Execute
		req := httptest.NewRequest("POST", "/auth/register", bytes.NewReader(invalidJSON))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		// Assert
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		var response ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		assert.Equal(t, "Invalid request body", response.Message)
	})

	t.Run("Email Already Used", func(t *testing.T) {
		// Setup
		mockService := new(MockAuthService)
		handler := handler.NewAuthHandler(mockService)
		app := setupApp()

		reqBody := dto.CreateUserRequest{
			Email:    "existing@example.com",
			Password: "password123",
			Name:     "Test",
		}

		jsonBody, _ := json.Marshal(reqBody)

		mockService.On("Register", mock.Anything, mock.MatchedBy(func(req *dto.CreateUserRequest) bool {
			return req.Email == reqBody.Email
		})).Return(nil, repository.ErrEmailAlreadyUsed)

		app.Post("/auth/register", handler.Register)

		// Execute
		req := httptest.NewRequest("POST", "/auth/register", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		// Assert
		assert.Equal(t, fiber.StatusConflict, resp.StatusCode)
		var response ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}
		assert.Equal(t, "Email already in use", response.Message)

		mockService.AssertExpectations(t)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		// Setup
		mockService := new(MockAuthService)
		handler := handler.NewAuthHandler(mockService)
		app := setupApp()

		reqBody := dto.CreateUserRequest{
			Email:    "test@example.com",
			Password: "password123",
			Name:     "Test",
		}
		jsonBody, _ := json.Marshal(reqBody)

		mockService.On("Register", mock.Anything, mock.MatchedBy(func(req *dto.CreateUserRequest) bool {
			return req.Email == reqBody.Email
		})).Return(nil, errors.New("database error"))

		app.Post("/auth/register", handler.Register)

		// Execute
		req := httptest.NewRequest("POST", "/auth/register", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		// Assert
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		var response ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}
		assert.Equal(t, "Failed to register user", response.Message)

		mockService.AssertExpectations(t)
	})
}

func TestAuthHandler_Login(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Setup
		mockService := new(MockAuthService)
		handler := handler.NewAuthHandler(mockService)
		app := setupApp()

		loginResp := &dto.LoginResponse{
			Token: "jwt.token.here",
			User: user.User{
				ID:    "user123",
				Email: "test@example.com",
				Name:  "Test",
			},
		}

		reqBody := dto.LoginRequest{
			Email:    "test@example.com",
			Password: "password123",
		}
		jsonBody, _ := json.Marshal(reqBody)

		mockService.On("Login", mock.Anything, mock.MatchedBy(func(req *dto.LoginRequest) bool {
			return req.Email == reqBody.Email
		})).Return(loginResp, nil)

		app.Post("/auth/login", handler.Login)

		// Execute
		req := httptest.NewRequest("POST", "/auth/login", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		// Assert
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		var response SuccessResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		responseData, ok := response.Data.(map[string]interface{})
		assert.True(t, ok)
		assert.Equal(t, loginResp.Token, responseData["token"])

		mockService.AssertExpectations(t)
	})

	t.Run("Invalid Request Body", func(t *testing.T) {
		// Setup
		mockService := new(MockAuthService)
		handler := handler.NewAuthHandler(mockService)
		app := setupApp()

		// Invalid JSON
		invalidJSON := []byte(`{"email": "invalid`)

		app.Post("/auth/login", handler.Login)

		// Execute
		req := httptest.NewRequest("POST", "/auth/login", bytes.NewReader(invalidJSON))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		// Assert
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		var response ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}
		assert.Equal(t, "Invalid request body", response.Message)
	})

	t.Run("Invalid Credentials", func(t *testing.T) {
		// Setup
		mockService := new(MockAuthService)
		handler := handler.NewAuthHandler(mockService)
		app := setupApp()

		reqBody := dto.LoginRequest{
			Email:    "test@example.com",
			Password: "wrongpassword",
		}
		jsonBody, _ := json.Marshal(reqBody)

		mockService.On("Login", mock.Anything, mock.MatchedBy(func(req *dto.LoginRequest) bool {
			return req.Email == reqBody.Email
		})).Return(nil, repository.ErrInvalidCredentials)

		app.Post("/auth/login", handler.Login)

		// Execute
		req := httptest.NewRequest("POST", "/auth/login", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		// Assert
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
		var response ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}
		assert.Equal(t, "Invalid email or password", response.Message)

		mockService.AssertExpectations(t)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		// Setup
		mockService := new(MockAuthService)
		handler := handler.NewAuthHandler(mockService)
		app := setupApp()

		reqBody := dto.LoginRequest{
			Email:    "test@example.com",
			Password: "password123",
		}
		jsonBody, _ := json.Marshal(reqBody)

		mockService.On("Login", mock.Anything, mock.MatchedBy(func(req *dto.LoginRequest) bool {
			return req.Email == reqBody.Email
		})).Return(nil, errors.New("database error"))

		app.Post("/auth/login", handler.Login)

		// Execute
		req := httptest.NewRequest("POST", "/auth/login", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		// Assert
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		var response ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}
		assert.Equal(t, "Failed to process login", response.Message)

		mockService.AssertExpectations(t)
	})
}

func TestAuthHandler_Logout(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Setup
		mockService := new(MockAuthService)
		handler := handler.NewAuthHandler(mockService)
		app := setupApp()

		app.Post("/auth/logout", handler.Logout)

		// Execute
		req := httptest.NewRequest("POST", "/auth/logout", nil)
		resp, _ := app.Test(req)

		// Assert
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		var response SuccessResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}
		assert.Equal(t, "Logged out successfully", response.Data)

		// Check cookie is cleared
		for _, cookie := range resp.Cookies() {
			if cookie.Name == "token" {
				assert.True(t, cookie.Expires.Before(time.Now()))
				break
			}
		}
	})
}
