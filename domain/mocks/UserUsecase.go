package mocks

import (
	"context"

	domain "github.com/dabarov/bank-auth-service/domain"
	mock "github.com/stretchr/testify/mock"
)

type UserUsecase struct {
	mock.Mock
}

func (_m *UserUsecase) SignUp(ctx context.Context, user *domain.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *UserUsecase) SignIn(ctx context.Context, login string, password string) (string, error) {
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

func (_m *UserUsecase) GetUserByIIN(ctx context.Context, requestedIIN string, currentUserIIN string) ([]byte, error) {
	ret := _m.Called(ctx, requestedIIN, currentUserIIN)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []byte); ok {
		r0 = rf(ctx, requestedIIN, currentUserIIN)
	} else {
		r0 = ret.Get(0).([]byte)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, requestedIIN, currentUserIIN)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *UserUsecase) GetUser(ctx context.Context, iin string) ([]byte, error) {
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

func (_m *UserUsecase) GetRedisValue(key string) (string, error) {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *UserUsecase) GetRedisSecret() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
