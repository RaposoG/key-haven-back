package user

import "time"

// User represents a user entity
// TODO: replace this with UserV2
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

type UserV2 struct {
	ID           string `json:"id" bson:"_id"`
	Name         string `json:"name" bson:"name"`
	Email        string `json:"email" bson:"email"`
	HashPassword string `json:"hash_password" bson:"hash_password"`

	TOTPKey       string   `json:"totpKey" bson:"totpKey"`
	RecoveryCodes []string `json:"recoveryCodes" bson:"recoveryCodes"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
