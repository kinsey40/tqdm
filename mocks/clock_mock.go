// Code generated by MockGen. DO NOT EDIT.
// Source: render/clock.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockClock is a mock of Clock interface
type MockClock struct {
	ctrl     *gomock.Controller
	recorder *MockClockMockRecorder
}

// MockClockMockRecorder is the mock recorder for MockClock
type MockClockMockRecorder struct {
	mock *MockClock
}

// NewMockClock creates a new mock instance
func NewMockClock(ctrl *gomock.Controller) *MockClock {
	mock := &MockClock{ctrl: ctrl}
	mock.recorder = &MockClockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClock) EXPECT() *MockClockMockRecorder {
	return m.recorder
}

// Now mocks base method
func (m *MockClock) Now() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now
func (mr *MockClockMockRecorder) Now() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockClock)(nil).Now))
}

// Subtract mocks base method
func (m *MockClock) Subtract(arg0 time.Time) time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subtract", arg0)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Subtract indicates an expected call of Subtract
func (mr *MockClockMockRecorder) Subtract(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subtract", reflect.TypeOf((*MockClock)(nil).Subtract), arg0)
}

// SetStart mocks base method
func (m *MockClock) SetStart(arg0 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStart", arg0)
}

// SetStart indicates an expected call of SetStart
func (mr *MockClockMockRecorder) SetStart(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStart", reflect.TypeOf((*MockClock)(nil).SetStart), arg0)
}

// Start mocks base method
func (m *MockClock) Start() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockClockMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockClock)(nil).Start))
}

// Seconds mocks base method
func (m *MockClock) Seconds(arg0 time.Duration) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seconds", arg0)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Seconds indicates an expected call of Seconds
func (mr *MockClockMockRecorder) Seconds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seconds", reflect.TypeOf((*MockClock)(nil).Seconds), arg0)
}

// Remaining mocks base method
func (m *MockClock) Remaining(arg0 float64) time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remaining", arg0)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Remaining indicates an expected call of Remaining
func (mr *MockClockMockRecorder) Remaining(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remaining", reflect.TypeOf((*MockClock)(nil).Remaining), arg0)
}

// Format mocks base method
func (m *MockClock) Format(arg0 time.Duration) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Format", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Format indicates an expected call of Format
func (mr *MockClockMockRecorder) Format(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Format", reflect.TypeOf((*MockClock)(nil).Format), arg0)
}
