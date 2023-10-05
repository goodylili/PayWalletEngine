package db

import (
	"PayWalletEngine/internal/accounts"
	"PayWalletEngine/internal/transactions"
	"PayWalletEngine/internal/users"
	"context"
	"gorm.io/gorm"
	"time"
)

type Transactions struct {
	TransactionID uint    `gorm:"primarykey"`
	Amount        float64 `gorm:"type:decimal(10,2);not null"`
	PaymentMethod string  `gorm:"type:varchar(50);not null"`
	Type          string  `gorm:"type:varchar(50);not null"`
	Status        string  `gorm:"type:varchar(50);not null"`
	Description   string  `gorm:"type:varchar(255)"`
	Reference     string  `gorm:"type:varchar(100);uniqueIndex"`

	SenderAccountNumber   uint `gorm:"type:integer"`
	ReceiverAccountNumber uint `gorm:"type:integer"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
			Username: usr.Username,
			Email:    usr.Email,
			Password: usr.Password,
			IsActive: false,
		}, &accounts.Account{
			ID:            acct.ID,
			AccountNumber: acct.AccountNumber,
			AccountType:   acct.AccountType,
			Balance:       acct.Balance,
			UserID:        acct.UserID,
		}, &transactions.Transactions{
			SenderAccountNumber:   txn.SenderAccountNumber,
			ReceiverAccountNumber: txn.ReceiverAccountNumber,
			Amount:                txn.Amount,
			Type:                  txn.Type,
			PaymentMethod:         txn.PaymentMethod,
			Status:                txn.Status,
			Description:           txn.Description,
			Reference:             txn.Reference,
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
			SenderAccountNumber:   txn.SenderAccountNumber,
			ReceiverAccountNumber: txn.ReceiverAccountNumber,
			Amount:                txn.Amount,
			Type:                  txn.Type,
			PaymentMethod:         txn.PaymentMethod,
			Status:                txn.Status,
			Description:           txn.Description,
			Reference:             txn.Reference,
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
		SenderAccountNumber:   t.SenderAccountNumber,
		ReceiverAccountNumber: t.ReceiverAccountNumber,
		Amount:                t.Amount,
		Type:                  t.Type,
		PaymentMethod:         t.PaymentMethod,
		Status:                t.Status,
		Description:           t.Description,
		Reference:             t.Reference,
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
			SenderAccountNumber:   transaction.SenderAccountNumber,
			ReceiverAccountNumber: transaction.ReceiverAccountNumber,
			Amount:                transaction.Amount,
			Type:                  transaction.Type,
			PaymentMethod:         transaction.PaymentMethod,
			Status:                transaction.Status,
			Description:           transaction.Description,
			Reference:             transaction.Reference,
		})
	}
	return transactionsList, nil
}

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
	receiverAccount, err := d.creditAccountHelper(tx, ctx, receiverAccountNumber, amount)
	if err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	t := transactions.Transactions{
		ReceiverAccountNumber: receiverAccount.AccountNumber,
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

	if err := tx.WithContext(ctx).Model(&transactions.Transactions{}).Where("transaction_id = ?", t.TransactionID).Update("Status", "Completed").Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	return t, nil
}

func (d *Database) DebitAccount(ctx context.Context, senderAccountNumber int64, amount float64, description string, paymentMethod string) (transactions.Transactions, error) {
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

	senderAccount, err := d.debitAccountHelper(tx, ctx, senderAccountNumber, amount)
	if err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	t := transactions.Transactions{
		SenderAccountNumber: senderAccount.AccountNumber,
		Amount:              amount,
		PaymentMethod:       paymentMethod,
		Status:              "Pending",
		Type:                "Debit",
		Description:         description,
		Reference:           reference,
	}

	if err := tx.WithContext(ctx).Create(&t).Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	if err := tx.WithContext(ctx).Model(&transactions.Transactions{}).Where("transaction_id = ?", t.TransactionID).Update("Status", "Completed").Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	return t, nil
}

func (d *Database) TransferFunds(ctx context.Context, senderAccountNumber int64, receiverAccountNumber int64, amount float64, description string, paymentMethod string) (transactions.Transactions, error) {
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

	senderAccount, err := d.debitAccountHelper(tx, ctx, senderAccountNumber, amount)
	if err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	receiverAccount, err := d.creditAccountHelper(tx, ctx, receiverAccountNumber, amount)
	if err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	t := transactions.Transactions{
		SenderAccountNumber:   senderAccount.AccountNumber,
		ReceiverAccountNumber: receiverAccount.AccountNumber,
		Amount:                amount,
		PaymentMethod:         paymentMethod,
		Status:                "Pending",
		Type:                  "Transfer",
		Description:           description,
		Reference:             reference,
	}

	if err := tx.WithContext(ctx).Create(&t).Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	if err := tx.WithContext(ctx).Model(&transactions.Transactions{}).Where("transaction_id = ?", t.TransactionID).Update("Status", "Completed").Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return transactions.Transactions{}, err
	}

	return t, nil
}
