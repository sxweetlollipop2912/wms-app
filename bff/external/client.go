package external

import (
	"context"
	"google.golang.org/grpc"

	cli "simple_warehouse/product_manager/api"
)

type clientT struct {
	*grpc.ClientConn
}

func (c *clientT) Import(ctx context.Context, request *cli.ImportRequest) (*cli.ImportResponse, error) {
	r, err := c.Import(ctx, request)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *clientT) Export(ctx context.Context, request *cli.ExportRequest) (*cli.ExportResponse, error) {
	r, err := c.Export(ctx, request)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *clientT) GetProduct(ctx context.Context, request *cli.GetProductRequest) (*cli.GetProductResponse, error) {
	r, err := c.GetProduct(ctx, request)
	if err != nil {
		return nil, err
	}
	return r, nil
}
