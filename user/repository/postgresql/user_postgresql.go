package postgresql

import (
	"context"

	"github.com/dabarov/bank-auth-service/domain"
	"gorm.io/gorm"
)

type userPostgresqlRepository struct {
	Conn *gorm.DB
}

func NewUserPostgresqlRepository(Conn *gorm.DB) domain.UserDBRepository {
	Conn.AutoMigrate(&domain.User{})
	Conn.Create(&domain.User{Login: "admin", Password: "admin", IIN: "000000000000", Role: domain.Admin_role})
	return &userPostgresqlRepository{Conn: Conn}
}

func (uPR *userPostgresqlRepository) SignUp(ctx context.Context, user *domain.User) error {
	var emptyUser *domain.User
	if notFound := uPR.Conn.Where(&domain.User{Login: user.Login}).First(&emptyUser).Error; notFound != nil {
	} else {
		return domain.ErrLoginTaken
	}
	if notFound := uPR.Conn.Where(&domain.User{IIN: user.IIN}).First(&emptyUser).Error; notFound != nil {
	} else {
		return domain.ErrIINTaken
	}

	if err := uPR.Conn.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (uPR *userPostgresqlRepository) SignIn(ctx context.Context, login string, password string) (string, error) {
	var currentUser *domain.User
	if notFound := uPR.Conn.Where(&domain.User{Login: login, Password: password}).First(&currentUser).Error; notFound != nil {
		return currentUser.IIN, domain.ErrInvalidLoginPassword
	}
	return currentUser.IIN, nil
}

func (uPR *userPostgresqlRepository) GetUserByIIN(ctx context.Context, iin string) (*domain.User, error) {
	var user *domain.User
	if notFound := uPR.Conn.Where(&domain.User{IIN: iin}).First(&user).Error; notFound != nil {
		return user, domain.ErrUserDoesntExist
	}
	return user, nil
}
