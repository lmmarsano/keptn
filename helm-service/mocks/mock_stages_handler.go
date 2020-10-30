// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/keptn/keptn/helm-service/pkg/types (interfaces: IStagesHandler)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/keptn/go-utils/pkg/api/models"
)

// MockIStagesHandler is a mock of IStagesHandler interface.
type MockIStagesHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIStagesHandlerMockRecorder
}

// MockIStagesHandlerMockRecorder is the mock recorder for MockIStagesHandler.
type MockIStagesHandlerMockRecorder struct {
	mock *MockIStagesHandler
}

// NewMockIStagesHandler creates a new mock instance.
func NewMockIStagesHandler(ctrl *gomock.Controller) *MockIStagesHandler {
	mock := &MockIStagesHandler{ctrl: ctrl}
	mock.recorder = &MockIStagesHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStagesHandler) EXPECT() *MockIStagesHandlerMockRecorder {
	return m.recorder
}

// CreateStage mocks base method.
func (m *MockIStagesHandler) CreateStage(arg0, arg1 string) (*models.EventContext, *models.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStage", arg0, arg1)
	ret0, _ := ret[0].(*models.EventContext)
	ret1, _ := ret[1].(*models.Error)
	return ret0, ret1
}

// CreateStage indicates an expected call of CreateStage.
func (mr *MockIStagesHandlerMockRecorder) CreateStage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStage", reflect.TypeOf((*MockIStagesHandler)(nil).CreateStage), arg0, arg1)
}

// GetAllStages mocks base method.
func (m *MockIStagesHandler) GetAllStages(arg0 string) ([]*models.Stage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllStages", arg0)
	ret0, _ := ret[0].([]*models.Stage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllStages indicates an expected call of GetAllStages.
func (mr *MockIStagesHandlerMockRecorder) GetAllStages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllStages", reflect.TypeOf((*MockIStagesHandler)(nil).GetAllStages), arg0)
}
