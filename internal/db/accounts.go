package db

import (
	"PayWalletEngine/internal/accounts"
	"context"
	"github.com/google/uuid"
)

type Account struct {
	AccountOwnerID uint    `gorm:"primaryKey;foreignKey:AccountOwnerID"` // Correct foreign key reference
	AccountNumber  string  `gorm:"type:varchar(100);uniqueIndex"`
	AccountType    string  `gorm:"type:varchar(50)"`
	Balance        float64 `gorm:"type:decimal(10,2)"`
}

// CreateAccount creates a new account in the database.
func (d *Database) CreateAccount(ctx context.Context, account *accounts.Account) error {
	accountNumber, err := accounts.GenerateAccountNumber()
	if err != nil {
		return err
	}
	account.AccountNumber = accountNumber
	account.AccountID = uuid.New().String()
	dbAccount := Account{
		AccountOwnerID: account.AccountOwnerID.ID,
		AccountNumber:  account.AccountNumber,
		AccountType:    account.AccountType,
		Balance:        account.Balance,
	}
	return d.Client.WithContext(ctx).Create(&dbAccount).Error
}

func (d *Database) GetAccountByID(ctx context.Context, id int64) (accounts.Account, error) {
	var a Account
	err := d.Client.WithContext(ctx).Where("account_id = ?", id).First(&a).Error
	if err != nil {
		return accounts.Account{}, err
	}
	return accounts.Account{
		AccountNumber: a.AccountNumber,
		AccountType:   a.AccountType,
		Balance:       a.Balance,
	}, nil
}

// DeleteAccountDetails deletes an account by its AccountID from the database.
func (d *Database) DeleteAccountDetails(ctx context.Context, accountID int64) error {
	return d.Client.WithContext(ctx).Where("account_id = ?", accountID).Delete(&Account{}).Error
}

func (d *Database) GetAccountByNumber(ctx context.Context, s int64) (accounts.Account, error) {
	var a Account
	err := d.Client.WithContext(ctx).Where("account_number = ?", s).First(&a).Error
	if err != nil {
		return accounts.Account{}, err
	}
	return accounts.Account{
		AccountNumber: a.AccountNumber,
		AccountType:   a.AccountType,
		Balance:       a.Balance,
	}, nil
}

// UpdateAccountBalance updates the balance of an account in the database.
func (d *Database) UpdateAccountBalance(ctx context.Context, accountID string, newBalance float64) error {
	var a Account
	err := d.Client.WithContext(ctx).Where("account_id = ?", accountID).First(&a).Error
	if err != nil {
		return err
	}
	return d.Client.WithContext(ctx).Model(&Account{}).Where("account_id = ?", accountID).Update("balance", newBalance).Error
}

// UpdateAccountDetails updates an account in the database.
func (d *Database) UpdateAccountDetails(ctx context.Context, account accounts.Account) error {
	var a Account
	err := d.Client.WithContext(ctx).Where("account_id = ?", account.AccountID).First(&a).Error
	if err != nil {
		return err
	}
	dbAccount := Account{
		AccountOwnerID: account.AccountOwnerID.ID,
		AccountNumber:  account.AccountNumber,
		AccountType:    account.AccountType,
		Balance:        account.Balance,
	}
	return d.Client.WithContext(ctx).Save(&dbAccount).Error
}
