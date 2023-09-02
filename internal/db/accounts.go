package db

import (
	"PayWalletEngine/internal/accounts"
	"PayWalletEngine/internal/users"
	"context"
	crypto "crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
	"math/rand"
	"time"
)

type Account struct {
	gorm.Model
	AccountOwnerID int     `gorm:"foreignkey:UserID"`
	AccountID      string  `gorm:"type:varchar(100);uniqueIndex"`
	AccountNumber  string  `gorm:"type:varchar(100);uniqueIndex"`
	AccountType    string  `gorm:"type:varchar(50)"`
	Balance        float64 `gorm:"type:decimal(10,2)"`
	Currency       string  `gorm:"type:varchar(50)"`
	AccountStatus  string  `gorm:"type:varchar(50)"`
}

// GenerateAccountNumber generates a unique 10-digit account number
func GenerateAccountNumber() (string, error) {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Int63n(1e10)
	uid, err := uuid.NewRandomFromReader(crypto.Reader)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256([]byte(fmt.Sprintf("%d%s", randomInt, uid.String())))
	hashInt := big.NewInt(0)
	hashInt.SetBytes(hash[:])
	accountNumber := hashInt.Mod(hashInt, big.NewInt(1e10)).Int64()
	return fmt.Sprintf("%010d", accountNumber), nil
}

// CreateAccount creates a new account in the database.
func (d *Database) CreateAccount(ctx context.Context, account *accounts.Account) error {
	accountNumber, err := GenerateAccountNumber()
	if err != nil {
		return err
	}
	account.AccountNumber = accountNumber
	account.AccountID = uuid.New().String()
	dbAccount := Account{
		AccountOwnerID: account.AccountOwner.UserID,
		AccountID:      account.AccountID,
		AccountNumber:  account.AccountNumber,
		AccountType:    account.AccountType,
		Balance:        account.Balance,
		Currency:       account.Currency,
		AccountStatus:  account.AccountStatus,
	}
	return d.Client.WithContext(ctx).Create(&dbAccount).Error
}

// UpdateAccountDetails updates an account in the database.
func (d *Database) UpdateAccountDetails(ctx context.Context, account accounts.Account) error {
	var a Account
	err := d.Client.WithContext(ctx).Where("account_id = ?", account.AccountID).First(&a).Error
	if err != nil {
		return err
	}
	dbAccount := Account{
		AccountOwnerID: account.AccountOwner.UserID,
		AccountID:      account.AccountID,
		AccountNumber:  account.AccountNumber,
		AccountType:    account.AccountType,
		Balance:        account.Balance,
		Currency:       account.Currency,
		AccountStatus:  account.AccountStatus,
	}
	return d.Client.WithContext(ctx).Save(&dbAccount).Error
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

// DeleteAccountDetails deletes an account by its AccountID from the database.
func (d *Database) DeleteAccountDetails(ctx context.Context, accountID int64) error {
	return d.Client.WithContext(ctx).Where("account_id = ?", accountID).Delete(&Account{}).Error
}

func (d *Database) GetAccountByID(ctx context.Context, s int64) (accounts.Account, error) {
	var a Account
	err := d.Client.WithContext(ctx).Where("account_id = ?", s).First(&a).Error
	if err != nil {
		return accounts.Account{}, err
	}
	return accounts.Account{
		AccountOwner: users.User{
			UserID: a.AccountOwnerID,
		},
		AccountID:     a.AccountID,
		AccountNumber: a.AccountNumber,
		AccountType:   a.AccountType,
		Balance:       a.Balance,
		Currency:      a.Currency,
		AccountStatus: a.AccountStatus,
	}, nil
}

func (d *Database) GetAccountByNumber(ctx context.Context, s int64) (accounts.Account, error) {
	var a Account
	err := d.Client.WithContext(ctx).Where("account_number = ?", s).First(&a).Error
	if err != nil {
		return accounts.Account{}, err
	}
	return accounts.Account{
		AccountOwner: users.User{
			UserID: a.AccountOwnerID,
		},
		AccountID:     a.AccountID,
		AccountNumber: a.AccountNumber,
		AccountType:   a.AccountType,
		Balance:       a.Balance,
		Currency:      a.Currency,
		AccountStatus: a.AccountStatus,
	}, nil
}
