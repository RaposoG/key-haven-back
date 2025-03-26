package dto

import (
	"key-haven-back/internal/domain/user"
	"time"
)

type CreatePasswordRequest struct {
	UserID      string `json:"user_id" validate:"required"`
	Password    string `json:"password" validate:"required,min=12"`
	Login       string `json:"login" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type CreatePasswordResponse struct {
	UserID   string `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required,min=12"`
	Login    string `json:"login" validate:"required"`
	Title    string `json:"title" validate:"required"`
	URL      string `json:"url"`
}

// NewPassword cria uma nova instância de Password a partir de um CreatePasswordRequest.
// Gera um UUID para o ID e define os timestamps de criação e atualização.
func NewPassword(req *CreatePasswordRequest) *user.Password {
	now := time.Now()
	return &user.Password{
		UserID:      req.UserID,
		Password:    req.Password,
		Login:       req.Login,
		Title:       req.Title,
		Description: req.Description,
		URL:         req.URL,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
