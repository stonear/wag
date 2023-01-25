// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package server is a generated GoMock package.
package server

import (
	context "context"
	reflect "reflect"

	models "github.com/Clever/wag/samples/gen-go-basic/models/v9"
	gomock "github.com/golang/mock/gomock"
)

// MockController is a mock of Controller interface.
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController.
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance.
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockController) CreateBook(ctx context.Context, i *models.Book) (*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", ctx, i)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockControllerMockRecorder) CreateBook(ctx, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockController)(nil).CreateBook), ctx, i)
}

// GetAuthors mocks base method.
func (m *MockController) GetAuthors(ctx context.Context, i *models.GetAuthorsInput) (*models.AuthorsResponse, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthors", ctx, i)
	ret0, _ := ret[0].(*models.AuthorsResponse)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAuthors indicates an expected call of GetAuthors.
func (mr *MockControllerMockRecorder) GetAuthors(ctx, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthors", reflect.TypeOf((*MockController)(nil).GetAuthors), ctx, i)
}

// GetAuthorsWithPut mocks base method.
func (m *MockController) GetAuthorsWithPut(ctx context.Context, i *models.GetAuthorsWithPutInput) (*models.AuthorsResponse, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorsWithPut", ctx, i)
	ret0, _ := ret[0].(*models.AuthorsResponse)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAuthorsWithPut indicates an expected call of GetAuthorsWithPut.
func (mr *MockControllerMockRecorder) GetAuthorsWithPut(ctx, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorsWithPut", reflect.TypeOf((*MockController)(nil).GetAuthorsWithPut), ctx, i)
}

// GetBookByID mocks base method.
func (m *MockController) GetBookByID(ctx context.Context, i *models.GetBookByIDInput) (*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookByID", ctx, i)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByID indicates an expected call of GetBookByID.
func (mr *MockControllerMockRecorder) GetBookByID(ctx, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByID", reflect.TypeOf((*MockController)(nil).GetBookByID), ctx, i)
}

// GetBookByID2 mocks base method.
func (m *MockController) GetBookByID2(ctx context.Context, id string) (*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookByID2", ctx, id)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByID2 indicates an expected call of GetBookByID2.
func (mr *MockControllerMockRecorder) GetBookByID2(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByID2", reflect.TypeOf((*MockController)(nil).GetBookByID2), ctx, id)
}

// GetBooks mocks base method.
func (m *MockController) GetBooks(ctx context.Context, i *models.GetBooksInput) ([]models.Book, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooks", ctx, i)
	ret0, _ := ret[0].([]models.Book)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBooks indicates an expected call of GetBooks.
func (mr *MockControllerMockRecorder) GetBooks(ctx, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooks", reflect.TypeOf((*MockController)(nil).GetBooks), ctx, i)
}

// HealthCheck mocks base method.
func (m *MockController) HealthCheck(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockControllerMockRecorder) HealthCheck(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockController)(nil).HealthCheck), ctx)
}

// PutBook mocks base method.
func (m *MockController) PutBook(ctx context.Context, i *models.Book) (*models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutBook", ctx, i)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutBook indicates an expected call of PutBook.
func (mr *MockControllerMockRecorder) PutBook(ctx, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutBook", reflect.TypeOf((*MockController)(nil).PutBook), ctx, i)
}