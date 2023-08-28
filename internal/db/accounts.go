package db

import (
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
	AccountOwner  users.User `gorm:"foreignkey:AccountOwnerID"`
	AccountID     string     `gorm:"type:varchar(100);uniqueIndex"`
	AccountNumber string     `gorm:"type:varchar(100);uniqueIndex"`
	AccountType   string     `gorm:"type:varchar(50)"`
	Balance       float64    `gorm:"type:decimal(10,2)"`
	Currency      string     `gorm:"type:varchar(50)"`
	AccountStatus string     `gorm:"type:varchar(50)"`
}

// GenerateAccountNumber generates a unique 10-digit account number
func GenerateAccountNumber(email string) (int, error) {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Int63n(1e10)
	emailPlusRandom := fmt.Sprintf("%s%d", email, randomInt)
	uid, err := uuid.NewRandomFromReader(crypto.Reader)
	if err != nil {
		return 0, err
	}
	hash := sha256.Sum256([]byte(emailPlusRandom + uid.String()))
	hashInt := big.NewInt(0)
	hashInt.SetBytes(hash[:])
	accountNumber := hashInt.Mod(hashInt, big.NewInt(1e10)).Int64()
	return int(accountNumber), nil
}

// Create creates a new account in the database.
func (d *Database) Create(ctx context.Context, account *Account) error {
	accountNumber, err := GenerateAccountNumber(account.AccountOwner.Email)
	if err != nil {
		return err
	}
	account.AccountNumber = fmt.Sprintf("%010d", accountNumber)
	return d.Client.WithContext(ctx).Create(account).Error
}

// GetByAccountID fetches an account by its AccountID from the database.
func (d *Database) GetByAccountID(ctx context.Context, accountID string) (*Account, error) {
	var a Account
	err := d.Client.WithContext(ctx).Where("account_id = ?", accountID).First(&a).Error
	return &a, err
}

// GetByAccountNumber fetches an account by its AccountNumber from the database.
func (d *Database) GetByAccountNumber(ctx context.Context, accountNumber string) (*Account, error) {
	var a Account
	err := d.Client.WithContext(ctx).Where("account_number = ?", accountNumber).First(&a).Error
	return &a, err
}

// UpdateAccountDetails updates an account in the database.
func (d *Database) UpdateAccountDetails(ctx context.Context, account *Account) error {
	return d.Client.WithContext(ctx).Save(account).Error
}

// DeleteAccountDetails deletes an account by its ID from the database.
func (d *Database) DeleteAccountDetails(ctx context.Context, ID uint) error {
	return d.Client.WithContext(ctx).Delete(&Account{}, ID).Error
}
