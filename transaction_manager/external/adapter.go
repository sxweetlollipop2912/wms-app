package external

import (
	sv "simple_warehouse/transaction_manager/api"
	"simple_warehouse/transaction_manager/domain"
)

func convertInsertToDomain(req *sv.InsertRequest) (dmTransaction *domain.Transaction, authorName string, error error) {
	return &domain.Transaction{
		Id:              -1,
		Action:          int(req.GetAction()),
		Sku:             req.GetSku(),
		ShelfName:       req.GetShelfName(),
		QuantityOnShelf: req.GetQuantityOnShelf(),
	}, req.GetAuthorName(), nil
}

func convertDomainToGetDuring(dmTransaction []*domain.Transaction) (*sv.GetDuringResponse, error) {
	var transactions []*sv.Transaction
	for _, t := range dmTransaction {
		transactions = append(transactions, &sv.Transaction{
			Action:          int32(t.Action),
			Date:            &t.Date,
			Sku:             t.Sku,
			ShelfName:       t.ShelfName,
			QuantityOnShelf: t.QuantityOnShelf,
			AuthorName:      t.Author.Name,
		})
	}
	return &sv.GetDuringResponse{
		Transactions: transactions,
	}, nil
}

func convertDomainToGetByUser(dmTransaction []*domain.Transaction) (*sv.GetByUserResponse, error) {
	var transactions []*sv.Transaction
	for _, t := range dmTransaction {
		transactions = append(transactions, &sv.Transaction{
			Action:          int32(t.Action),
			Date:            &t.Date,
			Sku:             t.Sku,
			ShelfName:       t.ShelfName,
			QuantityOnShelf: t.QuantityOnShelf,
		})
	}
	return &sv.GetByUserResponse{
		Transactions: transactions,
	}, nil
}
