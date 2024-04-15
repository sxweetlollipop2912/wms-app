package api

import "errors"

var (
	ErrImportNoShelfProvided            = errors.New("no shelf provided")
	ErrImportShelvesAndQuantityMismatch = errors.New("shelves and quantity mismatch")
	ErrImportSkuEmpty                   = errors.New("sku is empty")
	ErrImportNameEmpty                  = errors.New("name is empty")

	ErrExportNoShelfProvided            = errors.New("no shelf provided")
	ErrExportShelvesAndQuantityMismatch = errors.New("shelves and quantity mismatch")
	ErrExportSkuEmpty                   = errors.New("sku is empty")

	ErrGetProductSkuEmpty = errors.New("sku is empty")
)

func (req *ImportRequest) Validate() error {
	if req.GetShelfNames() == nil || len(req.GetShelfNames()) == 0 {
		return ErrImportNoShelfProvided
	}
	if req.GetQuantityOnShelf() == nil || len(req.GetQuantityOnShelf()) != len(req.GetShelfNames()) {
		return ErrImportShelvesAndQuantityMismatch
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
	if req.GetShelfNames() == nil || len(req.GetShelfNames()) == 0 {
		return ErrExportNoShelfProvided
	}
	if req.GetQuantityOnShelf() == nil || len(req.GetQuantityOnShelf()) != len(req.GetShelfNames()) {
		return ErrExportShelvesAndQuantityMismatch
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
