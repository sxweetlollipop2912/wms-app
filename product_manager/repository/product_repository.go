package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"simple_warehouse/product_manager/domain"
	"simple_warehouse/product_manager/repository/store"
)

func NewProductRepository(dbQuerier *store.Queries) ProductRepository {
	return &ProductRepositoryImpl{
		q: dbQuerier,
	}
}

type ProductRepository interface {
	GetById(ctx context.Context, id int) (*domain.Product, error)
	GetBySku(ctx context.Context, sku string) (*domain.Product, error)
	Create(ctx context.Context, product *domain.Product) (*domain.Product, error)
	Update(ctx context.Context, product *domain.Product) (*domain.Product, error)
}

type ProductRepositoryImpl struct {
	q *store.Queries
}

func (pr *ProductRepositoryImpl) GetById(ctx context.Context, id int) (*domain.Product, error) {
	dbProduct, err := pr.q.GetProductById(ctx, int32(id))
	if err != nil {
		return nil, err
	}
	dmProduct, err := convertDbProductToDmProduct(&dbProduct)
	if err != nil {
		return nil, err
	}
	return dmProduct, nil
}

func (pr *ProductRepositoryImpl) GetBySku(ctx context.Context, sku string) (*domain.Product, error) {
	dbProduct, err := pr.q.GetProductBySku(ctx, sku)
	if err != nil {
		return nil, err
	}
	dmProduct, err := convertDbProductToDmProduct(&dbProduct)
	if err != nil {
		return nil, err
	}
	return dmProduct, nil
}

func (pr *ProductRepositoryImpl) Create(ctx context.Context, inProduct *domain.Product) (*domain.Product, error) {
	dbProduct, err := pr.q.CreateProduct(ctx, store.CreateProductParams{
		Sku:         inProduct.Sku,
		Name:        inProduct.Name,
		Category:    pgtype.Text{String: inProduct.Category},
		ExpiredDate: pgtype.Timestamp{Time: inProduct.ExpiredDate.AsTime()},
	})
	if err != nil {
		return nil, err
	}
	dmProduct, err := convertDbProductToDmProduct(&dbProduct)
	if err != nil {
		return nil, err
	}
	return dmProduct, nil
}

func (pr *ProductRepositoryImpl) Update(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	dbProduct, err := pr.q.UpdateProduct(ctx, store.UpdateProductParams{
		ID:          int32(product.Id),
		Sku:         product.Sku,
		Name:        product.Name,
		Category:    pgtype.Text{String: product.Category},
		ExpiredDate: pgtype.Timestamp{Time: product.ExpiredDate.AsTime()},
	})
	if err != nil {
		return nil, err
	}
	dmProduct, err := convertDbProductToDmProduct(&dbProduct)
	if err != nil {
		return nil, err
	}
	return dmProduct, nil
}
