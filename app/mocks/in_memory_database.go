// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// InMemoryDatabase is an autogenerated mock type for the InMemoryDatabase type
type InMemoryDatabase struct {
	mock.Mock
}

// CreateSync provides a mock function with given fields: key, value
func (_m *InMemoryDatabase) CreateSync(key string, value string) {
	_m.Called(key, value)
}

// GetSync provides a mock function with given fields: key
func (_m *InMemoryDatabase) GetSync(key string) string {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
