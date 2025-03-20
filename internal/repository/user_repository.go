package repository

import (
	"context"
	"database/sql"
	"errors"
	"key-haven-back/internal/model"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyUsed   = errors.New("email is already in use")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	UpdatePassword(ctx context.Context, userID, hashedPassword string) error
}

type SQLUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &SQLUserRepository{db: db}
}

func (r *SQLUserRepository) Create(ctx context.Context, user *model.User) error {
	existingUser, err := r.FindByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return ErrEmailAlreadyUsed
	}

	query := `
		INSERT INTO users (id, email, password, name, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err = r.db.ExecContext(ctx, query,
		user.ID, user.Email, user.Password, user.Name, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *SQLUserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	query := `
		SELECT id, email, password, name, created_at, updated_at
		FROM users
		WHERE id = ?
	`

	user := &model.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (r *SQLUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `
		SELECT id, email, password, name, created_at, updated_at
		FROM users
		WHERE email = ?
	`

	user := &model.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (r *SQLUserRepository) UpdatePassword(ctx context.Context, userID, hashedPassword string) error {
	query := `UPDATE users SET password = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query, hashedPassword, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}
