// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	server "github.com/yunussandikci/go-pure-api/app/server"
)

// InMemoryController is an autogenerated mock type for the InMemoryController type
type InMemoryController struct {
	mock.Mock
}

// Create provides a mock function with given fields: request, response
func (_m *InMemoryController) Create(request *server.Request, response *server.Response) {
	_m.Called(request, response)
}

// GetRecords provides a mock function with given fields: request, response
func (_m *InMemoryController) GetRecords(request *server.Request, response *server.Response) {
	_m.Called(request, response)
}