package use_cases

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"simple_warehouse/transaction_manager/domain"
	"simple_warehouse/transaction_manager/repository"
	"simple_warehouse/transaction_manager/repository/store"
)

type UseCases struct {
	transactionRepository repository.TransactionRepository
	userRepository        repository.UserRepository
}

func NewUseCases(dbQuerier *store.Queries) *UseCases {
	return &UseCases{
		transactionRepository: repository.NewTransactionRepository(dbQuerier),
		userRepository:        repository.NewUserRepository(dbQuerier),
	}
}

func (uc *UseCases) InsertTransactionInBulk(ctx context.Context, inTransactions []*domain.Transaction, inAuthor *domain.User) error {
	author, err := uc.userRepository.GetByExactName(ctx, inAuthor.Name)
	if err != nil {
		return err
	}
	for _, transaction := range inTransactions {
		transaction.Author = author
		_, err := uc.transactionRepository.Create(ctx, transaction)
		if err != nil {
			return err
		}
	}
	return nil
}

func (uc *UseCases) GetTransactionDuring(ctx context.Context, start, end *timestamp.Timestamp) (outTransactions []*domain.Transaction, error error) {
	outTransactions, error = uc.transactionRepository.GetByStartAndEndDate(ctx, start, end)
	if error != nil {
		return nil, error
	}
	return outTransactions, nil
}

func (uc *UseCases) GetTransactionByUserId(ctx context.Context, userId int) (outTransactions []*domain.Transaction, error error) {
	user, err := uc.userRepository.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}
	outTransactions, error = uc.transactionRepository.GetByAuthorId(ctx, user.Id)
	if error != nil {
		return nil, error
	}
	return outTransactions, nil
}

func (uc *UseCases) GetTransactionByUserName(ctx context.Context, userName string) (outTransactions []*domain.Transaction, error error) {
	user, err := uc.userRepository.GetByExactName(ctx, userName)
	if err != nil {
		return nil, err
	}
	outTransactions, error = uc.transactionRepository.GetByAuthorId(ctx, user.Id)
	if error != nil {
		return nil, error
	}
	return outTransactions, nil
}

func (uc *UseCases) GetTransactionByProduct(ctx context.Context, sku string) (outTransactions []*domain.Transaction, error error) {
	outTransactions, error = uc.transactionRepository.GetBySku(ctx, sku)
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
