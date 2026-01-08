package repository

import (
	"context"
	"identity-service/models"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	GetAll(ctx context.Context) ([]models.User, error)
	Create(ctx context.Context, name, email string) (*models.User, error)
	EmailExists(ctx context.Context, email string) (bool, error)
}
