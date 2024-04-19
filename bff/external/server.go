package external

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	sv "simple_warehouse/bff/api"
	pm "simple_warehouse/product_manager/api"
	tm "simple_warehouse/transaction_manager/api"
)

type serverT struct {
	sv.UnimplementedBFFServer
	pm pm.ProductManagerClient
	tm tm.TransactionManagerClient
}

func (s *serverT) Import(ctx context.Context, req *sv.ImportRequest) (*emptypb.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	pmReq := convertImportSvToPM(req)
	_, err = s.pm.Import(ctx, pmReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	tmReq := convertImportSvToTM(req)
	_, err = s.tm.Insert(ctx, tmReq)

	return &emptypb.Empty{}, err
}

func (s *serverT) Export(ctx context.Context, req *sv.ExportRequest) (response *sv.ExportResponse, errorResponse error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	pmReq := convertExportSvToPM(req)
	pmRes, err := s.pm.Export(ctx, pmReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	tmReq := convertExportSvToTM(req)
	_, err = s.tm.Insert(ctx, tmReq)

	svRes := convertExportPMToSv(pmRes)
	return svRes, err
}

func (s *serverT) GetProduct(ctx context.Context, req *sv.GetProductRequest) (response *sv.GetProductResponse, errorResponse error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	pmReq := convertGetProductSvToPM(req)
	pmRes, err := s.pm.GetProduct(ctx, pmReq)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	svRes := convertGetProductPMToSv(pmRes)
	return svRes, nil
}
