// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package db is a generated GoMock package.
package db

import (
	context "context"
	models "github.com/Clever/wag/samples/gen-go-db/models"
	strfmt "github.com/go-openapi/strfmt"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// SaveNoRangeThingWithCompositeAttributes mocks base method
func (m_2 *MockInterface) SaveNoRangeThingWithCompositeAttributes(ctx context.Context, m models.NoRangeThingWithCompositeAttributes) error {
	ret := m_2.ctrl.Call(m_2, "SaveNoRangeThingWithCompositeAttributes", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveNoRangeThingWithCompositeAttributes indicates an expected call of SaveNoRangeThingWithCompositeAttributes
func (mr *MockInterfaceMockRecorder) SaveNoRangeThingWithCompositeAttributes(ctx, m interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveNoRangeThingWithCompositeAttributes", reflect.TypeOf((*MockInterface)(nil).SaveNoRangeThingWithCompositeAttributes), ctx, m)
}

// GetNoRangeThingWithCompositeAttributes mocks base method
func (m *MockInterface) GetNoRangeThingWithCompositeAttributes(ctx context.Context, name, branch string) (*models.NoRangeThingWithCompositeAttributes, error) {
	ret := m.ctrl.Call(m, "GetNoRangeThingWithCompositeAttributes", ctx, name, branch)
	ret0, _ := ret[0].(*models.NoRangeThingWithCompositeAttributes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNoRangeThingWithCompositeAttributes indicates an expected call of GetNoRangeThingWithCompositeAttributes
func (mr *MockInterfaceMockRecorder) GetNoRangeThingWithCompositeAttributes(ctx, name, branch interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNoRangeThingWithCompositeAttributes", reflect.TypeOf((*MockInterface)(nil).GetNoRangeThingWithCompositeAttributes), ctx, name, branch)
}

// DeleteNoRangeThingWithCompositeAttributes mocks base method
func (m *MockInterface) DeleteNoRangeThingWithCompositeAttributes(ctx context.Context, name, branch string) error {
	ret := m.ctrl.Call(m, "DeleteNoRangeThingWithCompositeAttributes", ctx, name, branch)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNoRangeThingWithCompositeAttributes indicates an expected call of DeleteNoRangeThingWithCompositeAttributes
func (mr *MockInterfaceMockRecorder) DeleteNoRangeThingWithCompositeAttributes(ctx, name, branch interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNoRangeThingWithCompositeAttributes", reflect.TypeOf((*MockInterface)(nil).DeleteNoRangeThingWithCompositeAttributes), ctx, name, branch)
}

// GetNoRangeThingWithCompositeAttributessByNameVersionAndDate mocks base method
func (m *MockInterface) GetNoRangeThingWithCompositeAttributessByNameVersionAndDate(ctx context.Context, input GetNoRangeThingWithCompositeAttributessByNameVersionAndDateInput, fn func(*models.NoRangeThingWithCompositeAttributes, bool) bool) error {
	ret := m.ctrl.Call(m, "GetNoRangeThingWithCompositeAttributessByNameVersionAndDate", ctx, input, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetNoRangeThingWithCompositeAttributessByNameVersionAndDate indicates an expected call of GetNoRangeThingWithCompositeAttributessByNameVersionAndDate
func (mr *MockInterfaceMockRecorder) GetNoRangeThingWithCompositeAttributessByNameVersionAndDate(ctx, input, fn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNoRangeThingWithCompositeAttributessByNameVersionAndDate", reflect.TypeOf((*MockInterface)(nil).GetNoRangeThingWithCompositeAttributessByNameVersionAndDate), ctx, input, fn)
}

// SaveSimpleThing mocks base method
func (m_2 *MockInterface) SaveSimpleThing(ctx context.Context, m models.SimpleThing) error {
	ret := m_2.ctrl.Call(m_2, "SaveSimpleThing", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveSimpleThing indicates an expected call of SaveSimpleThing
func (mr *MockInterfaceMockRecorder) SaveSimpleThing(ctx, m interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveSimpleThing", reflect.TypeOf((*MockInterface)(nil).SaveSimpleThing), ctx, m)
}

// GetSimpleThing mocks base method
func (m *MockInterface) GetSimpleThing(ctx context.Context, name string) (*models.SimpleThing, error) {
	ret := m.ctrl.Call(m, "GetSimpleThing", ctx, name)
	ret0, _ := ret[0].(*models.SimpleThing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSimpleThing indicates an expected call of GetSimpleThing
func (mr *MockInterfaceMockRecorder) GetSimpleThing(ctx, name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSimpleThing", reflect.TypeOf((*MockInterface)(nil).GetSimpleThing), ctx, name)
}

// DeleteSimpleThing mocks base method
func (m *MockInterface) DeleteSimpleThing(ctx context.Context, name string) error {
	ret := m.ctrl.Call(m, "DeleteSimpleThing", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSimpleThing indicates an expected call of DeleteSimpleThing
func (mr *MockInterfaceMockRecorder) DeleteSimpleThing(ctx, name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSimpleThing", reflect.TypeOf((*MockInterface)(nil).DeleteSimpleThing), ctx, name)
}

// SaveTeacherSharingRule mocks base method
func (m_2 *MockInterface) SaveTeacherSharingRule(ctx context.Context, m models.TeacherSharingRule) error {
	ret := m_2.ctrl.Call(m_2, "SaveTeacherSharingRule", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTeacherSharingRule indicates an expected call of SaveTeacherSharingRule
func (mr *MockInterfaceMockRecorder) SaveTeacherSharingRule(ctx, m interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTeacherSharingRule", reflect.TypeOf((*MockInterface)(nil).SaveTeacherSharingRule), ctx, m)
}

// GetTeacherSharingRule mocks base method
func (m *MockInterface) GetTeacherSharingRule(ctx context.Context, teacher, school, app string) (*models.TeacherSharingRule, error) {
	ret := m.ctrl.Call(m, "GetTeacherSharingRule", ctx, teacher, school, app)
	ret0, _ := ret[0].(*models.TeacherSharingRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeacherSharingRule indicates an expected call of GetTeacherSharingRule
func (mr *MockInterfaceMockRecorder) GetTeacherSharingRule(ctx, teacher, school, app interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeacherSharingRule", reflect.TypeOf((*MockInterface)(nil).GetTeacherSharingRule), ctx, teacher, school, app)
}

// GetTeacherSharingRulesByTeacherAndSchoolApp mocks base method
func (m *MockInterface) GetTeacherSharingRulesByTeacherAndSchoolApp(ctx context.Context, input GetTeacherSharingRulesByTeacherAndSchoolAppInput, fn func(*models.TeacherSharingRule, bool) bool) error {
	ret := m.ctrl.Call(m, "GetTeacherSharingRulesByTeacherAndSchoolApp", ctx, input, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetTeacherSharingRulesByTeacherAndSchoolApp indicates an expected call of GetTeacherSharingRulesByTeacherAndSchoolApp
func (mr *MockInterfaceMockRecorder) GetTeacherSharingRulesByTeacherAndSchoolApp(ctx, input, fn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeacherSharingRulesByTeacherAndSchoolApp", reflect.TypeOf((*MockInterface)(nil).GetTeacherSharingRulesByTeacherAndSchoolApp), ctx, input, fn)
}

// DeleteTeacherSharingRule mocks base method
func (m *MockInterface) DeleteTeacherSharingRule(ctx context.Context, teacher, school, app string) error {
	ret := m.ctrl.Call(m, "DeleteTeacherSharingRule", ctx, teacher, school, app)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTeacherSharingRule indicates an expected call of DeleteTeacherSharingRule
func (mr *MockInterfaceMockRecorder) DeleteTeacherSharingRule(ctx, teacher, school, app interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTeacherSharingRule", reflect.TypeOf((*MockInterface)(nil).DeleteTeacherSharingRule), ctx, teacher, school, app)
}

// GetTeacherSharingRulesByDistrictAndSchoolTeacherApp mocks base method
func (m *MockInterface) GetTeacherSharingRulesByDistrictAndSchoolTeacherApp(ctx context.Context, input GetTeacherSharingRulesByDistrictAndSchoolTeacherAppInput, fn func(*models.TeacherSharingRule, bool) bool) error {
	ret := m.ctrl.Call(m, "GetTeacherSharingRulesByDistrictAndSchoolTeacherApp", ctx, input, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetTeacherSharingRulesByDistrictAndSchoolTeacherApp indicates an expected call of GetTeacherSharingRulesByDistrictAndSchoolTeacherApp
func (mr *MockInterfaceMockRecorder) GetTeacherSharingRulesByDistrictAndSchoolTeacherApp(ctx, input, fn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeacherSharingRulesByDistrictAndSchoolTeacherApp", reflect.TypeOf((*MockInterface)(nil).GetTeacherSharingRulesByDistrictAndSchoolTeacherApp), ctx, input, fn)
}

// SaveThing mocks base method
func (m_2 *MockInterface) SaveThing(ctx context.Context, m models.Thing) error {
	ret := m_2.ctrl.Call(m_2, "SaveThing", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveThing indicates an expected call of SaveThing
func (mr *MockInterfaceMockRecorder) SaveThing(ctx, m interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveThing", reflect.TypeOf((*MockInterface)(nil).SaveThing), ctx, m)
}

// GetThing mocks base method
func (m *MockInterface) GetThing(ctx context.Context, name string, version int64) (*models.Thing, error) {
	ret := m.ctrl.Call(m, "GetThing", ctx, name, version)
	ret0, _ := ret[0].(*models.Thing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThing indicates an expected call of GetThing
func (mr *MockInterfaceMockRecorder) GetThing(ctx, name, version interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThing", reflect.TypeOf((*MockInterface)(nil).GetThing), ctx, name, version)
}

// ScanThings mocks base method
func (m *MockInterface) ScanThings(ctx context.Context, input ScanThingsInput, fn func(*models.Thing, bool) bool) error {
	ret := m.ctrl.Call(m, "ScanThings", ctx, input, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanThings indicates an expected call of ScanThings
func (mr *MockInterfaceMockRecorder) ScanThings(ctx, input, fn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanThings", reflect.TypeOf((*MockInterface)(nil).ScanThings), ctx, input, fn)
}

// GetThingsByNameAndVersion mocks base method
func (m *MockInterface) GetThingsByNameAndVersion(ctx context.Context, input GetThingsByNameAndVersionInput, fn func(*models.Thing, bool) bool) error {
	ret := m.ctrl.Call(m, "GetThingsByNameAndVersion", ctx, input, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetThingsByNameAndVersion indicates an expected call of GetThingsByNameAndVersion
func (mr *MockInterfaceMockRecorder) GetThingsByNameAndVersion(ctx, input, fn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingsByNameAndVersion", reflect.TypeOf((*MockInterface)(nil).GetThingsByNameAndVersion), ctx, input, fn)
}

// DeleteThing mocks base method
func (m *MockInterface) DeleteThing(ctx context.Context, name string, version int64) error {
	ret := m.ctrl.Call(m, "DeleteThing", ctx, name, version)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteThing indicates an expected call of DeleteThing
func (mr *MockInterfaceMockRecorder) DeleteThing(ctx, name, version interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteThing", reflect.TypeOf((*MockInterface)(nil).DeleteThing), ctx, name, version)
}

// GetThingByID mocks base method
func (m *MockInterface) GetThingByID(ctx context.Context, id string) (*models.Thing, error) {
	ret := m.ctrl.Call(m, "GetThingByID", ctx, id)
	ret0, _ := ret[0].(*models.Thing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThingByID indicates an expected call of GetThingByID
func (mr *MockInterfaceMockRecorder) GetThingByID(ctx, id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingByID", reflect.TypeOf((*MockInterface)(nil).GetThingByID), ctx, id)
}

// GetThingsByNameAndCreatedAt mocks base method
func (m *MockInterface) GetThingsByNameAndCreatedAt(ctx context.Context, input GetThingsByNameAndCreatedAtInput, fn func(*models.Thing, bool) bool) error {
	ret := m.ctrl.Call(m, "GetThingsByNameAndCreatedAt", ctx, input, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetThingsByNameAndCreatedAt indicates an expected call of GetThingsByNameAndCreatedAt
func (mr *MockInterfaceMockRecorder) GetThingsByNameAndCreatedAt(ctx, input, fn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingsByNameAndCreatedAt", reflect.TypeOf((*MockInterface)(nil).GetThingsByNameAndCreatedAt), ctx, input, fn)
}

// SaveThingWithCompositeAttributes mocks base method
func (m_2 *MockInterface) SaveThingWithCompositeAttributes(ctx context.Context, m models.ThingWithCompositeAttributes) error {
	ret := m_2.ctrl.Call(m_2, "SaveThingWithCompositeAttributes", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveThingWithCompositeAttributes indicates an expected call of SaveThingWithCompositeAttributes
func (mr *MockInterfaceMockRecorder) SaveThingWithCompositeAttributes(ctx, m interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveThingWithCompositeAttributes", reflect.TypeOf((*MockInterface)(nil).SaveThingWithCompositeAttributes), ctx, m)
}

// GetThingWithCompositeAttributes mocks base method
func (m *MockInterface) GetThingWithCompositeAttributes(ctx context.Context, name, branch string, date strfmt.DateTime) (*models.ThingWithCompositeAttributes, error) {
	ret := m.ctrl.Call(m, "GetThingWithCompositeAttributes", ctx, name, branch, date)
	ret0, _ := ret[0].(*models.ThingWithCompositeAttributes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThingWithCompositeAttributes indicates an expected call of GetThingWithCompositeAttributes
func (mr *MockInterfaceMockRecorder) GetThingWithCompositeAttributes(ctx, name, branch, date interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingWithCompositeAttributes", reflect.TypeOf((*MockInterface)(nil).GetThingWithCompositeAttributes), ctx, name, branch, date)
}

// GetThingWithCompositeAttributessByNameBranchAndDate mocks base method
func (m *MockInterface) GetThingWithCompositeAttributessByNameBranchAndDate(ctx context.Context, input GetThingWithCompositeAttributessByNameBranchAndDateInput, fn func(*models.ThingWithCompositeAttributes, bool) bool) error {
	ret := m.ctrl.Call(m, "GetThingWithCompositeAttributessByNameBranchAndDate", ctx, input, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetThingWithCompositeAttributessByNameBranchAndDate indicates an expected call of GetThingWithCompositeAttributessByNameBranchAndDate
func (mr *MockInterfaceMockRecorder) GetThingWithCompositeAttributessByNameBranchAndDate(ctx, input, fn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingWithCompositeAttributessByNameBranchAndDate", reflect.TypeOf((*MockInterface)(nil).GetThingWithCompositeAttributessByNameBranchAndDate), ctx, input, fn)
}

// DeleteThingWithCompositeAttributes mocks base method
func (m *MockInterface) DeleteThingWithCompositeAttributes(ctx context.Context, name, branch string, date strfmt.DateTime) error {
	ret := m.ctrl.Call(m, "DeleteThingWithCompositeAttributes", ctx, name, branch, date)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteThingWithCompositeAttributes indicates an expected call of DeleteThingWithCompositeAttributes
func (mr *MockInterfaceMockRecorder) DeleteThingWithCompositeAttributes(ctx, name, branch, date interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteThingWithCompositeAttributes", reflect.TypeOf((*MockInterface)(nil).DeleteThingWithCompositeAttributes), ctx, name, branch, date)
}

// GetThingWithCompositeAttributessByNameVersionAndDate mocks base method
func (m *MockInterface) GetThingWithCompositeAttributessByNameVersionAndDate(ctx context.Context, input GetThingWithCompositeAttributessByNameVersionAndDateInput, fn func(*models.ThingWithCompositeAttributes, bool) bool) error {
	ret := m.ctrl.Call(m, "GetThingWithCompositeAttributessByNameVersionAndDate", ctx, input, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetThingWithCompositeAttributessByNameVersionAndDate indicates an expected call of GetThingWithCompositeAttributessByNameVersionAndDate
func (mr *MockInterfaceMockRecorder) GetThingWithCompositeAttributessByNameVersionAndDate(ctx, input, fn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingWithCompositeAttributessByNameVersionAndDate", reflect.TypeOf((*MockInterface)(nil).GetThingWithCompositeAttributessByNameVersionAndDate), ctx, input, fn)
}

// SaveThingWithCompositeEnumAttributes mocks base method
func (m_2 *MockInterface) SaveThingWithCompositeEnumAttributes(ctx context.Context, m models.ThingWithCompositeEnumAttributes) error {
	ret := m_2.ctrl.Call(m_2, "SaveThingWithCompositeEnumAttributes", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveThingWithCompositeEnumAttributes indicates an expected call of SaveThingWithCompositeEnumAttributes
func (mr *MockInterfaceMockRecorder) SaveThingWithCompositeEnumAttributes(ctx, m interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveThingWithCompositeEnumAttributes", reflect.TypeOf((*MockInterface)(nil).SaveThingWithCompositeEnumAttributes), ctx, m)
}

// GetThingWithCompositeEnumAttributes mocks base method
func (m *MockInterface) GetThingWithCompositeEnumAttributes(ctx context.Context, name string, branchID models.Branch, date strfmt.DateTime) (*models.ThingWithCompositeEnumAttributes, error) {
	ret := m.ctrl.Call(m, "GetThingWithCompositeEnumAttributes", ctx, name, branchID, date)
	ret0, _ := ret[0].(*models.ThingWithCompositeEnumAttributes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThingWithCompositeEnumAttributes indicates an expected call of GetThingWithCompositeEnumAttributes
func (mr *MockInterfaceMockRecorder) GetThingWithCompositeEnumAttributes(ctx, name, branchID, date interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingWithCompositeEnumAttributes", reflect.TypeOf((*MockInterface)(nil).GetThingWithCompositeEnumAttributes), ctx, name, branchID, date)
}

// GetThingWithCompositeEnumAttributessByNameBranchAndDate mocks base method
func (m *MockInterface) GetThingWithCompositeEnumAttributessByNameBranchAndDate(ctx context.Context, input GetThingWithCompositeEnumAttributessByNameBranchAndDateInput, fn func(*models.ThingWithCompositeEnumAttributes, bool) bool) error {
	ret := m.ctrl.Call(m, "GetThingWithCompositeEnumAttributessByNameBranchAndDate", ctx, input, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetThingWithCompositeEnumAttributessByNameBranchAndDate indicates an expected call of GetThingWithCompositeEnumAttributessByNameBranchAndDate
func (mr *MockInterfaceMockRecorder) GetThingWithCompositeEnumAttributessByNameBranchAndDate(ctx, input, fn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingWithCompositeEnumAttributessByNameBranchAndDate", reflect.TypeOf((*MockInterface)(nil).GetThingWithCompositeEnumAttributessByNameBranchAndDate), ctx, input, fn)
}

// DeleteThingWithCompositeEnumAttributes mocks base method
func (m *MockInterface) DeleteThingWithCompositeEnumAttributes(ctx context.Context, name string, branchID models.Branch, date strfmt.DateTime) error {
	ret := m.ctrl.Call(m, "DeleteThingWithCompositeEnumAttributes", ctx, name, branchID, date)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteThingWithCompositeEnumAttributes indicates an expected call of DeleteThingWithCompositeEnumAttributes
func (mr *MockInterfaceMockRecorder) DeleteThingWithCompositeEnumAttributes(ctx, name, branchID, date interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteThingWithCompositeEnumAttributes", reflect.TypeOf((*MockInterface)(nil).DeleteThingWithCompositeEnumAttributes), ctx, name, branchID, date)
}

// SaveThingWithDateRange mocks base method
func (m_2 *MockInterface) SaveThingWithDateRange(ctx context.Context, m models.ThingWithDateRange) error {
	ret := m_2.ctrl.Call(m_2, "SaveThingWithDateRange", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveThingWithDateRange indicates an expected call of SaveThingWithDateRange
func (mr *MockInterfaceMockRecorder) SaveThingWithDateRange(ctx, m interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveThingWithDateRange", reflect.TypeOf((*MockInterface)(nil).SaveThingWithDateRange), ctx, m)
}

// GetThingWithDateRange mocks base method
func (m *MockInterface) GetThingWithDateRange(ctx context.Context, name string, date strfmt.DateTime) (*models.ThingWithDateRange, error) {
	ret := m.ctrl.Call(m, "GetThingWithDateRange", ctx, name, date)
	ret0, _ := ret[0].(*models.ThingWithDateRange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThingWithDateRange indicates an expected call of GetThingWithDateRange
func (mr *MockInterfaceMockRecorder) GetThingWithDateRange(ctx, name, date interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingWithDateRange", reflect.TypeOf((*MockInterface)(nil).GetThingWithDateRange), ctx, name, date)
}

// GetThingWithDateRangesByNameAndDate mocks base method
func (m *MockInterface) GetThingWithDateRangesByNameAndDate(ctx context.Context, input GetThingWithDateRangesByNameAndDateInput, fn func(*models.ThingWithDateRange, bool) bool) error {
	ret := m.ctrl.Call(m, "GetThingWithDateRangesByNameAndDate", ctx, input, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetThingWithDateRangesByNameAndDate indicates an expected call of GetThingWithDateRangesByNameAndDate
func (mr *MockInterfaceMockRecorder) GetThingWithDateRangesByNameAndDate(ctx, input, fn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingWithDateRangesByNameAndDate", reflect.TypeOf((*MockInterface)(nil).GetThingWithDateRangesByNameAndDate), ctx, input, fn)
}

// DeleteThingWithDateRange mocks base method
func (m *MockInterface) DeleteThingWithDateRange(ctx context.Context, name string, date strfmt.DateTime) error {
	ret := m.ctrl.Call(m, "DeleteThingWithDateRange", ctx, name, date)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteThingWithDateRange indicates an expected call of DeleteThingWithDateRange
func (mr *MockInterfaceMockRecorder) DeleteThingWithDateRange(ctx, name, date interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteThingWithDateRange", reflect.TypeOf((*MockInterface)(nil).DeleteThingWithDateRange), ctx, name, date)
}

// SaveThingWithDateTimeComposite mocks base method
func (m_2 *MockInterface) SaveThingWithDateTimeComposite(ctx context.Context, m models.ThingWithDateTimeComposite) error {
	ret := m_2.ctrl.Call(m_2, "SaveThingWithDateTimeComposite", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveThingWithDateTimeComposite indicates an expected call of SaveThingWithDateTimeComposite
func (mr *MockInterfaceMockRecorder) SaveThingWithDateTimeComposite(ctx, m interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveThingWithDateTimeComposite", reflect.TypeOf((*MockInterface)(nil).SaveThingWithDateTimeComposite), ctx, m)
}

// GetThingWithDateTimeComposite mocks base method
func (m *MockInterface) GetThingWithDateTimeComposite(ctx context.Context, typeVar, id string, created strfmt.DateTime, resource string) (*models.ThingWithDateTimeComposite, error) {
	ret := m.ctrl.Call(m, "GetThingWithDateTimeComposite", ctx, typeVar, id, created, resource)
	ret0, _ := ret[0].(*models.ThingWithDateTimeComposite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThingWithDateTimeComposite indicates an expected call of GetThingWithDateTimeComposite
func (mr *MockInterfaceMockRecorder) GetThingWithDateTimeComposite(ctx, typeVar, id, created, resource interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingWithDateTimeComposite", reflect.TypeOf((*MockInterface)(nil).GetThingWithDateTimeComposite), ctx, typeVar, id, created, resource)
}

// GetThingWithDateTimeCompositesByTypeIDAndCreatedResource mocks base method
func (m *MockInterface) GetThingWithDateTimeCompositesByTypeIDAndCreatedResource(ctx context.Context, input GetThingWithDateTimeCompositesByTypeIDAndCreatedResourceInput, fn func(*models.ThingWithDateTimeComposite, bool) bool) error {
	ret := m.ctrl.Call(m, "GetThingWithDateTimeCompositesByTypeIDAndCreatedResource", ctx, input, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetThingWithDateTimeCompositesByTypeIDAndCreatedResource indicates an expected call of GetThingWithDateTimeCompositesByTypeIDAndCreatedResource
func (mr *MockInterfaceMockRecorder) GetThingWithDateTimeCompositesByTypeIDAndCreatedResource(ctx, input, fn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingWithDateTimeCompositesByTypeIDAndCreatedResource", reflect.TypeOf((*MockInterface)(nil).GetThingWithDateTimeCompositesByTypeIDAndCreatedResource), ctx, input, fn)
}

// DeleteThingWithDateTimeComposite mocks base method
func (m *MockInterface) DeleteThingWithDateTimeComposite(ctx context.Context, typeVar, id string, created strfmt.DateTime, resource string) error {
	ret := m.ctrl.Call(m, "DeleteThingWithDateTimeComposite", ctx, typeVar, id, created, resource)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteThingWithDateTimeComposite indicates an expected call of DeleteThingWithDateTimeComposite
func (mr *MockInterfaceMockRecorder) DeleteThingWithDateTimeComposite(ctx, typeVar, id, created, resource interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteThingWithDateTimeComposite", reflect.TypeOf((*MockInterface)(nil).DeleteThingWithDateTimeComposite), ctx, typeVar, id, created, resource)
}

// SaveThingWithRequiredFields mocks base method
func (m_2 *MockInterface) SaveThingWithRequiredFields(ctx context.Context, m models.ThingWithRequiredFields) error {
	ret := m_2.ctrl.Call(m_2, "SaveThingWithRequiredFields", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveThingWithRequiredFields indicates an expected call of SaveThingWithRequiredFields
func (mr *MockInterfaceMockRecorder) SaveThingWithRequiredFields(ctx, m interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveThingWithRequiredFields", reflect.TypeOf((*MockInterface)(nil).SaveThingWithRequiredFields), ctx, m)
}

// GetThingWithRequiredFields mocks base method
func (m *MockInterface) GetThingWithRequiredFields(ctx context.Context, name string) (*models.ThingWithRequiredFields, error) {
	ret := m.ctrl.Call(m, "GetThingWithRequiredFields", ctx, name)
	ret0, _ := ret[0].(*models.ThingWithRequiredFields)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThingWithRequiredFields indicates an expected call of GetThingWithRequiredFields
func (mr *MockInterfaceMockRecorder) GetThingWithRequiredFields(ctx, name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingWithRequiredFields", reflect.TypeOf((*MockInterface)(nil).GetThingWithRequiredFields), ctx, name)
}

// DeleteThingWithRequiredFields mocks base method
func (m *MockInterface) DeleteThingWithRequiredFields(ctx context.Context, name string) error {
	ret := m.ctrl.Call(m, "DeleteThingWithRequiredFields", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteThingWithRequiredFields indicates an expected call of DeleteThingWithRequiredFields
func (mr *MockInterfaceMockRecorder) DeleteThingWithRequiredFields(ctx, name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteThingWithRequiredFields", reflect.TypeOf((*MockInterface)(nil).DeleteThingWithRequiredFields), ctx, name)
}

// SaveThingWithRequiredFields2 mocks base method
func (m_2 *MockInterface) SaveThingWithRequiredFields2(ctx context.Context, m models.ThingWithRequiredFields2) error {
	ret := m_2.ctrl.Call(m_2, "SaveThingWithRequiredFields2", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveThingWithRequiredFields2 indicates an expected call of SaveThingWithRequiredFields2
func (mr *MockInterfaceMockRecorder) SaveThingWithRequiredFields2(ctx, m interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveThingWithRequiredFields2", reflect.TypeOf((*MockInterface)(nil).SaveThingWithRequiredFields2), ctx, m)
}

// GetThingWithRequiredFields2 mocks base method
func (m *MockInterface) GetThingWithRequiredFields2(ctx context.Context, name, id string) (*models.ThingWithRequiredFields2, error) {
	ret := m.ctrl.Call(m, "GetThingWithRequiredFields2", ctx, name, id)
	ret0, _ := ret[0].(*models.ThingWithRequiredFields2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThingWithRequiredFields2 indicates an expected call of GetThingWithRequiredFields2
func (mr *MockInterfaceMockRecorder) GetThingWithRequiredFields2(ctx, name, id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingWithRequiredFields2", reflect.TypeOf((*MockInterface)(nil).GetThingWithRequiredFields2), ctx, name, id)
}

// GetThingWithRequiredFields2sByNameAndID mocks base method
func (m *MockInterface) GetThingWithRequiredFields2sByNameAndID(ctx context.Context, input GetThingWithRequiredFields2sByNameAndIDInput, fn func(*models.ThingWithRequiredFields2, bool) bool) error {
	ret := m.ctrl.Call(m, "GetThingWithRequiredFields2sByNameAndID", ctx, input, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetThingWithRequiredFields2sByNameAndID indicates an expected call of GetThingWithRequiredFields2sByNameAndID
func (mr *MockInterfaceMockRecorder) GetThingWithRequiredFields2sByNameAndID(ctx, input, fn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingWithRequiredFields2sByNameAndID", reflect.TypeOf((*MockInterface)(nil).GetThingWithRequiredFields2sByNameAndID), ctx, input, fn)
}

// DeleteThingWithRequiredFields2 mocks base method
func (m *MockInterface) DeleteThingWithRequiredFields2(ctx context.Context, name, id string) error {
	ret := m.ctrl.Call(m, "DeleteThingWithRequiredFields2", ctx, name, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteThingWithRequiredFields2 indicates an expected call of DeleteThingWithRequiredFields2
func (mr *MockInterfaceMockRecorder) DeleteThingWithRequiredFields2(ctx, name, id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteThingWithRequiredFields2", reflect.TypeOf((*MockInterface)(nil).DeleteThingWithRequiredFields2), ctx, name, id)
}

// SaveThingWithUnderscores mocks base method
func (m_2 *MockInterface) SaveThingWithUnderscores(ctx context.Context, m models.ThingWithUnderscores) error {
	ret := m_2.ctrl.Call(m_2, "SaveThingWithUnderscores", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveThingWithUnderscores indicates an expected call of SaveThingWithUnderscores
func (mr *MockInterfaceMockRecorder) SaveThingWithUnderscores(ctx, m interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveThingWithUnderscores", reflect.TypeOf((*MockInterface)(nil).SaveThingWithUnderscores), ctx, m)
}

// GetThingWithUnderscores mocks base method
func (m *MockInterface) GetThingWithUnderscores(ctx context.Context, iDApp string) (*models.ThingWithUnderscores, error) {
	ret := m.ctrl.Call(m, "GetThingWithUnderscores", ctx, iDApp)
	ret0, _ := ret[0].(*models.ThingWithUnderscores)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThingWithUnderscores indicates an expected call of GetThingWithUnderscores
func (mr *MockInterfaceMockRecorder) GetThingWithUnderscores(ctx, iDApp interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThingWithUnderscores", reflect.TypeOf((*MockInterface)(nil).GetThingWithUnderscores), ctx, iDApp)
}

// DeleteThingWithUnderscores mocks base method
func (m *MockInterface) DeleteThingWithUnderscores(ctx context.Context, iDApp string) error {
	ret := m.ctrl.Call(m, "DeleteThingWithUnderscores", ctx, iDApp)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteThingWithUnderscores indicates an expected call of DeleteThingWithUnderscores
func (mr *MockInterfaceMockRecorder) DeleteThingWithUnderscores(ctx, iDApp interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteThingWithUnderscores", reflect.TypeOf((*MockInterface)(nil).DeleteThingWithUnderscores), ctx, iDApp)
}
