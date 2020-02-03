// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/api/v1/rbac_event_loop.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/mesh-projects/pkg/api/v1"
)

// MockRbacSyncer is a mock of RbacSyncer interface
type MockRbacSyncer struct {
	ctrl     *gomock.Controller
	recorder *MockRbacSyncerMockRecorder
}

// MockRbacSyncerMockRecorder is the mock recorder for MockRbacSyncer
type MockRbacSyncerMockRecorder struct {
	mock *MockRbacSyncer
}

// NewMockRbacSyncer creates a new mock instance
func NewMockRbacSyncer(ctrl *gomock.Controller) *MockRbacSyncer {
	mock := &MockRbacSyncer{ctrl: ctrl}
	mock.recorder = &MockRbacSyncerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRbacSyncer) EXPECT() *MockRbacSyncerMockRecorder {
	return m.recorder
}

// Sync mocks base method
func (m *MockRbacSyncer) Sync(arg0 context.Context, arg1 *v1.RbacSnapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync
func (mr *MockRbacSyncerMockRecorder) Sync(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockRbacSyncer)(nil).Sync), arg0, arg1)
}