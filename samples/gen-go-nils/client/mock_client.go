// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package client is a generated GoMock package.
package client

import (
	context "context"
	models "github.com/Clever/wag/v6/samples/gen-go-nils/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// NilCheck mocks base method.
func (m *MockClient) NilCheck(ctx context.Context, i *models.NilCheckInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NilCheck", ctx, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// NilCheck indicates an expected call of NilCheck.
func (mr *MockClientMockRecorder) NilCheck(ctx, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NilCheck", reflect.TypeOf((*MockClient)(nil).NilCheck), ctx, i)
}
