package repository

import (
	"context"
	"database/sql"
	"identity-service/models"
	"strings"
)

type userRepository struct {
	DB *sql.DB
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) GetAll(ctx context.Context) ([]models.User, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, name, email FROM users ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) Create(ctx context.Context, name, email string) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRowContext(
		ctx,
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email",
		name, email,
	).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := r.DB.QueryRowContext(
		ctx,
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)",
		strings.ToLower(email),
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}
