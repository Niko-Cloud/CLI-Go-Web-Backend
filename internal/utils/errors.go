package utils

import "net/http"

type APIError struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return e.Message
}

var (
	ErrBadRequest = &APIError{
		Status:  "BAD_REQUEST",
		Code:    http.StatusBadRequest,
		Message: "Bad request",
	}

	ErrUnauthorized = &APIError{
		Status:  "UNAUTHORIZED",
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	}

	ErrForbidden = &APIError{
		Status:  "FORBIDDEN",
		Code:    http.StatusForbidden,
		Message: "Forbidden",
	}

	ErrNotFound = &APIError{
		Status:  "NOT_FOUND",
		Code:    http.StatusNotFound,
		Message: "Resource not found",
	}

	ErrInternal = &APIError{
		Status:  "INTERNAL_SERVER_ERROR",
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	}
)
