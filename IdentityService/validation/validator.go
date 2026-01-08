package validation

import (
	"fmt"
	"regexp"

	"identity-service/config"
)

// Validator provides validation functionality
type Validator struct {
	config *config.ValidationConfig
}

// NewValidator creates a new Validator instance
func NewValidator(cfg *config.ValidationConfig) *Validator {
	return &Validator{
		config: cfg,
	}
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationErrors represents multiple validation errors
type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	var msg string
	for _, err := range e {
		if msg != "" {
			msg += "; "
		}
		msg += err.Error()
	}
	return msg
}

// ValidateEmail validates an email address
func (v *Validator) ValidateEmail(email string) error {
	if email == "" {
		return ValidationError{Field: "email", Message: "is required"}
	}

	if len(email) > v.config.MaxEmailLength {
		return ValidationError{
			Field:   "email",
			Message: fmt.Sprintf("must be at most %d characters", v.config.MaxEmailLength),
		}
	}

	matched, err := regexp.MatchString(v.config.EmailRegex, email)
	if err != nil {
		return ValidationError{Field: "email", Message: "invalid email format"}
	}
	if !matched {
		return ValidationError{Field: "email", Message: "invalid email format"}
	}

	return nil
}

// ValidateName validates a name
func (v *Validator) ValidateName(name string) error {
	if name == "" {
		return ValidationError{Field: "name", Message: "is required"}
	}

	if len(name) > v.config.MaxNameLength {
		return ValidationError{
			Field:   "name",
			Message: fmt.Sprintf("must be at most %d characters", v.config.MaxNameLength),
		}
	}

	return nil
}

// ValidateCreateUserRequest validates a create user request
func (v *Validator) ValidateCreateUserRequest(name, email string) error {
	var errors ValidationErrors

	if err := v.ValidateName(name); err != nil {
		if validationErr, ok := err.(ValidationError); ok {
			errors = append(errors, validationErr)
		}
	}

	if err := v.ValidateEmail(email); err != nil {
		if validationErr, ok := err.(ValidationError); ok {
			errors = append(errors, validationErr)
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
