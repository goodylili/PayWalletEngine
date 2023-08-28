package db

import (
	"PayWalletEngine/internal/accounts"
	"context"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Sender        accounts.Account `gorm:"foreignkey:AccountNumber"`
	Receiver      accounts.Account `gorm:"foreignkey:AccountNumber"`
	TransactionID string           `gorm:"type:varchar(100);unique_index"`
	Amount        float64          `gorm:"type:decimal(10,2);not null"`
	Currency      string           `gorm:"type:varchar(10);not null"`
	PaymentMethod string           `gorm:"type:varchar(50);not null"`
	Status        string           `gorm:"type:varchar(50);not null"`
	Description   string           `gorm:"type:varchar(255)"`
}

// CreateTransaction creates a new transaction in the database.
func (d *Database) CreateTransaction(ctx context.Context, transaction *Transaction) error {
	return d.Client.WithContext(ctx).Create(transaction).Error
}

// GetByTransactionID fetches a transaction by its TransactionID from the database.
func (d *Database) GetByTransactionID(ctx context.Context, transactionID string) (*Transaction, error) {
	var t Transaction
	err := d.Client.WithContext(ctx).Where("transaction_id = ?", transactionID).First(&t).Error
	return &t, err
}

// UpdateTransaction updates a transaction in the database.
func (d *Database) UpdateTransaction(ctx context.Context, transaction *Transaction) error {
	return d.Client.WithContext(ctx).Save(transaction).Error
}

// DeleteTransaction deletes a transaction by its TransactionID from the database.
func (d *Database) DeleteTransaction(ctx context.Context, transactionID string) error {
	return d.Client.WithContext(ctx).Where("transaction_id = ?", transactionID).Delete(&Transaction{}).Error
}

// GetTransactionsBySender fetches all transactions made by a specific sender from the database.
func (d *Database) GetTransactionsBySender(ctx context.Context, senderAccountNumber string) ([]Transaction, error) {
	var transactions []Transaction
	err := d.Client.WithContext(ctx).Where("sender_account_number = ?", senderAccountNumber).Find(&transactions).Error
	return transactions, err
}

// GetTransactionsByReceiver fetches all transactions received by a specific receiver from the database.
func (d *Database) GetTransactionsByReceiver(ctx context.Context, receiverAccountNumber string) ([]Transaction, error) {
	var transactions []Transaction
	err := d.Client.WithContext(ctx).Where("receiver_account_number = ?", receiverAccountNumber).Find(&transactions).Error
	return transactions, err
}
