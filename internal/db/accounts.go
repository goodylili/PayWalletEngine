package db

import (
	"PayWalletEngine/internal/accounts"
	"PayWalletEngine/internal/users"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountNumber string  `gorm:"type:varchar(100);uniqueIndex;column:account_number"`
	AccountType   string  `gorm:"type:varchar(50)"`
	Balance       float64 `gorm:"type:decimal(10,2)"`
	UserID        uint    `gorm:"column:user_id"`
}

// CreateAccount updates an existing account in the database within a transaction.
func (d *Database) CreateAccount(ctx context.Context, account *accounts.Account) error {
	if account.UserID == 0 {
		return fmt.Errorf("UserID is required to update an account")
	}

	// Start a database transaction
	tx := d.Client.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Check if an account with the provided UserID already exists within the transaction
	var dbAccount Account
	err := tx.WithContext(ctx).Where("user_id = ?", account.UserID).First(&dbAccount).Error
	if err != nil {
		tx.Rollback() // Rollback the transaction on error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("account with the provided UserID does not exist")
		}
		return err
	}

	// Update the existing account within the transaction
	dbAccount.AccountType = account.AccountType
	dbAccount.UserID = account.UserID
	dbAccount.Balance = account.Balance
	accountNumber, err := accounts.GenerateAccountNumber()
	if err != nil {
		tx.Rollback() // Rollback the transaction on error
		return err
	}
	account.AccountNumber = accountNumber

	// Save the changes to the database within the transaction
	if err := tx.WithContext(ctx).Save(&dbAccount).Error; err != nil {
		tx.Rollback() // Rollback the transaction on error
		return err
	}

	// Commit the transaction if all operations were successful
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

// UpdateAccountDetails updates an existing account in the database within a transaction.
func (d *Database) UpdateAccountDetails(ctx context.Context, account accounts.Account) error {
	tx := d.Client.WithContext(ctx).Begin() // Start a new transaction

	var a Account
	err := tx.Where("id = ?", account.ID).First(&a).Error
	if err != nil {
		tx.Rollback() // Rollback transaction on error
		return err
	}

	// Update account details
	a.AccountNumber = account.AccountNumber
	a.AccountType = account.AccountType
	a.Balance = account.Balance
	a.UserID = account.UserID

	err = tx.Save(&a).Error
	if err != nil {
		tx.Rollback() // Rollback transaction on error
		return err
	}

	tx.Commit() // Commit the transaction
	return nil
}

// GetAccountByID retrieves an account by its ID
func (d *Database) GetAccountByID(ctx context.Context, id int64) (accounts.Account, error) {
	var a Account
	err := d.Client.WithContext(ctx).Where("id = ?", id).First(&a).Error
	if err != nil {
		return accounts.Account{}, err
	}
	return accounts.Account{
		ID:            a.ID,
		AccountNumber: a.AccountNumber,
		AccountType:   a.AccountType,
		Balance:       a.Balance,
		UserID:        a.UserID,
	}, nil
}

// GetAccountByNumber retrieves an account by its account number
func (d *Database) GetAccountByNumber(ctx context.Context, accountNumber int64) (accounts.Account, error) {
	var a Account
	err := d.Client.WithContext(ctx).Where("account_number = ?", accountNumber).First(&a).Error
	if err != nil {
		return accounts.Account{}, err
	}
	return accounts.Account{
		ID:            a.ID,
		AccountNumber: a.AccountNumber,
		AccountType:   a.AccountType,
		Balance:       a.Balance,
		UserID:        a.UserID,
	}, nil
}

// GetUserByAccountNumber retrieves a user by their account details
func (d *Database) GetUserByAccountNumber(ctx context.Context, accountNumber uint) (*users.User, error) {
	var acct Account
	err := d.Client.WithContext(ctx).Where("account_number = ?", accountNumber).First(&acct).Error
	if err != nil {
		return nil, err
	}

	var user User
	err = d.Client.WithContext(ctx).First(&user, acct.UserID).Error
	if err != nil {
		return nil, err
	}

	return &users.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		IsActive: user.IsActive,
	}, nil
}

// GetAccountsByUserID retrieves all accounts associated with a user
func (d *Database) GetAccountsByUserID(ctx context.Context, userID uint) ([]*accounts.Account, error) {
	var userAccounts []*accounts.Account

	// Retrieve all accounts associated with the provided userID
	err := d.Client.WithContext(ctx).Where("user_id = ?", userID).Find(&userAccounts).Error
	if err != nil {
		return nil, err
	}

	return userAccounts, nil
}
