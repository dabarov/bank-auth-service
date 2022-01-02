package middleware

import (
	"fmt"
	"log"

	"github.com/dabarov/bank-auth-service/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
)

type userAuthMiddleware struct {
	userUsecase domain.UserUsecase
}

func NewCORSMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "*")
		next(ctx)
	}
}

func NewUserAuthMiddleware(userUsecase domain.UserUsecase, next fasthttp.RequestHandler) fasthttp.RequestHandler {
	middleware := &userAuthMiddleware{
		userUsecase: userUsecase,
	}
	return func(ctx *fasthttp.RequestCtx) {

		token, err := middleware.ExtractToken(ctx)
		if err != nil {
			log.Printf("Extract token error: %v", err)
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			return
		}

		iin, err := middleware.ParseToken(token)
		if err != nil {
			log.Printf("Parse token error: %v", err)
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			return
		}

		ok := middleware.FindToken(token, iin)
		if !ok {
			log.Printf("Getting token failed")
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			return
		}

		ctx.SetUserValue("userIIN", iin)

		next(ctx)
	}
}

func (u *userAuthMiddleware) FindToken(token string, iin string) bool {
	key := fmt.Sprintf("user:%s", iin)

	value, err := u.userUsecase.GetRedisValue(key)
	if err != nil {
		return false
	}

	return token == value
}

func (u *userAuthMiddleware) ExtractToken(ctx *fasthttp.RequestCtx) (token string, err error) {
	token = string(ctx.Request.Header.Cookie("AuthToken"))
	if token == "" {
		err = fmt.Errorf("authorization header not found")
		return
	}
	return
}

func (u *userAuthMiddleware) ParseToken(token string) (string, error) {
	JWTToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("failed to extract token metadata, unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(u.userUsecase.GetRedisSecret()), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := JWTToken.Claims.(jwt.MapClaims)

	var userINN string

	if ok && JWTToken.Valid {
		userINN, ok = claims["iin"].(string)
		if !ok {
			return "", fmt.Errorf("field id not found")
		}
		return userINN, nil
	}

	return "", fmt.Errorf("invalid token")
}
