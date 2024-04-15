package repository

import (
	"context"
	"simple_warehouse/transaction_manager/domain"
)

type UserRepository interface {
	GetById(ctx context.Context, id int64) (*domain.User, error)
	GetConditional(ctx context.Context, predicate func(*domain.User) bool) ([]*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
}
