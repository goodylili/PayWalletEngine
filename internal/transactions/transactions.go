package transactions

import (
	"PayWalletEngine/internal/accounts"
	"context"
	"log"
)

type Transaction struct {
	Sender        accounts.Account `json:"sender"`
	Receiver      accounts.Account `json:"receiver"`
	TransactionID string           `json:"transaction_id"`
	Amount        float64          `json:"amount"`
	Currency      string           `json:"currency"`
	PaymentMethod string           `json:"payment_method"`
	Status        string           `json:"status"`
	Description   string           `json:"description"`
}

type TransactionStore interface {
	CreateTransaction(ctx context.Context, transaction *Transaction) error
	GetTransactionByTransactionID(ctx context.Context, transactionID string) (*Transaction, error)
	UpdateTransaction(ctx context.Context, transaction *Transaction) error
	DeleteTransactionByID(ctx context.Context, transactionID string) error
	GetTransactionsBySender(ctx context.Context, senderAccountNumber string) ([]Transaction, error)
	GetTransactionsByReceiver(ctx context.Context, receiverAccountNumber string) ([]Transaction, error)
}

type TransactionService struct {
	Store TransactionStore
}

func NewTransactionService(store TransactionStore) *TransactionService {
	return &TransactionService{
		Store: store,
	}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, transaction *Transaction) error {
	if err := s.Store.CreateTransaction(ctx, transaction); err != nil {
		log.Printf("Error creating transaction: %v", err)
		return err
	}
	return nil
}

func (s *TransactionService) GetTransactionByTransactionID(ctx context.Context, transactionID string) (*Transaction, error) {
	transaction, err := s.Store.GetTransactionByTransactionID(ctx, transactionID)
	if err != nil {
		log.Printf("Error getting transaction by ID: %v", err)
		return nil, err
	}
	return transaction, nil
}

func (s *TransactionService) UpdateTransaction(ctx context.Context, transaction *Transaction) error {
	if err := s.Store.UpdateTransaction(ctx, transaction); err != nil {
		log.Printf("Error updating transaction: %v", err)
		return err
	}
	return nil
}

func (s *TransactionService) DeleteTransactionByID(ctx context.Context, transactionID string) error {
	if err := s.Store.DeleteTransactionByID(ctx, transactionID); err != nil {
		log.Printf("Error deleting transaction by ID: %v", err)
		return err
	}
	return nil
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
