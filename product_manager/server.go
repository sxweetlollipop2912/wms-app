package main

import (
	"context"
	sv "simple_warehouse/product_manager/api"
)

type server struct {
	sv.UnimplementedProductManagerServer
}

func (s *server) Import(ctx context.Context, req *sv.ImportRequest) (response *sv.ImportResponse, errorResponse error) {
	// TODO
	return &sv.ImportResponse{
		Ok: true,
	}, nil
}

func (s *server) Export(ctx context.Context, req *sv.ExportRequest) (response *sv.ExportResponse, errorResponse error) {
	// TODO
	return &sv.ExportResponse{
		Ok:              true,
		Sku:             req.GetSku(),
		ShelfNames:      req.GetShelfNames(),
		QuantityOnShelf: []int64{1, 2, 3},
		ExpiredDate:     1e9,
		Category:        "category",
	}, nil
}

func (s *server) GetProduct(ctx context.Context, req *sv.GetProductRequest) (response *sv.GetProductResponse, errorResponse error) {
	// TODO
	return &sv.GetProductResponse{
		Ok:              true,
		Sku:             req.GetSku(),
		Name:            "name",
		ShelfNames:      []string{"shelf1", "shelf2"},
		QuantityOnShelf: []int64{1, 2},
		ExpiredDate:     1e9,
		Category:        "category",
	}, nil
}
