package service

import (
	"context"
	"key-haven-back/internal/domain/user"
	"key-haven-back/internal/repository"
	"key-haven-back/internal/service/dto"
	"key-haven-back/pkg/secret"
)

// UserService defines the interface for user-related operations
type UserService interface {
	CreateUser(ctx context.Context, request *dto.CreateUserRequest) (*user.User, error)
	GetUserByID(ctx context.Context, id string) (*user.User, error)
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
	UpdatePassword(ctx context.Context, userID, password string) error
}

// userService implements the UserService interface
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new instance of the user service
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user with the provided data
func (s *userService) CreateUser(ctx context.Context, request *dto.CreateUserRequest) (*user.User, error) {
	// Create user model from request
	user := dto.NewUser(request)

	// Hash the password
	hashedPassword, err := secret.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	// Save user to repository
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	// Don't return the password
	user.Password = ""
	return user, nil
}

// GetUserByID retrieves a user by their ID
func (s *userService) GetUserByID(ctx context.Context, id string) (*user.User, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Don't return the password
	user.Password = ""
	return user, nil
}

// GetUserByEmail retrieves a user by their email
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	return s.userRepo.FindByEmail(ctx, email)
}

// UpdatePassword updates a user's password
func (s *userService) UpdatePassword(ctx context.Context, userID, password string) error {
	hashedPassword, err := secret.HashPassword(password)
	if err != nil {
		return err
	}

	return s.userRepo.UpdatePassword(ctx, userID, hashedPassword)
}
