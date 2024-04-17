package use_cases

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"simple_warehouse/transaction_manager/domain"
	"simple_warehouse/transaction_manager/repository"
)

type UseCases struct {
	transactionRepository repository.TransactionRepository
	userRepository        repository.UserRepository
}

func NewUseCases() *UseCases {
	return &UseCases{
		transactionRepository: repository.NewTransactionRepository(),
		userRepository:        repository.NewUserRepository(),
	}
}

func (uc *UseCases) InsertTransactionInBulk(ctx context.Context, inTransactions []*domain.Transaction, inAuthor *domain.User) error {
	dbAuthor, err := uc.userRepository.GetByName(ctx, inAuthor.Name)
	if err != nil {
		return err
	}
	for _, transaction := range inTransactions {
		transaction.Author = dbAuthor
		_, err := uc.transactionRepository.Create(ctx, transaction)
		if err != nil {
			return err
		}
	}
	return nil
}

func (uc *UseCases) GetTransactionDuring(ctx context.Context, start, end *timestamp.Timestamp) (outTransactions []*domain.Transaction, error error) {
	outTransactions, error = uc.transactionRepository.GetConditional(ctx, func(transaction *domain.Transaction) bool {
		if start != nil && transaction.Date.AsTime().Before(start.AsTime()) {
			return false
		}
		if end != nil && transaction.Date.AsTime().After(end.AsTime()) {
			return false
		}
		return true
	})
	if error != nil {
		return nil, error
	}
	return outTransactions, nil
}

func (uc *UseCases) GetTransactionByUser(ctx context.Context, userName string) (outTransactions []*domain.Transaction, error error) {
	dbUser, err := uc.userRepository.GetByName(ctx, userName)
	if err != nil {
		return nil, err
	}
	outTransactions, error = uc.transactionRepository.GetConditional(ctx, func(transaction *domain.Transaction) bool {
		return transaction.Author.Id == dbUser.Id
	})
	if error != nil {
		return nil, error
	}
	return outTransactions, nil
}

func (uc *UseCases) GetTransactionByProduct(ctx context.Context, sku string) (outTransactions []*domain.Transaction, error error) {
	outTransactions, error = uc.transactionRepository.GetConditional(ctx, func(transaction *domain.Transaction) bool {
		return transaction.Sku == sku
	})
	if error != nil {
		return nil, error
	}
	return outTransactions, nil
}

func (uc *UseCases) AddUser(ctx context.Context, user *domain.User) error {
	_, err := uc.userRepository.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
