package errors

import (
	"errors"
	"fmt"
	"net/http"
)

// ErrorType holds a type string and integer code for the error
type ErrorType string

// "Set" of valid errorTypes
const (
	Authorization ErrorType = "AUTHORIZATION" // Authentication Failures -
	BadRequest    ErrorType = "BADREQUEST"    // Validation errors / BadInput
	Conflict      ErrorType = "CONFLICT"      // Already exists (eg, create account with existent email) - 409
	Internal      ErrorType = "INTERNAL"      // Server (500) and fallback errors
	NotFound      ErrorType = "NOTFOUND"      // For not finding resource
)

// Error holds a custom error for the application
type Error struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
}

// Error satisfies standard error interface
func (e *Error) Error() string {
	return e.Message
}

// Status is a mapping errors to status codes
func (e *Error) Status() int {
	switch e.Type {
	case Authorization:
		return http.StatusUnauthorized
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

// Status checks the runtime type
// of the error and returns a http
// status code if the error is model.Error
func Status(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Status()
	}
	return http.StatusInternalServerError
}

// NewAuthorization to create a 401
func NewAuthorization(reason string) *Error {
	return &Error{
		Type:    Authorization,
		Message: reason,
	}
}

// NewBadRequest to create 400 errors (validation, for example)
func NewBadRequest(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad request. Reason: %v", reason),
	}
}

// NewConflict to create an error for 409
func NewConflict(name string, value string) *Error {
	return &Error{
		Type:    Conflict,
		Message: fmt.Sprintf("resource: %v with value: %v already exists", name, value),
	}
}

// NewInternal for 500 errors and unknown errors
func NewInternal() *Error {
	return &Error{
		Type:    Internal,
		Message: fmt.Sprintf("Internal server error."),
	}
}

// NewNotFound to create an error for 404
func NewNotFound(name string, value string) *Error {
	return &Error{
		Type:    NotFound,
		Message: fmt.Sprintf("resource: %v with value: %v not found", name, value),
	}
}
