// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package client is a generated GoMock package.
package client

import (
	context "context"
	models "github.com/Clever/wag/v7/samples/gen-go/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetAuthors mocks base method
func (m *MockClient) GetAuthors(ctx context.Context, i *models.GetAuthorsInput) (*models.AuthorsResponse, error) {
	ret := m.ctrl.Call(m, "GetAuthors", ctx, i)
	ret0, _ := ret[0].(*models.AuthorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthors indicates an expected call of GetAuthors
func (mr *MockClientMockRecorder) GetAuthors(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthors", reflect.TypeOf((*MockClient)(nil).GetAuthors), ctx, i)
}

// NewGetAuthorsIter mocks base method
func (m *MockClient) NewGetAuthorsIter(ctx context.Context, i *models.GetAuthorsInput) (GetAuthorsIter, error) {
	ret := m.ctrl.Call(m, "NewGetAuthorsIter", ctx, i)
	ret0, _ := ret[0].(GetAuthorsIter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewGetAuthorsIter indicates an expected call of NewGetAuthorsIter
func (mr *MockClientMockRecorder) NewGetAuthorsIter(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGetAuthorsIter", reflect.TypeOf((*MockClient)(nil).NewGetAuthorsIter), ctx, i)
}

// GetAuthorsWithPut mocks base method
func (m *MockClient) GetAuthorsWithPut(ctx context.Context, i *models.GetAuthorsWithPutInput) (*models.AuthorsResponse, error) {
	ret := m.ctrl.Call(m, "GetAuthorsWithPut", ctx, i)
	ret0, _ := ret[0].(*models.AuthorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorsWithPut indicates an expected call of GetAuthorsWithPut
func (mr *MockClientMockRecorder) GetAuthorsWithPut(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorsWithPut", reflect.TypeOf((*MockClient)(nil).GetAuthorsWithPut), ctx, i)
}

// NewGetAuthorsWithPutIter mocks base method
func (m *MockClient) NewGetAuthorsWithPutIter(ctx context.Context, i *models.GetAuthorsWithPutInput) (GetAuthorsWithPutIter, error) {
	ret := m.ctrl.Call(m, "NewGetAuthorsWithPutIter", ctx, i)
	ret0, _ := ret[0].(GetAuthorsWithPutIter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewGetAuthorsWithPutIter indicates an expected call of NewGetAuthorsWithPutIter
func (mr *MockClientMockRecorder) NewGetAuthorsWithPutIter(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGetAuthorsWithPutIter", reflect.TypeOf((*MockClient)(nil).NewGetAuthorsWithPutIter), ctx, i)
}

// GetBooks mocks base method
func (m *MockClient) GetBooks(ctx context.Context, i *models.GetBooksInput) ([]models.Book, error) {
	ret := m.ctrl.Call(m, "GetBooks", ctx, i)
	ret0, _ := ret[0].([]models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooks indicates an expected call of GetBooks
func (mr *MockClientMockRecorder) GetBooks(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooks", reflect.TypeOf((*MockClient)(nil).GetBooks), ctx, i)
}

// NewGetBooksIter mocks base method
func (m *MockClient) NewGetBooksIter(ctx context.Context, i *models.GetBooksInput) (GetBooksIter, error) {
	ret := m.ctrl.Call(m, "NewGetBooksIter", ctx, i)
	ret0, _ := ret[0].(GetBooksIter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewGetBooksIter indicates an expected call of NewGetBooksIter
func (mr *MockClientMockRecorder) NewGetBooksIter(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGetBooksIter", reflect.TypeOf((*MockClient)(nil).NewGetBooksIter), ctx, i)
}

// CreateBook mocks base method
func (m *MockClient) CreateBook(ctx context.Context, i *models.Book) (*models.Book, error) {
	ret := m.ctrl.Call(m, "CreateBook", ctx, i)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook
func (mr *MockClientMockRecorder) CreateBook(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockClient)(nil).CreateBook), ctx, i)
}

// PutBook mocks base method
func (m *MockClient) PutBook(ctx context.Context, i *models.Book) (*models.Book, error) {
	ret := m.ctrl.Call(m, "PutBook", ctx, i)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutBook indicates an expected call of PutBook
func (mr *MockClientMockRecorder) PutBook(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutBook", reflect.TypeOf((*MockClient)(nil).PutBook), ctx, i)
}

// GetBookByID mocks base method
func (m *MockClient) GetBookByID(ctx context.Context, i *models.GetBookByIDInput) (*models.Book, error) {
	ret := m.ctrl.Call(m, "GetBookByID", ctx, i)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByID indicates an expected call of GetBookByID
func (mr *MockClientMockRecorder) GetBookByID(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByID", reflect.TypeOf((*MockClient)(nil).GetBookByID), ctx, i)
}

// GetBookByID2 mocks base method
func (m *MockClient) GetBookByID2(ctx context.Context, id string) (*models.Book, error) {
	ret := m.ctrl.Call(m, "GetBookByID2", ctx, id)
	ret0, _ := ret[0].(*models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByID2 indicates an expected call of GetBookByID2
func (mr *MockClientMockRecorder) GetBookByID2(ctx, id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByID2", reflect.TypeOf((*MockClient)(nil).GetBookByID2), ctx, id)
}

// HealthCheck mocks base method
func (m *MockClient) HealthCheck(ctx context.Context) error {
	ret := m.ctrl.Call(m, "HealthCheck", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck
func (mr *MockClientMockRecorder) HealthCheck(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockClient)(nil).HealthCheck), ctx)
}

// MockGetAuthorsIter is a mock of GetAuthorsIter interface
type MockGetAuthorsIter struct {
	ctrl     *gomock.Controller
	recorder *MockGetAuthorsIterMockRecorder
}

// MockGetAuthorsIterMockRecorder is the mock recorder for MockGetAuthorsIter
type MockGetAuthorsIterMockRecorder struct {
	mock *MockGetAuthorsIter
}

// NewMockGetAuthorsIter creates a new mock instance
func NewMockGetAuthorsIter(ctrl *gomock.Controller) *MockGetAuthorsIter {
	mock := &MockGetAuthorsIter{ctrl: ctrl}
	mock.recorder = &MockGetAuthorsIterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetAuthorsIter) EXPECT() *MockGetAuthorsIterMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *MockGetAuthorsIter) Next(arg0 *models.Author) bool {
	ret := m.ctrl.Call(m, "Next", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockGetAuthorsIterMockRecorder) Next(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockGetAuthorsIter)(nil).Next), arg0)
}

// Err mocks base method
func (m *MockGetAuthorsIter) Err() error {
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockGetAuthorsIterMockRecorder) Err() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockGetAuthorsIter)(nil).Err))
}

// MockGetAuthorsWithPutIter is a mock of GetAuthorsWithPutIter interface
type MockGetAuthorsWithPutIter struct {
	ctrl     *gomock.Controller
	recorder *MockGetAuthorsWithPutIterMockRecorder
}

// MockGetAuthorsWithPutIterMockRecorder is the mock recorder for MockGetAuthorsWithPutIter
type MockGetAuthorsWithPutIterMockRecorder struct {
	mock *MockGetAuthorsWithPutIter
}

// NewMockGetAuthorsWithPutIter creates a new mock instance
func NewMockGetAuthorsWithPutIter(ctrl *gomock.Controller) *MockGetAuthorsWithPutIter {
	mock := &MockGetAuthorsWithPutIter{ctrl: ctrl}
	mock.recorder = &MockGetAuthorsWithPutIterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetAuthorsWithPutIter) EXPECT() *MockGetAuthorsWithPutIterMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *MockGetAuthorsWithPutIter) Next(arg0 *models.Author) bool {
	ret := m.ctrl.Call(m, "Next", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockGetAuthorsWithPutIterMockRecorder) Next(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockGetAuthorsWithPutIter)(nil).Next), arg0)
}

// Err mocks base method
func (m *MockGetAuthorsWithPutIter) Err() error {
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockGetAuthorsWithPutIterMockRecorder) Err() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockGetAuthorsWithPutIter)(nil).Err))
}

// MockGetBooksIter is a mock of GetBooksIter interface
type MockGetBooksIter struct {
	ctrl     *gomock.Controller
	recorder *MockGetBooksIterMockRecorder
}

// MockGetBooksIterMockRecorder is the mock recorder for MockGetBooksIter
type MockGetBooksIterMockRecorder struct {
	mock *MockGetBooksIter
}

// NewMockGetBooksIter creates a new mock instance
func NewMockGetBooksIter(ctrl *gomock.Controller) *MockGetBooksIter {
	mock := &MockGetBooksIter{ctrl: ctrl}
	mock.recorder = &MockGetBooksIterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetBooksIter) EXPECT() *MockGetBooksIterMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *MockGetBooksIter) Next(arg0 *models.Book) bool {
	ret := m.ctrl.Call(m, "Next", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockGetBooksIterMockRecorder) Next(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockGetBooksIter)(nil).Next), arg0)
}

// Err mocks base method
func (m *MockGetBooksIter) Err() error {
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockGetBooksIterMockRecorder) Err() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockGetBooksIter)(nil).Err))
}
