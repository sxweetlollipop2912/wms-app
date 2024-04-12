package main

import (
	"context"
	"log"
	sv "simple_warehouse/bff/api"
	cli "simple_warehouse/product_manager/api"
)

type serverT struct {
	sv.UnimplementedBFFServer
	client cli.ProductManagerClient
}

func (s *serverT) Import(ctx context.Context, req *sv.ImportRequest) (response *sv.ImportResponse, errorResponse error) {
	pmReq := &cli.ImportRequest{
		Sku:             req.GetSku(),
		Name:            req.GetName(),
		ExpiredDate:     req.GetExpiredDate(),
		Category:        req.GetCategory(),
		ShelfNames:      []string{"A"},
		QuantityOnShelf: []int64{1},
	}

	pmRes, err := s.client.Import(ctx, pmReq)
	if err != nil {
		return nil, err
	}
	log.Printf("Import response: %v", pmRes)
	res := sv.ImportResponse{
		Ok:     pmRes.GetOk(),
		Reason: pmRes.GetReason(),
	}
	return &res, nil
}

func (s *serverT) Export(ctx context.Context, req *sv.ExportRequest) (response *sv.ExportResponse, errorResponse error) {
	pmReq := &cli.ExportRequest{
		Sku:             req.GetSku(),
		ShelfNames:      req.GetShelfNames(),
		QuantityOnShelf: req.GetQuantityOnShelf(),
	}
	pmRes, err := s.client.Export(ctx, pmReq)
	if err != nil {
		return nil, err
	}

	res := sv.ExportResponse{
		Ok:     pmRes.GetOk(),
		Reason: pmRes.GetReason(),
	}
	return &res, nil
}

func (s *serverT) GetProduct(ctx context.Context, req *sv.GetProductRequest) (response *sv.GetProductResponse, errorResponse error) {
	pmReq := &cli.GetProductRequest{
		Sku: req.GetSku(),
	}
	pmRes, err := s.client.GetProduct(ctx, pmReq)
	if err != nil {
		return nil, err
	}
	if !pmRes.GetOk() {
		return &sv.GetProductResponse{
			Ok:     false,
			Reason: pmRes.GetReason(),
		}, nil
	}

	res := sv.GetProductResponse{
		Ok:              true,
		Sku:             pmRes.GetSku(),
		Name:            pmRes.GetName(),
		ShelfNames:      pmRes.GetShelfNames(),
		QuantityOnShelf: pmRes.GetQuantityOnShelf(),
		ExpiredDate:     pmRes.GetExpiredDate(),
		Category:        pmRes.GetCategory(),
	}
	return &res, nil
}
