// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/yunussandikci/go-pure-api/app/models"

	time "time"
)

// MongoRecordsRepository is an autogenerated mock type for the MongoRecordsRepository type
type MongoRecordsRepository struct {
	mock.Mock
}

// GetRecordsWith provides a mock function with given fields: minCount, maxCount, startDate, endDate
func (_m *MongoRecordsRepository) GetRecordsWith(minCount int, maxCount int, startDate time.Time, endDate time.Time) ([]models.MongoRecord, error) {
	ret := _m.Called(minCount, maxCount, startDate, endDate)

	var r0 []models.MongoRecord
	if rf, ok := ret.Get(0).(func(int, int, time.Time, time.Time) []models.MongoRecord); ok {
		r0 = rf(minCount, maxCount, startDate, endDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.MongoRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, time.Time, time.Time) error); ok {
		r1 = rf(minCount, maxCount, startDate, endDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}