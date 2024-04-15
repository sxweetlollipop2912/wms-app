package repository

import (
	"context"
	"simple_warehouse/transaction_manager/domain"
)

type TransactionRepository interface {
	GetById(ctx context.Context, id int64) (*domain.Transaction, error)
	GetConditional(ctx context.Context, predicate func(*domain.Transaction) bool) ([]*domain.Transaction, error)
	Create(ctx context.Context, transaction *domain.Transaction) error
}
