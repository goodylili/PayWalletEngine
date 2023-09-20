package db

import (
	"PayWalletEngine/internal/accounts"
	"PayWalletEngine/internal/transactions"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Amount        float64 `gorm:"type:decimal(10,2);not null"`
	PaymentMethod string  `gorm:"type:varchar(50);not null"`
	Type          string  `gorm:"type:varchar(50);not null"` // "credit", "debit", or "transfer"
	Status        string  `gorm:"type:varchar(50);not null"`
	Description   string  `gorm:"type:varchar(255)"`
	Reference     string  `gorm:"type:varchar(100);uniqueIndex"`

	// Add Sender and Receiver IDs
	SenderID   uint
	ReceiverID uint

	// Specify the relationship
	Sender   Account `gorm:"foreignKey:sender_id"`
	Receiver Account `gorm:"foreignKey:receiver_id"`
}

// GetTransactionByReference retrieves a transaction by its reference
func (d *Database) GetTransactionByReference(ctx context.Context, reference string) (*transactions.Transaction, error) {
	var t Transaction
	err := d.Client.WithContext(ctx).Where("reference = ?", reference).First(&t).Error
	if err != nil {
		return nil, err
	}
	return &transactions.Transaction{
		SenderID:      t.Sender.ID,
		ReceiverID:    t.Receiver.ID,
		Amount:        t.Amount,
		Type:          t.Type,
		PaymentMethod: t.PaymentMethod,
		Status:        t.Status,
		Description:   t.Description,
		Reference:     t.Reference,
	}, nil
}

// GetTransactionsFromAccount retrieves the transactions a specific account made to the database
func (d *Database) GetTransactionsFromAccount(ctx context.Context, accountNumber int64) ([]transactions.Transaction, error) {
	var t []Transaction
	err := d.Client.WithContext(ctx).Joins("Sender").Where("account_number = ?", accountNumber).Find(&t).Error
	if err != nil {
		return nil, err
	}
	var transactionsList []transactions.Transaction
	for _, transaction := range t {
		transactionsList = append(transactionsList, transactions.Transaction{
			SenderID:      transaction.Sender.ID,
			ReceiverID:    transaction.Receiver.ID,
			Amount:        transaction.Amount,
			Type:          transaction.Type,
			PaymentMethod: transaction.PaymentMethod,
			Status:        transaction.Status,
			Description:   transaction.Description,
			Reference:     transaction.Reference,
		})
	}
	return transactionsList, nil
}

// CreditAccount credits an account for a transaction
func (d *Database) CreditAccount(ctx context.Context, receiverAccountNumber int64, amount float64, description string, paymentMethod string) (transactions.Transaction, error) {

	reference, err := transactions.GenerateTransactionRef()
	if err != nil {
		return transactions.Transaction{}, err
	}

	tx := d.Client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var receiverAccount accounts.Account
	if err := tx.WithContext(ctx).Where("account_number = ?", receiverAccountNumber).First(&receiverAccount).Error; err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}

	t := Transaction{
		ReceiverID:    receiverAccount.ID,
		Amount:        amount,
		PaymentMethod: paymentMethod,
		Status:        "Pending",
		Type:          "Credit",
		Description:   description,
		Reference:     reference,
	}

	if err := tx.WithContext(ctx).Create(&t).Error; err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}

	receiverAccount.Balance += amount
	if err := tx.WithContext(ctx).Save(&receiverAccount).Error; err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}

	t.Status = "Completed"
	if err := tx.WithContext(ctx).Save(&t).Error; err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}

	return transactions.Transaction{
		SenderID:      t.Sender.ID,
		ReceiverID:    t.Receiver.ID,
		Amount:        t.Amount,
		Type:          t.Type,
		PaymentMethod: t.PaymentMethod,
		Status:        t.Status,
		Description:   t.Description,
		Reference:     t.Reference,
	}, nil
}

