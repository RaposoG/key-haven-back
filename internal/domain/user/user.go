package user

import "time"

// User represents a user entity
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
