package external

import (
	"simple_warehouse/product_manager/api"
	"simple_warehouse/product_manager/domain"
)

func convertDomainToGetProduct(inProduct *domain.Product, inShelves []*domain.Shelf) (*api.GetProductResponse, error) {
	var shelfNames []string
	var quantityOnShelf []int64
	for _, shelf := range inShelves {
		if shelf.Product.Sku == inProduct.Sku {
			shelfNames = append(shelfNames, shelf.Name)
			quantityOnShelf = append(quantityOnShelf, shelf.Quantity)
		}
	}
	return &api.GetProductResponse{
		Sku:             inProduct.Sku,
		Name:            inProduct.Name,
		ShelfNames:      shelfNames,
		QuantityOnShelf: quantityOnShelf,
		ExpiredDate:     &inProduct.ExpiredDate,
		Category:        inProduct.Category,
	}, nil
}

func convertDomainToExport(inProduct *domain.Product, inShelves []*domain.Shelf) (*api.ExportResponse, error) {
	var shelfNames []string
	var quantityOnShelf []int64
	for _, shelf := range inShelves {
		if shelf.Product.Sku == inProduct.Sku {
			shelfNames = append(shelfNames, shelf.Name)
			quantityOnShelf = append(quantityOnShelf, shelf.Quantity)
		}
	}
	return &api.ExportResponse{
		Sku:             inProduct.Sku,
		ShelfNames:      shelfNames,
		QuantityOnShelf: quantityOnShelf,
		ExpiredDate:     &inProduct.ExpiredDate,
		Category:        inProduct.Category,
	}, nil
}

func convertImportToDomain(in *api.ImportRequest) (*domain.Product, []*domain.Shelf, error) {
	outProduct := domain.Product{
		Id:          -1,
		Sku:         in.GetSku(),
		Name:        in.GetName(),
		ExpiredDate: *in.GetExpiredDate(),
		Category:    in.GetCategory(),
	}
	outShelves := make([]*domain.Shelf, 0)
	for i, shelfName := range in.GetShelfNames() {
		outShelves = append(outShelves, &domain.Shelf{
			Id:       -1,
			Name:     shelfName,
			Product:  &outProduct,
			Quantity: in.GetQuantityOnShelf()[i],
		})
	}
	return &outProduct, outShelves, nil
}
