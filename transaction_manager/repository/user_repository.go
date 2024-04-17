package repository

import (
	"context"
	"simple_warehouse/transaction_manager/domain"
)

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

type UserRepository interface {
	GetById(ctx context.Context, id int) (*domain.User, error)
	GetByName(ctx context.Context, name string) (*domain.User, error)
	GetConditional(ctx context.Context, predicate func(*domain.User) bool) ([]*domain.User, error)
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
}

type UserRepositoryImpl struct{}

func (ur *UserRepositoryImpl) GetById(ctx context.Context, id int) (*domain.User, error) {
	return &domain.User{}, nil
}

func (ur *UserRepositoryImpl) GetByName(ctx context.Context, name string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (ur *UserRepositoryImpl) GetConditional(ctx context.Context, predicate func(*domain.User) bool) ([]*domain.User, error) {
	return []*domain.User{}, nil
}

func (ur *UserRepositoryImpl) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	return &domain.User{}, nil
}
