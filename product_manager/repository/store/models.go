// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package store

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Product struct {
	ID          int32            `json:"id"`
	Sku         string           `json:"sku"`
	Name        string           `json:"name"`
	ExpiredDate pgtype.Timestamp `json:"expired_date"`
	Category    pgtype.Text      `json:"category"`
}

type Shelf struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	ProductID int32  `json:"product_id"`
	Quantity  int32  `json:"quantity"`
}
