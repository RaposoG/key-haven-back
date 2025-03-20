package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                  string    `json:"id" bson:"_id,omitempty"`
	Email               string    `json:"email" bson:"email"`
	Password            string    `json:"-" bson:"password"`
	Name                string    `json:"name" bson:"name"`
	CreatedAt           time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" bson:"updated_at"`
	LastLoginAt         time.Time `json:"last_login_at,omitempty" bson:"last_login_at"`
	FailedLoginAttempts int       `json:"-" bson:"failed_login_attempts"`
}

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=12"`
	Name     string `json:"name" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

func NewUser(req *CreateUserRequest) *User {
	now := time.Now()
	return &User{
		ID:          uuid.New().String(),
		Email:       req.Email,
		Name:        req.Name,
		CreatedAt:   now,
		UpdatedAt:   now,
		LastLoginAt: now,
	}
}
