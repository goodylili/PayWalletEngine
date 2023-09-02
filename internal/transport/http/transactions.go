package http

import (
	"PayWalletEngine/internal/transactions"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, transaction *transactions.Transaction) error
	GetTransactionByTransactionID(ctx context.Context, transactionID int64) (*transactions.Transaction, error)
	GetTransactionsBySender(ctx context.Context, senderAccountNumber string) ([]transactions.Transaction, error)
	GetTransactionsByReceiver(ctx context.Context, receiverAccountNumber string) ([]transactions.Transaction, error)
	UpdateTransaction(ctx context.Context, transaction *transactions.Transaction) error
	DeleteTransactionByID(ctx context.Context, transactionID int64) error
	GetTransactionByReference(ctx context.Context, reference int64) (*transactions.Transaction, error)
}

// CreateTransaction handles the creation of a new transaction.
func (h *Handler) CreateTransaction(writer http.ResponseWriter, request *http.Request) {
	var txn transactions.Transaction
	if err := json.NewDecoder(request.Body).Decode(&txn); err != nil {
		return
	}
	err := h.Transaction.CreateTransaction(request.Context(), &txn)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.NewEncoder(writer).Encode(txn); err != nil {
		log.Panicln(err)
	}
}

// GetTransactionByTransactionID handles the retrieval of a single transaction
func (h *Handler) GetTransactionByTransactionID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	stringID := vars["transaction_id"]
	if stringID == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		writer.Write([]byte("There's an error with the transaction id"))
	}

	txn, err := h.Transaction.GetTransactionByTransactionID(request.Context(), id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(txn)
	if err != nil {
		log.Panic(err)
	}
}

// GetTransactionsBySender handles the retrieval of all transactions made by a specific sender.
func (h *Handler) GetTransactionsBySender(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	senderAccountNumber := vars["accountNumber"]
	if senderAccountNumber == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	txns, err := h.Transaction.GetTransactionsBySender(request.Context(), senderAccountNumber)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(txns)
	if err != nil {
		log.Panic(err)
	}

}

// GetTransactionsByReceiver handles the retrieval of all transactions received by a specific receiver.
func (h *Handler) GetTransactionsByReceiver(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	receiverAccountNumber := vars["accountNumber"]
	if receiverAccountNumber == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	txns, err := h.Transaction.GetTransactionsByReceiver(request.Context(), receiverAccountNumber)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(txns)
	if err != nil {
		log.Panic(err)
	}
}

// UpdateTransaction handles the updating of an existing transaction.
func (h *Handler) UpdateTransaction(writer http.ResponseWriter, request *http.Request) {
	var txn transactions.Transaction
	if err := json.NewDecoder(request.Body).Decode(&txn); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Transaction.UpdateTransaction(request.Context(), &txn)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(txn)
	if err != nil {
		log.Panic(err)
	}
}

// DeleteTransactionByID handles the deletion of a transaction identified by its transaction ID.
func (h *Handler) DeleteTransactionByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	stringID := vars["transaction_id"]
	if stringID == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		writer.Write([]byte("There's an error with the transaction id"))
		return
	}

	err = h.Transaction.DeleteTransactionByID(request.Context(), id)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

// GetTransactionByReference handles the retrieval of a single transaction by its reference number.
func (h *Handler) GetTransactionByReference(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	stringReference := vars["transaction_reference"]
	if stringReference == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	reference, err := strconv.ParseInt(stringReference, 10, 64)
	if err != nil {
		writer.Write([]byte("There's an error with the reference number"))
		return
	}

	txn, err := h.Transaction.GetTransactionByReference(request.Context(), reference)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(txn)
	if err != nil {
		log.Panic(err)
	}
}
