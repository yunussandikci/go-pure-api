package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBadRequestError(t *testing.T) {
	//given

	//when
	err := NewBadRequestError()

	//then
	assert.Error(t, err)
	assert.Equal(t, err.(*ApiError).StatusCode, 400)
	assert.Equal(t, err.(*ApiError).Code, 400)
}

func TestNewBadRequestErrorWithMessage(t *testing.T) {
	//given

	//when
	err := NewBadRequestErrorWithMessage("message")

	//then
	assert.Error(t, err)
	assert.Equal(t, err.(*ApiError).StatusCode, 400)
	assert.Equal(t, err.(*ApiError).Code, 400)
	assert.Equal(t, err.(*ApiError).Message, "message")
}

func TestNewApiError(t *testing.T) {
	//given

	//when
	err := NewApiError(100, 200, "message")

	//then
	assert.Error(t, err)
	assert.Equal(t, err.(*ApiError).Code, 200)
	assert.Equal(t, err.(*ApiError).StatusCode, 100)
	assert.Equal(t, err.(*ApiError).Message, "message")
}