// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SCRestAPIHandler is an autogenerated mock type for the SCRestAPIHandler type
type SCRestAPIHandler struct {
	mock.Mock
}

// Execute provides a mock function with given fields: response, numSharders, err
func (_m *SCRestAPIHandler) Execute(response map[string][]byte, numSharders int, err error) {
	_m.Called(response, numSharders, err)
}
