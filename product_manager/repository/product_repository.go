package repository

import (
	"context"
	"simple_warehouse/product_manager/domain"
)

type ProductRepository interface {
	GetById(ctx context.Context, id int64) (*domain.Product, error)
	GetBySku(ctx context.Context, sku string) (*domain.Product, error)
	GetConditional(ctx context.Context, predicate func(*domain.Product) bool) ([]*domain.Product, error)
	Create(ctx context.Context, product *domain.Product) error
	Update(ctx context.Context, product *domain.Product) error
	DeleteById(ctx context.Context, id int64) error
	DeleteBySku(ctx context.Context, sku string) error
}
