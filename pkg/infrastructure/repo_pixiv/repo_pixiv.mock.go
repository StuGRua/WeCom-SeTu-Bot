// Code generated by MockGen. DO NOT EDIT.
// Source: repo_pixiv.go

// Package repo_pixiv is a generated GoMock package.
package repo_pixiv

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepoPixiv is a mock of RepoPixiv interface.
type MockRepoPixiv struct {
	ctrl     *gomock.Controller
	recorder *MockRepoPixivMockRecorder
}

// MockRepoPixivMockRecorder is the mock recorder for MockRepoPixiv.
type MockRepoPixivMockRecorder struct {
	mock *MockRepoPixiv
}

// NewMockRepoPixiv creates a new mock instance.
func NewMockRepoPixiv(ctrl *gomock.Controller) *MockRepoPixiv {
	mock := &MockRepoPixiv{ctrl: ctrl}
	mock.recorder = &MockRepoPixivMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepoPixiv) EXPECT() *MockRepoPixivMockRecorder {
	return m.recorder
}

// FetchPixivPictureToMem mocks base method.
func (m *MockRepoPixiv) FetchPixivPictureToMem(ctx context.Context, url string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchPixivPictureToMem", ctx, url)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchPixivPictureToMem indicates an expected call of FetchPixivPictureToMem.
func (mr *MockRepoPixivMockRecorder) FetchPixivPictureToMem(ctx, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchPixivPictureToMem", reflect.TypeOf((*MockRepoPixiv)(nil).FetchPixivPictureToMem), ctx, url)
}
