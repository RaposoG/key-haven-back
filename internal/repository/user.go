package repository

import (
	"context"
	"errors"
	"key-haven-back/internal/domain/user"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyUsed   = errors.New("email is already in use")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// UserRepository defines the interface for user-related database operations
type UserRepository interface {
	Create(ctx context.Context, user *user.User) error
	FindByID(ctx context.Context, id string) (*user.User, error)
	FindByEmail(ctx context.Context, email string) (*user.User, error)
	UpdatePassword(ctx context.Context, userID, hashedPassword string) error
}

// MongoUserRepository implements UserRepository interface using MongoDB
type MongoUserRepository struct {
	repo *MongoRepository[user.User]
}

// NewUserRepository creates a new user repository with MongoDB implementation
func NewUserRepository(database *mongo.Database) UserRepository {
	// Create the generic repository
	collection := database.Collection("users")
	repo := NewMongoRepository[user.User](collection)

	// Create unique index on email field
	err := repo.CreateIndex("email", true)
	if err != nil {
		log.Printf("Warning: Error creating email index: %v", err)
	}

	return &MongoUserRepository{
		repo: repo,
	}
}

// Create adds a new user to the database
func (r *MongoUserRepository) Create(ctx context.Context, user *user.User) error {
	// Check if email already exists
	existingUser, err := r.FindByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return ErrEmailAlreadyUsed
	}

	// Use the generic repository to insert
	err = r.repo.Create(ctx, *user)
	if err != nil {
		if err == ErrDuplicateKey {
			return ErrEmailAlreadyUsed
		}
		return err
	}

	return nil
}

// FindByID retrieves a user by their ID
func (r *MongoUserRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	user, err := r.repo.FindByID(ctx, id, "_id")
	if err != nil {
		if err == ErrDocumentNotFound {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

// FindByEmail retrieves a user by their email address
func (r *MongoUserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	user, err := r.repo.FindOne(ctx, bson.M{"email": email})
	if err != nil {
		if err == ErrDocumentNotFound {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

// UpdatePassword updates a user's password
func (r *MongoUserRepository) UpdatePassword(ctx context.Context, userID, hashedPassword string) error {
	update := bson.M{
		"$set": bson.M{
			"password":   hashedPassword,
			"updated_at": time.Now(),
		},
	}

	err := r.repo.Update(ctx, userID, "_id", update)
	if err != nil {
		if err == ErrDocumentNotFound {
			return ErrUserNotFound
		}
		return err
	}
	return nil
}