// DebitAccount debits the specified account
func (d *Database) DebitAccount(ctx context.Context, senderAccountNumber int64, amount float64, description string, paymentMethod string) (transactions.Transaction, error) {

	reference, err := transactions.GenerateTransactionRef()

	if err != nil {
		return transactions.Transaction{}, err
	}

	tx := d.Client.Begin()
	// Fetch the sender's account details
	var senderAccount accounts.Account
	if err := tx.WithContext(ctx).Where("account_number = ?", senderAccountNumber).First(&senderAccount).Error; err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}

	// Check if there are sufficient funds in the sender's account
	if senderAccount.Balance < amount {
		tx.Rollback()
		return transactions.Transaction{}, fmt.Errorf("insufficient funds in account")
	}

	t := Transaction{
		SenderID:      senderAccount.ID,
		Amount:        amount,
		PaymentMethod: paymentMethod,
		Status:        "Pending",
		Type:          "Debit",
		Description:   description,
		Reference:     reference,
	}

	// Save transaction in database
	if err := tx.WithContext(ctx).Create(&t).Error; err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}

	// Deduct the amount from the sender's account balance
	senderAccount.Balance -= amount
	if err := tx.WithContext(ctx).Save(&senderAccount).Error; err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}

	t.Status = "Completed"
	if err := tx.WithContext(ctx).Save(&t).Error; err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}

	tx.Commit()
	return transactions.Transaction{
		SenderID:      t.Sender.ID,
		ReceiverID:    t.Receiver.ID,
		Amount:        t.Amount,
		Type:          t.Type,
		PaymentMethod: t.PaymentMethod,
		Status:        t.Status,
		Description:   t.Description,
		Reference:     t.Reference,
	}, nil

}

// TransferFunds transfers funds by crediting and debiting specified users
func (d *Database) TransferFunds(ctx context.Context, senderAccountNumber int64, receiverAccountNumber int64, amount float64, description string, paymentMethod string) (transactions.Transaction, error) {
	// Begin transactions and generate ID and reference
	tx := d.Client.Begin()

	reference, err := transactions.GenerateTransactionRef()
	if err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}

	// Fetch the sender's account details
	var senderAccount accounts.Account
	if err := tx.WithContext(ctx).Where("account_number = ?", senderAccountNumber).First(&senderAccount).Error; err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}

	// Check if there are sufficient funds in the sender's account
	if senderAccount.Balance < amount {
		tx.Rollback()
		return transactions.Transaction{}, fmt.Errorf("insufficient funds in account")
	}

	// Debit the sender's account
	senderAccount.Balance -= amount
	if err := tx.WithContext(ctx).Save(&senderAccount).Error; err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}
	// Fetch the receiver's account details
	var receiverAccount accounts.Account
	if err := tx.WithContext(ctx).Where("account_number = ?", receiverAccountNumber).First(&receiverAccount).Error; err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}

	// Create a transaction with status "Pending"
	t := Transaction{
		SenderID:      senderAccount.ID,
		ReceiverID:    receiverAccount.ID,
		Amount:        amount,
		PaymentMethod: paymentMethod, // or any other required method
		Status:        "Pending",
		Type:          "Transfer",
		Description:   description,
		Reference:     reference,
	}

	// Create a new transaction
	if err := tx.WithContext(ctx).Create(&t).Error; err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}

	// Credit the receiver's account
	receiverAccount.Balance += amount
	if err := tx.WithContext(ctx).Save(&receiverAccount).Error; err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}

	// Update the transaction status to "Completed"
	t.Status = "Completed"
	if err := tx.WithContext(ctx).Save(&t).Error; err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}

	tx.Commit()

	return transactions.Transaction{
		SenderID:      t.Sender.ID,
		ReceiverID:    t.Receiver.ID,
		Amount:        t.Amount,
		Type:          t.Type,
		PaymentMethod: t.PaymentMethod,
		Status:        t.Status,
		Description:   t.Description,
		Reference:     t.Reference,
	}, nil
}
