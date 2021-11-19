package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/dabarov/online-banking/domain"
	"github.com/valyala/fasthttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	username = "postgres"
	password = "postgres"
	hostname = "localhost"
	port     = 5432
	db       = "postgres"
)

type MyRESTServer struct {
	db *gorm.DB
}

func extractCredential(ctx *fasthttp.RequestCtx) (login string, pass string) {
	return string(ctx.FormValue("login")), string(ctx.FormValue("password"))
}

func (s *MyRESTServer) SignUp(ctx *fasthttp.RequestCtx) {
}

func (s *MyRESTServer) SignIn(ctx *fasthttp.RequestCtx) {
	login, password := extractCredential(ctx)

	var user domain.User
	s.db.Where(&domain.User{Login: login, Password: password}).First(&user)

	if user.Name == "" {
		log.Println("Invalid credentials!")
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		fmt.Fprintf(ctx, "Invalid credentials!")
		return
	}
}

func main() {
	DSN := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", username, password, hostname, port, db)
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	server := &MyRESTServer{
		db: db,
	}
	db.AutoMigrate(&domain.User{})

	router := fasthttprouter.New()
	router.POST("/signup", server.SignUp)
	router.POST("/signin", server.SignIn)
	fasthttp.ListenAndServe(":8080", router.Handler)
}
