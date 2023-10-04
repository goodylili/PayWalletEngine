package accounts

import (
	"PayWalletEngine/internal/users"
	"context"
	"gorm.io/gorm"
	"log"
)

type Account struct {
	gorm.Model    `json:"-"`
	ID            uint    `json:"id"`
	AccountNumber string  `json:"account_number"`
	AccountType   string  `json:"account_type"`
	Balance       float64 `json:"balance"`
	UserID        uint    `json:"user_id"`
}

type AccountStore interface {
	CreateAccount(ctx context.Context, account *Account) error
	GetAccountByID(context.Context, int64) (Account, error)
	GetAccountByNumber(context.Context, int64) (Account, error)
	UpdateAccountDetails(context.Context, Account) error
	GetUserByAccountNumber(context.Context, uint) (*users.User, error)
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

// GetUserByAccountDetails retrieves a user by their account details
func (s *AccountService) GetUserByAccountNumber(ctx context.Context, accountNumber uint) (*users.User, error) {
	user, err := s.Store.GetUserByAccountNumber(ctx, accountNumber)
	if err != nil {
		log.Printf("Error fetching user by account details: %v", err)
		return nil, err
	}

	return user, nil
}

func (s *AccountService) CreateAccount(ctx context.Context, account *Account) error {
	if err := s.Store.CreateAccount(ctx, account); err != nil {
		log.Printf("Error creating account: %v", err)
		return err
	}

	return nil
}

func (s *AccountService) GetAccountByID(ctx context.Context, accountID int64) (Account, error) {

	account, err := s.Store.GetAccountByID(ctx, accountID)
	if err != nil {
		log.Printf("Error fetching account with ID %v: %v", accountID, err)
		return account, err
	}

	return account, nil
}

func (s *AccountService) GetAccountByNumber(ctx context.Context, accountNumber int64) (Account, error) {
	account, err := s.Store.GetAccountByNumber(ctx, accountNumber)
	if err != nil {
		log.Printf("Error fetching account with number %d: %v", accountNumber, err)
		return account, err
	}

	return account, nil
}

func (s *AccountService) UpdateAccountDetails(ctx context.Context, account Account) error {
	if err := s.Store.UpdateAccountDetails(ctx, account); err != nil {
		log.Printf("Error updating account: %v", err)
		return err
	}
	return nil
}
