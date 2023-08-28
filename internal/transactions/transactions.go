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
	CreateTransaction(context.Context, *Transaction) error
	GetTransaction(context.Context, string) (Transaction, error)
	UpdateTransaction(context.Context, Transaction) error
	GetTransactionsBySender(context.Context, string) ([]Transaction, error)
	GetTransactionsByReceiver(context.Context, string) ([]Transaction, error)
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

func (s *TransactionService) GetTransaction(ctx context.Context, transactionID string) (Transaction, error) {
	transaction, err := s.Store.GetTransaction(ctx, transactionID)
	if err != nil {
		log.Printf("Error fetching transaction with ID %s: %v", transactionID, err)
		return transaction, err
	}
	return transaction, nil
}

func (s *TransactionService) UpdateTransaction(ctx context.Context, transaction Transaction) error {
	if err := s.Store.UpdateTransaction(ctx, transaction); err != nil {
		log.Printf("Error updating transaction: %v", err)
		return err
	}
	return nil
}

func (s *TransactionService) GetTransactionsBySender(ctx context.Context, senderID string) ([]Transaction, error) {
	transactions, err := s.Store.GetTransactionsBySender(ctx, senderID)
	if err != nil {
		log.Printf("Error fetching transactions for sender with ID %s: %v", senderID, err)
		return nil, err
	}
	return transactions, nil
}

func (s *TransactionService) GetTransactionsByReceiver(ctx context.Context, receiverID string) ([]Transaction, error) {
	transactions, err := s.Store.GetTransactionsByReceiver(ctx, receiverID)
	if err != nil {
		log.Printf("Error fetching transactions for receiver with ID %s: %v", receiverID, err)
		return nil, err
	}
	return transactions, nil
}
