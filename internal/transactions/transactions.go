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
	GetTransactionsBySender(ctx context.Context, senderAccountNumber string) ([]Transaction, error)
	GetTransactionsByReceiver(ctx context.Context, receiverAccountNumber string) ([]Transaction, error)
	GetTransactionByReference(ctx context.Context, reference string) (*Transaction, error)
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

func (s *TransactionService) GetTransactionsBySender(ctx context.Context, senderAccountNumber string) ([]Transaction, error) {
	transactions, err := s.Store.GetTransactionsBySender(ctx, senderAccountNumber)
	if err != nil {
		log.Printf("Error getting transactions by sender: %v", err)
		return nil, err
	}
	return transactions, nil
}

func (s *TransactionService) GetTransactionsByReceiver(ctx context.Context, receiverAccountNumber string) ([]Transaction, error) {
	transactions, err := s.Store.GetTransactionsByReceiver(ctx, receiverAccountNumber)
	if err != nil {
		log.Printf("Error getting transactions by receiver: %v", err)
		return nil, err
	}
	return transactions, nil
}
