package errs

import "net/http"

type AppError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

func (e *AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnprocessableEntityError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
