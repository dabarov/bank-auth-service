package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/dabarov/bank-auth-service/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
)

type userRedisRepository struct {
	redisClient *redis.Client
	timeout     time.Duration
	secret      string
}

func NewUserRedisRepository(redisClient *redis.Client, timeout int, secret string) domain.UserRedisRepository {
	return &userRedisRepository{
		redisClient: redisClient,
		timeout:     time.Duration(timeout) * time.Second,
		secret:      secret,
	}
}

func (u *userRedisRepository) GetAccessToken(ctx context.Context, iin string) (string, error) {
	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["iin"] = iin
	accessTokenClaims["iat"] = time.Now().Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	signedToken, err := accessToken.SignedString([]byte(u.secret))
	if err != nil {
		return signedToken, err
	}

	if err := u.InsertToken(signedToken, iin); err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (u *userRedisRepository) InsertToken(token string, iin string) error {
	key := fmt.Sprintf("user:%s", iin)
	return u.redisClient.Set(key, token, u.timeout).Err()
}

func (u *userRedisRepository) GetValue(key string) (string, error) {
	return u.redisClient.Get(key).Result()
}

func (u *userRedisRepository) GetSecret() string {
	return u.secret
}
