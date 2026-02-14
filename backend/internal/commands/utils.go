package commandUtils

import (
	"Groundwork/backend/internal/database"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ErrorResponse(err error) *Response {
	return &Response{
		Status:  MapErrorToStatus(err),
		Message: err.Error(),
		Data:    nil,
	}
}

func SuccessResponse(message string, status int, data any) *Response {
	return &Response{
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
