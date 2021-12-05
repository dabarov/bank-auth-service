package domain

import "context"

type User struct {
	IIN       uint64 `json:"iin"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type UserUsecase interface {
	GetByIIN(ctx context.Context, IIN uint64) (User, error)
	SignUp(ctx context.Context, user *User) error
	SignIn(ctx context.Context, user *User) error
}

type UserRepository interface {
	GetByIIN(ctx context.Context, IIN uint64) (User, error)
	SignUp(ctx context.Context, user *User) error
	SignIn(ctx context.Context, user *User) error
}
