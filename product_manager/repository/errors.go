package repository

import (
	"errors"
)

var (
	ErrProductNotExists     = errors.New("product does not exist")
	ErrProductAlreadyExists = errors.New("product already exists")

	ErrProductNotOnShelf       = errors.New("product is not on shelf")
	ErrProductNotEnoughOnShelf = errors.New("product is not enough on shelf")

	ErrFieldViolation = errors.New("field violation")

	DbErrCodeAlreadyExists = []string{
		"23505", // unique_violation
	}
	DbErrCodeFieldViolation = []string{
		"23000", // integrity_constraint_violation
		"23501", // restrict_violation
		"23502", // not_null_violation
		"23503", // foreign_key_violation
		"23514", // check_violation
		"23P01", // exclusion_violation
	}
	DbErrCodeNotFound = []string{
		"20000", // case_not_found
	}
)
