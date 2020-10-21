// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/keptn/kubernetes-utils/pkg (interfaces: ChartStorer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockChartStorer is a mock of ChartStorer interface.
type MockChartStorer struct {
	ctrl     *gomock.Controller
	recorder *MockChartStorerMockRecorder
}

// MockChartStorerMockRecorder is the mock recorder for MockChartStorer.
type MockChartStorerMockRecorder struct {
	mock *MockChartStorer
}

// NewMockChartStorer creates a new mock instance.
func NewMockChartStorer(ctrl *gomock.Controller) *MockChartStorer {
	mock := &MockChartStorer{ctrl: ctrl}
	mock.recorder = &MockChartStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChartStorer) EXPECT() *MockChartStorerMockRecorder {
	return m.recorder
}

// StoreChart mocks base method.
func (m *MockChartStorer) StoreChart(arg0, arg1, arg2, arg3 string, arg4 []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreChart", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreChart indicates an expected call of StoreChart.
func (mr *MockChartStorerMockRecorder) StoreChart(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreChart", reflect.TypeOf((*MockChartStorer)(nil).StoreChart), arg0, arg1, arg2, arg3, arg4)
}
