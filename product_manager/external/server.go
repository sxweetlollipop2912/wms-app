package external

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	sv "simple_warehouse/product_manager/api"
	"simple_warehouse/product_manager/repository"
	"simple_warehouse/product_manager/repository/store"
	"simple_warehouse/product_manager/use_cases"
)

type Server struct {
	sv.UnimplementedProductManagerServer
	uc *use_cases.UseCases
}

func NewServer(dbQuerier *store.Queries) *Server {
	return &Server{uc: use_cases.NewUseCases(dbQuerier)}
}

func (s *Server) Import(ctx context.Context, req *sv.ImportRequest) (*emptypb.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dmProduct, dmShelves := convertImportToDomain(req)

	err = s.uc.Import(ctx, dmProduct, dmShelves)
	if err != nil {
		if errors.Is(err, repository.ErrProductAlreadyExists) {
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

	sku, dmShelves := convertExportToDomain(req)
	dmShelves, err = s.uc.Export(ctx, sku, dmShelves)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}

	res := convertDomainToExport(dmShelves)
	return res, nil
}

func (s *Server) GetProduct(ctx context.Context, req *sv.GetProductRequest) (response *sv.GetProductResponse, errorResponse error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dmProduct, dmShelves, err := s.uc.GetProductBySku(ctx, req.GetSku())
	if err != nil {
		if errors.Is(err, repository.ErrProductNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}

	res := convertDomainToGetProduct(dmProduct, dmShelves)
	return res, nil
}
