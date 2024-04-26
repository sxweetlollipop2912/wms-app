package repository

import (
	"context"
	"errors"
	"github.com/google/martian/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/exp/slices"
	"simple_warehouse/product_manager/domain"
	"simple_warehouse/product_manager/repository/store"
)

func NewProductRepository(dbQuerier *store.Queries) ProductRepository {
	return &ProductRepositoryImpl{
		q: dbQuerier,
	}
}

type ProductRepository interface {
	GetById(ctx context.Context, id int) (*domain.Product, error)
	GetBySku(ctx context.Context, sku string) (*domain.Product, error)
	FindByName(ctx context.Context, name string) ([]*domain.Product, error)
	Create(ctx context.Context, product *domain.Product) (*domain.Product, error)
	Update(ctx context.Context, product *domain.Product) (*domain.Product, error)
}

type ProductRepositoryImpl struct {
	q *store.Queries
}

func (pr *ProductRepositoryImpl) GetById(ctx context.Context, id int) (*domain.Product, error) {
	dbProduct, err := pr.q.GetProductById(ctx, int32(id))
	if err != nil {
		log.Errorf("[Product Repo] GetById query: %v", err)
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
	dmProduct := convertDbProductToDmProduct(&dbProduct)
	return dmProduct, nil
}

func (pr *ProductRepositoryImpl) GetBySku(ctx context.Context, sku string) (*domain.Product, error) {
	dbProduct, err := pr.q.GetProductBySku(ctx, sku)
	if err != nil {
		log.Errorf("[Product Repo] GetBySku query: %v", err)
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
	dmProduct := convertDbProductToDmProduct(&dbProduct)
	return dmProduct, nil
}

func (pr *ProductRepositoryImpl) FindByName(ctx context.Context, name string) ([]*domain.Product, error) {
	dbProducts, err := pr.q.FindProductByName(ctx, pgtype.Text{String: name})
	if err != nil {
		log.Errorf("[Product Repo] FindByName query: %v", err)
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
	var dmProducts []*domain.Product
	for _, dbProduct := range dbProducts {
		dmProduct := convertDbProductToDmProduct(&dbProduct)
		dmProducts = append(dmProducts, dmProduct)
	}
	return dmProducts, nil

}

func (pr *ProductRepositoryImpl) Create(ctx context.Context, inProduct *domain.Product) (*domain.Product, error) {
	dbProduct, err := pr.q.CreateProduct(ctx, store.CreateProductParams{
		Sku:      inProduct.Sku,
		Name:     inProduct.Name,
		Category: pgtype.Text{String: inProduct.Category},
	})
	if err != nil {
		log.Errorf("[Product Repo] Create query: %v", err)
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
	dmProduct := convertDbProductToDmProduct(&dbProduct)
	return dmProduct, nil
}

func (pr *ProductRepositoryImpl) Update(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	dbProduct, err := pr.q.UpdateProduct(ctx, store.UpdateProductParams{
		ID:       int32(product.Id),
		Sku:      product.Sku,
		Name:     product.Name,
		Category: pgtype.Text{String: product.Category},
	})
	if err != nil {
		log.Errorf("[Product Repo] Update query: %v", err)
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
	dmProduct := convertDbProductToDmProduct(&dbProduct)
	return dmProduct, nil
}
