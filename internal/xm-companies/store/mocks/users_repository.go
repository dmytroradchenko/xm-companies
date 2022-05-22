// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"
	model "xm-companies/internal/xm-companies/model"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// UsersRepository is an autogenerated mock type for the UsersRepository type
type UsersRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, user
func (_m *UsersRepository) Create(ctx context.Context, user *model.User) (string, error) {
	ret := _m.Called(ctx, user)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) string); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, userName
func (_m *UsersRepository) Find(ctx context.Context, userName string) (*model.User, error) {
	ret := _m.Called(ctx, userName)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.User); ok {
		r0 = rf(ctx, userName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUsersRepository creates a new instance of UsersRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsersRepository(t testing.TB) *UsersRepository {
	mock := &UsersRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
