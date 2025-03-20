package service

import (
	"context"
	"key-haven-back/internal/model"
	"key-haven-back/internal/repository"
	"key-haven-back/pkg/secret"
	"time"
)

type AuthService interface {
	Register(ctx context.Context, request *model.CreateUserRequest) (*model.User, error)
	Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error)
}

type authService struct {
	userService UserService
}

func NewAuthService(userService UserService) AuthService {
	return &authService{
		userService: userService,
	}
}

func (s *authService) Register(ctx context.Context, request *model.CreateUserRequest) (*model.User, error) {
	return s.userService.CreateUser(ctx, request)
}

func (s *authService) Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error) {
	user, err := s.userService.GetUserByEmail(ctx, request.Email)
	if err != nil {
		if err == repository.ErrUserNotFound {
			return nil, repository.ErrInvalidCredentials
		}
		return nil, err
	}

	valid, err := secret.VerifyPassword(user.Password, request.Password)
	if err != nil || !valid {
		return nil, repository.ErrInvalidCredentials
	}

	token, err := secret.GenerateToken(user.ID, user.Email, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	// Return login response
	return &model.LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}
