package domain

import (
	"context"
)

type User struct {
	IIN       string `json:"iin"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type UserUsecase interface {
	SignUp(ctx context.Context, user *User) error
	SignIn(ctx context.Context, login string, password string) (string, error)
	GetUserByIIN(ctx context.Context, iin string) ([]byte, error)
	GetRedisValue(key string) (string, error)
	GetRedisSecret() string
}

type UserDBRepository interface {
	SignUp(ctx context.Context, user *User) error
	SignIn(ctx context.Context, login string, password string) (string, error)
	GetUserByIIN(ctx context.Context, iin string) (*User, error)
}

type UserRedisRepository interface {
	GetAccessToken(ctx context.Context, iin string) (string, error)
	InsertToken(token string, iin string) error
	GetValue(key string) (string, error)
	GetSecret() string
}
