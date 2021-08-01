package services

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/yunussandikci/go-pure-api/app/common"
	"github.com/yunussandikci/go-pure-api/app/mocks"
	"github.com/yunussandikci/go-pure-api/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestMongoRecordsService_GetRecords(t *testing.T) {
	//given
	recordsRepository := new(mocks.MongoRecordsRepository)
	recordsService := NewMongoRecordsServiceWith(recordsRepository)
	startDate, _ := time.Parse(common.DefaultTimeFormat, "2021-01-01")
	endDate, _ := time.Parse(common.DefaultTimeFormat, "2021-01-02")
	recordsRepository.On("GetRecordsWith", 100, 200, startDate, endDate).
		Return([]models.MongoRecord{{
			Id:         primitive.ObjectID{100},
			Key:        "200",
			Value:      "300",
			CreatedAt:  startDate,
			TotalCount: 400,
		}}, nil)

	//when
	records, getRecordsErr := recordsService.GetRecords(100, 200, startDate, endDate)

	//then
	assert.NoError(t, getRecordsErr)
	assert.Equal(t, len(records), 1)
	assert.Equal(t, records[0].Key, "200")
	assert.Equal(t, records[0].CreatedAt, startDate)
	assert.Equal(t, records[0].TotalCount, 400)
}


func TestMongoRecordsService_RepositoryError(t *testing.T) {
	//given
	recordsRepository := new(mocks.MongoRecordsRepository)
	recordsService := NewMongoRecordsServiceWith(recordsRepository)
	startDate, _ := time.Parse(common.DefaultTimeFormat, "2021-01-01")
	endDate, _ := time.Parse(common.DefaultTimeFormat, "2021-01-02")
	unexpectedError := errors.New("connection timeout")
	recordsRepository.On("GetRecordsWith", 100, 200, startDate, endDate).
		Return(nil, unexpectedError)

	//when
	records, getRecordsErr := recordsService.GetRecords(100, 200, startDate, endDate)

	//then
	assert.Empty(t, records)
	assert.Error(t, getRecordsErr)
	assert.Equal(t, getRecordsErr, unexpectedError)
}
