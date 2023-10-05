package transactions

import (
	"PayWalletEngine/internal/accounts"
	"PayWalletEngine/internal/users"
	"context"
	"github.com/google/uuid"
	"log"
)

type Transactions struct {
	TransactionID         uuid.UUID `json:"transaction_id"`
	Amount                float64   `json:"amount"`
	PaymentMethod         string    `json:"paymentMethod"`
	Type                  string    `json:"type"`
	Status                string    `json:"status"`
	Description           string    `json:"description"`
	Reference             string    `json:"reference"`
	SenderAccountNumber   int64     `json:"sender_account_number"`
	ReceiverAccountNumber int64     `json:"receiver_account_number"`
}

type TransactionStore interface {
	GetTransactionsFromAccount(ctx context.Context, accountNumber int64) ([]Transactions, error)
	GetTransactionByReference(ctx context.Context, reference string) (*Transactions, error)
	DebitAccount(ctx context.Context, senderAccountNumber int64, amount float64, description string, paymentMethod string) (Transactions, error)
	CreditAccount(ctx context.Context, retrieveAccountNumber int64, amount float64, description string, paymentMethod string) (Transactions, error)
	TransferFunds(ctx context.Context, senderAccountNumber int64, receiverAccountNumber int64, amount float64, description string, paymentMethod string) (Transactions, error)
	GetUserAccountAndTransactionByTransactionID(ctx context.Context, transactionID string) (*users.User, *accounts.Account, *Transactions, error)
	GetAccountandTransactionByTransactionID(ctx context.Context, transactionID string) (*accounts.Account, *Transactions, error)
}

type TransactionService struct {
	Store TransactionStore
}

func NewTransactionService(store TransactionStore) TransactionService {
	return TransactionService{
		Store: store,
	}
}

// GetUserAccountAndTransactionByTransactionID retrieves the user, account and transaction by transaction TransactionID
func (s *TransactionService) GetUserAccountAndTransactionByTransactionID(ctx context.Context, transactionID string) (*users.User, *accounts.Account, *Transactions, error) {
	user, account, transaction, err := s.Store.GetUserAccountAndTransactionByTransactionID(ctx, transactionID)
	if err != nil {
		log.Printf("Error getting user, account and transaction by transaction TransactionID: %v", err)
		return nil, nil, nil, err
	}
	return user, account, transaction, nil
}

// GetAccountByTransactionID retrieves the account and transaction by transaction TransactionID
func (s *TransactionService) GetAccountByTransactionID(ctx context.Context, transactionID string) (*accounts.Account, *Transactions, error) {
	account, transaction, err := s.Store.GetAccountandTransactionByTransactionID(ctx, transactionID)
	if err != nil {
		log.Printf("Error getting account and transaction by transaction TransactionID: %v", err)
		return nil, nil, err
	}
	return account, transaction, nil
}

// GetTransactionByReference retrieves a transaction by reference.
func (s *TransactionService) GetTransactionByReference(ctx context.Context, reference string) (*Transactions, error) {
	transaction, err := s.Store.GetTransactionByReference(ctx, reference)
	if err != nil {
		log.Printf("Error getting transaction by reference: %v", err)
		return nil, err
	}
	return transaction, nil
}

// DebitAccount debits the specified account.
func (s *TransactionService) DebitAccount(ctx context.Context, senderAccountNumber int64, amount float64, description string, paymentMethod string) (*Transactions, error) {
	transaction, err := s.Store.DebitAccount(ctx, senderAccountNumber, amount, description, paymentMethod)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// CreditAccount credits an account for a transaction.
func (s *TransactionService) CreditAccount(ctx context.Context, receiverAccountNumber int64, amount float64, description string, paymentMethod string) (*Transactions, error) {
	transaction, err := s.Store.CreditAccount(ctx, receiverAccountNumber, amount, description, paymentMethod)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// TransferFunds transfers funds by crediting and debiting specified users.
func (s *TransactionService) TransferFunds(ctx context.Context, senderAccountNumber int64, receiverAccountNumber int64, amount float64, description string, paymentMethod string) (*Transactions, error) {
	transaction, err := s.Store.TransferFunds(ctx, senderAccountNumber, receiverAccountNumber, amount, description, paymentMethod)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// GetTransactionsFromAccount retrieves the transactions a specific account made.
func (s *TransactionService) GetTransactionsFromAccount(ctx context.Context, accountNumber int64) ([]Transactions, error) {
	transactions, err := s.Store.GetTransactionsFromAccount(ctx, accountNumber)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
