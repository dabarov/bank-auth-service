package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/dabarov/online-banking/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
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
	db        *gorm.DB
	secret    string
	expTime   time.Duration
	redisConn *redis.Client
}

func extractCredential(ctx *fasthttp.RequestCtx) (login string, pass string) {
	return string(ctx.FormValue("login")), string(ctx.FormValue("password"))
}

func (s *MyRESTServer) insertToken(token string, iin uint64) error {
	key := fmt.Sprintf("user:%d", iin)
	return s.redisConn.Set(key, token, s.expTime).Err()
}

func (s *MyRESTServer) SignUp(ctx *fasthttp.RequestCtx) {
	s.db.Create(&domain.User{
		IIN:       binary.BigEndian.Uint64(ctx.FormValue("iin")),
		Login:     string(ctx.FormValue("login")),
		Password:  string(ctx.FormValue("password")),
		Name:      string(ctx.FormValue("name")),
		Surname:   string(ctx.FormValue("surname")),
		Phone:     string(ctx.FormValue("phone")),
		Role:      string(ctx.FormValue("password")),
		CreatedAt: time.Now().String(),
	})
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

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["id"] = user.IIN
	accessTokenClaims["iat"] = time.Now().Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	signedToken, err := accessToken.SignedString([]byte(s.secret))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Coundn't create token. Error: %v", err)
		return
	}

	if err := s.insertToken(signedToken, user.IIN); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "Coundn't insert token in redis. Error: %v", err)
		return
	}

	fmt.Fprintf(ctx, "Token: %s", signedToken)
}

func main() {
	DSN := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", username, password, hostname, port, db)
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Ping error: %v", err)
	}
	log.Println(pong)

	server := &MyRESTServer{
		db:        db,
		secret:    "18DbJX9NR0WApJtB9OgmQkdlmHLwaHpK",
		expTime:   20 * time.Second,
		redisConn: client,
	}
	db.AutoMigrate(&domain.User{})

	router := fasthttprouter.New()
	router.POST("/signup", server.SignUp)
	router.POST("/signin", server.SignIn)
	fasthttp.ListenAndServe(":8080", router.Handler)
}
