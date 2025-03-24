package repository

import (
	"context"
	"key-haven-back/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type PasswordRepository interface {
	Create(ctx context.Context, password *model.Password) error
	FindByUserID(ctx context.Context, userID string) (*model.Password, error)
	Update(ctx context.Context, password *model.Password) error
}

type MongoPasswordRepository struct {
	repo *MongoRepository[model.Password]
}

// Create implements PasswordRepository.
func (m *MongoPasswordRepository) Create(ctx context.Context, password *model.Password) error {
	panic("unimplemented")
}

// FindByUserID implements PasswordRepository.
func (m *MongoPasswordRepository) FindByUserID(ctx context.Context, userID string) (*model.Password, error) {
	panic("unimplemented")
}

// Update implements PasswordRepository.
func (m *MongoPasswordRepository) Update(ctx context.Context, password *model.Password) error {
	panic("unimplemented")
}

func NewPasswordRepository(database *mongo.Database) PasswordRepository {
	// Create the generic repository
	collection := database.Collection("passwords")
	repo := NewMongoRepository[model.Password](collection)

	return &MongoPasswordRepository{
		repo: repo,
	}
}
