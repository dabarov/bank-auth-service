package mocks

import (
	"context"

	mock "github.com/stretchr/testify/mock"
)

type UserRedisRepository struct {
	mock.Mock
}

func (_m *UserRedisRepository) GetAccessToken(ctx context.Context, iin string) (string, error) {
	ret := _m.Called(ctx, iin)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, iin)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, iin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *UserRedisRepository) InsertToken(token string, iin string) error {
	ret := _m.Called(token, iin)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(token, iin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *UserRedisRepository) GetValue(key string) (string, error) {
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

func (_m *UserRedisRepository) GetSecret() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
