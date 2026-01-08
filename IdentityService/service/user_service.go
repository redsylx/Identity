package service

import (
	"context"

	"identity-service/models"
	"identity-service/repository"
	apperrors "identity-service/errors"
	"identity-service/validation"
)

// ValidationErrors is exported for handlers to use
type ValidationErrors = validation.ValidationErrors

// userService implements the UserService interface
type userService struct {
	repo      repository.UserRepository
	validator *validation.Validator
}

// NewUserService creates a new UserService instance
func NewUserService(repo repository.UserRepository, validator *validation.Validator) UserService {
	return &userService{
		repo:      repo,
		validator: validator,
	}
}

func (s *userService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternalServerError("failed to retrieve users", err)
	}
	return users, nil
}

func (s *userService) CreateUser(ctx context.Context, name, email string) (*models.User, error) {
	// Validate input
	if err := s.validator.ValidateCreateUserRequest(name, email); err != nil {
		return nil, apperrors.NewBadRequestError("validation failed", err)
	}

	// Check if email already exists (business logic)
	existingUser, err := s.repo.EmailExists(ctx, email)
	if err != nil {
		return nil, apperrors.NewInternalServerError("failed to check email existence", err)
	}
	if existingUser {
		return nil, apperrors.NewConflictError("user with this email already exists", nil)
	}

	// Create user
	user, err := s.repo.Create(ctx, name, email)
	if err != nil {
		return nil, apperrors.NewInternalServerError("failed to create user", err)
	}

	return user, nil
}

// Ensure userService implements UserService interface
var _ UserService = (*userService)(nil)
