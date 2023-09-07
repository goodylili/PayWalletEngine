package db

import (
	"PayWalletEngine/internal/accounts"
	"PayWalletEngine/internal/transactions"
	"context"
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
	Reference     int64            `gorm:"type:varchar(100);unique_index"`
}

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

func (d *Database) GetTransactionByTransactionID(ctx context.Context, transactionID int64) (*transactions.Transaction, error) {
	var t Transaction
	err := d.Client.WithContext(ctx).Where("reference = ?", transactionID).First(&t).Error
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

func (d *Database) GetTransactionsBySender(ctx context.Context, senderAccountNumber string) ([]transactions.Transaction, error) {
	var t []Transaction
	err := d.Client.WithContext(ctx).Joins("Sender").Where("sender_account_number = ?", senderAccountNumber).Find(&t).Error
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

func (d *Database) GetTransactionsByReceiver(ctx context.Context, receiverAccountNumber string) ([]transactions.Transaction, error) {
	var t []Transaction
	err := d.Client.WithContext(ctx).Joins("Receiver").Where("receiver_account_number = ?", receiverAccountNumber).Find(&t).Error
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
