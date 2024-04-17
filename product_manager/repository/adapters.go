package repository

import (
	"errors"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jackc/pgx/v5/pgtype"
	"simple_warehouse/product_manager/domain"
	"simple_warehouse/product_manager/repository/store"
)

var (
	ErrorProductIsNil = errors.New("product is nil")
)

func convertDbProductToDmProduct(dbProduct *store.Product) (*domain.Product, error) {
	return &domain.Product{
		Id:          int(dbProduct.ID),
		Sku:         dbProduct.Sku,
		Name:        dbProduct.Name,
		ExpiredDate: timestamp.Timestamp{Seconds: dbProduct.ExpiredDate.Time.Unix()},
		Category:    dbProduct.Category.String,
	}, nil
}

func convertDmProductToDbProduct(dmProduct *domain.Product) (*store.Product, error) {
	return &store.Product{
		ID:          int32(dmProduct.Id),
		Sku:         dmProduct.Sku,
		Name:        dmProduct.Name,
		ExpiredDate: pgtype.Timestamp{Time: dmProduct.ExpiredDate.AsTime()},
		Category:    pgtype.Text{String: dmProduct.Category},
	}, nil
}

func convertDbShelfToDmShelf(dbShelf *store.Shelf) (*domain.Shelf, error) {
	return &domain.Shelf{
		Id:   int(dbShelf.ID),
		Name: dbShelf.Name,
		Product: &domain.Product{
			Id: int(dbShelf.ProductID),
		},
		Quantity: int64(dbShelf.Quantity),
	}, nil
}

func convertDmShelfToDbShelf(dmShelf *domain.Shelf) (*store.Shelf, error) {
	if dmShelf.Product == nil {
		return nil, ErrorProductIsNil
	}
	return &store.Shelf{
		ID:        int32(dmShelf.Id),
		Name:      dmShelf.Name,
		ProductID: int32(dmShelf.Product.Id),
		Quantity:  int32(dmShelf.Quantity),
	}, nil
}
