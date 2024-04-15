package repository

import "errors"

var (
	ErrTransactionNotExists     = errors.New("ErrTransactionNotExists")
	ErrTransactionAlreadyExists = errors.New("ErrTransactionAlreadyExists")
	ErrUserNotExists            = errors.New("ErrUserNotExists")
	ErrUserAlreadyExists        = errors.New("ErrUserAlreadyExists")
)
