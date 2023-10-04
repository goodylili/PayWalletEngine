package db

import (
	"PayWalletEngine/internal/accounts"
	"PayWalletEngine/internal/transactions"
	"PayWalletEngine/internal/users"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model            `json:"-"`
	Amount                float64 `gorm:"type:decimal(10,2);not null"`
	PaymentMethod         string  `gorm:"type:varchar(50);not null"`
	Type                  string  `gorm:"type:varchar(50);not null"` // "credit", "debit", or "transfer"
	Status                string  `gorm:"type:varchar(50);not null"`
	Description           string  `gorm:"type:varchar(255)"`
	Reference             string  `gorm:"type:varchar(100);uniqueIndex"`
	SenderAccountNumber   int64   `gorm:"column:sender_account_number"`
	ReceiverAccountNumber int64   `gorm:"column:receiver_account_number"`
}

func (d *Database) GetUserAccountAndTransactionByTransactionID(ctx context.Context, transactionID uint) (*users.User, *accounts.Account, *transactions.Transactions, error) {
	var txn Transactions
	err := d.Client.WithContext(ctx).First(&txn, transactionID).Error
	if err != nil {
		return nil, nil, nil, err
	}

	var acct Account
	err = d.Client.WithContext(ctx).Where("account_number = ?", txn.ReceiverAccountNumber).First(&acct).Error
	if err != nil {
		return nil, nil, nil, err
	}

	var usr User
	err = d.Client.WithContext(ctx).First(&usr, acct.UserID).Error
	if err != nil {
		return nil, nil, nil, err
	}

	return &users.User{
			Model:    gorm.Model{},
			Username: "",
			Email:    "",
			Password: "",
			IsActive: false,
			Account:  nil,
		}, &accounts.Account{
			ID:            acct.ID,
			AccountNumber: acct.AccountNumber,
			AccountType:   acct.AccountType,
			Balance:       acct.Balance,
			UserID:        acct.UserID,
		}, &transactions.Transactions{
			SenderAccountID:   txn.SenderAccountNumber,
			ReceiverAccountID: txn.ReceiverAccountNumber,
			Amount:            txn.Amount,
			Type:              txn.Type,
			PaymentMethod:     txn.PaymentMethod,
			Status:            txn.Status,
			Description:       txn.Description,
			Reference:         txn.Reference,
		}, nil
}

func (d *Database) GetAccountByTransactionID(ctx context.Context, transactionID uint) (*accounts.Account, *transactions.Transactions, error) {
	var txn Transactions
	err := d.Client.WithContext(ctx).First(&txn, transactionID).Error
	if err != nil {
		return nil, nil, err
	}

	var acct Account
	err = d.Client.WithContext(ctx).Where("account_number = ?", txn.ReceiverAccountNumber).First(&acct).Error
	if err != nil {
		return nil, nil, err
	}

	return &accounts.Account{
			ID:            acct.ID,
			AccountNumber: acct.AccountNumber,
			AccountType:   acct.AccountType,
			Balance:       acct.Balance,
			UserID:        acct.UserID,
		}, &transactions.Transactions{
			SenderAccountID:   txn.SenderAccountNumber,
			ReceiverAccountID: txn.ReceiverAccountNumber,
			Amount:            txn.Amount,
			Type:              txn.Type,
			PaymentMethod:     txn.PaymentMethod,
			Status:            txn.Status,
			Description:       txn.Description,
			Reference:         txn.Reference,
		}, nil
}

// GetTransactionByReference retrieves a transaction by its reference
func (d *Database) GetTransactionByReference(ctx context.Context, reference string) (*transactions.Transactions, error) {
	var t Transactions
	err := d.Client.WithContext(ctx).Where("reference = ?", reference).First(&t).Error
	if err != nil {
		return nil, err
	}
	return &transactions.Transactions{
		SenderAccountID:   t.SenderAccountNumber,
		ReceiverAccountID: t.ReceiverAccountNumber,
		Amount:            t.Amount,
		Type:              t.Type,
		PaymentMethod:     t.PaymentMethod,
		Status:            t.Status,
		Description:       t.Description,
		Reference:         t.Reference,
	}, nil
}

// GetTransactionsFromAccount retrieves the transactions a specific account made to the database
func (d *Database) GetTransactionsFromAccount(ctx context.Context, accountNumber int64) ([]transactions.Transactions, error) {
	var t []Transactions
	err := d.Client.WithContext(ctx).Where("sender_account_number = ?", accountNumber).Or("receiver_account_number = ?", accountNumber).Find(&t).Error
	if err != nil {
		return nil, err
	}
	var transactionsList []transactions.Transactions
	for _, transaction := range t {
		transactionsList = append(transactionsList, transactions.Transactions{
			SenderAccountID:   transaction.SenderAccountNumber,
			ReceiverAccountID: transaction.ReceiverAccountNumber,
			Amount:            transaction.Amount,
			Type:              transaction.Type,
			PaymentMethod:     transaction.PaymentMethod,
			Status:            transaction.Status,
			Description:       transaction.Description,
			Reference:         transaction.Reference,
		})
	}
	return transactionsList, nil
}

