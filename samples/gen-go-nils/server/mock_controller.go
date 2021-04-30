// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package server is a generated GoMock package.
package server

import (
	context "context"
	models "github.com/Clever/wag/v7/samples/gen-go-nils/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockController is a mock of Controller interface
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// NilCheck mocks base method
func (m *MockController) NilCheck(ctx context.Context, i *models.NilCheckInput) error {
	ret := m.ctrl.Call(m, "NilCheck", ctx, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// NilCheck indicates an expected call of NilCheck
func (mr *MockControllerMockRecorder) NilCheck(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NilCheck", reflect.TypeOf((*MockController)(nil).NilCheck), ctx, i)
}
