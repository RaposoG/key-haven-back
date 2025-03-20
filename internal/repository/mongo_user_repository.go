package repository

import (
	"context"
	"errors"
	"key-haven-back/internal/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(collection *mongo.Collection) UserRepository {
	// Create unique index on email field
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	// Create the index in the background
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Printf("Error creating index: %v", err)
	}

	return &MongoUserRepository{
		collection: collection,
	}
}

func (r *MongoUserRepository) Create(ctx context.Context, user *model.User) error {
	// Check if email already exists
	existingUser, err := r.FindByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return ErrEmailAlreadyUsed
	}

	// Insert user document
	_, err = r.collection.InsertOne(ctx, user)
	if err != nil {
		// Check for duplicate key error
		if mongo.IsDuplicateKeyError(err) {
			return ErrEmailAlreadyUsed
		}
		return err
	}

	return nil
}

func (r *MongoUserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User

	filter := bson.M{"_id": id}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *MongoUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	filter := bson.M{"email": email}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *MongoUserRepository) UpdatePassword(ctx context.Context, userID, hashedPassword string) error {
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$set": bson.M{
			"password":   hashedPassword,
			"updated_at": time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return ErrUserNotFound
	}

	return nil
}
