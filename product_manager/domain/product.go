package domain

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Product struct {
	Id          int                 `json:"id"`
	Name        string              `json:"name"`
	Sku         string              `json:"sku"`
	ExpiredDate timestamp.Timestamp `json:"expired_date"`
	Category    string              `json:"category"`
}
