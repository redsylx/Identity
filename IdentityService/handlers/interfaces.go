package handlers

import (
	"net/http"
)

// Handler defines the common interface for all HTTP handlers
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// UserHandlerInterface defines the interface for user HTTP handlers
type UserHandlerInterface interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
}
