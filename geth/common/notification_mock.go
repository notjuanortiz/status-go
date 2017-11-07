// Code generated by MockGen. DO NOT EDIT.
// Source: geth/common/notification.go

// Package common is a generated GoMock package.
package common

import (
	go_fcm "github.com/NaySoftware/go-fcm"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNotifier is a mock of Notifier interface
type MockNotifier struct {
	ctrl     *gomock.Controller
	recorder *MockNotifierMockRecorder
}

// MockNotifierMockRecorder is the mock recorder for MockNotifier
type MockNotifierMockRecorder struct {
	mock *MockNotifier
}

// NewMockNotifier creates a new mock instance
func NewMockNotifier(ctrl *gomock.Controller) *MockNotifier {
	mock := &MockNotifier{ctrl: ctrl}
	mock.recorder = &MockNotifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNotifier) EXPECT() *MockNotifierMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockNotifier) Send(body string, payload go_fcm.NotificationPayload, tokens ...string) error {
	varargs := []interface{}{body, payload}
	for _, a := range tokens {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Send", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockNotifierMockRecorder) Send(body, payload interface{}, tokens ...interface{}) *gomock.Call {
	varargs := append([]interface{}{body, payload}, tokens...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNotifier)(nil).Send), varargs...)
}