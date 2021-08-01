package v1

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yunussandikci/go-pure-api/app/common"
	"github.com/yunussandikci/go-pure-api/app/mocks"
	"github.com/yunussandikci/go-pure-api/app/models/response"
	"github.com/yunussandikci/go-pure-api/app/server"
	"testing"
	"time"
)

func TestMongoGetRecords(t *testing.T) {
	//given
	recordsService := new(mocks.MongoRecordsService)
	inMemoryController := NewMongoControllerWith(recordsService)
	req := &server.Request{
		Body: []byte("{\"startDate\": \"2011-01-20\",\"endDate\":\"2021-01-02\",\"minCount\":2000,\"maxCount\":3000}"),
	}
	resp := &server.Response{}
	startDate, _ := time.Parse(common.DefaultTimeFormat, "2011-02-20")
	records := []response.MongoRecordResponse{{
		Key:        "100",
		CreatedAt:  startDate,
		TotalCount: 101,
	}}
	recordsService.On("GetRecords", 2000, 3000, mock.AnythingOfType("time.Time"),
		mock.AnythingOfType("time.Time")).
		Return(records, nil)

	//when
	inMemoryController.GetRecords(req, resp)

	//then
	assert.NoError(t, resp.Error)
	assert.Equal(t, (resp.Body).(response.MongoRecordsResponse).Code, 0)
	assert.Equal(t, (resp.Body).(response.MongoRecordsResponse).Message, "Success")
	assert.Equal(t, (resp.Body).(response.MongoRecordsResponse).Records[0].Key, "100")
	assert.Equal(t, (resp.Body).(response.MongoRecordsResponse).Records[0].TotalCount, 101)
	assert.Equal(t, (resp.Body).(response.MongoRecordsResponse).Records[0].CreatedAt, startDate)
}

func TestMongoGetRecords_BadRequest(t *testing.T) {
	//given
	recordsService := new(mocks.MongoRecordsService)
	inMemoryController := NewMongoControllerWith(recordsService)
	req := &server.Request{
		Body: []byte("bad request"),
	}
	resp := &server.Response{}

	//when
	inMemoryController.GetRecords(req, resp)

	//then
	assert.Error(t, resp.Error)
	assert.Equal(t, resp.Error.(*common.ApiError).Code, 400)
	assert.Equal(t, resp.Error.(*common.ApiError).StatusCode, 400)
	assert.Equal(t, resp.Error.(*common.ApiError).Message, "Bad Request")
}

func TestMongoGetRecords_ValidationError(t *testing.T) {
	//given
	recordsService := new(mocks.MongoRecordsService)
	inMemoryController := NewMongoControllerWith(recordsService)
	req := &server.Request{
		Body: []byte("{\"startDate\": \"2011-01-20\",\"endDate\":\"2021-01-02\",\"minCount\":3000,\"maxCount\":0}"),
	}
	resp := &server.Response{}

	//when
	inMemoryController.GetRecords(req, resp)

	//then
	assert.Error(t, resp.Error)
	assert.Equal(t, resp.Error.(*common.ApiError).StatusCode, 400)
	assert.Equal(t, resp.Error.(*common.ApiError).Code, 400)
}

func TestMongoGetRecords_ServiceError(t *testing.T) {
	//given
	recordsService := new(mocks.MongoRecordsService)
	inMemoryController := NewMongoControllerWith(recordsService)
	req := &server.Request{
		Body: []byte("{\"startDate\": \"2011-01-20\",\"endDate\":\"2021-01-02\",\"minCount\":2000,\"maxCount\":3000}"),
	}
	resp := &server.Response{}
	startDate, _ := time.Parse(common.DefaultTimeFormat, "2011-02-20")
	records := []response.MongoRecordResponse{{
		Key:        "100",
		CreatedAt:  startDate,
		TotalCount: 101,
	}}

	serviceError := errors.New("an error occurred")
	recordsService.On("GetRecords", 2000, 3000, mock.AnythingOfType("time.Time"),
		mock.AnythingOfType("time.Time")).
		Return(records, serviceError)

	//when
	inMemoryController.GetRecords(req, resp)

	//then
	assert.Error(t, resp.Error)
	assert.Equal(t, resp.Error, serviceError)
}