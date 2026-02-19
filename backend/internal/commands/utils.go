package commandUtils

import (
	"Groundwork/backend/internal/database"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// Existing ErrorResponse: Good for commands returning *Response[any] (e.g., Delete)
func ErrorResponse(err error) *Response[any] {
	return &Response[any]{
		Status:  MapErrorToStatus(err),
		Message: err.Error(),
		Data:    nil,
	}
}

// ErrorResponseTyped: Essential for commands returning specific types
func ErrorResponseTyped[T any](err error) *Response[T] {
	var zero T // This will be nil for pointers and slices
	return &Response[T]{
		Status:  MapErrorToStatus(err),
		Message: err.Error(),
		Data:    zero,
	}
}

// SuccessResponse takes a type T and returns a Response of that type
func SuccessResponse[T any](message string, status int, data T) *Response[T] {
	return &Response[T]{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func MapErrorToStatus(err error) int {
	switch {
	// Validation Errors (400)
	case errors.As(err, &validator.ValidationErrors{}):
		return http.StatusBadRequest

	// Conflict Errors (409)
	case errors.Is(err, database.ErrDuplicateCompany):
		return http.StatusConflict

	// Not Found Errors (404)
	case errors.Is(err, database.ErrNotFound):
		return http.StatusNotFound

	default:
		return http.StatusInternalServerError
	}
}