// CreditAccount credits an account for a transaction
func (d *Database) CreditAccount(ctx context.Context, receiverAccountNumber int64, amount float64, description string, paymentMethod string) (transactions.Transactions, error) {
	reference, err := transactions.GenerateTransactionRef()
	if err != nil {
		return transactions.Transactions{}, err
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
		return transactions.Transactions{}, err
	}

	t := Transactions{
		ReceiverAccountNumber: int64(receiverAccount.ID),
		Amount:                amount,
		PaymentMethod:         paymentMethod,
		Status:                "Pending",
		Type:                  "Credit",
		Description:           description,
		Reference:             reference,
	}

	if err := tx.WithContext(ctx).Create(&t).Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	receiverAccount.Balance += amount
	if err := tx.WithContext(ctx).Save(&receiverAccount).Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	t.Status = "Completed"
	if err := tx.WithContext(ctx).Save(&t).Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	return transactions.Transactions{
		SenderAccountID:   t.SenderAccountNumber,
		ReceiverAccountID: t.ReceiverAccountNumber,
		Amount:            t.Amount,
		Type:              t.Type,
		PaymentMethod:     t.PaymentMethod,
		Status:            t.Status,
		Description:       t.Description,
		Reference:         t.Reference,
	}, nil
}

// DebitAccount debits the specified account
func (d *Database) DebitAccount(ctx context.Context, senderAccountNumber int64, amount float64, description string, paymentMethod string) (transactions.Transactions, error) {

	reference, err := transactions.GenerateTransactionRef()

	if err != nil {
		return transactions.Transactions{}, err
	}

	tx := d.Client.Begin()
	// Fetch the sender's account details
	var senderAccount accounts.Account
	if err := tx.WithContext(ctx).Where("account_number = ?", senderAccountNumber).First(&senderAccount).Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	// Check if there are sufficient funds in the sender's account
	if senderAccount.Balance < amount {
		tx.Rollback()
		return transactions.Transactions{}, fmt.Errorf("insufficient funds in account")
	}

	t := Transactions{
		ReceiverAccountNumber: int64(senderAccount.ID),
		Amount:                amount,
		PaymentMethod:         paymentMethod,
		Status:                "Pending",
		Type:                  "Debit",
		Description:           description,
		Reference:             reference,
	}

	// Save transaction in database
	if err := tx.WithContext(ctx).Create(&t).Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	// Deduct the amount from the sender's account balance
	senderAccount.Balance -= amount
	if err := tx.WithContext(ctx).Save(&senderAccount).Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	t.Status = "Completed"
	if err := tx.WithContext(ctx).Save(&t).Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	tx.Commit()
	return transactions.Transactions{
		SenderAccountID:   t.SenderAccountNumber,
		ReceiverAccountID: t.ReceiverAccountNumber,
		Amount:            t.Amount,
		Type:              t.Type,
		PaymentMethod:     t.PaymentMethod,
		Status:            t.Status,
		Description:       t.Description,
		Reference:         t.Reference,
	}, nil

}

// TransferFunds transfers funds by crediting and debiting specified users
func (d *Database) TransferFunds(ctx context.Context, senderAccountNumber int64, receiverAccountNumber int64, amount float64, description string, paymentMethod string) (transactions.Transactions, error) {
	// Begin transactions and generate ID and reference
	tx := d.Client.Begin()

	reference, err := transactions.GenerateTransactionRef()
	if err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	// Fetch the sender's account details
	var senderAccount accounts.Account
	if err := tx.WithContext(ctx).Where("account_number = ?", senderAccountNumber).First(&senderAccount).Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	// Check if there are sufficient funds in the sender's account
	if senderAccount.Balance < amount {
		tx.Rollback()
		return transactions.Transactions{}, fmt.Errorf("insufficient funds in account")
	}

	// Debit the sender's account
	senderAccount.Balance -= amount
	if err := tx.WithContext(ctx).Save(&senderAccount).Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}
	// Fetch the receiver's account details
	var receiverAccount accounts.Account
	if err := tx.WithContext(ctx).Where("account_number = ?", receiverAccountNumber).First(&receiverAccount).Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	// Create a transaction with status "Pending"
	t := Transactions{
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
		return transactions.Transactions{}, err
	}

	// Credit the receiver's account
	receiverAccount.Balance += amount
	if err := tx.WithContext(ctx).Save(&receiverAccount).Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	// Update the transaction status to "Completed"
	t.Status = "Completed"
	if err := tx.WithContext(ctx).Save(&t).Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	tx.Commit()

	return transactions.Transactions{
		SenderAccountID:   t.SenderAccountNumber,
		ReceiverAccountID: t.ReceiverAccountNumber,
		Amount:            t.Amount,
		Type:              t.Type,
		PaymentMethod:     t.PaymentMethod,
		Status:            t.Status,
		Description:       t.Description,
		Reference:         t.Reference,
	}, nil
}
