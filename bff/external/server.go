package external

import (
	"context"
	"github.com/google/martian/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	sv "simple_warehouse/bff/api"
	pm "simple_warehouse/product_manager/api"
	tm "simple_warehouse/transaction_manager/api"
)

type ServerT struct {
	sv.UnimplementedBFFServer
	Pm pm.ProductManagerClient
	Tm tm.TransactionManagerClient
}

func (s *ServerT) Import(ctx context.Context, req *sv.ImportRequest) (*emptypb.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	pmReq := convertImportSvToPM(req)
	log.Infof("[Server] Import PM: %v", pmReq)
	_, err = s.Pm.Import(ctx, pmReq)
	if err != nil {
		log.Errorf("[Server] Import PM: %v", err)
		return nil, err
	}

	tmReq := convertImportSvToTM(req)
	log.Infof("[Server] Import TM: %v", tmReq)
	_, err = s.Tm.Insert(ctx, tmReq)
	if err != nil {
		log.Errorf("[Server] Import TM: %v", err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ServerT) Export(ctx context.Context, req *sv.ExportRequest) (response *sv.ExportResponse, errorResponse error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	pmReq := convertExportSvToPM(req)
	log.Infof("[Server] Export PM: %v", pmReq)
	pmRes, err := s.Pm.Export(ctx, pmReq)
	if err != nil {
		log.Errorf("[Server] Export PM: %v", err)
		return nil, err
	}

	tmReq := convertExportSvToTM(req)
	log.Infof("[Server] Export TM: %v", tmReq)
	_, err = s.Tm.Insert(ctx, tmReq)
	if err != nil {
		log.Errorf("[Server] Export TM: %v", err)
		return nil, err
	}

	svRes := convertExportPMToSv(pmRes)
	return svRes, nil
}

func (s *ServerT) GetProduct(ctx context.Context, req *sv.GetProductRequest) (response *sv.GetProductResponse, errorResponse error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	pmReq := convertGetProductSvToPM(req)
	log.Infof("[Server] GetProduct PM: %v", pmReq)
	pmRes, err := s.Pm.GetProduct(ctx, pmReq)
	if err != nil {
		log.Errorf("[Server] GetProduct PM: %v", err)
		return nil, err
	}

	svRes := convertGetProductPMToSv(pmRes)
	return svRes, nil
}
