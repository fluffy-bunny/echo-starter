// Code generated by MockGen. DO NOT EDIT.
// Source: echo-starter/internal/contracts/claimsprovider (interfaces: IClaimsProvider)

// Package claimsprovider is a generated GoMock package.
package claimsprovider

import (
	reflect "reflect"

	claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	gomock "github.com/golang/mock/gomock"
)

// MockIClaimsProvider is a mock of IClaimsProvider interface.
type MockIClaimsProvider struct {
	ctrl     *gomock.Controller
	recorder *MockIClaimsProviderMockRecorder
}

// MockIClaimsProviderMockRecorder is the mock recorder for MockIClaimsProvider.
type MockIClaimsProviderMockRecorder struct {
	mock *MockIClaimsProvider
}

// NewMockIClaimsProvider creates a new mock instance.
func NewMockIClaimsProvider(ctrl *gomock.Controller) *MockIClaimsProvider {
	mock := &MockIClaimsProvider{ctrl: ctrl}
	mock.recorder = &MockIClaimsProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIClaimsProvider) EXPECT() *MockIClaimsProviderMockRecorder {
	return m.recorder
}

// GetClaims mocks base method.
func (m *MockIClaimsProvider) GetClaims(arg0 string) ([]*claimsprincipal.Claim, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClaims", arg0)
	ret0, _ := ret[0].([]*claimsprincipal.Claim)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClaims indicates an expected call of GetClaims.
func (mr *MockIClaimsProviderMockRecorder) GetClaims(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClaims", reflect.TypeOf((*MockIClaimsProvider)(nil).GetClaims), arg0)
}