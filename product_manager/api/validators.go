package api

import "errors"

var (
	ErrImportNoShelfProvided = errors.New("no shelf provided")
	ErrImportSkuEmpty        = errors.New("sku is empty")
	ErrImportNameEmpty       = errors.New("name is empty")

	ErrExportNoShelfProvided = errors.New("no shelf provided")
	ErrExportSkuEmpty        = errors.New("sku is empty")

	ErrGetProductSkuEmpty = errors.New("sku is empty")
)

func (req *ImportRequest) Validate() error {
	if req.GetShelfQuantities() == nil || len(req.GetShelfQuantities()) == 0 {
		return ErrImportNoShelfProvided
	}
	if req.GetSku() == "" {
		return ErrImportSkuEmpty
	}
	if req.GetName() == "" {
		return ErrImportNameEmpty
	}
	return nil
}

func (req *ExportRequest) Validate() error {
	if req.GetShelfQuantities() == nil || len(req.GetShelfQuantities()) == 0 {
		return ErrExportNoShelfProvided
	}
	if req.GetSku() == "" {
		return ErrExportSkuEmpty
	}
	return nil
}

func (req *GetProductRequest) Validate() error {
	if req.GetSku() == "" {
		return ErrGetProductSkuEmpty
	}
	return nil
}
