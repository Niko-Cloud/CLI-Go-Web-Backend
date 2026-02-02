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

func NewBadRequest(message string) *APIError {
	return &APIError{
		Code:    http.StatusBadRequest,
		Status:  "BAD_REQUEST",
		Message: message,
	}
}

func NewUnauthorized(message string) *APIError {
	return &APIError{
		Code:    http.StatusUnauthorized,
		Status:  "UNAUTHORIZED",
		Message: message,
	}
}

func NewForbidden(message string) *APIError {
	return &APIError{
		Code:    http.StatusForbidden,
		Status:  "FORBIDDEN",
		Message: message,
	}
}

func NewNotFound(message string) *APIError {
	return &APIError{
		Code:    http.StatusNotFound,
		Status:  "NOT_FOUND",
		Message: message,
	}
}

func NewInternal(message string) *APIError {
	return &APIError{
		Code:    http.StatusInternalServerError,
		Status:  "INTERNAL_SERVER_ERROR",
		Message: message,
	}
}
