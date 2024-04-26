package external

import (
	"context"
	"errors"
	"github.com/google/martian/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	sv "simple_warehouse/transaction_manager/api"
	"simple_warehouse/transaction_manager/repository"
	"simple_warehouse/transaction_manager/repository/store"
	"simple_warehouse/transaction_manager/use_cases"
)

type Server struct {
	*sv.UnimplementedTransactionManagerServer
	uc *use_cases.UseCases
}

func NewServer(dbQuerier *store.Queries) *Server {
	return &Server{uc: use_cases.NewUseCases(dbQuerier)}
}

func (s *Server) Insert(ctx context.Context, req *sv.InsertRequest) (*emptypb.Empty, error) {
	log.Infof("[Server] Insert request: %v", req)

	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dmTransactions, dmAuthor := convertInsertToDomain(req)
	err = s.uc.InsertTransactionInBulk(ctx, dmTransactions, dmAuthor)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		if errors.Is(err, repository.ErrTransactionAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		if errors.Is(err, repository.ErrFieldViolation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) GetDuring(ctx context.Context, req *sv.GetDuringRequest) (*sv.GetDuringResponse, error) {
	log.Infof("[Server] GetDuring request: %v", req)

	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dmTransactions, err := s.uc.GetTransactionDuring(ctx, req.GetStartDate(), req.GetEndDate())
	if err != nil {
		if errors.Is(err, repository.ErrFieldViolation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}

	res := convertDomainToGetDuring(dmTransactions)
	return res, nil
}

func (s *Server) GetByUser(ctx context.Context, req *sv.GetByUserRequest) (*sv.GetByUserResponse, error) {
	log.Infof("[Server] GetByUser request: %v", req)

	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dmTransactions, err := s.uc.GetTransactionByUserName(ctx, req.GetAuthorName())
	if err != nil {
		if errors.Is(err, repository.ErrFieldViolation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}

	res := convertDomainToGetByUser(dmTransactions)
	return res, nil
}

func (s *Server) AddUser(ctx context.Context, req *sv.AddUserRequest) (*emptypb.Empty, error) {
	log.Infof("[Server] AddUser request: %v", req)

	err := req.Validate()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dmUser := convertAddUserToDomain(req)
	err = s.uc.AddUser(ctx, dmUser)
	if err != nil {
		if errors.Is(err, repository.ErrUserAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		if errors.Is(err, repository.ErrFieldViolation) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &emptypb.Empty{}, nil
}
