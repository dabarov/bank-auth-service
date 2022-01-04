package postgresql_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dabarov/bank-auth-service/domain"
	"github.com/dabarov/bank-auth-service/user/repository/postgresql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SignUp tests:
func TestSuccessSignUp(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := &domain.User{
		Login:     "login",
		Password:  "password",
		Role:      domain.User_role,
		IIN:       "123123123123",
		CreatedAt: "heello",
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent)})
	repo := postgresql.NewUserPostgresqlRepository(gormDB)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm connection", err)
	}

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users" ("iin","login","password","role","created_at") VALUES ($1,$2,$3,$4,$5)`)).
		WithArgs(user.IIN, user.Login, user.Password, user.Role, user.CreatedAt).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	testCtx := context.Background()
	if err := repo.SignUp(testCtx, user); err != nil {
		t.Fatal("should be sign up")
	}
}

func TestUnsuccessSignUp(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := &domain.User{
		Login:     "login",
		Password:  "password",
		Role:      domain.User_role,
		IIN:       "123123123123",
		CreatedAt: "heello",
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent)})
	repo := postgresql.NewUserPostgresqlRepository(gormDB)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm connection", err)
	}

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users" ("iin","login","password","role","created_at") VALUES ($1,$2,$3,$4,$5)`)).
		WithArgs(user.IIN, user.Login, user.Password, user.Role, user.CreatedAt).
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectCommit()

	testCtx := context.Background()
	if err := repo.SignUp(testCtx, user); err == nil {
		t.Fatal("should not been signed up")
	}
}

func TestUsedLoginSignUp(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := &domain.User{
		Login:     "login",
		Password:  "password",
		Role:      domain.User_role,
		IIN:       "123123123123",
		CreatedAt: "heello",
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent)})

	repo := postgresql.NewUserPostgresqlRepository(gormDB)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm connection", err)
	}

	row := sqlmock.NewRows([]string{"iin", "login", "password", "role", "created_at"})
	row.AddRow("iin", "login", "password", "role", "created_at")
	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WithArgs(user.Login).WillReturnRows(row)
	mock.ExpectCommit()

	testCtx := context.Background()
	if err := repo.SignUp(testCtx, user); err != domain.ErrLoginTaken {
		t.Fatal(err)
	}
}

func TestUsedIINSignUp(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := &domain.User{
		Login:     "login",
		Password:  "password",
		Role:      domain.User_role,
		IIN:       "123123123123",
		CreatedAt: "heello",
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent)})

	repo := postgresql.NewUserPostgresqlRepository(gormDB)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm connection", err)
	}

	row := sqlmock.NewRows([]string{"iin", "login", "password", "role", "created_at"})
	row.AddRow("iin", "different", "password", "role", "created_at")
	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WithArgs(user.Login)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WithArgs(user.IIN).WillReturnRows(row)
	mock.ExpectCommit()

	testCtx := context.Background()
	if err := repo.SignUp(testCtx, user); err != domain.ErrIINTaken {
		t.Fatal(err)
	}
}

func TestSuccessSignIn(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := &domain.User{
		Login:    "login",
		Password: "password",
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent)})
	repo := postgresql.NewUserPostgresqlRepository(gormDB)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm connection", err)
	}

	row := sqlmock.NewRows([]string{"iin", "login", "password", "role", "created_at"})
	row.AddRow("iin", "different", "password", "role", "created_at")
	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WithArgs(user.Login, user.Password).WillReturnRows(row)
	mock.ExpectCommit()

	testCtx := context.Background()
	if _, err := repo.SignIn(testCtx, user.Login, user.Password); err != nil {
		t.Fatal("should have logged in")
	}
}

func TestUnsuccessSignIn(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := &domain.User{
		Login:    "login",
		Password: "password",
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent)})
	repo := postgresql.NewUserPostgresqlRepository(gormDB)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm connection", err)
	}

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WithArgs(user.Login, user.Password)
	mock.ExpectCommit()

	testCtx := context.Background()
	if _, err := repo.SignIn(testCtx, user.Login, user.Password); err == nil {
		t.Fatal("should not have logged in")
	}
}

func TestSuccessGetByIIN(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := &domain.User{
		Login:    "login",
		Password: "password",
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent)})
	repo := postgresql.NewUserPostgresqlRepository(gormDB)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm connection", err)
	}

	row := sqlmock.NewRows([]string{"iin", "login", "password", "role", "created_at"})
	row.AddRow("iin", "different", "password", "role", "created_at")
	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).WillReturnRows(row)
	mock.ExpectCommit()

	testCtx := context.Background()
	if _, err := repo.GetUserByIIN(testCtx, user.IIN); err != nil {
		t.Fatal(err)
		t.Fatal("should have found user")
	}
}

func TestUnsuccessGetByIIN(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := &domain.User{
		Login:    "login",
		Password: "password",
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent)})
	repo := postgresql.NewUserPostgresqlRepository(gormDB)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm connection", err)
	}

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT"))
	mock.ExpectCommit()

	testCtx := context.Background()
	if _, err := repo.GetUserByIIN(testCtx, user.IIN); err == nil {
		t.Fatal("should not have found user")
	}
}
