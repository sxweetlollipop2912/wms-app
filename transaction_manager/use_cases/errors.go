package use_cases

import "errors"

var (
	ErrUserNotExists     = errors.New("user does not exists")
	ErrUserAlreadyExists = errors.New("user already exists")
)
