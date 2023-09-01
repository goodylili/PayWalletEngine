package accounts

import (
	"PayWalletEngine/internal/users"
	"context"
	"gorm.io/gorm"
	"log"
)

type Account struct {
	gorm.Model
	AccountOwner  users.User `json:"account_owner"`
	AccountID     string     `json:"account_id"`
	AccountNumber string     `json:"account_number"`
	AccountType   string     `json:"account_type"`
	Balance       float64    `json:"balance"`
	Currency      string     `json:"currency"`
	AccountStatus string     `json:"account_status"`
}

type AccountStore interface {
	CreateAccount(ctx context.Context, account *Account) error
	GetAccountByID(context.Context, int64) (Account, error)
	GetAccountByNumber(context.Context, int64) (Account, error)
	UpdateAccountDetails(context.Context, Account) error
	UpdateAccountBalance(context.Context, string, float64) error
	DeleteAccountDetails(context.Context, int64) error
}

// AccountService is the blueprint for the account logic
type AccountService struct {
	Store AccountStore
}

// NewAccountService creates a new service
func NewAccountService(store AccountStore) *AccountService {
	return &AccountService{
		Store: store,
	}
}

func (s *AccountService) GetAccountByID(ctx context.Context, accountID int64) (Account, error) {
	account, err := s.Store.GetAccountByID(ctx, accountID)
	if err != nil {
		log.Printf("Error fetching account with ID %s: %v", accountID, err)
		return account, err
	}

	return account, nil
}

func (s *AccountService) CreateAccount(ctx context.Context, account *Account) error {
	if err := s.Store.CreateAccount(ctx, account); err != nil {
		log.Printf("Error creating account: %v", err)
		return err
	}

	return nil
}

func (s *AccountService) UpdateAccountDetails(ctx context.Context, account Account) error {
	if err := s.Store.UpdateAccountDetails(ctx, account); err != nil {
		log.Printf("Error updating account: %v", err)
		return err
	}

	return nil
}

func (s *AccountService) DeleteAccountDetails(ctx context.Context, accountID int) error {
	if err := s.Store.DeleteAccountDetails(ctx, int64(accountID)); err != nil {
		log.Printf("Error deleting account with ID %s: %v", accountID, err)
		return err
	}

	return nil
}

func (s *AccountService) UpdateAccountBalance(ctx context.Context, accountID int64, newBalance float64) error {
	// Get the account by accountID
	account, err := s.Store.GetAccountByID(ctx, accountID)
	if err != nil {
		log.Printf("Error fetching account with ID %s: %v", accountID, err)
		return err
	}

	// Update the balance
	account.Balance = newBalance

	// Update the account in the store
	if err := s.Store.UpdateAccountDetails(ctx, account); err != nil {
		log.Printf("Error updating balance for account with ID %s: %v", accountID, err)
		return err
	}

	return nil
}

func (s *AccountService) GetAccountByNumber(ctx context.Context, accountNumber int64) (Account, error) {
	account, err := s.Store.GetAccountByNumber(ctx, accountNumber)
	if err != nil {
		log.Printf("Error fetching account with number %d: %v", accountNumber, err)
		return account, err
	}

	return account, nil
}
