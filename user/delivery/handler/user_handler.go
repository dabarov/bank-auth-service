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
	router.POST("/signin", handler.SignIn)
	router.GET("/user/:iin", handler.GetUserByIIN)
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

func (u *UserHandler) SignIn(ctx *fasthttp.RequestCtx) {
	login := string(ctx.FormValue("login"))
	password := string(ctx.FormValue("password"))

	token, err := u.UserUsecase.SignIn(ctx, login, password)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Server error: %v", err)
	}

	ctx.Response.Header.Set("auth", token)
}

func (u *UserHandler) GetUserByIIN(ctx *fasthttp.RequestCtx) {
	iin := fmt.Sprintf("%s", ctx.UserValue("iin"))
	user, err := u.UserUsecase.GetUserByIIN(ctx, iin)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Server error: %v", err)
		return
	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Write(user)
}
