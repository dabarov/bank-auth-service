package mocks

import (
	"context"

	domain "github.com/dabarov/bank-auth-service/domain"
	mock "github.com/stretchr/testify/mock"
)

type UserDBRepository struct {
	mock.Mock
}

func (_m *UserDBRepository) SignUp(ctx context.Context, user *domain.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *UserDBRepository) SignIn(ctx context.Context, login string, password string) (string, error) {
	ret := _m.Called(ctx, login, password)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, login, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, login, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *UserDBRepository) GetUserByIIN(ctx context.Context, iin string) ([]byte, error) {
	ret := _m.Called(ctx, iin)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, string) []byte); ok {
		r0 = rf(ctx, iin)
	} else {
		r0 = ret.Get(0).([]byte)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, iin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
