package utils

type APIError struct {
    Message    string
    StatusCode int
}

func NewAPIError(message string, statusCode int) *APIError {
    return &APIError{
        Message:    message,
        StatusCode: statusCode,
    }
}