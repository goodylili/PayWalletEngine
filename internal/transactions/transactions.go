package transactions

import (
	"PayWalletEngine/internal/accounts"
	"context"
	"log"
)

type Transaction struct {
	Sender        accounts.Account `json:"sender"`
	Receiver      accounts.Account `json:"receiver"`
	TransactionID int64            `json:"transaction_id"`
	Amount        float64          `json:"amount"`
	PaymentMethod string           `json:"payment_method"`
	Status        string           `json:"status"`
	Description   string           `json:"description"`
	Reference     string           `json:"reference"`
}

type TransactionStore interface {
	GetTransactionByTransactionID(ctx context.Context, transactionID int64) (*Transaction, error)
	GetTransactionsFromAccount(ctx context.Context, accountNumber int64) ([]Transaction, error)
	GetTransactionByReference(ctx context.Context, reference string) (*Transaction, error)
	DebitAccount(ctx context.Context, senderAccountNumber int64, amount float64, description string, paymentMethod string) (Transaction, error)
	CreditAccount(ctx context.Context, retrieveAccountNumber int64, amount float64, description string, paymentMethod string) (Transaction, error)
	TransferFunds(ctx context.Context, senderAccountNumber int64, receiverAccountNumber int64, amount float64, description string, paymentMethod string) (Transaction, error)
}

type TransactionService struct {
	Store TransactionStore
}

func NewTransactionService(store TransactionStore) TransactionService {
	return TransactionService{
		Store: store,
	}
}

func (s *TransactionService) GetTransactionByTransactionID(ctx context.Context, transactionID int64) (*Transaction, error) {
	transaction, err := s.Store.GetTransactionByTransactionID(ctx, transactionID)
	if err != nil {
		log.Printf("Error getting transaction by ID: %v", err)
		return nil, err
	}
	return transaction, nil
}

func (s *TransactionService) GetTransactionByReference(ctx context.Context, reference string) (*Transaction, error) {
	transaction, err := s.Store.GetTransactionByReference(ctx, reference)
	if err != nil {
		log.Printf("Error getting transaction by reference: %v", err)
		return nil, err
	}
	return transaction, nil
}
