package response

import "net/http"

type ApiError struct {
	StatusCode int
	Message    string
}

func (e ApiError) Error() string {
	return e.Message
}

func NewApiError(statusCode int, message string) ApiError {
	return ApiError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func InvalidRequestError(message string) ApiError {
	return NewApiError(http.StatusBadRequest, message)
}

func InternalServerError(message string) ApiError {
	return NewApiError(http.StatusInternalServerError, message)
}

func NotFoundError(message string) ApiError {
	return NewApiError(http.StatusNotFound, message)
}
