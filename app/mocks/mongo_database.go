// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

// MongoDatabase is an autogenerated mock type for the MongoDatabase type
type MongoDatabase struct {
	mock.Mock
}

// Disconnect provides a mock function with given fields:
func (_m *MongoDatabase) Disconnect() {
	_m.Called()
}

// GetCollection provides a mock function with given fields: collection
func (_m *MongoDatabase) GetCollection(collection string) *mongo.Collection {
	ret := _m.Called(collection)

	var r0 *mongo.Collection
	if rf, ok := ret.Get(0).(func(string) *mongo.Collection); ok {
		r0 = rf(collection)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Collection)
		}
	}

	return r0
}
