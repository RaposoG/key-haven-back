package model

import (
	"time"

	"github.com/google/uuid"
)

type Password struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	UserID      string    `json:"user_id" bson:"user_id"`
	Password    string    `json:"password" bson:"password"`
	Login       string    `json:"login" bson:"login"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Url         string    `json:"url" bson:"url"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

type CreatePasswordRequest struct {
	UserID      string `json:"user_id" validate:"required"`
	Password    string `json:"password" validate:"required,min=12"`
	Login       string `json:"login" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

func NewPassword(req *CreatePasswordRequest) *Password {
	now := time.Now()
	return &Password{
		ID:          uuid.New().String(),
		UserID:      req.UserID,
		Password:    req.Password,
		Login:       req.Login,
		Title:       req.Title,
		Description: req.Description,
		Url:         req.Url,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
