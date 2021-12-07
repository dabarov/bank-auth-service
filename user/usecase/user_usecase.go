package usecase

import (
	"context"

	"github.com/dabarov/bank-auth-service/domain"
)

type userUsecase struct {
	userDBRepository    domain.UserDBRepository
	userRedisRepository domain.UserRedisRepository
}

func NewUserUsecase(uDBR domain.UserDBRepository, uRR domain.UserRedisRepository) domain.UserUsecase {
	return &userUsecase{
		userDBRepository:    uDBR,
		userRedisRepository: uRR,
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

func (u *userUsecase) SignIn(ctx context.Context, login string, password string) (string, error) {
	if InvalidField(login) || InvalidField(password) {
		return "", domain.ErrEmptyField
	}

	iin, dbErr := u.userDBRepository.SignIn(ctx, login, password)
	if dbErr != nil {
		return "", dbErr
	}

	token, redisErr := u.userRedisRepository.GetAccessToken(ctx, iin)
	if redisErr != nil {
		return token, redisErr
	}
	return token, nil
}
