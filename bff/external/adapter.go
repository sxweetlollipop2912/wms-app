package external

import (
	sv "simple_warehouse/bff/api"
	pm "simple_warehouse/product_manager/api"
	tm "simple_warehouse/transaction_manager/api"
	tmDomain "simple_warehouse/transaction_manager/domain"
)

func convertImportSvToPM(svReq *sv.ImportRequest) *pm.ImportRequest {
	pmReq := &pm.ImportRequest{
		Sku:             svReq.GetSku(),
		Name:            svReq.GetName(),
		Category:        svReq.GetCategory(),
		ShelfQuantities: make([]*pm.ShelfQuantity, 0),
	}
	for _, sq := range svReq.GetShelfQuantities() {
		pmReq.ShelfQuantities = append(pmReq.ShelfQuantities, &pm.ShelfQuantity{
			ShelfName: sq.GetShelfName(),
			Quantity:  sq.GetQuantity(),
		})
	}
	return pmReq
}

func convertImportSvToTM(svReq *sv.ImportRequest) *tm.InsertRequest {
	tmReq := &tm.InsertRequest{
		Action:          int32(tmDomain.ImportAction),
		Sku:             svReq.GetSku(),
		ShelfQuantities: make([]*tm.ShelfQuantity, 0),
		AuthorName:      svReq.GetAuthorName(),
	}
	for _, sq := range svReq.GetShelfQuantities() {
		tmReq.ShelfQuantities = append(tmReq.ShelfQuantities, &tm.ShelfQuantity{
			ShelfName: sq.GetShelfName(),
			Quantity:  sq.GetQuantity(),
		})
	}
	return tmReq
}

func convertExportSvToPM(svReq *sv.ExportRequest) *pm.ExportRequest {
	pmReq := &pm.ExportRequest{
		Sku:             svReq.GetSku(),
		ShelfQuantities: make([]*pm.ShelfQuantity, 0),
	}
	for _, sq := range svReq.GetShelfQuantities() {
		pmReq.ShelfQuantities = append(pmReq.ShelfQuantities, &pm.ShelfQuantity{
			ShelfName: sq.GetShelfName(),
			Quantity:  sq.GetQuantity(),
		})
	}
	return pmReq
}

func convertExportSvToTM(svReq *sv.ExportRequest) *tm.InsertRequest {
	tmReq := &tm.InsertRequest{
		Action:          int32(tmDomain.ExportAction),
		Sku:             svReq.GetSku(),
		ShelfQuantities: make([]*tm.ShelfQuantity, 0),
		AuthorName:      svReq.GetAuthorName(),
	}
	for _, sq := range svReq.GetShelfQuantities() {
		tmReq.ShelfQuantities = append(tmReq.ShelfQuantities, &tm.ShelfQuantity{
			ShelfName: sq.GetShelfName(),
			Quantity:  sq.GetQuantity(),
		})
	}
	return tmReq
}

func convertExportPMToSv(pmRes *pm.ExportResponse) *sv.ExportResponse {
	svRes := &sv.ExportResponse{
		ShelfQuantities: make([]*sv.ShelfQuantity, 0),
	}
	for _, sq := range pmRes.GetShelfQuantities() {
		svRes.ShelfQuantities = append(svRes.ShelfQuantities, &sv.ShelfQuantity{
			ShelfName: sq.GetShelfName(),
			Quantity:  sq.GetQuantity(),
		})
	}
	return svRes
}

func convertGetProductSvToPM(svReq *sv.GetProductRequest) *pm.GetProductRequest {
	return &pm.GetProductRequest{
		Sku: svReq.GetSku(),
	}
}

func convertGetProductPMToSv(pmProduct *pm.GetProductResponse) *sv.GetProductResponse {
	svProduct := &sv.GetProductResponse{
		Sku:             pmProduct.GetSku(),
		Name:            pmProduct.GetName(),
		Category:        pmProduct.GetCategory(),
		ShelfQuantities: make([]*sv.ShelfQuantity, 0),
	}
	for _, sq := range pmProduct.GetShelfQuantities() {
		svProduct.ShelfQuantities = append(svProduct.ShelfQuantities, &sv.ShelfQuantity{
			ShelfName: sq.GetShelfName(),
			Quantity:  sq.GetQuantity(),
		})
	}
	return svProduct
}
