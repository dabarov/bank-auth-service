package usecase_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/dabarov/bank-auth-service/domain"
	"github.com/dabarov/bank-auth-service/domain/mocks"
	"github.com/dabarov/bank-auth-service/user/usecase"
	"github.com/stretchr/testify/mock"
)

func TestSignUp(t *testing.T) {
	dbRepo := new(mocks.UserDBRepository)
	redisRepo := new(mocks.UserRedisRepository)

	dbRepo.On("SignUp", mock.Anything, mock.Anything).Return(nil).Once()
	mockUser := &domain.User{
		IIN:       "123123123123",
		Login:     "login",
		Password:  "password",
		Role:      domain.User_role,
		CreatedAt: "whens",
	}
	mockContext := context.Background()
	mockUseCase := usecase.NewUserUsecase(dbRepo, redisRepo)

	if err := mockUseCase.SignUp(mockContext, mockUser); err != nil {
		t.Fatalf("incorrect signup %v", err)
	}
}

func TestSignUpWrongIINFormat(t *testing.T) {
	dbRepo := new(mocks.UserDBRepository)
	redisRepo := new(mocks.UserRedisRepository)

	dbRepo.On("SignUp", mock.Anything, mock.Anything).Return(nil).Once()
	mockUser := &domain.User{
		IIN:       "--",
		Login:     "login",
		Password:  "password",
		Role:      domain.User_role,
		CreatedAt: "whens",
	}
	mockContext := context.Background()
	mockUseCase := usecase.NewUserUsecase(dbRepo, redisRepo)

	if err := mockUseCase.SignUp(mockContext, mockUser); err != domain.ErrIINIncorect {
		t.Fatalf("should return iin incorrect error but: %v", err)
	}
}

func TestSignUpEmptyPassword(t *testing.T) {
	dbRepo := new(mocks.UserDBRepository)
	redisRepo := new(mocks.UserRedisRepository)

	dbRepo.On("SignUp", mock.Anything, mock.Anything).Return(nil).Once()
	mockUser := &domain.User{
		IIN:       "132123123123",
		Login:     "login",
		Password:  "",
		Role:      domain.User_role,
		CreatedAt: "whens",
	}
	mockContext := context.Background()
	mockUseCase := usecase.NewUserUsecase(dbRepo, redisRepo)

	if err := mockUseCase.SignUp(mockContext, mockUser); err != domain.ErrEmptyField {
		t.Fatalf("should return empty field error but: %v", err)
	}
}

func TestSignIn(t *testing.T) {
	dbRepo := new(mocks.UserDBRepository)
	redisRepo := new(mocks.UserRedisRepository)
	mockUser := &domain.User{
		IIN:       "123123123123",
		Login:     "login",
		Password:  "password",
		Role:      domain.User_role,
		CreatedAt: "whens",
	}

	dbRepo.On("SignIn", mock.Anything, mock.Anything, mock.Anything).Return(mockUser.IIN, nil).Once()
	redisRepo.On("GetAccessToken", mock.Anything, mock.Anything).Return("token", nil).Once()
	mockContext := context.Background()
	mockUseCase := usecase.NewUserUsecase(dbRepo, redisRepo)

	if _, err := mockUseCase.SignIn(mockContext, mockUser.Login, mockUser.Password); err != nil {
		t.Fatalf("incorrect signin %v", err)
	}
}

func TestSignInEmptyPassword(t *testing.T) {
	dbRepo := new(mocks.UserDBRepository)
	redisRepo := new(mocks.UserRedisRepository)
	mockUser := &domain.User{
		IIN:       "123123123123",
		Login:     "login",
		Password:  "",
		Role:      domain.User_role,
		CreatedAt: "whens",
	}

	dbRepo.On("SignIn", mock.Anything, mock.Anything, mock.Anything).Return(mockUser.IIN, nil).Once()
	redisRepo.On("GetAccessToken", mock.Anything, mock.Anything).Return("token", nil).Once()
	mockContext := context.Background()
	mockUseCase := usecase.NewUserUsecase(dbRepo, redisRepo)

	if _, err := mockUseCase.SignIn(mockContext, mockUser.Login, mockUser.Password); err != domain.ErrEmptyField {
		t.Fatalf("incorrect signin %v", err)
	}
}

