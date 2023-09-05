package http

import (
	"PayWalletEngine/internal/accounts"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type AccountService interface {
	CreateAccount(ctx context.Context, account *accounts.Account) error
	GetAccountByID(context.Context, int64) (accounts.Account, error)
	GetAccountByNumber(context.Context, int64) (accounts.Account, error)
	UpdateAccountDetails(context.Context, accounts.Account) error
	UpdateAccountBalance(context.Context, int, float64) error
	DeleteAccountDetails(context.Context, int64) error
}

// CreateAccount decodes an Account object from the HTTP request body and then tries to create a new account in the database using the CreateAccount method of the AccountService interface. If the account is successfully created, it encodes and sends the created account as a response.
func (h *Handler) CreateAccount(writer http.ResponseWriter, request *http.Request) {
	var acct accounts.Account
	if err := json.NewDecoder(request.Body).Decode(&acct); err != nil {
		return
	}
	err := h.Accounts.CreateAccount(request.Context(), &acct)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.NewEncoder(writer).Encode(acct); err != nil {
		log.Panicln(err)
	}
}

// GetAccountByID extracts the id from the URL parameters and then fetches the account with that id from the database using the GetAccountByID method of the AccountService interface. If the account is found, it encodes and sends the account as a response.
func (h *Handler) GetAccountByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	account, err := h.Accounts.GetAccountByID(request.Context(), id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(writer).Encode(account); err != nil {
		log.Panicln(err)
	}
}

// GetAccountByNumber extracts the number from the URL parameters and then fetches the account with that number from the database using the GetAccountByNumber method of the AccountService interface. If the account is found, it encodes and sends the account as a response.
func (h *Handler) GetAccountByNumber(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	number, err := strconv.ParseInt(vars["number"], 10, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	account, err := h.Accounts.GetAccountByNumber(request.Context(), number)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(writer).Encode(account); err != nil {
		log.Panicln(err)
	}
}

// UpdateAccountDetails decodes an Account object from the HTTP request body and then updates the account in the database using the UpdateAccountDetails method of the AccountService interface. If the account is successfully updated, it responds with a status code 200 OK.
func (h *Handler) UpdateAccountDetails(writer http.ResponseWriter, request *http.Request) {
	var acct accounts.Account
	if err := json.NewDecoder(request.Body).Decode(&acct); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.Accounts.UpdateAccountDetails(request.Context(), acct)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

// UpdateAccountBalance extracts the number and amount from the URL parameters and then updates the balance of the account with that number in the database using the UpdateAccountBalance method of the AccountService interface. If the balance is successfully updated, it responds with a status code 200 OK.
func (h *Handler) UpdateAccountBalance(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	stringAccountNumber := vars["account_number"]
	stringAmount := vars["amount"]
	amount, err := strconv.ParseFloat(stringAmount, 64)
	accountNumber, err := strconv.ParseInt(stringAccountNumber, 10, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.Accounts.UpdateAccountBalance(request.Context(), accountNumber, amount)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
