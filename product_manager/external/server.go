package external

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	sv "simple_warehouse/product_manager/api"
	"simple_warehouse/product_manager/use_cases"
)

type Server struct {
	sv.UnimplementedProductManagerServer
	uc *use_cases.UseCases
}

func (s *Server) Import(ctx context.Context, req *sv.ImportRequest) (*emptypb.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dmProduct, dmShelves, err := convertImportToDomain(req)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	err = s.uc.Import(ctx, dmProduct, dmShelves)
	if err != nil {
		if errors.Is(err, use_cases.ErrProductAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) Export(ctx context.Context, req *sv.ExportRequest) (*sv.ExportResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dmProduct, dmShelves, err := s.uc.Export(ctx, req.GetSku(), req.GetShelfNames(), req.GetQuantityOnShelf())
	if err != nil {
		if errors.Is(err, use_cases.ErrProductNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}

	res, err := convertDomainToExport(dmProduct, dmShelves)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return res, nil
}

func (s *Server) GetProduct(ctx context.Context, req *sv.GetProductRequest) (response *sv.GetProductResponse, errorResponse error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dmProduct, dmShelves, err := s.uc.GetProductBySku(ctx, req.GetSku())
	if err != nil {
		if errors.Is(err, use_cases.ErrProductNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}

	res, err := convertDomainToGetProduct(dmProduct, dmShelves)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return res, nil
}
