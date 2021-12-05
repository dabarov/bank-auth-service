package redis

import (
	"context"

	"github.com/dabarov/online-banking/domain"
	"github.com/go-redis/redis"
)

type redisUserRepository struct {
	redisClient *redis.Client
}

func NewMysqlArticleRepository(redisClient *redis.Client) domain.UserRepository {
	return &redisUserRepository{redisClient}
}

func (r *redisUserRepository) GetByIIN(ctx context.Context, IIN uint64) (domain.User, error) {
	return domain.User{}, nil
}

func (r *redisUserRepository) SignUp(ctx context.Context, user *domain.User) error {

	return nil
}

func (r *redisUserRepository) SignIn(ctx context.Context, user *domain.User) error {
	return nil
}
