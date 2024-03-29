// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/ips.go

// Package mock_store is a generated GoMock package.
package mock_store

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/maxheckel/censys-assessment/internal/domain"
)

// MockIPStore is a mock of IPStore interface.
type MockIPStore struct {
	ctrl     *gomock.Controller
	recorder *MockIPStoreMockRecorder
}

// MockIPStoreMockRecorder is the mock recorder for MockIPStore.
type MockIPStoreMockRecorder struct {
	mock *MockIPStore
}

// NewMockIPStore creates a new mock instance.
func NewMockIPStore(ctrl *gomock.Controller) *MockIPStore {
	mock := &MockIPStore{ctrl: ctrl}
	mock.recorder = &MockIPStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPStore) EXPECT() *MockIPStoreMockRecorder {
	return m.recorder
}

// GetIP mocks base method.
func (m *MockIPStore) GetIP(address string) (domain.IP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIP", address)
	ret0, _ := ret[0].(domain.IP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIP indicates an expected call of GetIP.
func (mr *MockIPStoreMockRecorder) GetIP(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIP", reflect.TypeOf((*MockIPStore)(nil).GetIP), address)
}
