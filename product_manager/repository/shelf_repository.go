package repository

import (
	"context"
	"errors"
	"github.com/google/martian/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/exp/slices"
	"simple_warehouse/product_manager/domain"
	"simple_warehouse/product_manager/repository/store"
)

func NewShelfRepository(dbQuerier *store.Queries) ShelfRepository {
	return &ShelfRepositoryImpl{
		q: dbQuerier,
	}
}

type ShelfRepository interface {
	GetById(ctx context.Context, id int) (*domain.Shelf, error)
	GetByProduct(ctx context.Context, product *domain.Product) ([]*domain.Shelf, error)
	GetByProductId(ctx context.Context, productId int) ([]*domain.Shelf, error)
	GetByNameAndSku(ctx context.Context, name, sku string) (*domain.Shelf, error)
	Create(ctx context.Context, shelf *domain.Shelf) (*domain.Shelf, error)
	Update(ctx context.Context, shelf *domain.Shelf) (*domain.Shelf, error)
	DeleteById(ctx context.Context, id int) (*domain.Shelf, error)
}

type ShelfRepositoryImpl struct {
	q *store.Queries
}

func (sr *ShelfRepositoryImpl) GetById(ctx context.Context, id int) (*domain.Shelf, error) {
	dbShelf, err := sr.q.GetShelfById(ctx, int32(id))
	if err != nil {
		log.Errorf("[Shelf Repo] GetById query: %v", err)
		var dbErr *pgconn.PgError
		if errors.As(err, &dbErr) {
			if slices.Contains(DbErrCodeNotFound, dbErr.Code) {
				return nil, ErrProductNotExists
			}
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProductNotExists
		}
		return nil, err
	}
	dmShelf := convertDbShelfToDmShelf(&dbShelf)
	return dmShelf, nil
}

func (sr *ShelfRepositoryImpl) GetByProduct(ctx context.Context, inProduct *domain.Product) ([]*domain.Shelf, error) {
	dbShelves, err := sr.q.GetShelfByProductId(ctx, int32(inProduct.Id))
	if err != nil {
		log.Errorf("[Shelf Repo] GetByProduct query: %v", err)
		var dbErr *pgconn.PgError
		if errors.As(err, &dbErr) {
			if slices.Contains(DbErrCodeNotFound, dbErr.Code) {
				return nil, ErrProductNotExists
			}
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProductNotExists
		}
		return nil, err
	}
	var dmShelves []*domain.Shelf
	for _, dbShelf := range dbShelves {
		dmShelf := convertDbShelfToDmShelf(&dbShelf)
		dmShelf.Product = inProduct
		dmShelves = append(dmShelves, dmShelf)
	}
	return dmShelves, nil
}

func (sr *ShelfRepositoryImpl) GetByProductId(ctx context.Context, productId int) ([]*domain.Shelf, error) {
	dbShelves, err := sr.q.GetShelfByProductId(ctx, int32(productId))
	if err != nil {
		log.Errorf("[Shelf Repo] GetByProductId query: %v", err)
		var dbErr *pgconn.PgError
		if errors.As(err, &dbErr) {
			if slices.Contains(DbErrCodeNotFound, dbErr.Code) {
				return nil, ErrProductNotExists
			}
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProductNotExists
		}
		return nil, err
	}
	var dmShelves []*domain.Shelf
	for _, dbShelf := range dbShelves {
		dmShelf := convertDbShelfToDmShelf(&dbShelf)
		dmShelves = append(dmShelves, dmShelf)
	}
	return dmShelves, nil
}

func (sr *ShelfRepositoryImpl) GetByNameAndSku(ctx context.Context, name, sku string) (*domain.Shelf, error) {
	dbShelf, err := sr.q.GetShelfBySkuAndName(ctx, store.GetShelfBySkuAndNameParams{
		Name: name,
		Sku:  sku,
	})
	if err != nil {
		log.Errorf("[Shelf Repo] GetByNameAndSku query: %v", err)
		var dbErr *pgconn.PgError
		if errors.As(err, &dbErr) {
			if slices.Contains(DbErrCodeNotFound, dbErr.Code) {
				return nil, ErrProductNotExists
			}
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProductNotExists
		}
		return nil, err
	}
	dmShelf := convertDbShelfToDmShelf(&dbShelf)
	return dmShelf, nil
}

func (sr *ShelfRepositoryImpl) Create(ctx context.Context, inShelf *domain.Shelf) (*domain.Shelf, error) {
	dbShelf, err := sr.q.CreateShelf(ctx, store.CreateShelfParams{
		Name:      inShelf.Name,
		ProductID: int32(inShelf.Product.Id),
		Quantity:  int32(inShelf.Quantity),
	})
	if err != nil {
		log.Errorf("[Shelf Repo] Create query: %v", err)
		var dbErr *pgconn.PgError
		if errors.As(err, &dbErr) {
			if slices.Contains(DbErrCodeAlreadyExists, dbErr.Code) {
				return nil, ErrProductAlreadyExists
			}
			if slices.Contains(DbErrCodeFieldViolation, dbErr.Code) {
				return nil, ErrFieldViolation
			}
		}
		return nil, err
	}
	dmShelf := convertDbShelfToDmShelf(&dbShelf)
	return dmShelf, nil
}

func (sr *ShelfRepositoryImpl) Update(ctx context.Context, inShelf *domain.Shelf) (*domain.Shelf, error) {
	dbShelf, err := sr.q.UpdateShelf(ctx, store.UpdateShelfParams{
		ID:        int32(inShelf.Id),
		Name:      inShelf.Name,
		ProductID: int32(inShelf.Product.Id),
		Quantity:  int32(inShelf.Quantity),
	})
	if err != nil {
		log.Errorf("[Shelf Repo] Update query: %v", err)
		var dbErr *pgconn.PgError
		if errors.As(err, &dbErr) {
			if slices.Contains(DbErrCodeNotFound, dbErr.Code) {
				return nil, ErrProductNotExists
			}
			if slices.Contains(DbErrCodeAlreadyExists, dbErr.Code) {
				return nil, ErrFieldViolation
			}
			if slices.Contains(DbErrCodeFieldViolation, dbErr.Code) {
				return nil, ErrFieldViolation
			}
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProductNotExists
		}
		return nil, err
	}
	dmShelf := convertDbShelfToDmShelf(&dbShelf)
	return dmShelf, nil
}

func (sr *ShelfRepositoryImpl) DeleteById(ctx context.Context, id int) (*domain.Shelf, error) {
	dbShelf, err := sr.q.DeleteShelfById(ctx, int32(id))
	if err != nil {
		log.Errorf("[Shelf Repo] DeleteById query: %v", err)
		var dbErr *pgconn.PgError
		if errors.As(err, &dbErr) {
			if slices.Contains(DbErrCodeNotFound, dbErr.Code) {
				return nil, ErrProductNotExists
			}
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProductNotExists
		}
		return nil, err
	}
	dmShelf := convertDbShelfToDmShelf(&dbShelf)
	return dmShelf, nil
}
