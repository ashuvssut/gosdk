// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	zcncore "github.com/0chain/gosdk/zcncore"
	mock "github.com/stretchr/testify/mock"
)

// QueryResultHandle is an autogenerated mock type for the QueryResultHandle type
type QueryResultHandle struct {
	mock.Mock
}

// Execute provides a mock function with given fields: result
func (_m *QueryResultHandle) Execute(result zcncore.QueryResult) bool {
	ret := _m.Called(result)

	var r0 bool
	if rf, ok := ret.Get(0).(func(zcncore.QueryResult) bool); ok {
		r0 = rf(result)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewQueryResultHandle creates a new instance of QueryResultHandle. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQueryResultHandle(t interface {
	mock.TestingT
	Cleanup(func())
}) *QueryResultHandle {
	mock := &QueryResultHandle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
