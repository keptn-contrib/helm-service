// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/keptn-contrib/helm-service/pkg/types (interfaces: IProjectHandler)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/keptn/go-utils/pkg/api/models"
)

// MockIProjectHandler is a mock of IProjectHandler interface.
type MockIProjectHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIProjectHandlerMockRecorder
}

// MockIProjectHandlerMockRecorder is the mock recorder for MockIProjectHandler.
type MockIProjectHandlerMockRecorder struct {
	mock *MockIProjectHandler
}

// NewMockIProjectHandler creates a new mock instance.
func NewMockIProjectHandler(ctrl *gomock.Controller) *MockIProjectHandler {
	mock := &MockIProjectHandler{ctrl: ctrl}
	mock.recorder = &MockIProjectHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProjectHandler) EXPECT() *MockIProjectHandlerMockRecorder {
	return m.recorder
}

// CreateProject mocks base method.
func (m *MockIProjectHandler) CreateProject(arg0 models.Project) (*models.EventContext, *models.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0)
	ret0, _ := ret[0].(*models.EventContext)
	ret1, _ := ret[1].(*models.Error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockIProjectHandlerMockRecorder) CreateProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockIProjectHandler)(nil).CreateProject), arg0)
}

// DeleteProject mocks base method.
func (m *MockIProjectHandler) DeleteProject(arg0 models.Project) (*models.EventContext, *models.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", arg0)
	ret0, _ := ret[0].(*models.EventContext)
	ret1, _ := ret[1].(*models.Error)
	return ret0, ret1
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockIProjectHandlerMockRecorder) DeleteProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockIProjectHandler)(nil).DeleteProject), arg0)
}

// GetAllProjects mocks base method.
func (m *MockIProjectHandler) GetAllProjects() ([]*models.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProjects")
	ret0, _ := ret[0].([]*models.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProjects indicates an expected call of GetAllProjects.
func (mr *MockIProjectHandlerMockRecorder) GetAllProjects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProjects", reflect.TypeOf((*MockIProjectHandler)(nil).GetAllProjects))
}

// GetProject mocks base method.
func (m *MockIProjectHandler) GetProject(arg0 models.Project) (*models.Project, *models.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", arg0)
	ret0, _ := ret[0].(*models.Project)
	ret1, _ := ret[1].(*models.Error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject.
func (mr *MockIProjectHandlerMockRecorder) GetProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockIProjectHandler)(nil).GetProject), arg0)
}

// UpdateConfigurationServiceProject mocks base method.
func (m *MockIProjectHandler) UpdateConfigurationServiceProject(arg0 models.Project) (*models.EventContext, *models.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfigurationServiceProject", arg0)
	ret0, _ := ret[0].(*models.EventContext)
	ret1, _ := ret[1].(*models.Error)
	return ret0, ret1
}

// UpdateConfigurationServiceProject indicates an expected call of UpdateConfigurationServiceProject.
func (mr *MockIProjectHandlerMockRecorder) UpdateConfigurationServiceProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfigurationServiceProject", reflect.TypeOf((*MockIProjectHandler)(nil).UpdateConfigurationServiceProject), arg0)
}
