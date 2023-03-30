// Code generated by MockGen. DO NOT EDIT.
// Source: repository/post.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "blog/api/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPostRepository is a mock of PostRepository interface.
type MockPostRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPostRepositoryMockRecorder
}

// MockPostRepositoryMockRecorder is the mock recorder for MockPostRepository.
type MockPostRepositoryMockRecorder struct {
	mock *MockPostRepository
}

// NewMockPostRepository creates a new mock instance.
func NewMockPostRepository(ctrl *gomock.Controller) *MockPostRepository {
	mock := &MockPostRepository{ctrl: ctrl}
	mock.recorder = &MockPostRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostRepository) EXPECT() *MockPostRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockPostRepository) Create(ctx context.Context, m *model.PostModel) (*model.PostModel, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, m)
	ret0, _ := ret[0].(*model.PostModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPostRepositoryMockRecorder) Create(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPostRepository)(nil).Create), ctx, m)
}

// FindAll mocks base method.
func (m *MockPostRepository) FindAll(ctx context.Context) ([]*model.PostModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].([]*model.PostModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockPostRepositoryMockRecorder) FindAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockPostRepository)(nil).FindAll), ctx)
}

// FindByCategoryID mocks base method.
func (m *MockPostRepository) FindByCategoryID(ctx context.Context, categoryID int64) ([]*model.PostModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByCategoryID", ctx, categoryID)
	ret0, _ := ret[0].([]*model.PostModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByCategoryID indicates an expected call of FindByCategoryID.
func (mr *MockPostRepositoryMockRecorder) FindByCategoryID(ctx, categoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCategoryID", reflect.TypeOf((*MockPostRepository)(nil).FindByCategoryID), ctx, categoryID)
}
