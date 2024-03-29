// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/keptn-contrib/helm-service/pkg/types (interfaces: IServiceHandler)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/keptn/go-utils/pkg/api/models"
)

// MockIServiceHandler is a mock of IServiceHandler interface.
type MockIServiceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceHandlerMockRecorder
}

// MockIServiceHandlerMockRecorder is the mock recorder for MockIServiceHandler.
type MockIServiceHandlerMockRecorder struct {
	mock *MockIServiceHandler
}

// NewMockIServiceHandler creates a new mock instance.
func NewMockIServiceHandler(ctrl *gomock.Controller) *MockIServiceHandler {
	mock := &MockIServiceHandler{ctrl: ctrl}
	mock.recorder = &MockIServiceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIServiceHandler) EXPECT() *MockIServiceHandlerMockRecorder {
	return m.recorder
}

// CreateServiceInStage mocks base method.
func (m *MockIServiceHandler) CreateServiceInStage(arg0, arg1, arg2 string) (*models.EventContext, *models.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServiceInStage", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.EventContext)
	ret1, _ := ret[1].(*models.Error)
	return ret0, ret1
}

// CreateServiceInStage indicates an expected call of CreateServiceInStage.
func (mr *MockIServiceHandlerMockRecorder) CreateServiceInStage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServiceInStage", reflect.TypeOf((*MockIServiceHandler)(nil).CreateServiceInStage), arg0, arg1, arg2)
}

// DeleteServiceFromStage mocks base method.
func (m *MockIServiceHandler) DeleteServiceFromStage(arg0, arg1, arg2 string) (*models.EventContext, *models.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServiceFromStage", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.EventContext)
	ret1, _ := ret[1].(*models.Error)
	return ret0, ret1
}

// DeleteServiceFromStage indicates an expected call of DeleteServiceFromStage.
func (mr *MockIServiceHandlerMockRecorder) DeleteServiceFromStage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServiceFromStage", reflect.TypeOf((*MockIServiceHandler)(nil).DeleteServiceFromStage), arg0, arg1, arg2)
}

// GetAllServices mocks base method.
func (m *MockIServiceHandler) GetAllServices(arg0, arg1 string) ([]*models.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllServices", arg0, arg1)
	ret0, _ := ret[0].([]*models.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllServices indicates an expected call of GetAllServices.
func (mr *MockIServiceHandlerMockRecorder) GetAllServices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllServices", reflect.TypeOf((*MockIServiceHandler)(nil).GetAllServices), arg0, arg1)
}

// GetService mocks base method.
func (m *MockIServiceHandler) GetService(arg0, arg1, arg2 string) (*models.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockIServiceHandlerMockRecorder) GetService(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockIServiceHandler)(nil).GetService), arg0, arg1, arg2)
}