func TestSignInDBError(t *testing.T) {
	dbRepo := new(mocks.UserDBRepository)
	redisRepo := new(mocks.UserRedisRepository)
	mockUser := &domain.User{
		IIN:       "123123123123",
		Login:     "login",
		Password:  "password",
		Role:      domain.User_role,
		CreatedAt: "whens",
	}

	dbRepo.On("SignIn", mock.Anything, mock.Anything, mock.Anything).Return(mockUser.IIN, fmt.Errorf("someError")).Once()
	redisRepo.On("GetAccessToken", mock.Anything, mock.Anything).Return("token", nil).Once()
	mockContext := context.Background()
	mockUseCase := usecase.NewUserUsecase(dbRepo, redisRepo)

	if _, err := mockUseCase.SignIn(mockContext, mockUser.Login, mockUser.Password); err == nil {
		t.Fatalf("should have gotten err from db")
	}
}

func TestGetRedisSecret(t *testing.T) {
	dbRepo := new(mocks.UserDBRepository)
	redisRepo := new(mocks.UserRedisRepository)
	redisRepo.On("GetSecret").Return("secret").Once()
	mockUseCase := usecase.NewUserUsecase(dbRepo, redisRepo)
	if secret := mockUseCase.GetRedisSecret(); secret != "secret" {
		t.Fatalf("should have gotten secret")
	}
}

func TestGetRedisValue(t *testing.T) {
	dbRepo := new(mocks.UserDBRepository)
	redisRepo := new(mocks.UserRedisRepository)
	key := "key"
	value := "val"
	redisRepo.On("GetValue", key).Return(value, nil).Once()
	mockUseCase := usecase.NewUserUsecase(dbRepo, redisRepo)
	if valCheck, err := mockUseCase.GetRedisValue(key); valCheck != value || err != nil {
		t.Fatalf("should have gotten value")
	}
}

