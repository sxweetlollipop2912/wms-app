package external

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	sv "simple_warehouse/transaction_manager/api"
	"simple_warehouse/transaction_manager/repository"
	"simple_warehouse/transaction_manager/use_cases"
)

type Server struct {
	*sv.UnimplementedTransactionManagerServer
	uc *use_cases.UseCases
}

func NewServer() *Server {
	return &Server{uc: use_cases.NewUseCases()}
}

func (s *Server) Insert(ctx context.Context, req *sv.InsertRequest) (*emptypb.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dmTransactions, dmAuthor, err := convertInsertToDomain(req)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	err = s.uc.InsertTransactionInBulk(ctx, dmTransactions, dmAuthor)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) GetDuring(ctx context.Context, req *sv.GetDuringRequest) (*sv.GetDuringResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dmTransactions, err := s.uc.GetTransactionDuring(ctx, req.GetStartDate(), req.GetEndDate())
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	res, err := convertDomainToGetDuring(dmTransactions)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return res, nil
}

func (s *Server) GetByUser(ctx context.Context, req *sv.GetByUserRequest) (*sv.GetByUserResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dmTransactions, err := s.uc.GetTransactionByUser(ctx, req.GetAuthorName())
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	res, err := convertDomainToGetByUser(dmTransactions)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return res, nil
}

func (s *Server) AddUser(ctx context.Context, req *sv.AddUserRequest) (*emptypb.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dmUser, err := convertAddUserToDomain(req)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	err = s.uc.AddUser(ctx, dmUser)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &emptypb.Empty{}, nil
}
