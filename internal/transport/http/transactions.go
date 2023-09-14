package http

import (
	"PayWalletEngine/internal/transactions"
	"PayWalletEngine/utils"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type TransactionService interface {
	GetTransactionByTransactionID(ctx context.Context, transactionID int64) (*transactions.Transaction, error)
	GetTransactionsFromAccount(ctx context.Context, accountNumber int64) ([]transactions.Transaction, error)
	GetTransactionByReference(ctx context.Context, reference string) (*transactions.Transaction, error)
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
func (h *Handler) GetTransactionsFromAccount(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	senderAccountNumberStr := vars["account_number"]
	if senderAccountNumberStr == "" {
		utils.WriteErrorResponse(writer, http.StatusBadRequest, "account number is required")
		return
	}

	senderAccountNumber, err := strconv.ParseInt(senderAccountNumberStr, 10, 64)
	if err != nil {
		utils.WriteErrorResponse(writer, http.StatusBadRequest, "invalid account number format")
		return
	}

	txns, err := h.Transaction.GetTransactionsBySender(request.Context(), senderAccountNumber)
	if err != nil {
		log.Println(err)
		utils.WriteErrorResponse(writer, http.StatusInternalServerError, "internal server error")
		return
	}
	err = json.NewEncoder(writer).Encode(txns)
	if err != nil {
		log.Println("Error encoding response:", err)
		utils.WriteErrorResponse(writer, http.StatusInternalServerError, "internal server error")
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
