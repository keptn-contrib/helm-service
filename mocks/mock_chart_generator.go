// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/keptn-contrib/helm-service/pkg/helm (interfaces: ChartGenerator)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	chart "helm.sh/helm/v3/pkg/chart"
)

// MockChartGenerator is a mock of ChartGenerator interface.
type MockChartGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockChartGeneratorMockRecorder
}

// MockChartGeneratorMockRecorder is the mock recorder for MockChartGenerator.
type MockChartGeneratorMockRecorder struct {
	mock *MockChartGenerator
}

// NewMockChartGenerator creates a new mock instance.
func NewMockChartGenerator(ctrl *gomock.Controller) *MockChartGenerator {
	mock := &MockChartGenerator{ctrl: ctrl}
	mock.recorder = &MockChartGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChartGenerator) EXPECT() *MockChartGeneratorMockRecorder {
	return m.recorder
}

// GenerateDuplicateChart mocks base method.
func (m *MockChartGenerator) GenerateDuplicateChart(arg0, arg1, arg2, arg3 string) (*chart.Chart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateDuplicateChart", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*chart.Chart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateDuplicateChart indicates an expected call of GenerateDuplicateChart.
func (mr *MockChartGeneratorMockRecorder) GenerateDuplicateChart(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateDuplicateChart", reflect.TypeOf((*MockChartGenerator)(nil).GenerateDuplicateChart), arg0, arg1, arg2, arg3)
}

// GenerateMeshChart mocks base method.
func (m *MockChartGenerator) GenerateMeshChart(arg0, arg1, arg2, arg3 string) (*chart.Chart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateMeshChart", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*chart.Chart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateMeshChart indicates an expected call of GenerateMeshChart.
func (mr *MockChartGeneratorMockRecorder) GenerateMeshChart(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateMeshChart", reflect.TypeOf((*MockChartGenerator)(nil).GenerateMeshChart), arg0, arg1, arg2, arg3)
}
