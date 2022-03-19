// Code generated by MockGen. DO NOT EDIT.
// Source: k8sclient.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
)

// MockManagerer is a mock of Managerer interface.
type MockManagerer struct {
	ctrl     *gomock.Controller
	recorder *MockManagererMockRecorder
}

// MockManagererMockRecorder is the mock recorder for MockManagerer.
type MockManagererMockRecorder struct {
	mock *MockManagerer
}

// NewMockManagerer creates a new mock instance.
func NewMockManagerer(ctrl *gomock.Controller) *MockManagerer {
	mock := &MockManagerer{ctrl: ctrl}
	mock.recorder = &MockManagererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManagerer) EXPECT() *MockManagererMockRecorder {
	return m.recorder
}

// GetDeployments mocks base method.
func (m *MockManagerer) GetDeployments(ctx context.Context, namespace string) ([]v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployments", ctx, namespace)
	ret0, _ := ret[0].([]v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployments indicates an expected call of GetDeployments.
func (mr *MockManagererMockRecorder) GetDeployments(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployments", reflect.TypeOf((*MockManagerer)(nil).GetDeployments), ctx, namespace)
}

// GetDeploymentsPerLabel mocks base method.
func (m *MockManagerer) GetDeploymentsPerLabel(ctx context.Context, namespace, label, value string) ([]v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentsPerLabel", ctx, namespace, label, value)
	ret0, _ := ret[0].([]v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeploymentsPerLabel indicates an expected call of GetDeploymentsPerLabel.
func (mr *MockManagererMockRecorder) GetDeploymentsPerLabel(ctx, namespace, label, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentsPerLabel", reflect.TypeOf((*MockManagerer)(nil).GetDeploymentsPerLabel), ctx, namespace, label, value)
}