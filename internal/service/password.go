package service

import (
	"context"
	"key-haven-back/internal/domain/user"
	"key-haven-back/internal/repository"
	"key-haven-back/internal/service/dto"
)

type PasswordService interface {
	Register(ctx context.Context, request *dto.CreatePasswordRequest) (*user.Password, error)
}

type passwordService struct {
	repo repository.PasswordRepository
}

func NewPasswordService(repo repository.PasswordRepository) PasswordService {
	return &passwordService{
		repo: repo,
	}
}

func (s *passwordService) Register(ctx context.Context, request *dto.CreatePasswordRequest) (*user.Password, error) {
	return nil, nil
}
