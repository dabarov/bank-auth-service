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

func (u *userRedisRepository) ParseToken(token string) (int64, error) {
	JWTToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("failed to extract token metadata, unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(u.secret), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := JWTToken.Claims.(jwt.MapClaims)

	var userId float64

	if ok && JWTToken.Valid {
		userId, ok = claims["id"].(float64)
		if !ok {
			return 0, fmt.Errorf("field id not found")
		}
		return int64(userId), nil
	}

	return 0, fmt.Errorf("invalid token")
}

func (u *userRedisRepository) FindToken(token string, iin string) bool {
	key := fmt.Sprintf("user:%s", iin)

	value, err := u.redisClient.Get(key).Result()
	if err != nil {
		return false
	}

	return token == value
}

func (u *userRedisRepository) InsertToken(token string, iin string) error {
	key := fmt.Sprintf("user:%s", iin)
	return u.redisClient.Set(key, token, u.timeout).Err()
}
