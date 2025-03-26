package user

import "time"

// Password represents a password entity
type Password struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	UserID      string    `json:"user_id" bson:"user_id"`
	Password    string    `json:"-" bson:"password"`
	Login       string    `json:"login" bson:"login"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	URL         string    `json:"url" bson:"url"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}
