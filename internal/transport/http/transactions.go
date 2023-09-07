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
	GetTransactionByTransactionID(ctx context.Context, transactionID int64) (*transactions.Transaction, error)
	GetTransactionsBySender(ctx context.Context, senderAccountNumber int64) ([]transactions.Transaction, error)
	GetTransactionsByReceiver(ctx context.Context, receiverAccountNumber int64) ([]transactions.Transaction, error)
	GetTransactionByReference(ctx context.Context, reference int64) (*transactions.Transaction, error)
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
