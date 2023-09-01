package http

import (
	"PayWalletEngine/internal/users"
	"context"
	"net/http"
)

type UserService interface {
	CreateUser(context.Context, *users.User) error
	GetUser(context.Context, string) (users.User, error)
	GetByEmail(context.Context, string) (*users.User, error)
	GetByUsername(context.Context, string) (*users.User, error)
	UpdateUser(context.Context, users.User) error
	DeleteUser(context.Context, string) error
	Ping(ctx context.Context) error
}

func (h *Handler) CreateUser(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) GetUser(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) GetByEmail(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) GetByUsername(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) UpdateUser(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) DeleteUser(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) Ping(writer http.ResponseWriter, request *http.Request) {

}
