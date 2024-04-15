package use_cases

import "errors"

var (
	ErrProductNotExists     = errors.New("product does not exist")
	ErrProductAlreadyExists = errors.New("product already exists")

	ErrProductNotOnShelf       = errors.New("product is not on shelf")
	ErrProductNotEnoughOnShelf = errors.New("product is not enough on shelf")

	// ErrShelfProductMismatch is internal error
	ErrShelfProductMismatch = errors.New("ErrShelfProductMismatch")
)
