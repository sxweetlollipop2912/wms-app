package repository

import (
	"context"
	"simple_warehouse/transaction_manager/domain"
)

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

type TransactionRepository interface {
	GetById(ctx context.Context, id int) (*domain.Transaction, error)
	GetConditional(ctx context.Context, predicate func(*domain.Transaction) bool) ([]*domain.Transaction, error)
	Create(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error)
}

type TransactionRepositoryImpl struct{}

func (tr *TransactionRepositoryImpl) GetById(ctx context.Context, id int) (*domain.Transaction, error) {
	return &domain.Transaction{}, nil
}

func (tr *TransactionRepositoryImpl) GetConditional(ctx context.Context, predicate func(*domain.Transaction) bool) ([]*domain.Transaction, error) {
	return []*domain.Transaction{}, nil
}

func (tr *TransactionRepositoryImpl) Create(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error) {
	return &domain.Transaction{}, nil
}
