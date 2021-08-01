package v1

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/yunussandikci/go-pure-api/app/common"
	"github.com/yunussandikci/go-pure-api/app/mocks"
	"github.com/yunussandikci/go-pure-api/app/models/response"
	"github.com/yunussandikci/go-pure-api/app/server"
	"testing"
)

func TestCreate(t *testing.T) {
	//given
	recordsService := new(mocks.InMemoryRecordsService)
	inMemoryController := NewInMemoryControllerWith(recordsService)
	req := &server.Request{
		Body: []byte("{\"key\":\"100\",\"value\":\"101\"\n}"),
	}
	resp := &server.Response{}
	recordsService.On("Create", "100", "101").Return(nil)

	//when
	inMemoryController.Create(req, resp)

	//then
	assert.NoError(t, resp.Error)
	assert.Equal(t, (resp.Body).(response.InMemoryRecordResponse).Key, "100")
	assert.Equal(t, (resp.Body).(response.InMemoryRecordResponse).Value, "101")
}

func TestCreate_BadRequest(t *testing.T) {
	//given
	recordsService := new(mocks.InMemoryRecordsService)
	inMemoryController := NewInMemoryControllerWith(recordsService)
	req := &server.Request{
		Body: []byte("invalid json"),
	}
	resp := &server.Response{}
	recordsService.On("Create", "100", "101").Return(nil)

	//when
	inMemoryController.Create(req, resp)

	//then
	assert.Error(t, resp.Error)
	assert.Equal(t, resp.Error.(*common.ApiError).Code, 400)
	assert.Equal(t, resp.Error.(*common.ApiError).StatusCode, 400)
	assert.Equal(t, resp.Error.(*common.ApiError).Message, "Bad Request")
}

func TestCreate_ValidationError(t *testing.T) {
	//given
	recordsService := new(mocks.InMemoryRecordsService)
	inMemoryController := NewInMemoryControllerWith(recordsService)
	req := &server.Request{
		Body: []byte("{\"key\":\"100\"}"),
	}
	resp := &server.Response{}

	//when
	inMemoryController.Create(req, resp)

	//then
	assert.Error(t, resp.Error)
	assert.Equal(t, resp.Error.(*common.ApiError).StatusCode, 400)
	assert.Equal(t, resp.Error.(*common.ApiError).Code, 400)
}

func TestCreate_ServiceError(t *testing.T) {
	//given
	recordsService := new(mocks.InMemoryRecordsService)
	inMemoryController := NewInMemoryControllerWith(recordsService)
	req := &server.Request{
		Body: []byte("{\"key\":\"100\",\"value\":\"101\"\n}"),
	}
	resp := &server.Response{}
	serviceError := errors.New("error occurred")
	recordsService.On("Create", "100", "101").Return(serviceError)

	//when
	inMemoryController.Create(req, resp)

	//then
	assert.Error(t, resp.Error)
	assert.Equal(t, resp.Error, serviceError)
}

func TestGetRecords(t *testing.T) {
	//given
	recordsService := new(mocks.InMemoryRecordsService)
	inMemoryController := NewInMemoryControllerWith(recordsService)
	req := &server.Request{
		Parameters: map[string][]string{"key": {"100"}},
	}
	resp := &server.Response{}
	recordsService.On("Get", "100").Return("101", nil)

	//when
	inMemoryController.GetRecords(req, resp)

	//then
	assert.NoError(t, resp.Error)
	assert.Equal(t, (resp.Body).(response.InMemoryRecordResponse).Key, "100")
	assert.Equal(t, (resp.Body).(response.InMemoryRecordResponse).Value, "101")
}

func TestGetRecords_BadRequest(t *testing.T) {
	//given
	recordsService := new(mocks.InMemoryRecordsService)
	inMemoryController := NewInMemoryControllerWith(recordsService)
	req := &server.Request{}
	resp := &server.Response{}
	recordsService.On("Get", "100").Return("101", nil)

	//when
	inMemoryController.GetRecords(req, resp)

	//then
	//then
	assert.Error(t, resp.Error)
	assert.Equal(t, resp.Error.(*common.ApiError).Code, 400)
	assert.Equal(t, resp.Error.(*common.ApiError).StatusCode, 400)
	assert.Equal(t, resp.Error.(*common.ApiError).Message, "Bad Request")
}

func TestGetRecords_ValidationError(t *testing.T) {
	//given
	recordsService := new(mocks.InMemoryRecordsService)
	inMemoryController := NewInMemoryControllerWith(recordsService)
	req := &server.Request{
		Parameters: map[string][]string{"key": {""}},
	}
	resp := &server.Response{}
	recordsService.On("Get", "100").Return("101", nil)

	//when
	inMemoryController.GetRecords(req, resp)

	//then
	assert.Error(t, resp.Error)
	assert.Equal(t, resp.Error.(*common.ApiError).StatusCode, 400)
	assert.Equal(t, resp.Error.(*common.ApiError).Code, 400)
}

func TestGetRecords_ServiceError(t *testing.T) {
	//given
	recordsService := new(mocks.InMemoryRecordsService)
	inMemoryController := NewInMemoryControllerWith(recordsService)
	req := &server.Request{
		Parameters: map[string][]string{"key": {"100"}},
	}
	resp := &server.Response{}
	serviceError := errors.New("error occurred")
	recordsService.On("Get", "100").Return("", serviceError)

	//when
	inMemoryController.GetRecords(req, resp)

	//then
	assert.Error(t, resp.Error)
	assert.Equal(t, resp.Error, serviceError)
}
