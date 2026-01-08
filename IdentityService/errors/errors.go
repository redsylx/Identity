package errors

import (
	"fmt"
)

// AppError represents an application error with code and message
type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// Common error constructors
func NewBadRequestError(message string, err error) *AppError {
	return &AppError{
		Code:    400,
		Message: message,
		Err:     err,
	}
}

func NewInternalServerError(message string, err error) *AppError {
	return &AppError{
		Code:    500,
		Message: message,
		Err:     err,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    404,
		Message: message,
	}
}

func NewConflictError(message string, err error) *AppError {
	return &AppError{
		Code:    409,
		Message: message,
		Err:     err,
	}
}
