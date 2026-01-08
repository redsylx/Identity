package handlers

import (
	"context"
	"encoding/json"
	stderrors "errors"
	"identity-service/config"
	"identity-service/models"
	"identity-service/service"
	apperrors "identity-service/errors"
	"log/slog"
	"net/http"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	service service.UserService
	config  *config.Config
	log     *slog.Logger
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(svc service.UserService, cfg *config.Config, log *slog.Logger) *UserHandler {
	return &UserHandler{
		service: svc,
		config:  cfg,
		log:     log,
	}
}

// GetAllUsers handles GET requests to retrieve all users
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Validate method
	if r.Method != http.MethodGet {
		h.handleError(w, r, apperrors.NewBadRequestError("method not allowed", nil))
		return
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(r.Context(), h.config.Timeouts.Handler)
	defer cancel()

	// Call service layer
	users, err := h.service.GetAllUsers(ctx)
	if err != nil {
		h.handleError(w, r, err)
		return
	}

	// Send response
	h.writeJSONResponse(w, http.StatusOK, users)
}

// CreateUser handles POST requests to create a new user
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Validate method
	if r.Method != http.MethodPost {
		h.handleError(w, r, apperrors.NewBadRequestError("method not allowed", nil))
		return
	}

	// Parse request body
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.WarnContext(r.Context(), "invalid request body",
			slog.String("error", err.Error()),
			slog.String("remote_addr", r.RemoteAddr),
		)
		h.handleError(w, r, apperrors.NewBadRequestError("invalid request body", err))
		return
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(r.Context(), h.config.Timeouts.Handler)
	defer cancel()

	// Call service layer
	user, err := h.service.CreateUser(ctx, req.Name, req.Email)
	if err != nil {
		h.handleError(w, r, err)
		return
	}

	// Send response
	h.writeJSONResponse(w, http.StatusCreated, user)
}

// handleError handles errors and sends appropriate HTTP responses
func (h *UserHandler) handleError(w http.ResponseWriter, r *http.Request, err error) {
	var appErr *apperrors.AppError

	// Log error with context
	if err != nil {
		h.log.ErrorContext(r.Context(), "request error",
			slog.String("error", err.Error()),
			slog.String("path", r.URL.Path),
			slog.String("method", r.Method),
		)
	}

	// Check if it's an AppError
	if stderrors.As(err, &appErr) {
		h.writeJSONResponse(w, appErr.Code, map[string]string{
			"error": appErr.Message,
		})
		return
	}

	// Check if it's a validation error
	var validationErrors service.ValidationErrors
	if stderrors.As(err, &validationErrors) {
		errors := make([]map[string]string, len(validationErrors))
		for i, verr := range validationErrors {
			errors[i] = map[string]string{
				verr.Field: verr.Message,
			}
		}
		h.writeJSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"errors": errors,
		})
		return
	}

	// Fallback for unknown errors
	h.writeJSONResponse(w, http.StatusInternalServerError, map[string]string{
		"error": "internal server error",
	})
}

// writeJSONResponse writes a JSON response with the given status code
func (h *UserHandler) writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.log.ErrorContext(context.Background(), "failed to encode response",
			slog.String("error", err.Error()),
		)
	}
}

// Ensure UserHandler implements UserHandlerInterface
var _ UserHandlerInterface = (*UserHandler)(nil)
