package transactions

import (
	"context"
	"log"
)

type Transaction struct {
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"paymentMethod"`
	Type          string  `json:"type"`
	Status        string  `json:"status"`
	Description   string  `json:"description"`
	Reference     string  `json:"reference"`

	// Add Sender and Receiver IDs
	SenderID   uint `json:"sender_id"`
	ReceiverID uint `json:"receiver_id"`
}

type TransactionStore interface {
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

func (s *TransactionService) GetTransactionByReference(ctx context.Context, reference string) (*Transaction, error) {
	transaction, err := s.Store.GetTransactionByReference(ctx, reference)
	if err != nil {
		log.Printf("Error getting transaction by reference: %v", err)
		return nil, err
	}
	return transaction, nil
}

// DebitAccount debits the specified account.
func (s *TransactionService) DebitAccount(ctx context.Context, senderAccountNumber int64, amount float64, description string, paymentMethod string) (*Transaction, error) {
	transaction, err := s.Store.DebitAccount(ctx, senderAccountNumber, amount, description, paymentMethod)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// CreditAccount credits an account for a transaction.
func (s *TransactionService) CreditAccount(ctx context.Context, receiverAccountNumber int64, amount float64, description string, paymentMethod string) (*Transaction, error) {
	transaction, err := s.Store.CreditAccount(ctx, receiverAccountNumber, amount, description, paymentMethod)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// TransferFunds transfers funds by crediting and debiting specified users.
func (s *TransactionService) TransferFunds(ctx context.Context, senderAccountNumber int64, receiverAccountNumber int64, amount float64, description string, paymentMethod string) (*Transaction, error) {
	transaction, err := s.Store.TransferFunds(ctx, senderAccountNumber, receiverAccountNumber, amount, description, paymentMethod)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// GetTransactionsFromAccount retrieves the transactions a specific account made.
func (s *TransactionService) GetTransactionsFromAccount(ctx context.Context, accountNumber int64) ([]Transaction, error) {
	transactions, err := s.Store.GetTransactionsFromAccount(ctx, accountNumber)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
