// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/keptn-contrib/helm-service/pkg/namespacemanager (interfaces: INamespaceManager)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockINamespaceManager is a mock of INamespaceManager interface.
type MockINamespaceManager struct {
	ctrl     *gomock.Controller
	recorder *MockINamespaceManagerMockRecorder
}

// MockINamespaceManagerMockRecorder is the mock recorder for MockINamespaceManager.
type MockINamespaceManagerMockRecorder struct {
	mock *MockINamespaceManager
}

// NewMockINamespaceManager creates a new mock instance.
func NewMockINamespaceManager(ctrl *gomock.Controller) *MockINamespaceManager {
	mock := &MockINamespaceManager{ctrl: ctrl}
	mock.recorder = &MockINamespaceManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockINamespaceManager) EXPECT() *MockINamespaceManagerMockRecorder {
	return m.recorder
}

// CreateNamespaceIfNotExists mocks base method.
func (m *MockINamespaceManager) CreateNamespaceIfNotExists(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNamespaceIfNotExists", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNamespaceIfNotExists indicates an expected call of CreateNamespaceIfNotExists.
func (mr *MockINamespaceManagerMockRecorder) CreateNamespaceIfNotExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNamespaceIfNotExists", reflect.TypeOf((*MockINamespaceManager)(nil).CreateNamespaceIfNotExists), arg0)
}

// InjectIstio mocks base method.
func (m *MockINamespaceManager) InjectIstio(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InjectIstio", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InjectIstio indicates an expected call of InjectIstio.
func (mr *MockINamespaceManagerMockRecorder) InjectIstio(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectIstio", reflect.TypeOf((*MockINamespaceManager)(nil).InjectIstio), arg0, arg1)
}
