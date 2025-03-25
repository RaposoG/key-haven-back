package model

import "time"

type Key struct {
	ID        string `json:"id" bson:"_id"`
	MasterKey string `json:"master_key" bson:"master_key"`

	UserID string `json:"user_id" bson:"user_id"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
