package usecase

import (
	"context"
	"time"

	"github.com/dabarov/online-banking/domain"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(u domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: u,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) GetByIIN(ctx context.Context, IIN uint64) (domain.User, error) {
	user, err := u.userRepository.GetByIIN(ctx, IIN)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userUsecase) SignUp(ctx context.Context, user *domain.User) error {
	return u.userRepository.SignUp(ctx, user)
}

func (u *userUsecase) SignIn(ctx context.Context, user *domain.User) error {
	return u.userRepository.SignIn(ctx, user)
}
