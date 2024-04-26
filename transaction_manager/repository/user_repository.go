package repository

import (
	"context"
	"errors"
	"github.com/google/martian/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/exp/slices"
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
		log.Errorf("[User Repo] GetById query: %v", err)
		var dbErr *pgconn.PgError
		if errors.As(err, &dbErr) {
			if slices.Contains(DbErrCodeNotFound, dbErr.Code) {
				return nil, ErrUserNotExists
			}
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotExists
		}
		return nil, err
	}
	dmUser := convertDbUserToDmUser(&dbUser)
	return dmUser, nil
}

func (ur *UserRepositoryImpl) GetByExactName(ctx context.Context, name string) (*domain.User, error) {
	dbUser, err := ur.q.GetUserByExactName(ctx, name)
	if err != nil {
		log.Errorf("[User Repo] GetByExactName query: %v", err)
		var dbErr *pgconn.PgError
		if errors.As(err, &dbErr) {
			if slices.Contains(DbErrCodeNotFound, dbErr.Code) {
				return nil, ErrUserNotExists
			}
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotExists
		}
		return nil, err
	}
	dmUser := convertDbUserToDmUser(&dbUser)
	return dmUser, nil
}

func (ur *UserRepositoryImpl) FindByName(ctx context.Context, name string) ([]*domain.User, error) {
	dbUsers, err := ur.q.FindUserByName(ctx, pgtype.Text{String: name})
	if err != nil {
		log.Errorf("[User Repo] FindByName query: %v", err)
		return nil, err
	}
	dmUsers := make([]*domain.User, 0, len(dbUsers))
	for _, dbUser := range dbUsers {
		dmUser := convertDbUserToDmUser(&dbUser)
		dmUsers = append(dmUsers, dmUser)
	}
	return dmUsers, nil
}

func (ur *UserRepositoryImpl) Create(ctx context.Context, inUser *domain.User) (*domain.User, error) {
	createdDbUser, err := ur.q.CreateUser(ctx, inUser.Name)
	if err != nil {
		log.Errorf("[User Repo] Create query: %v", err)
		var dbErr *pgconn.PgError
		if errors.As(err, &dbErr) {
			if slices.Contains(DbErrCodeAlreadyExists, dbErr.Code) {
				return nil, ErrUserAlreadyExists
			}
			if slices.Contains(DbErrCodeFieldViolation, dbErr.Code) {
				return nil, ErrFieldViolation
			}
		}
		return nil, err
	}
	createdDmUser := convertDbUserToDmUser(&createdDbUser)
	return createdDmUser, nil
}
