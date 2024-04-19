package repository

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jackc/pgx/v5/pgtype"
	"simple_warehouse/transaction_manager/domain"
	"simple_warehouse/transaction_manager/repository/store"
)

func NewTransactionRepository(dbQuerier *store.Queries) TransactionRepository {
	return &TransactionRepositoryImpl{
		q: dbQuerier,
	}
}

type TransactionRepository interface {
	GetById(ctx context.Context, id int) (*domain.Transaction, error)
	GetBySku(ctx context.Context, sku string) ([]*domain.Transaction, error)
	GetByAuthorId(ctx context.Context, authorId int) ([]*domain.Transaction, error)
	GetByAuthorName(ctx context.Context, authorName string) ([]*domain.Transaction, error)
	GetByStartAndEndDate(ctx context.Context, startDate, endDate *timestamp.Timestamp) ([]*domain.Transaction, error)
	Create(ctx context.Context, inTransaction *domain.Transaction) (*domain.Transaction, error)
}

type TransactionRepositoryImpl struct {
	q *store.Queries
}

func (tr *TransactionRepositoryImpl) GetById(ctx context.Context, id int) (*domain.Transaction, error) {
	dbTransaction, err := tr.q.GetTransactionById(ctx, int32(id))
	if err != nil {
		return nil, err
	}
	dmTransaction, err := convertDbTransactionToDmTransaction(&dbTransaction)
	if err != nil {
		return nil, err
	}
	return dmTransaction, nil
}

func (tr *TransactionRepositoryImpl) GetBySku(ctx context.Context, sku string) ([]*domain.Transaction, error) {
	dbTransactions, err := tr.q.GetTransactionsBySku(ctx, sku)
	if err != nil {
		return nil, err
	}
	dmTransactions := make([]*domain.Transaction, 0, len(dbTransactions))
	for _, dbTransaction := range dbTransactions {
		dmTransaction, err := convertDbTransactionToDmTransaction(&dbTransaction)
		if err != nil {
			return nil, err
		}
		dmTransactions = append(dmTransactions, dmTransaction)
	}
	return dmTransactions, nil
}

func (tr *TransactionRepositoryImpl) GetByAuthorId(ctx context.Context, authorId int) ([]*domain.Transaction, error) {
	dbTransactions, err := tr.q.GetTransactionsByAuthorId(ctx, int32(authorId))
	if err != nil {
		return nil, err
	}
	dmTransactions := make([]*domain.Transaction, 0, len(dbTransactions))
	for _, dbTransaction := range dbTransactions {
		dmTransaction, err := convertDbTransactionToDmTransaction(&dbTransaction)
		if err != nil {
			return nil, err
		}
		dmTransactions = append(dmTransactions, dmTransaction)
	}
	return dmTransactions, nil
}

func (tr *TransactionRepositoryImpl) GetByAuthorName(ctx context.Context, authorName string) ([]*domain.Transaction, error) {
	dbTransactions, err := tr.q.GetTransactionsByAuthorName(ctx, authorName)
	if err != nil {
		return nil, err
	}
	dmTransactions := make([]*domain.Transaction, 0, len(dbTransactions))
	for _, dbTransaction := range dbTransactions {
		dmTransaction, err := convertDbTransactionToDmTransaction(&dbTransaction)
		if err != nil {
			return nil, err
		}
		dmTransactions = append(dmTransactions, dmTransaction)
	}
	return dmTransactions, nil

}

func (tr *TransactionRepositoryImpl) GetByStartAndEndDate(ctx context.Context, startDate, endDate *timestamp.Timestamp) ([]*domain.Transaction, error) {
	dbTransactions, err := tr.q.GetTransactionsByStartAndEndDate(ctx, store.GetTransactionsByStartAndEndDateParams{
		StartDate: pgtype.Timestamp{Time: startDate.AsTime()},
		EndDate:   pgtype.Timestamp{Time: endDate.AsTime()},
	})
	if err != nil {
		return nil, err
	}
	dmTransactions := make([]*domain.Transaction, 0, len(dbTransactions))
	for _, dbTransaction := range dbTransactions {
		dmTransaction, err := convertDbTransactionToDmTransaction(&dbTransaction)
		if err != nil {
			return nil, err
		}
		dmTransactions = append(dmTransactions, dmTransaction)
	}
	return dmTransactions, nil
}

func (tr *TransactionRepositoryImpl) Create(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error) {
	dbTransaction, err := convertDmTransactionToDbTransaction(transaction)
	if err != nil {
		return nil, err
	}
	createdDbTransaction, err := tr.q.CreateTransaction(ctx, store.CreateTransactionParams{
		Action:    dbTransaction.Action,
		Sku:       dbTransaction.Sku,
		ShelfName: dbTransaction.ShelfName,
		Quantity:  dbTransaction.Quantity,
		AuthorID:  dbTransaction.AuthorID,
	})
	if err != nil {
		return nil, err
	}
	createdDmTransaction, err := convertDbTransactionToDmTransaction(&createdDbTransaction)
	if err != nil {
		return nil, err
	}
	return createdDmTransaction, nil
}
