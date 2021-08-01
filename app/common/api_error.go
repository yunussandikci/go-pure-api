package common

import "net/http"

//NewBadRequestError Returns a new API Error with Bad Request specified in it.
func NewBadRequestError() error {
	return NewApiError(http.StatusBadRequest, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
}

//NewBadRequestErrorWithMessage Returns a new API Error with Bad Request specified in it with custom message.
func NewBadRequestErrorWithMessage(message string) error {
	return NewApiError(http.StatusBadRequest, http.StatusBadRequest, message)
}

//NewApiError Returns a new API Error with custom statusCode, code and message
func NewApiError(statusCode, code int, message string) error {
	return &ApiError{StatusCode: statusCode, Code: code, Message: message}
}

type ApiError struct {
	StatusCode int
	Code       int
	Message    string
}

func (e *ApiError) Error() string {
	return e.Message
}
