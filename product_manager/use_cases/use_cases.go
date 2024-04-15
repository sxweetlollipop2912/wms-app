package use_cases

import (
	"context"
	"simple_warehouse/product_manager/domain"
	"simple_warehouse/product_manager/repository"
)

type UseCases struct {
	productRepo repository.ProductRepository
	shelfRepo   repository.ShelfRepository
}

func (uc *UseCases) Import(ctx context.Context, in *domain.Product, inShelves []*domain.Shelf) error {
	// TODO
	return nil
}

func (uc *UseCases) Export(ctx context.Context, inSku string, inShelfNames []string, inShelfQuantity []int64) (*domain.Product, []*domain.Shelf, error) {
	// TODO
	return &domain.Product{}, []*domain.Shelf{}, nil
}

func (uc *UseCases) GetProductBySku(ctx context.Context, inSku string) (*domain.Product, []*domain.Shelf, error) {
	// TODO
	return &domain.Product{}, []*domain.Shelf{}, nil
}
