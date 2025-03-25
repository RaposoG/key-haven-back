package model

type Credential struct {
	ID              string `json:"id" bson:"_id"`
	Name            string `json:"name" bson:"name"`
	Description     string `json:"description" bson:"description"`
	Username        string `json:"username" bson:"username"`
	EncryptPassword string `json:"encrypt_password" bson:"encrypt_password"`
	URL             string `json:"url" bson:"url"`

	ValtID string `json:"valt_id" bson:"valt_id"`

	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
}
