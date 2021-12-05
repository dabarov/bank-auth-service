package http

import (
	"encoding/binary"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/dabarov/online-banking/domain"
	"github.com/valyala/fasthttp"
)

type UserHandler struct {
	UUsecase domain.UserUsecase
}

func NewUserUsecase(router *fasthttprouter.Router, ucase domain.UserUsecase) {
	handler := &UserHandler{
		UUsecase: ucase,
	}
	router.POST("/signup", handler.SignUp)
	router.POST("/signin", handler.SignIn)
	router.GET("/user/:iin", handler.GetByIIN)
}

func (u *UserHandler) SignUp(ctx *fasthttp.RequestCtx) {
	user := &domain.User{
		IIN:       binary.BigEndian.Uint64(ctx.FormValue("iin")),
		Login:     string(ctx.FormValue("login")),
		Password:  string(ctx.FormValue("password")),
		CreatedAt: time.Now().String(),
	}

	if err := u.UUsecase.SignUp(ctx, user); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}
}

func (u *UserHandler) SignIn(ctx *fasthttp.RequestCtx) {
}

func (u *UserHandler) GetByIIN(ctx *fasthttp.RequestCtx) {
}
