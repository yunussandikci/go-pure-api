package services

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/yunussandikci/go-pure-api/app/common"
	"github.com/yunussandikci/go-pure-api/app/mocks"
	in_memory "github.com/yunussandikci/go-pure-api/app/repositories/in-memory"
	"net/http"
	"testing"
)

func TestInMemoryRecordService_Create(t *testing.T) {
	//given
	recordsRepository := new(mocks.InMemoryRecordsRepository)
	recordsService := NewInMemoryRecordsServiceWith(recordsRepository)
	recordsRepository.On("Create", "100", "200").Return(nil)

	//when
	createErr := recordsService.Create("100", "200")

	//then
	assert.NoError(t, createErr)
}

func TestInMemoryRecordService_Create_KeyAlreadyExistError(t *testing.T) {
	//given
	recordsRepository := new(mocks.InMemoryRecordsRepository)
	recordsService := NewInMemoryRecordsServiceWith(recordsRepository)
	keyExistErr := errors.New(in_memory.KeyAlreadyExistError)
	recordsRepository.On("Create", "100", "200").Return(keyExistErr)

	//when
	createErr := recordsService.Create("100", "200")

	//then
	assert.Error(t, createErr)
	assert.Equal(t, createErr.(*common.ApiError).Code, 1)
	assert.Equal(t, createErr.(*common.ApiError).StatusCode, http.StatusConflict)
	assert.Equal(t, createErr.(*common.ApiError).Message, "Record with key already exist.")
}

func TestInMemoryRecordService_Create_UnexpectedError(t *testing.T) {
	//given
	recordsRepository := new(mocks.InMemoryRecordsRepository)
	recordsService := NewInMemoryRecordsServiceWith(recordsRepository)
	unexpectedError := errors.New("connection timeout")
	recordsRepository.On("Create", "100", "200").Return(unexpectedError)

	//when
	createErr := recordsService.Create("100", "200")

	//then
	assert.Error(t, createErr)
	assert.Equal(t, createErr, unexpectedError)
}

func TestInMemoryRecordService_Get(t *testing.T) {
	//given
	recordsRepository := new(mocks.InMemoryRecordsRepository)
	recordsService := NewInMemoryRecordsServiceWith(recordsRepository)
	recordsRepository.On("Get", "100").Return("200", nil)

	//when
	value, getErr := recordsService.Get("100")

	//then
	assert.NoError(t, getErr)
	assert.Equal(t, "200", value)
}

func TestInMemoryRecordService_Get_KeyNotExistError(t *testing.T) {
	//given
	recordsRepository := new(mocks.InMemoryRecordsRepository)
	recordsService := NewInMemoryRecordsServiceWith(recordsRepository)
	keyNotExistError := errors.New(in_memory.KeyNotExistError)
	recordsRepository.On("Get", "100").Return("", keyNotExistError)

	//when
	value, getErr := recordsService.Get("100")

	//then
	assert.Error(t, getErr)
	assert.Empty(t, value)
	assert.Equal(t, getErr.(*common.ApiError).Code, 1)
	assert.Equal(t, getErr.(*common.ApiError).StatusCode, http.StatusNotFound)
	assert.Equal(t, getErr.(*common.ApiError).Message, "Record not found")
}

func TestInMemoryRecordService_Get_UnexpectedError(t *testing.T) {
	//given
	recordsRepository := new(mocks.InMemoryRecordsRepository)
	recordsService := NewInMemoryRecordsServiceWith(recordsRepository)
	unexpectedError := errors.New("connection timeout")
	recordsRepository.On("Get", "100").Return("", unexpectedError)

	//when
	value, getErr := recordsService.Get("100")

	//then
	assert.Error(t, getErr)
	assert.Empty(t, value)
	assert.Equal(t, unexpectedError, getErr)
}
