package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/buaazp/fasthttprouter"
	"github.com/dabarov/bank-auth-service/user/delivery/handler"
	"github.com/dabarov/bank-auth-service/user/repository/postgresql"
	"github.com/dabarov/bank-auth-service/user/usecase"
	"github.com/valyala/fasthttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbusername := os.Getenv("POSTGRES_USERNAME")
	dbpassword := os.Getenv("POSTGRES_PASSWORD")
	dbhostname := os.Getenv("POSTGRES_HOSTNAME")
	dbport, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	dbname := os.Getenv("POSTGRES_DB")
	httpport := os.Getenv("HTTP_PORT")

	DSN := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", dbusername, dbpassword, dbhostname, dbport, dbname)
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	dbRepository := postgresql.NewUserPostgresqlRepository(db)
	userUsecase := usecase.NewUserUsecase(dbRepository)
	router := fasthttprouter.New()
	handler.NewUserHandler(router, userUsecase)
	log.Fatal(fasthttp.ListenAndServe(":"+httpport, router.Handler))
}
