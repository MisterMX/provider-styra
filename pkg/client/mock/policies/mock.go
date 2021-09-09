// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mistermx/styra-go-client/pkg/client/policies (interfaces: ClientService)

// Package policies is a generated GoMock package.
package policies

import (
	reflect "reflect"

	runtime "github.com/go-openapi/runtime"
	gomock "github.com/golang/mock/gomock"
	policies "github.com/mistermx/styra-go-client/pkg/client/policies"
)

// MockClientService is a mock of ClientService interface.
type MockClientService struct {
	ctrl     *gomock.Controller
	recorder *MockClientServiceMockRecorder
}

// MockClientServiceMockRecorder is the mock recorder for MockClientService.
type MockClientServiceMockRecorder struct {
	mock *MockClientService
}

// NewMockClientService creates a new mock instance.
func NewMockClientService(ctrl *gomock.Controller) *MockClientService {
	mock := &MockClientService{ctrl: ctrl}
	mock.recorder = &MockClientServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientService) EXPECT() *MockClientServiceMockRecorder {
	return m.recorder
}

// BulkUploadPolicies mocks base method.
func (m *MockClientService) BulkUploadPolicies(arg0 *policies.BulkUploadPoliciesParams, arg1 ...policies.ClientOption) (*policies.BulkUploadPoliciesOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BulkUploadPolicies", varargs...)
	ret0, _ := ret[0].(*policies.BulkUploadPoliciesOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkUploadPolicies indicates an expected call of BulkUploadPolicies.
func (mr *MockClientServiceMockRecorder) BulkUploadPolicies(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkUploadPolicies", reflect.TypeOf((*MockClientService)(nil).BulkUploadPolicies), varargs...)
}

// BulkUploadSystemPolicies mocks base method.
func (m *MockClientService) BulkUploadSystemPolicies(arg0 *policies.BulkUploadSystemPoliciesParams, arg1 ...policies.ClientOption) (*policies.BulkUploadSystemPoliciesOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BulkUploadSystemPolicies", varargs...)
	ret0, _ := ret[0].(*policies.BulkUploadSystemPoliciesOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkUploadSystemPolicies indicates an expected call of BulkUploadSystemPolicies.
func (mr *MockClientServiceMockRecorder) BulkUploadSystemPolicies(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkUploadSystemPolicies", reflect.TypeOf((*MockClientService)(nil).BulkUploadSystemPolicies), varargs...)
}

// DeletePolicy mocks base method.
func (m *MockClientService) DeletePolicy(arg0 *policies.DeletePolicyParams, arg1 ...policies.ClientOption) (*policies.DeletePolicyOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeletePolicy", varargs...)
	ret0, _ := ret[0].(*policies.DeletePolicyOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePolicy indicates an expected call of DeletePolicy.
func (mr *MockClientServiceMockRecorder) DeletePolicy(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePolicy", reflect.TypeOf((*MockClientService)(nil).DeletePolicy), varargs...)
}

// GetPolicy mocks base method.
func (m *MockClientService) GetPolicy(arg0 *policies.GetPolicyParams, arg1 ...policies.ClientOption) (*policies.GetPolicyOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPolicy", varargs...)
	ret0, _ := ret[0].(*policies.GetPolicyOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicy indicates an expected call of GetPolicy.
func (mr *MockClientServiceMockRecorder) GetPolicy(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicy", reflect.TypeOf((*MockClientService)(nil).GetPolicy), varargs...)
}

// ListPolicies mocks base method.
func (m *MockClientService) ListPolicies(arg0 *policies.ListPoliciesParams, arg1 ...policies.ClientOption) (*policies.ListPoliciesOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPolicies", varargs...)
	ret0, _ := ret[0].(*policies.ListPoliciesOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPolicies indicates an expected call of ListPolicies.
func (mr *MockClientServiceMockRecorder) ListPolicies(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPolicies", reflect.TypeOf((*MockClientService)(nil).ListPolicies), varargs...)
}

// ListSystemPolicies mocks base method.
func (m *MockClientService) ListSystemPolicies(arg0 *policies.ListSystemPoliciesParams, arg1 ...policies.ClientOption) (*policies.ListSystemPoliciesOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSystemPolicies", varargs...)
	ret0, _ := ret[0].(*policies.ListSystemPoliciesOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSystemPolicies indicates an expected call of ListSystemPolicies.
func (mr *MockClientServiceMockRecorder) ListSystemPolicies(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSystemPolicies", reflect.TypeOf((*MockClientService)(nil).ListSystemPolicies), varargs...)
}

// SetTransport mocks base method.
func (m *MockClientService) SetTransport(arg0 runtime.ClientTransport) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTransport", arg0)
}

// SetTransport indicates an expected call of SetTransport.
func (mr *MockClientServiceMockRecorder) SetTransport(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTransport", reflect.TypeOf((*MockClientService)(nil).SetTransport), arg0)
}

// UpdatePolicy mocks base method.
func (m *MockClientService) UpdatePolicy(arg0 *policies.UpdatePolicyParams, arg1 ...policies.ClientOption) (*policies.UpdatePolicyOK, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdatePolicy", varargs...)
	ret0, _ := ret[0].(*policies.UpdatePolicyOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePolicy indicates an expected call of UpdatePolicy.
func (mr *MockClientServiceMockRecorder) UpdatePolicy(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePolicy", reflect.TypeOf((*MockClientService)(nil).UpdatePolicy), varargs...)
}