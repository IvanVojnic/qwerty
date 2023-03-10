// Code generated by mockery v2.18.0. DO NOT EDIT.

package service

/*
import (
	models "EFpractic2/models"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockBookAct is an autogenerated mock type for the BookAct type
type MockBookAct struct {
	mock.Mock
}

// CreateBook provides a mock function with given fields: _a0, _a1
func (_m *MockBookAct) CreateBook(_a0 context.Context, _a1 *models.Book) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Book) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBook provides a mock function with given fields: _a0, _a1
func (_m *MockBookAct) DeleteBook(_a0 context.Context, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllBooks provides a mock function with given fields: _a0
func (_m *MockBookAct) GetAllBooks(_a0 context.Context) ([]models.Book, error) {
	ret := _m.Called(_a0)

	var r0 []models.Book
	if rf, ok := ret.Get(0).(func(context.Context) []models.Book); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBook provides a mock function with given fields: _a0, _a1
func (_m *MockBookAct) GetBook(_a0 context.Context, _a1 int) (models.Book, error) {
	ret := _m.Called(_a0, _a1)

	var r0 models.Book
	if rf, ok := ret.Get(0).(func(context.Context, int) models.Book); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(models.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBookId provides a mock function with given fields: ctx, bookName
func (_m *MockBookAct) GetBookId(ctx context.Context, bookName string) (int, error) {
	ret := _m.Called(ctx, bookName)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, bookName)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, bookName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBook provides a mock function with given fields: _a0, _a1
func (_m *MockBookAct) UpdateBook(_a0 context.Context, _a1 models.Book) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Book) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockBookAct interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockBookAct creates a new instance of MockBookAct. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockBookAct(t mockConstructorTestingTNewMockBookAct) *MockBookAct {
	mock := &MockBookAct{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}*/
