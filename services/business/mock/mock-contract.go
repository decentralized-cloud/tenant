// Code generated by MockGen. DO NOT EDIT.
// Source: ../contract.go

// Package mock_business is a generated GoMock package.
package mock_business

import (
	context "context"
	business "github.com/decentralized-cloud/tenant/services/business"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBusinessContract is a mock of BusinessContract interface
type MockBusinessContract struct {
	ctrl     *gomock.Controller
	recorder *MockBusinessContractMockRecorder
}

// MockBusinessContractMockRecorder is the mock recorder for MockBusinessContract
type MockBusinessContractMockRecorder struct {
	mock *MockBusinessContract
}

// NewMockBusinessContract creates a new mock instance
func NewMockBusinessContract(ctrl *gomock.Controller) *MockBusinessContract {
	mock := &MockBusinessContract{ctrl: ctrl}
	mock.recorder = &MockBusinessContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBusinessContract) EXPECT() *MockBusinessContractMockRecorder {
	return m.recorder
}

// CreateTenant mocks base method
func (m *MockBusinessContract) CreateTenant(ctx context.Context, request *business.CreateTenantRequest) (*business.CreateTenantResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTenant", ctx, request)
	ret0, _ := ret[0].(*business.CreateTenantResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTenant indicates an expected call of CreateTenant
func (mr *MockBusinessContractMockRecorder) CreateTenant(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTenant", reflect.TypeOf((*MockBusinessContract)(nil).CreateTenant), ctx, request)
}

// ReadTenant mocks base method
func (m *MockBusinessContract) ReadTenant(ctx context.Context, request *business.ReadTenantRequest) (*business.ReadTenantResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadTenant", ctx, request)
	ret0, _ := ret[0].(*business.ReadTenantResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadTenant indicates an expected call of ReadTenant
func (mr *MockBusinessContractMockRecorder) ReadTenant(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadTenant", reflect.TypeOf((*MockBusinessContract)(nil).ReadTenant), ctx, request)
}

// UpdateTenant mocks base method
func (m *MockBusinessContract) UpdateTenant(ctx context.Context, request *business.UpdateTenantRequest) (*business.UpdateTenantResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTenant", ctx, request)
	ret0, _ := ret[0].(*business.UpdateTenantResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTenant indicates an expected call of UpdateTenant
func (mr *MockBusinessContractMockRecorder) UpdateTenant(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTenant", reflect.TypeOf((*MockBusinessContract)(nil).UpdateTenant), ctx, request)
}

// DeleteTenant mocks base method
func (m *MockBusinessContract) DeleteTenant(ctx context.Context, request *business.DeleteTenantRequest) (*business.DeleteTenantResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTenant", ctx, request)
	ret0, _ := ret[0].(*business.DeleteTenantResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTenant indicates an expected call of DeleteTenant
func (mr *MockBusinessContractMockRecorder) DeleteTenant(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTenant", reflect.TypeOf((*MockBusinessContract)(nil).DeleteTenant), ctx, request)
}
