package redis_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/dabarov/bank-auth-service/domain"
	"github.com/dabarov/bank-auth-service/user/repository/redis"
	goredis "github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
)

var (
	r        domain.UserRedisRepository
	db       *goredis.Client
	mock     redismock.ClientMock
	secret   = "secret"
	exp_time = 10
)

func TestMain(m *testing.M) {
	db, mock = redismock.NewClientMock()
	r = redis.NewUserRedisRepository(db, exp_time, secret)
	os.Exit(m.Run())
}
func TestGetValue(t *testing.T) {
	key := "key"
	val := "val"
	mock.ExpectGet(key).SetVal(val)
	if value, err := r.GetValue(key); err != nil || value != val {
		t.Fatalf("got error on get value: %s", err)
	}
}

func TestGetSecret(t *testing.T) {
	if value := r.GetSecret(); value != secret {
		t.Fatalf("did not get proper secret")
	}
}

func TestInsertToken(t *testing.T) {
	iin := "user:iin"
	var token interface{} = "sometoken"
	mock.ExpectSet(iin, token, time.Duration(exp_time)*time.Second).SetVal("success")
	if err := r.InsertToken(token.(string), "iin"); err != nil {
		t.Fatalf("failed to insert token %s", err)
	}
}

func TestGetAccessToken(t *testing.T) {
	iin := "user:iin"
	testCtx := context.Background()
	var err error
	mock.CustomMatch(func(expected, actual []interface{}) error {
		return nil
	}).ExpectSet("", iin, time.Duration(exp_time)*time.Second).SetVal("string")
	if _, err = r.GetAccessToken(testCtx, iin); err != nil {
		t.Fatal(err)
	}
}

func TestGetAccessTokenFailInsert(t *testing.T) {
	iin := "user:iin"
	testCtx := context.Background()
	var err error
	mock.CustomMatch(func(expected, actual []interface{}) error {
		return nil
	}).ExpectSet("", iin, time.Duration(exp_time)*time.Second).SetErr(fmt.Errorf("some error"))
	if _, err = r.GetAccessToken(testCtx, iin); err == nil {
		t.Fatal(err)
	}
}
