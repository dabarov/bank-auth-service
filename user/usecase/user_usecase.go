package usecase

import (
	"context"

	"github.com/dabarov/bank-auth-service/domain"
)

type userUsecase struct {
	userDBRepository domain.UserDBRepository
}

func NewUserUsecase(uDBR domain.UserDBRepository) domain.UserUsecase {
	return &userUsecase{
		userDBRepository: uDBR,
	}
}

func (u *userUsecase) SignUp(ctx context.Context, user *domain.User) error {
	if InvalidIIN(user.IIN) {
		return domain.ErrIINIncorect
	}
	if InvalidField(user.Login) || InvalidField(user.Password) {
		return domain.ErrEmptyField
	}
	return u.userDBRepository.SignUp(ctx, user)
}
