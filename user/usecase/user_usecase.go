package usecase

import (
	"context"
	"encoding/json"

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
	err := u.userDBRepository.SignUp(ctx, user)
	return err
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
	return token, redisErr
}

func (u *userUsecase) GetUserByIIN(ctx context.Context, iin string) ([]byte, error) {
	if InvalidIIN(iin) {
		return []byte{}, domain.ErrIINIncorect
	}
	user, err := u.userDBRepository.GetUserByIIN(ctx, iin)
	if err != nil {
		return []byte{}, err
	}
	responseJSON, jsonErr := json.Marshal(user)
	return responseJSON, jsonErr
}

func (u *userUsecase) GetRedisValue(key string) (string, error) {
	return u.userRedisRepository.GetValue(key)
}

func (u *userUsecase) GetRedisSecret() string {
	return u.userRedisRepository.GetSecret()
}
