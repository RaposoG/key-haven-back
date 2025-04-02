package repository

import (
	"context"
	"errors"
	"key-haven-back/internal/domain/user"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrVaultNotFound      = errors.New("vault not found")
	ErrVaultNameExists    = errors.New("vault name already exists for this user")
	ErrDefaultVaultExists = errors.New("default vault already exists for this user")
)

type VaultRepository interface {
	Create(ctx context.Context, vault *user.Vault) error
	FindByID(ctx context.Context, id string) (*user.Vault, error)
	FindByName(ctx context.Context, userID, name string) (*user.Vault, error)
	FindDefaultByUserID(ctx context.Context, userID string) (*user.Vault, error)
	FindAllByUserID(ctx context.Context, userID string) ([]*user.Vault, error)
	Update(ctx context.Context, vault *user.Vault) error
	Delete(ctx context.Context, id string) error
}

type MongoVaultRepository struct {
	repo *MongoRepository[user.Vault]
}

func NewVaultRepository(database *mongo.Database) VaultRepository {
	collection := database.Collection("vaults")
	repo := NewMongoRepository[user.Vault](collection)

	_, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "user_id", Value: 1}, {Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Printf("Warning: Error creating compound index for vaults: %v", err)
	}

	return &MongoVaultRepository{
		repo: repo,
	}
}

func (r *MongoVaultRepository) Create(ctx context.Context, vault *user.Vault) error {
	if vault.ID == "" {
		vault.ID = uuid.New().String()
	}

	existingVault, err := r.FindByName(ctx, vault.UserID, vault.Name)
	if err == nil && existingVault != nil {
		return ErrVaultNameExists
	}

	if vault.Name == "Default" {
		defaultVault, err := r.FindDefaultByUserID(ctx, vault.UserID)
		if err == nil && defaultVault != nil {
			return ErrDefaultVaultExists
		}
	}

	return r.repo.Create(ctx, *vault)
}

func (r *MongoVaultRepository) FindByID(ctx context.Context, id string) (*user.Vault, error) {
	vault, err := r.repo.FindByID(ctx, id, "_id")
	if err != nil {
		if errors.Is(err, ErrDocumentNotFound) {
			return nil, ErrVaultNotFound
		}
		return nil, err
	}
	return vault, nil
}

// FindByName retrieves a vault by user ID and vault name
func (r *MongoVaultRepository) FindByName(ctx context.Context, userID, name string) (*user.Vault, error) {
	filter := bson.M{"user_id": userID, "name": name}
	vault, err := r.repo.FindOne(ctx, filter)
	if err != nil {
		if errors.Is(err, ErrDocumentNotFound) {
			return nil, ErrVaultNotFound
		}
		return nil, err
	}
	return vault, nil
}

// FindDefaultByUserID retrieves the default vault for a user
func (r *MongoVaultRepository) FindDefaultByUserID(ctx context.Context, userID string) (*user.Vault, error) {
	return r.FindByName(ctx, userID, "Default")
}

// FindAllByUserID retrieves all vaults belonging to a user
func (r *MongoVaultRepository) FindAllByUserID(ctx context.Context, userID string) ([]*user.Vault, error) {
	filter := bson.M{"user_id": userID}
	vaults, err := r.repo.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	result := make([]*user.Vault, len(vaults))
	for i, vault := range vaults {
		vault := vault // create a new variable to avoid issues with the closure
		result[i] = &vault
	}
	return result, nil
}

// Update updates a vault's details
func (r *MongoVaultRepository) Update(ctx context.Context, vault *user.Vault) error {
	// Check if the updated name conflicts with an existing vault
	if existingVault, err := r.FindByName(ctx, vault.UserID, vault.Name); err == nil && existingVault != nil && existingVault.ID != vault.ID {
		return ErrVaultNameExists
	}

	update := bson.M{
		"$set": bson.M{
			"name":        vault.Name,
			"description": vault.Description,
			"updated_at":  vault.UpdatedAt,
		},
	}

	return r.repo.Update(ctx, vault.ID, "_id", update)
}

// Delete removes a vault from the database
func (r *MongoVaultRepository) Delete(ctx context.Context, id string) error {
	// NOTE: In a real application, you might want to check if the vault has any credentials
	// and either prevent deletion or implement a cascading delete
	return r.repo.Delete(ctx, id, "_id")
}
