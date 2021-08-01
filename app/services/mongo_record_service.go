package services

import (
	"github.com/yunussandikci/go-pure-api/app/database"
	"github.com/yunussandikci/go-pure-api/app/models/response"
	"github.com/yunussandikci/go-pure-api/app/repositories/mongo"
	"time"
)

type MongoRecordsService interface {
	GetRecords(minCount, maxCount int, startDate, endDate time.Time) ([]response.MongoRecordResponse, error)
}

type mongoRecordsService struct {
	recordsRepository mongo.MongoRecordsRepository
}

func NewMongoRecordsService(database database.MongoDatabase) MongoRecordsService {
	return &mongoRecordsService{
		recordsRepository: mongo.NewMongoRecordsRepository(database),
	}
}

func NewMongoRecordsServiceWith(recordsRepository mongo.MongoRecordsRepository) MongoRecordsService {
	return &mongoRecordsService{
		recordsRepository: recordsRepository,
	}
}


func (m *mongoRecordsService) GetRecords(minCount, maxCount int, startDate, endDate time.Time) ([]response.MongoRecordResponse, error) {
	records, recordsErr := m.recordsRepository.GetRecordsWith(minCount, maxCount, startDate, endDate)
	if recordsErr != nil {
		return nil, recordsErr
	}
	recordResponse := []response.MongoRecordResponse{}
	for _, v := range records {
		recordResponse = append(recordResponse, response.NewMongoRecordResponse(v.Key, v.CreatedAt, v.TotalCount))
	}
	return recordResponse, nil
}
