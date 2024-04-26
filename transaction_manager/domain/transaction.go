package domain

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Transaction struct {
	Id              int                 `json:"id"`
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
