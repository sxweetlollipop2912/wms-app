package use_cases

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"simple_warehouse/transaction_manager/domain"
	"simple_warehouse/transaction_manager/repository"
)

type UseCases struct {
	transactionRepository repository.TransactionRepository
	userRepository        repository.UserRepository
}

func (uc *UseCases) InsertTransaction(ctx context.Context, transaction *domain.Transaction, authorName string) error {
	// TODO
	return nil
}

func (uc *UseCases) GetTransactionDuring(ctx context.Context, start, end *timestamp.Timestamp) ([]*domain.Transaction, error) {
	// TODO
	return []*domain.Transaction{}, nil
}

func (uc *UseCases) GetTransactionByUser(ctx context.Context, userName string) ([]*domain.Transaction, error) {
	// TODO
	return []*domain.Transaction{}, nil
}

func (uc *UseCases) GetTransactionByProduct(ctx context.Context, sku string) ([]*domain.Transaction, error) {
	// TODO
	return []*domain.Transaction{}, nil
}

func (uc *UseCases) AddUser(ctx context.Context, user *domain.User) error {
	// TODO
	return nil
}
