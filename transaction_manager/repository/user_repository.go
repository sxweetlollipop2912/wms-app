package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"simple_warehouse/transaction_manager/domain"
	"simple_warehouse/transaction_manager/repository/store"
)

func NewUserRepository(dbQuerier *store.Queries) UserRepository {
	return &UserRepositoryImpl{
		q: dbQuerier,
	}
}

type UserRepository interface {
	GetById(ctx context.Context, id int) (*domain.User, error)
	GetByExactName(ctx context.Context, name string) (*domain.User, error)
	FindByName(ctx context.Context, name string) ([]*domain.User, error)
	Create(ctx context.Context, inUser *domain.User) (*domain.User, error)
}

type UserRepositoryImpl struct {
	q *store.Queries
}

func (ur *UserRepositoryImpl) GetById(ctx context.Context, id int) (*domain.User, error) {
	dbUser, err := ur.q.GetUserById(ctx, int32(id))
	if err != nil {
		return nil, err
	}
	dmUser, err := convertDbUserToDmUser(&dbUser)
	if err != nil {
		return nil, err
	}
	return dmUser, nil
}

func (ur *UserRepositoryImpl) GetByExactName(ctx context.Context, name string) (*domain.User, error) {
	dbUser, err := ur.q.GetUserByExactName(ctx, name)
	if err != nil {
		return nil, err
	}
	dmUser, err := convertDbUserToDmUser(&dbUser)
	if err != nil {
		return nil, err
	}
	return dmUser, nil
}

func (ur *UserRepositoryImpl) FindByName(ctx context.Context, name string) ([]*domain.User, error) {
	dbUsers, err := ur.q.FindUserByName(ctx, pgtype.Text{String: name})
	if err != nil {
		return nil, err
	}
	dmUsers := make([]*domain.User, 0, len(dbUsers))
	for _, dbUser := range dbUsers {
		dmUser, err := convertDbUserToDmUser(&dbUser)
		if err != nil {
			return nil, err
		}
		dmUsers = append(dmUsers, dmUser)
	}
	return dmUsers, nil
}

func (ur *UserRepositoryImpl) Create(ctx context.Context, inUser *domain.User) (*domain.User, error) {
	createdDbUser, err := ur.q.CreateUser(ctx, inUser.Name)
	if err != nil {
		return nil, err
	}
	createdDmUser, err := convertDbUserToDmUser(&createdDbUser)
	if err != nil {
		return nil, err
	}
	return createdDmUser, nil
}
