package accounts

import (
	"PayWalletEngine/internal/users"
	"context"
	"fmt"
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
	ActivateAccount(context.Context, int64) error
	DeActivateAccount(context.Context, int64) error
	CreditAccount(context.Context, int64) error
	DebitAccount(context.Context, int64) error
}

// AccountService is the blueprint for the account logic
type AccountService struct {
	Store AccountStore
}

// NewAccountService creates a new service
func NewAccountService(store AccountStore) AccountService {
	return AccountService{
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

func (s *AccountService) DeleteAccountDetails(ctx context.Context, accountID int64) error {
	if err := s.Store.DeleteAccountDetails(ctx, accountID); err != nil {
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

// ActivateAccount activates an account by setting its status to 'Active'
func (s *AccountService) ActivateAccount(ctx context.Context, accountID int64) error {
	account, err := s.Store.GetAccountByID(ctx, accountID)
	if err != nil {
		log.Printf("Error fetching account with ID %d: %v", accountID, err)
		return err
	}
	account.AccountStatus = "Active"

	if err := s.Store.UpdateAccountDetails(ctx, account); err != nil {
		log.Printf("Error activating account with ID %d: %v", accountID, err)
		return err
	}
	return nil
}

// DeActivateAccount deactivates an account by setting its status to 'Inactive'
func (s *AccountService) DeActivateAccount(ctx context.Context, accountID int64) error {
	account, err := s.Store.GetAccountByID(ctx, accountID)
	if err != nil {
		log.Printf("Error fetching account with ID %d: %v", accountID, err)
		return err
	}
	account.AccountStatus = "Inactive"

	if err := s.Store.UpdateAccountDetails(ctx, account); err != nil {
		log.Printf("Error deactivating account with ID %d: %v", accountID, err)
		return err
	}
	return nil
}

// CreditAccount adds the specified amount to the account balance
func (s *AccountService) CreditAccount(ctx context.Context, accountID int64, amount float64) error {
	account, err := s.Store.GetAccountByID(ctx, accountID)
	if err != nil {
		log.Printf("Error fetching account with ID %d: %v", accountID, err)
		return err
	}
	account.Balance += amount

	if err := s.Store.UpdateAccountDetails(ctx, account); err != nil {
		log.Printf("Error crediting account with ID %d: %v", accountID, err)
		return err
	}
	return nil
}

// DebitAccount subtracts the specified amount from the account balance
func (s *AccountService) DebitAccount(ctx context.Context, accountID int64, amount float64) error {
	account, err := s.Store.GetAccountByID(ctx, accountID)
	if err != nil {
		log.Printf("Error fetching account with ID %d: %v", accountID, err)
		return err
	}
	if account.Balance < amount {
		return fmt.Errorf("insufficient funds in account with ID %d", accountID)
	}
	account.Balance -= amount

	if err := s.Store.UpdateAccountDetails(ctx, account); err != nil {
		log.Printf("Error debiting account with ID %d: %v", accountID, err)
		return err
	}
	return nil
}
