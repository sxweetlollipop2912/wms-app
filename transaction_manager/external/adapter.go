package external

import (
	sv "simple_warehouse/transaction_manager/api"
	"simple_warehouse/transaction_manager/domain"
)

func convertInsertToDomain(req *sv.InsertRequest) (dmTransaction []*domain.Transaction, dmAuthor *domain.User, error error) {
	dmTransaction = make([]*domain.Transaction, 0)
	for _, shelfQuantity := range req.ShelfQuantities {
		dmTransaction = append(dmTransaction, &domain.Transaction{
			Id:              -1,
			Action:          int(req.GetAction()),
			Sku:             req.GetSku(),
			ShelfName:       shelfQuantity.GetShelfName(),
			QuantityOnShelf: shelfQuantity.GetQuantity(),
		})
	}
	dmAuthor = &domain.User{
		Id:   -1,
		Name: req.GetAuthorName(),
	}

	return dmTransaction, dmAuthor, nil
}

func convertAddUserToDomain(req *sv.AddUserRequest) (*domain.User, error) {
	return &domain.User{
		Id:   -1,
		Name: req.GetUserName(),
	}, nil
}

func convertDomainToGetDuring(dmTransaction []*domain.Transaction) (*sv.GetDuringResponse, error) {
	var transactions []*sv.Transaction
	for _, t := range dmTransaction {
		transactions = append(transactions, &sv.Transaction{
			Action: int32(t.Action),
			Date:   &t.Date,
			Sku:    t.Sku,
			ShelfQuantities: []*sv.ShelfQuantity{
				{
					ShelfName: t.ShelfName,
					Quantity:  t.QuantityOnShelf,
				},
			},
			AuthorName: t.Author.Name,
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
			Action: int32(t.Action),
			Date:   &t.Date,
			Sku:    t.Sku,
			ShelfQuantities: []*sv.ShelfQuantity{
				{
					ShelfName: t.ShelfName,
					Quantity:  t.QuantityOnShelf,
				},
			},
			AuthorName: t.Author.Name,
		})
	}
	return &sv.GetByUserResponse{
		Transactions: transactions,
	}, nil
}
