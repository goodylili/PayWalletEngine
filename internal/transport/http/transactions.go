package http

import (
	"PayWalletEngine/internal/transactions"
	"context"
	"net/http"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, transaction *transactions.Transaction) error
	GetTransactionByTransactionID(ctx context.Context, transactionID int64) (*transactions.Transaction, error)
	GetTransactionsBySender(ctx context.Context, senderAccountNumber string) ([]transactions.Transaction, error)
	GetTransactionsByReceiver(ctx context.Context, receiverAccountNumber string) ([]transactions.Transaction, error)
	UpdateTransaction(ctx context.Context, transaction *transactions.Transaction) error
	DeleteTransactionByID(ctx context.Context, transactionID int64) error
}

func (h *Handler) CreateTransaction(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) GetTransactionByTransactionID(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) GetTransactionsBySender(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) GetTransactionsByReceiver(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) UpdateTransaction(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) DeleteTransactionByID(writer http.ResponseWriter, request *http.Request) {

}
