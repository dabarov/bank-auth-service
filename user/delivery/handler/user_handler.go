package handler

import (
	"fmt"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/dabarov/bank-auth-service/domain"
	"github.com/valyala/fasthttp"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(router *fasthttprouter.Router, userUsecase domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: userUsecase,
	}
	router.POST("/signup", handler.SignUp)
}

func (u *UserHandler) SignUp(ctx *fasthttp.RequestCtx) {
	user := &domain.User{
		IIN:       string(ctx.FormValue("iin")),
		Login:     string(ctx.FormValue("login")),
		Password:  string(ctx.FormValue("password")),
		CreatedAt: time.Now().String(),
	}

	if err := u.UserUsecase.SignUp(ctx, user); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Server error: %v", err)
	}
}
