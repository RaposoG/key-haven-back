package model

import "time"

type User struct {
	ID           string `json:"id" bson:"_id"`
	Name         string `json:"name" bson:"name"`
	Email        string `json:"email" bson:"email"`
	HashPassword string `json:"hash_password" bson:"hash_password"`

	TOTPKey       string   `json:"totpKey" bson:"totpKey"`
	RecoveryCodes []string `json:"recoveryCodes" bson:"recoveryCodes"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
