package repository

import (
	"errors"
)

var (
	ErrUserExist    = errors.New("user already exists")
	ErrUserNotExist = errors.New("user does not exist")
)
