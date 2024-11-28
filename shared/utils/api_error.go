package utils

type APIError struct {
    Message    any
    StatusCode int
}

func NewAPIError(message any, statusCode int) *APIError {
    return &APIError{
        Message:    message,
        StatusCode: statusCode,
    }
}