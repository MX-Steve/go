package model

import "errors"

var (
	ErrUserNotExist  = errors.New("user not exist.")
	ErrInvalidPasswd = errors.New("Passwd or username not right.")
	ErrInvalidParams = errors.New("Invalid user params")
	ErrUserExist     = errors.New("user exist.")
)
