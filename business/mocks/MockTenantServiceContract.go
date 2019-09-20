// Code generated by MockGen. DO NOT EDIT.
// Source: ../contracts/TenantServiceContract.go

// Package mock_contracts is a generated GoMock package.
package mock_contracts

import (
	context "context"
	reflect "reflect"

	contracts "github.com/decentralized-cloud/tenant/business/contracts"
	gomock "github.com/golang/mock/gomock"
)

// MockTenantServiceContract is a mock of TenantServiceContract interface
type MockTenantServiceContract struct {
	ctrl     *gomock.Controller
	recorder *MockTenantServiceContractMockRecorder
}

// MockTenantServiceContractMockRecorder is the mock recorder for MockTenantServiceContract
type MockTenantServiceContractMockRecorder struct {
	mock *MockTenantServiceContract
}

// NewMockTenantServiceContract creates a new mock instance
func NewMockTenantServiceContract(ctrl *gomock.Controller) *MockTenantServiceContract {
	mock := &MockTenantServiceContract{ctrl: ctrl}
	mock.recorder = &MockTenantServiceContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTenantServiceContract) EXPECT() *MockTenantServiceContractMockRecorder {
	return m.recorder
}

// CreateTenant mocks base method
func (m *MockTenantServiceContract) CreateTenant(ctx context.Context, request *contracts.CreateTenantRequest) (*contracts.CreateTenantResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTenant", ctx, request)
	ret0, _ := ret[0].(*contracts.CreateTenantResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTenant indicates an expected call of CreateTenant
func (mr *MockTenantServiceContractMockRecorder) CreateTenant(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTenant", reflect.TypeOf((*MockTenantServiceContract)(nil).CreateTenant), ctx, request)
}

// ReadTenant mocks base method
func (m *MockTenantServiceContract) ReadTenant(ctx context.Context, request *contracts.ReadTenantRequest) (*contracts.ReadTenantResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadTenant", ctx, request)
	ret0, _ := ret[0].(*contracts.ReadTenantResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadTenant indicates an expected call of ReadTenant
func (mr *MockTenantServiceContractMockRecorder) ReadTenant(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadTenant", reflect.TypeOf((*MockTenantServiceContract)(nil).ReadTenant), ctx, request)
}

// UpdateTenant mocks base method
func (m *MockTenantServiceContract) UpdateTenant(ctx context.Context, request *contracts.UpdateTenantRequest) (*contracts.UpdateTenantResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTenant", ctx, request)
	ret0, _ := ret[0].(*contracts.UpdateTenantResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTenant indicates an expected call of UpdateTenant
func (mr *MockTenantServiceContractMockRecorder) UpdateTenant(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTenant", reflect.TypeOf((*MockTenantServiceContract)(nil).UpdateTenant), ctx, request)
}

// DeleteTenant mocks base method
func (m *MockTenantServiceContract) DeleteTenant(ctx context.Context, request *contracts.DeleteTenantRequest) (*contracts.DeleteTenantResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTenant", ctx, request)
	ret0, _ := ret[0].(*contracts.DeleteTenantResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTenant indicates an expected call of DeleteTenant
func (mr *MockTenantServiceContractMockRecorder) DeleteTenant(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTenant", reflect.TypeOf((*MockTenantServiceContract)(nil).DeleteTenant), ctx, request)
}
