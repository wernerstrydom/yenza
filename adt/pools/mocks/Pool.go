// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Pool is an autogenerated mock type for the Pool type
type Pool struct {
	mock.Mock
}

// Get provides a mock function with given fields:
func (_m *Pool) Get() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Return provides a mock function with given fields: obj
func (_m *Pool) Return(obj interface{}) {
	_m.Called(obj)
}
