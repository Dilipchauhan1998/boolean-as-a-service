// Code generated by MockGen. DO NOT EDIT.
// Source: ./auth/token_auth.go

// Package mocks is a generated GoMock package.
package mocks

import (
	auth "boolean-as-a-service/auth"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MocktokenRepoInterface is a mock of tokenRepoInterface interface
type MocktokenRepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MocktokenRepoInterfaceMockRecorder
}

// MocktokenRepoInterfaceMockRecorder is the mock recorder for MocktokenRepoInterface
type MocktokenRepoInterfaceMockRecorder struct {
	mock *MocktokenRepoInterface
}

// NewMocktokenRepoInterface creates a new mock instance
func NewMocktokenRepoInterface(ctrl *gomock.Controller) *MocktokenRepoInterface {
	mock := &MocktokenRepoInterface{ctrl: ctrl}
	mock.recorder = &MocktokenRepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MocktokenRepoInterface) EXPECT() *MocktokenRepoInterfaceMockRecorder {
	return m.recorder
}

// CreateToken mocks base method
func (m *MocktokenRepoInterface) CreateToken(token auth.Token) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateToken", token)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateToken indicates an expected call of CreateToken
func (mr *MocktokenRepoInterfaceMockRecorder) CreateToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MocktokenRepoInterface)(nil).CreateToken), token)
}

// ExistToken mocks base method
func (m *MocktokenRepoInterface) ExistToken(id string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistToken", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ExistToken indicates an expected call of ExistToken
func (mr *MocktokenRepoInterfaceMockRecorder) ExistToken(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistToken", reflect.TypeOf((*MocktokenRepoInterface)(nil).ExistToken), id)
}
