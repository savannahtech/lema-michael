// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dilly3/houdini/internal/storage (interfaces: ICommitStore)
//
// Generated by this command:
//
//	mockgen -destination=api/server/mocks/commit_mock.go -package=mocks github.com/dilly3/houdini/internal/storage ICommitStore
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/dilly3/houdini/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockICommitStore is a mock of ICommitStore interface.
type MockICommitStore struct {
	ctrl     *gomock.Controller
	recorder *MockICommitStoreMockRecorder
}

// MockICommitStoreMockRecorder is the mock recorder for MockICommitStore.
type MockICommitStoreMockRecorder struct {
	mock *MockICommitStore
}

// NewMockICommitStore creates a new mock instance.
func NewMockICommitStore(ctrl *gomock.Controller) *MockICommitStore {
	mock := &MockICommitStore{ctrl: ctrl}
	mock.recorder = &MockICommitStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICommitStore) EXPECT() *MockICommitStoreMockRecorder {
	return m.recorder
}

// GetCommitByID mocks base method.
func (m *MockICommitStore) GetCommitByID(arg0 context.Context, arg1 string) (*model.CommitInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommitByID", arg0, arg1)
	ret0, _ := ret[0].(*model.CommitInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommitByID indicates an expected call of GetCommitByID.
func (mr *MockICommitStoreMockRecorder) GetCommitByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommitByID", reflect.TypeOf((*MockICommitStore)(nil).GetCommitByID), arg0, arg1)
}

// GetCommitsByRepoName mocks base method.
func (m *MockICommitStore) GetCommitsByRepoName(arg0 context.Context, arg1 string, arg2 int) ([]model.CommitInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommitsByRepoName", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.CommitInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommitsByRepoName indicates an expected call of GetCommitsByRepoName.
func (mr *MockICommitStoreMockRecorder) GetCommitsByRepoName(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommitsByRepoName", reflect.TypeOf((*MockICommitStore)(nil).GetCommitsByRepoName), arg0, arg1, arg2)
}

// SaveCommit mocks base method.
func (m *MockICommitStore) SaveCommit(arg0 context.Context, arg1 *model.CommitInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCommit", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCommit indicates an expected call of SaveCommit.
func (mr *MockICommitStoreMockRecorder) SaveCommit(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCommit", reflect.TypeOf((*MockICommitStore)(nil).SaveCommit), arg0, arg1)
}

// SaveCommits mocks base method.
func (m *MockICommitStore) SaveCommits(arg0 context.Context, arg1 []model.CommitInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCommits", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCommits indicates an expected call of SaveCommits.
func (mr *MockICommitStoreMockRecorder) SaveCommits(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCommits", reflect.TypeOf((*MockICommitStore)(nil).SaveCommits), arg0, arg1)
}