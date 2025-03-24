package service

import (
	"context"
	"key-haven-back/internal/model"
	"key-haven-back/internal/repository"
)

type PasswordService interface {
	Register(ctx context.Context, request *model.CreatePasswordRequest) (*model.Password, error)
}

type passwordService struct {
	repo repository.PasswordRepository
}

func NewPasswordService(repo repository.PasswordRepository) PasswordService {
	return &passwordService{
		repo: repo,
	}
}

func (s *passwordService) Register(ctx context.Context, request *model.CreatePasswordRequest) (*model.Password, error) {
	return nil, nil
}
