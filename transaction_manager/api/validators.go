package api

import "errors"

var (
	ErrInsertSkuMissing        = errors.New("sku is missing")
	ErrInsertShelfNameMissing  = errors.New("shelf name is missing")
	ErrInsertQuantityInvalid   = errors.New("quantity is invalid")
	ErrInsertAuthorNameMissing = errors.New("author is missing")

	ErrGetByUserAuthorNameMissing = errors.New("author name is missing")

	ErrAddUserUserNameMissing = errors.New("user name is missing")
)

func (req *InsertRequest) Validate() error {
	if req.GetSku() == "" {
		return ErrInsertSkuMissing
	}
	if req.GetShelfName() == "" {
		return ErrInsertShelfNameMissing
	}
	if req.GetQuantityOnShelf() <= 0 {
		return ErrInsertQuantityInvalid
	}
	if req.GetAuthorName() == "" {
		return ErrInsertAuthorNameMissing
	}
	return nil
}

func (req *GetDuringRequest) Validate() error {
	return nil
}

func (req *GetByUserRequest) Validate() error {
	if req.GetAuthorName() == "" {
		return ErrGetByUserAuthorNameMissing
	}
	return nil
}

func (req *AddUserRequest) Validate() error {
	if req.GetUserName() == "" {
		return ErrAddUserUserNameMissing
	}
	return nil
}
