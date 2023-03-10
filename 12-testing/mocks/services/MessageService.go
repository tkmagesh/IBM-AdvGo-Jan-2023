// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MessageService is an autogenerated mock type for the MessageService type
type MessageService struct {
	mock.Mock
}

// Send provides a mock function with given fields: msg
func (_m *MessageService) Send(msg string) bool {
	ret := _m.Called(msg)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewMessageService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMessageService creates a new instance of MessageService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMessageService(t mockConstructorTestingTNewMessageService) *MessageService {
	mock := &MessageService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
