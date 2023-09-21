package db

import (
	"PayWalletEngine/internal/accounts"
	"context"
	"errors"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountNumber string  `gorm:"type:varchar(100);uniqueIndex;column:account_number"`
	AccountType   string  `gorm:"type:varchar(50)"`
	Balance       float64 `gorm:"type:decimal(10,2)"`
	UserID        uint    `gorm:"column:user_id"`
}

// CreateAccount updates an existing account in the database.
func (d *Database) CreateAccount(ctx context.Context, account *accounts.Account) error {
	if account.UserID == 0 {
		return errors.New("UserID is required to update an account")
	}

	// Check if an account with the provided UserID already exists
	var dbAccount Account
	err := d.Client.WithContext(ctx).Where("user_id = ?", account.UserID).First(&dbAccount).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("account with the provided UserID does not exist")
		}
		return err
	}

	// Update the existing account
	dbAccount.AccountType = account.AccountType
	dbAccount.Balance = account.Balance
	accountNumber, err := accounts.GenerateAccountNumber()
	if err != nil {
		return err
	}
	account.AccountNumber = accountNumber

	return d.Client.WithContext(ctx).Save(&dbAccount).Error
}

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
