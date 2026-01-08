package service

import (
	"context"
	"identity-service/models"
)

// UserService defines the business logic interface for user operations
type UserService interface {
	GetAllUsers(ctx context.Context) ([]models.User, error)
	CreateUser(ctx context.Context, name, email string) (*models.User, error)
}
