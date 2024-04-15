package repository

import (
	"context"
	"simple_warehouse/product_manager/domain"
)

type ShelfRepository interface {
	GetById(ctx context.Context, id int64) (*domain.Shelf, error)
	GetByProduct(ctx context.Context, product *domain.Product) ([]*domain.Shelf, error)
	GetByproductId(ctx context.Context, productId int64) ([]*domain.Shelf, error)
	GetConditional(ctx context.Context, predicate func(*domain.Shelf) bool) ([]*domain.Shelf, error)
	Create(ctx context.Context, shelf *domain.Shelf) error
	Update(ctx context.Context, shelf *domain.Shelf) error
	DeleteById(ctx context.Context, id int64) error
	DeleteByProductId(ctx context.Context, productId int64) error
}
