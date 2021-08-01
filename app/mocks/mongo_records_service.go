// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	response "github.com/yunussandikci/go-pure-api/app/models/response"

	time "time"
)

// MongoRecordsService is an autogenerated mock type for the MongoRecordsService type
type MongoRecordsService struct {
	mock.Mock
}

// GetRecords provides a mock function with given fields: minCount, maxCount, startDate, endDate
func (_m *MongoRecordsService) GetRecords(minCount int, maxCount int, startDate time.Time, endDate time.Time) ([]response.MongoRecordResponse, error) {
	ret := _m.Called(minCount, maxCount, startDate, endDate)

	var r0 []response.MongoRecordResponse
	if rf, ok := ret.Get(0).(func(int, int, time.Time, time.Time) []response.MongoRecordResponse); ok {
		r0 = rf(minCount, maxCount, startDate, endDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]response.MongoRecordResponse)
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
