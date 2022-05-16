// Code generated by MockGen. DO NOT EDIT.
// Source: ./defs.go

// Package repositories is a generated GoMock package.
package repositories

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBooksRepository is a mock of BooksRepository interface.
type MockBooksRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBooksRepositoryMockRecorder
}

// MockBooksRepositoryMockRecorder is the mock recorder for MockBooksRepository.
type MockBooksRepositoryMockRecorder struct {
	mock *MockBooksRepository
}

// NewMockBooksRepository creates a new mock instance.
func NewMockBooksRepository(ctrl *gomock.Controller) *MockBooksRepository {
	mock := &MockBooksRepository{ctrl: ctrl}
	mock.recorder = &MockBooksRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBooksRepository) EXPECT() *MockBooksRepositoryMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockBooksRepository) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockBooksRepositoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockBooksRepository)(nil).Close))
}

// SearchByAuthor mocks base method.
func (m *MockBooksRepository) SearchByAuthor(arg0 context.Context, arg1 string) ([]Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchByAuthor", arg0, arg1)
	ret0, _ := ret[0].([]Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchByAuthor indicates an expected call of SearchByAuthor.
func (mr *MockBooksRepositoryMockRecorder) SearchByAuthor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchByAuthor", reflect.TypeOf((*MockBooksRepository)(nil).SearchByAuthor), arg0, arg1)
}

// SearchByContent mocks base method.
func (m *MockBooksRepository) SearchByContent(arg0 context.Context, arg1 string) ([]Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchByContent", arg0, arg1)
	ret0, _ := ret[0].([]Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchByContent indicates an expected call of SearchByContent.
func (mr *MockBooksRepositoryMockRecorder) SearchByContent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchByContent", reflect.TypeOf((*MockBooksRepository)(nil).SearchByContent), arg0, arg1)
}

// SearchByTitle mocks base method.
func (m *MockBooksRepository) SearchByTitle(arg0 context.Context, arg1 string) ([]Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchByTitle", arg0, arg1)
	ret0, _ := ret[0].([]Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchByTitle indicates an expected call of SearchByTitle.
func (mr *MockBooksRepositoryMockRecorder) SearchByTitle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchByTitle", reflect.TypeOf((*MockBooksRepository)(nil).SearchByTitle), arg0, arg1)
}