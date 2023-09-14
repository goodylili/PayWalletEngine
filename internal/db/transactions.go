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
	Sender        accounts.Account `gorm:"foreignkey:AccountNumber"`
	Receiver      accounts.Account `gorm:"foreignkey:AccountNumber"`
	Amount        float64          `gorm:"type:decimal(10,2);not null"`
	PaymentMethod string           `gorm:"type:varchar(50);not null"`
	Status        string           `gorm:"type:varchar(50);not null"`
	Description   string           `gorm:"type:varchar(255)"`
	TransactionID int64            `gorm:"type:varchar(100);unique_index"`
	Reference     string           `gorm:"type:varchar(100);unique_index"`
}

// GetTransactionByReference retrieves a transaction by its reference
func (d *Database) GetTransactionByReference(ctx context.Context, reference int64) (*transactions.Transaction, error) {
	var t Transaction
	err := d.Client.WithContext(ctx).Where("reference = ?", reference).First(&t).Error
	if err != nil {
		return nil, err
	}
	return &transactions.Transaction{
		Sender:        t.Sender,
		Receiver:      t.Receiver,
		TransactionID: t.TransactionID,
		Amount:        t.Amount,
		PaymentMethod: t.PaymentMethod,
		Status:        t.Status,
		Description:   t.Description,
		Reference:     t.Reference,
	}, nil
}

// GetTransactionByTransactionID retrieves a transaction by it's ID
func (d *Database) GetTransactionByTransactionID(ctx context.Context, transactionID int64) (*transactions.Transaction, error) {
	var t Transaction
	err := d.Client.WithContext(ctx).Where("transaction_id = ?", transactionID).First(&t).Error
	if err != nil {
		return nil, err
	}
	return &transactions.Transaction{
		Sender:        t.Sender,
		Receiver:      t.Receiver,
		TransactionID: t.TransactionID,
		Amount:        t.Amount,
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
			Sender:        transaction.Sender,
			Receiver:      transaction.Receiver,
			TransactionID: transaction.TransactionID,
			Amount:        transaction.Amount,
			PaymentMethod: transaction.PaymentMethod,
			Status:        transaction.Status,
			Description:   transaction.Description,
			Reference:     transaction.Reference,
		})
	}
	return transactionsList, nil
}

// CreditAccount credits an account for a transaction
func (d *Database) CreditAccount(ctx context.Context, receiverAccountNumber int64, amount float64, description string, paymentMethod string) (Transaction, error) {
	// Begin the transaction and generate the ID
	tx := d.Client.Begin()
	transactionID, err := transactions.GenerateTransactionID()
	if err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	reference, err := transactions.GenerateTransactionRef(transactionID)
	if err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	var receiverAccount accounts.Account
	if err := tx.WithContext(ctx).Where("account_number = ?", receiverAccountNumber).First(&receiverAccount).Error; err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	t := Transaction{
		Receiver:      receiverAccount,
		Amount:        amount,
		PaymentMethod: paymentMethod,
		Status:        "Pending",
		Description:   description,
		TransactionID: transactionID,
		Reference:     reference,
	}

	if err := tx.WithContext(ctx).Create(&t).Error; err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	// Update the receiver's account balance
	receiverAccount.Balance += amount
	if err := tx.WithContext(ctx).Save(&receiverAccount).Error; err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	t.Status = "Completed"
	if err := tx.WithContext(ctx).Save(&t).Error; err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	tx.Commit()
	return t, nil
}

// DebitAccount debits the specified account
func (d *Database) DebitAccount(ctx context.Context, senderAccountNumber int64, amount float64, description string, paymentMethod string) (Transaction, error) {
	tx := d.Client.Begin()

	// Generate transaction ID and Reference
	transactionID, err := transactions.GenerateTransactionID()
	if err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	reference, err := transactions.GenerateTransactionRef(transactionID)
	if err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	// Fetch the sender's account details
	var senderAccount accounts.Account
	if err := tx.WithContext(ctx).Where("account_number = ?", senderAccountNumber).First(&senderAccount).Error; err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	// Check if there are sufficient funds in the sender's account
	if senderAccount.Balance < amount {
		tx.Rollback()
		return Transaction{}, fmt.Errorf("insufficient funds in account")
	}

	t := Transaction{
		Sender:        senderAccount,
		Amount:        amount,
		PaymentMethod: paymentMethod,
		Status:        "Pending",
		Description:   description,
		TransactionID: transactionID,
		Reference:     reference,
	}

	// Save transaction in database
	if err := tx.WithContext(ctx).Create(&t).Error; err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	// Deduct the amount from the sender's account balance
	senderAccount.Balance -= amount
	if err := tx.WithContext(ctx).Save(&senderAccount).Error; err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	t.Status = "Completed"
	if err := tx.WithContext(ctx).Save(&t).Error; err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	tx.Commit()
	return t, nil
}

// TransferFunds transfers funds by crediting and debiting specified users
func (d *Database) TransferFunds(ctx context.Context, senderAccountNumber int64, receiverAccountNumber int64, amount float64, description string, paymentMethod string) (Transaction, error) {
	// Begin transactions and generate ID and reference
	tx := d.Client.Begin()

	// Generate transaction ID and Reference
	transactionID, err := transactions.GenerateTransactionID()
	if err != nil {
		tx.Rollback()
		return Transaction{}, err
	}
	reference, err := transactions.GenerateTransactionRef(transactionID)
	if err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	// Fetch the sender's account details
	var senderAccount accounts.Account
	if err := tx.WithContext(ctx).Where("account_number = ?", senderAccountNumber).First(&senderAccount).Error; err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	// Check if there are sufficient funds in the sender's account
	if senderAccount.Balance < amount {
		tx.Rollback()
		return Transaction{}, fmt.Errorf("insufficient funds in account")
	}

	// Debit the sender's account
	senderAccount.Balance -= amount
	if err := tx.WithContext(ctx).Save(&senderAccount).Error; err != nil {
		tx.Rollback()
		return Transaction{}, err
	}
	// Fetch the receiver's account details
	var receiverAccount accounts.Account
	if err := tx.WithContext(ctx).Where("account_number = ?", receiverAccountNumber).First(&receiverAccount).Error; err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	// Create a transaction with status "Pending"
	t := Transaction{
		Sender:        senderAccount,
		Receiver:      receiverAccount,
		Amount:        amount,
		PaymentMethod: paymentMethod, // or any other required method
		Status:        "Pending",
		Description:   description,
		TransactionID: transactionID,
		Reference:     reference,
	}

	// Create a new transaction
	if err := tx.WithContext(ctx).Create(&t).Error; err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	// Credit the receiver's account
	receiverAccount.Balance += amount
	if err := tx.WithContext(ctx).Save(&receiverAccount).Error; err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	// Update the transaction status to "Completed"
	t.Status = "Completed"
	if err := tx.WithContext(ctx).Save(&t).Error; err != nil {
		tx.Rollback()
		return Transaction{}, err
	}

	tx.Commit()
	return t, nil
}
