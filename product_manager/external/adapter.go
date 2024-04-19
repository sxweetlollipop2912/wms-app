package external

import (
	"simple_warehouse/product_manager/api"
	"simple_warehouse/product_manager/domain"
)

func convertDomainToGetProduct(inProduct *domain.Product, inShelves []*domain.Shelf) *api.GetProductResponse {
	var shelfQuantityObjs []*api.ShelfQuantity
	for _, shelf := range inShelves {
		if shelf.Product.Sku == inProduct.Sku {
			shelfQuantityObjs = append(shelfQuantityObjs, &api.ShelfQuantity{
				ShelfName: shelf.Name,
				Quantity:  shelf.Quantity,
			})
		}
	}
	return &api.GetProductResponse{
		Sku:             inProduct.Sku,
		Name:            inProduct.Name,
		ShelfQuantities: shelfQuantityObjs,
		ExpiredDate:     &inProduct.ExpiredDate,
		Category:        inProduct.Category,
	}
}

func convertDomainToExport(inShelves []*domain.Shelf) *api.ExportResponse {
	var shelfQuantityObjs []*api.ShelfQuantity
	for _, shelf := range inShelves {
		shelfQuantityObjs = append(shelfQuantityObjs, &api.ShelfQuantity{
			ShelfName: shelf.Name,
			Quantity:  shelf.Quantity,
		})
	}
	return &api.ExportResponse{
		ShelfQuantities: shelfQuantityObjs,
	}
}

func convertImportToDomain(in *api.ImportRequest) (*domain.Product, []*domain.Shelf) {
	outProduct := domain.Product{
		Id:          -1,
		Sku:         in.GetSku(),
		Name:        in.GetName(),
		ExpiredDate: *in.GetExpiredDate(),
		Category:    in.GetCategory(),
	}
	outShelves := make([]*domain.Shelf, 0)
	for _, shelfQuantity := range in.GetShelfQuantities() {
		outShelves = append(outShelves, &domain.Shelf{
			Id:       -1,
			Name:     shelfQuantity.ShelfName,
			Product:  &outProduct,
			Quantity: shelfQuantity.Quantity,
		})
	}
	return &outProduct, outShelves
}

func convertExportToDomain(in *api.ExportRequest) (string, []*domain.Shelf) {
	outShelves := make([]*domain.Shelf, 0)
	for _, shelfQuantity := range in.GetShelfQuantities() {
		outShelves = append(outShelves, &domain.Shelf{
			Id:       -1,
			Name:     shelfQuantity.ShelfName,
			Quantity: shelfQuantity.Quantity,
		})
	}
	return in.GetSku(), outShelves
}
