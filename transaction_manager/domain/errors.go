package domain

import "errors"

var (
	ErrTransactionActionInvalid          = errors.New("ErrTransactionActionInvalid")
	ErrTransactionSkuNotMissing          = errors.New("ErrTransactionSkuNotMissing")
	ErrTransactionShelfNameMissing       = errors.New("ErrTransactionShelfNameMissing")
	ErrTransactionQuantityOnShelfInvalid = errors.New("ErrTransactionQuantityOnShelfInvalid")
	ErrTransactionAuthorNotExists        = errors.New("ErrTransactionAuthorNotExists")
)
