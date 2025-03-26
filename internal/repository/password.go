package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"key-haven-back/internal/domain/user"
)

type PasswordRepository interface {
	Create(ctx context.Context, password *user.Password) error
	FindByUserID(ctx context.Context, userID string) (*user.Password, error)
	Update(ctx context.Context, password *user.Password) error
}

type MongoPasswordRepository struct {
	repo *MongoRepository[user.Password]
}

// Create implements PasswordRepository.
func (m *MongoPasswordRepository) Create(ctx context.Context, password *user.Password) error {
	panic("unimplemented")
}

// FindByUserID implements PasswordRepository.
func (m *MongoPasswordRepository) FindByUserID(ctx context.Context, userID string) (*user.Password, error) {
	panic("unimplemented")
}

// Update implements PasswordRepository.
func (m *MongoPasswordRepository) Update(ctx context.Context, password *user.Password) error {
	panic("unimplemented")
}

func NewPasswordRepository(database *mongo.Database) PasswordRepository {
	// Create the generic repository
	collection := database.Collection("passwords")
	repo := NewMongoRepository[user.Password](collection)

	return &MongoPasswordRepository{
		repo: repo,
	}
}
