// Code generated by MockGen. DO NOT EDIT.
// Source: tokens.go

// Package mock_tokens is a generated GoMock package.
package mock_tokens

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTokenManagerI is a mock of TokenManagerI interface.
type MockTokenManagerI struct {
	ctrl     *gomock.Controller
	recorder *MockTokenManagerIMockRecorder
}

// MockTokenManagerIMockRecorder is the mock recorder for MockTokenManagerI.
type MockTokenManagerIMockRecorder struct {
	mock *MockTokenManagerI
}

// NewMockTokenManagerI creates a new mock instance.
func NewMockTokenManagerI(ctrl *gomock.Controller) *MockTokenManagerI {
	mock := &MockTokenManagerI{ctrl: ctrl}
	mock.recorder = &MockTokenManagerIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenManagerI) EXPECT() *MockTokenManagerIMockRecorder {
	return m.recorder
}

// CreateAccessToken mocks base method.
func (m *MockTokenManagerI) CreateAccessToken(userId int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessToken", userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessToken indicates an expected call of CreateAccessToken.
func (mr *MockTokenManagerIMockRecorder) CreateAccessToken(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessToken", reflect.TypeOf((*MockTokenManagerI)(nil).CreateAccessToken), userId)
}

// CreateEmailToken mocks base method.
func (m *MockTokenManagerI) CreateEmailToken(email string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmailToken", email)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEmailToken indicates an expected call of CreateEmailToken.
func (mr *MockTokenManagerIMockRecorder) CreateEmailToken(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmailToken", reflect.TypeOf((*MockTokenManagerI)(nil).CreateEmailToken), email)
}

// CreateRefreshToken mocks base method.
func (m *MockTokenManagerI) CreateRefreshToken() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRefreshToken")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRefreshToken indicates an expected call of CreateRefreshToken.
func (mr *MockTokenManagerIMockRecorder) CreateRefreshToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRefreshToken", reflect.TypeOf((*MockTokenManagerI)(nil).CreateRefreshToken))
}

// ParseAccessToken mocks base method.
func (m *MockTokenManagerI) ParseAccessToken(tokenString string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseAccessToken", tokenString)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseAccessToken indicates an expected call of ParseAccessToken.
func (mr *MockTokenManagerIMockRecorder) ParseAccessToken(tokenString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseAccessToken", reflect.TypeOf((*MockTokenManagerI)(nil).ParseAccessToken), tokenString)
}

// ParseEmailToken mocks base method.
func (m *MockTokenManagerI) ParseEmailToken(tokenString string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseEmailToken", tokenString)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseEmailToken indicates an expected call of ParseEmailToken.
func (mr *MockTokenManagerIMockRecorder) ParseEmailToken(tokenString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseEmailToken", reflect.TypeOf((*MockTokenManagerI)(nil).ParseEmailToken), tokenString)
}