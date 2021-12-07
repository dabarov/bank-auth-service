package domain

import "errors"

var (
	ErrLoginTaken  = errors.New("user with such username already exists")
	ErrEmptyField  = errors.New("login or password field is empty")
	ErrIINTaken    = errors.New("user with such IIN already exists")
	ErrIINIncorect = errors.New("incorect fomat of IIN")
)