func TestGetUser(t *testing.T) {
	dbRepo := new(mocks.UserDBRepository)
	redisRepo := new(mocks.UserRedisRepository)
	mockUser := &domain.User{
		IIN:       "123123123123",
		Login:     "login",
		Password:  "password",
		Role:      domain.User_role,
		CreatedAt: "whens",
	}
	mockContext := context.Background()
	dbRepo.On("GetUserByIIN", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
	mockUseCase := usecase.NewUserUsecase(dbRepo, redisRepo)
	if _, jsonErr := mockUseCase.GetUser(mockContext, mockUser.IIN); jsonErr != nil {
		t.Fatalf("unsuccessful get user")
	}
}

func TestGetUserInvalidIIN(t *testing.T) {
	dbRepo := new(mocks.UserDBRepository)
	redisRepo := new(mocks.UserRedisRepository)
	mockUser := &domain.User{
		IIN:       "123",
		Login:     "login",
		Password:  "password",
		Role:      domain.User_role,
		CreatedAt: "whens",
	}
	mockContext := context.Background()
	mockUseCase := usecase.NewUserUsecase(dbRepo, redisRepo)
	if _, jsonErr := mockUseCase.GetUser(mockContext, mockUser.IIN); jsonErr != domain.ErrIINIncorect {
		t.Fatalf("passed incorrect iin")
	}
}

func TestGetUserDBError(t *testing.T) {
	dbRepo := new(mocks.UserDBRepository)
	redisRepo := new(mocks.UserRedisRepository)
	mockUser := &domain.User{
		IIN:       "123123123123",
		Login:     "login",
		Password:  "password",
		Role:      domain.User_role,
		CreatedAt: "whens",
	}
	mockContext := context.Background()
	dbRepo.On("GetUserByIIN", mock.Anything, mock.Anything).Return(mockUser, fmt.Errorf("some error")).Once()
	mockUseCase := usecase.NewUserUsecase(dbRepo, redisRepo)
	if _, jsonErr := mockUseCase.GetUser(mockContext, mockUser.IIN); jsonErr == nil {
		t.Fatalf("should have gotten error from DB")
	}
}

func TestGetUserByIIN(t *testing.T) {
	dbRepo := new(mocks.UserDBRepository)
	redisRepo := new(mocks.UserRedisRepository)
	mockUser := &domain.User{
		IIN:       "123123123123",
		Login:     "login",
		Password:  "password",
		Role:      domain.Admin_role,
		CreatedAt: "whens",
	}
	mockContext := context.Background()

	dbRepo.On("GetUserByIIN", mock.Anything, mock.Anything).Return(mockUser, nil)
	mockUseCase := usecase.NewUserUsecase(dbRepo, redisRepo)

	if _, err := mockUseCase.GetUserByIIN(mockContext, mockUser.IIN, mockUser.IIN); err != nil {
		t.Fatal("unsuccessful get user by iin")
	}
}

func TestGetUserByIINInvalidIIN(t *testing.T) {
	dbRepo := new(mocks.UserDBRepository)
	redisRepo := new(mocks.UserRedisRepository)
	mockUser := &domain.User{
		IIN:       "123123123123",
		Login:     "login",
		Password:  "password",
		Role:      domain.Admin_role,
		CreatedAt: "whens",
	}
	mockContext := context.Background()
	mockUseCase := usecase.NewUserUsecase(dbRepo, redisRepo)

	if _, err := mockUseCase.GetUserByIIN(mockContext, mockUser.IIN, "invalidIIN"); err != domain.ErrIINIncorect {
		t.Fatalf("wrong error %v", err)
	}
}

func TestGetUserByIINNoAccess(t *testing.T) {
	dbRepo := new(mocks.UserDBRepository)
	redisRepo := new(mocks.UserRedisRepository)
	mockUser := &domain.User{
		IIN:       "123123123123",
		Login:     "login",
		Password:  "password",
		Role:      domain.User_role,
		CreatedAt: "whens",
	}
	mockContext := context.Background()
	dbRepo.On("GetUserByIIN", mock.Anything, mock.Anything).Return(mockUser, nil)
	mockUseCase := usecase.NewUserUsecase(dbRepo, redisRepo)

	if _, err := mockUseCase.GetUserByIIN(mockContext, mockUser.IIN, "999999999999"); err != domain.ErrNoAccessToRequestedIIN {
		t.Fatalf("wrong error %v", err)
	}
}

func TestGetUserByIINDBError(t *testing.T) {
	dbRepo := new(mocks.UserDBRepository)
	redisRepo := new(mocks.UserRedisRepository)
	mockUser := &domain.User{
		IIN:       "123123123123",
		Login:     "login",
		Password:  "password",
		Role:      domain.Admin_role,
		CreatedAt: "whens",
	}
	requestedIIN := "999999999999"
	dbError := errors.New("DB error")
	mockContext := context.Background()
	dbRepo.On("GetUserByIIN", mock.Anything, mock.Anything).Return(mockUser, dbError).Once()
	dbRepo.On("GetUserByIIN", mock.Anything, mock.Anything).Return(mockUser, nil)
	mockUseCase := usecase.NewUserUsecase(dbRepo, redisRepo)

	if _, err := mockUseCase.GetUserByIIN(mockContext, mockUser.IIN, requestedIIN); err != dbError {
		t.Fatalf("should get db error %v", err)
	}
}

func TestGetUserByIINSecondDBError(t *testing.T) {
	dbRepo := new(mocks.UserDBRepository)
	redisRepo := new(mocks.UserRedisRepository)
	mockUser := &domain.User{
		IIN:       "123123123123",
		Login:     "login",
		Password:  "password",
		Role:      domain.Admin_role,
		CreatedAt: "whens",
	}
	requestedIIN := "999999999999"
	dbError := errors.New("DB error")
	mockContext := context.Background()
	dbRepo.On("GetUserByIIN", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
	dbRepo.On("GetUserByIIN", mock.Anything, mock.Anything).Return(mockUser, dbError)
	mockUseCase := usecase.NewUserUsecase(dbRepo, redisRepo)

	if _, err := mockUseCase.GetUserByIIN(mockContext, mockUser.IIN, requestedIIN); err != dbError {
		t.Fatalf("should get db error %v", err)
	}
}
