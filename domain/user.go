package domain

import "context"

type User struct {
	IIN       string `json:"iin"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type UserUsecase interface {
	SignUp(ctx context.Context, user *User) error
}

type UserDBRepository interface {
	SignUp(ctx context.Context, user *User) error
}

type UserRedisRepository interface {
}
