package http

import (
	"PayWalletEngine/internal/accounts"
	"context"
	"net/http"
)

type AccountService interface {
	CreateAccount(ctx context.Context, account *accounts.Account) error
	GetAccountByID(context.Context, int64) (accounts.Account, error)
	GetAccountByNumber(context.Context, int64) (accounts.Account, error)
	UpdateAccountDetails(context.Context, accounts.Account) error
	UpdateAccountBalance(context.Context, string, float64) error
	DeleteAccountDetails(context.Context, int64) error
}

func (h *Handler) CreateAccount(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) GetAccountByID(writer http.ResponseWriter, request *http.Request) {

}
func (h *Handler) GetAccountByNumber(writer http.ResponseWriter, request *http.Request) {

}
func (h *Handler) UpdateAccountDetails(writer http.ResponseWriter, request *http.Request) {

}
func (h *Handler) UpdateAccountBalance(writer http.ResponseWriter, request *http.Request) {

}
func (h *Handler) DeleteAccountDetails(writer http.ResponseWriter, request *http.Request) {

}
