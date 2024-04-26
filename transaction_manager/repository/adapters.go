package repository

import (
	"errors"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/martian/log"
	"github.com/jackc/pgx/v5/pgtype"
	"simple_warehouse/transaction_manager/domain"
	"simple_warehouse/transaction_manager/repository/store"
)

var (
	ErrorAuthorIsNil = errors.New("author is nil")
)

func convertDbTransactionToDmTransaction(dbTransaction *store.Transaction) *domain.Transaction {
	transaction := domain.Transaction{
		Id:              int(dbTransaction.ID),
		Action:          int(dbTransaction.Action),
		Date:            timestamp.Timestamp{Seconds: dbTransaction.Date.Time.Unix()},
		Sku:             dbTransaction.Sku,
		ShelfName:       dbTransaction.ShelfName,
		QuantityOnShelf: int64(dbTransaction.Quantity),
		Author: &domain.User{
			Id: int(dbTransaction.AuthorID),
		},
	}
	return &transaction
}

func convertDmTransactionToDbTransaction(dmTransaction *domain.Transaction) (*store.Transaction, error) {
	if dmTransaction.Author == nil {
		log.Errorf("[Transaction Repo] convertDmTransactionToDbTransaction: %v", ErrorAuthorIsNil)
		return nil, ErrorAuthorIsNil
	}
	return &store.Transaction{
		ID:        int32(dmTransaction.Id),
		Action:    int32(dmTransaction.Action),
		Date:      pgtype.Timestamp{Time: dmTransaction.Date.AsTime()},
		Sku:       dmTransaction.Sku,
		ShelfName: dmTransaction.ShelfName,
		Quantity:  int32(dmTransaction.QuantityOnShelf),
		AuthorID:  int32(dmTransaction.Author.Id),
	}, nil
}

func convertDbUserToDmUser(dbUser *store.User) *domain.User {
	return &domain.User{
		Id:   int(dbUser.ID),
		Name: dbUser.Name,
	}
}

func convertDmUserToDbUser(dmUser *domain.User) *store.User {
	return &store.User{
		ID:   int32(dmUser.Id),
		Name: dmUser.Name,
	}
}
