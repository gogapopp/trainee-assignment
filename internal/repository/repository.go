package repository

import (
	"errors"
)

var (
	ErrUserExist        = errors.New("user already exists")
	ErrUserNotExist     = errors.New("user does not exist")
	ErrBannerNotExist   = errors.New("banner does not exist")
	ErrNoFieldsToUpdate = errors.New("no fields to update")
)
