package use_cases

import (
	"context"
	"simple_warehouse/product_manager/domain"
	"simple_warehouse/product_manager/repository"
	"simple_warehouse/product_manager/repository/store"
)

type UseCases struct {
	productRepo repository.ProductRepository
	shelfRepo   repository.ShelfRepository
}

func NewUseCases(dbQuerier *store.Queries) *UseCases {
	return &UseCases{
		productRepo: repository.NewProductRepository(dbQuerier),
		shelfRepo:   repository.NewShelfRepository(dbQuerier),
	}
}

func (uc *UseCases) Import(ctx context.Context, inProduct *domain.Product, inShelves []*domain.Shelf) error {
	dbProduct, err := uc.productRepo.Create(ctx, inProduct)
	if err != nil {
		return err
	}
	for _, shelf := range inShelves {
		shelf.Product = dbProduct
		_, err := uc.shelfRepo.Create(ctx, shelf)
		if err != nil {
			return err
		}
	}
	return nil
}

func (uc *UseCases) Export(ctx context.Context, inSku string, inShelves []*domain.Shelf) (outShelves []*domain.Shelf, error error) {
	outShelves = []*domain.Shelf{}
	for _, shelfQuantity := range inShelves {
		shelfName := shelfQuantity.Name
		quantity := shelfQuantity.Quantity

		dbShelf, err := uc.shelfRepo.GetByNameAndSku(ctx, shelfName, inSku)
		if err != nil {
			return nil, err
		}
		if dbShelf == nil {
			return nil, repository.ErrProductNotOnShelf
		}
		if dbShelf.Quantity < quantity {
			return nil, repository.ErrProductNotEnoughOnShelf
		}

		dbShelf.Quantity -= quantity
		if dbShelf.Quantity == 0 {
			dbShelf, err = uc.shelfRepo.DeleteById(ctx, dbShelf.Id)
			if err != nil {
				return nil, err
			}
		} else {
			dbShelf, err = uc.shelfRepo.Update(ctx, dbShelf)
			if err != nil {
				return nil, err
			}
		}
		outShelves = append(outShelves, dbShelf)
	}
	return outShelves, nil
}

func (uc *UseCases) GetProductBySku(ctx context.Context, inSku string) (*domain.Product, []*domain.Shelf, error) {
	dbProduct, err := uc.productRepo.GetBySku(ctx, inSku)
	if err != nil {
		return nil, nil, err
	}

	dbShelves, err := uc.shelfRepo.GetByProduct(ctx, dbProduct)
	if err != nil {
		return nil, nil, err
	}

	return dbProduct, dbShelves, nil
}
