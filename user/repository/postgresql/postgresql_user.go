package postgresql

import (
	"context"

	"github.com/dabarov/online-banking/domain"
	"gorm.io/gorm"
)

type postgresqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlArticleRepository(Conn *gorm.DB) domain.UserRepository {
	return &postgresqlUserRepository{Conn}
}

func (p *postgresqlUserRepository) GetByIIN(ctx context.Context, IIN uint64) (domain.User, error) {
	var user domain.User
	if err := p.Conn.Where(&domain.User{IIN: IIN}).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (p *postgresqlUserRepository) SignUp(ctx context.Context, user *domain.User) error {
	if err := p.Conn.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (p *postgresqlUserRepository) SignIn(ctx context.Context, user *domain.User) error {
	return nil
}
