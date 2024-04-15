package domain

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Transaction struct {
	Id              int64               `json:"id"`
	Action          int                 `json:"action"`
	Date            timestamp.Timestamp `json:"date"`
	Sku             string              `json:"sku"`
	ShelfName       string              `json:"shelf_name"`
	QuantityOnShelf int64               `json:"quantity_on_shelf"`
	Author          *User               `json:"author"`
}

const (
	ImportAction  int = iota
	ExportAction  int = iota
	DiscardAction int = iota
)

func (t *Transaction) Validate() error {
	if t.Sku == "" {
		return ErrTransactionSkuNotMissing
	}
	if t.ShelfName == "" {
		return ErrTransactionShelfNameMissing
	}
	if t.QuantityOnShelf <= 0 {
		return ErrTransactionQuantityOnShelfInvalid
	}
	if t.Author == nil {
		return ErrTransactionAuthorNotExists
	}
	if t.Action < ImportAction || t.Action > DiscardAction {
		return ErrTransactionActionInvalid
	}
	return nil
}
