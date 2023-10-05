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
	AccountNumber uint    `json:"account_number"`
	AccountType   string  `json:"account_type"`
	Balance       float64 `json:"balance"`
	UserID        uint    `json:"user_id"`
}

type AccountStore interface {
	CreateAccount(ctx context.Context, account *Account) error
	GetAccountByID(ctx context.Context, accountID uint) (Account, error)
	GetAccountByNumber(ctx context.Context, accountNumber uint) (Account, error)
	UpdateAccountDetails(ctx context.Context, account Account) error
	GetUserByAccountNumber(ctx context.Context, accountNumber uint) (*users.User, error)
	GetAccountsByUserID(ctx context.Context, userID uint) ([]*Account, error)
}

// AccountService is the blueprint for the account logic
type AccountService struct {
	Store AccountStore
}

func NewAccountService(store AccountStore) AccountService {
	return AccountService{
		Store: store,
	}
}

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

func (s *AccountService) GetAccountByID(ctx context.Context, accountID uint) (Account, error) {
	account, err := s.Store.GetAccountByID(ctx, accountID)
	if err != nil {
		log.Printf("Error fetching account with ID %v: %v", accountID, err)
		return account, err
	}
	return account, nil
}

func (s *AccountService) GetAccountByNumber(ctx context.Context, accountNumber uint) (Account, error) {
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

func (s *AccountService) GetAccountsByUserID(ctx context.Context, userID uint) ([]*Account, error) {
	account, err := s.Store.GetAccountsByUserID(ctx, userID)
	if err != nil {
		log.Printf("Error fetching accounts with userID %v: %v", userID, err)
		return nil, err
	}
	return account, nil
}
