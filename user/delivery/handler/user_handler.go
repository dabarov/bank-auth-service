package handler

import (
	"fmt"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/dabarov/bank-auth-service/domain"
	"github.com/dabarov/bank-auth-service/user/delivery/middleware"
	"github.com/valyala/fasthttp"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(router *fasthttprouter.Router, userUsecase domain.UserUsecase) {
	handler := &UserHandler{
		userUsecase: userUsecase,
	}

	getByIINWithAuth := middleware.NewUserAuthMiddleware(userUsecase, handler.GetUserByIIN)
	router.POST("/signup", handler.SignUp)
	router.POST("/signin", handler.SignIn)
	router.GET("/user/:iin", getByIINWithAuth)
}

func (u *UserHandler) SignUp(ctx *fasthttp.RequestCtx) {
	user := &domain.User{
		IIN:       string(ctx.FormValue("iin")),
		Login:     string(ctx.FormValue("login")),
		Password:  string(ctx.FormValue("password")),
		CreatedAt: time.Now().String(),
	}

	if err := u.userUsecase.SignUp(ctx, user); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Server error: %v", err)
	}
}

func (u *UserHandler) SignIn(ctx *fasthttp.RequestCtx) {
	login := string(ctx.FormValue("login"))
	password := string(ctx.FormValue("password"))

	token, err := u.userUsecase.SignIn(ctx, login, password)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Server error: %v", err)
	}
	cookie := fasthttp.Cookie{}
	cookie.SetKey("AuthToken")
	cookie.SetValue(token)
	ctx.Response.Header.SetCookie(&cookie)
}

func (u *UserHandler) GetUserByIIN(ctx *fasthttp.RequestCtx) {
	iin := fmt.Sprintf("%s", ctx.UserValue("iin"))
	user, err := u.userUsecase.GetUserByIIN(ctx, iin)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Server error: %v", err)
		return
	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Write(user)
}
