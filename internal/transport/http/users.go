package http

import (
	"PayWalletEngine/internal/users"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type UserService interface {
	CreateUser(context.Context, *users.User) error
	GetUserByID(context.Context, string) (users.User, error)
	GetByEmail(context.Context, string) (*users.User, error)
	GetByUsername(context.Context, string) (*users.User, error)
	UpdateUser(context.Context, users.User) error
	DeactivateUser(context.Context, string) error
	ResetPassword(context.Context, users.User) error
	Ping(ctx context.Context) error
}

// CreateUser decodes a User object from the HTTP request body and then tries to create a new user in the database using the CreateUser method of the UserService interface. If the user is successfully created, it encodes and sends the created user as a response.
func (h *Handler) CreateUser(writer http.ResponseWriter, request *http.Request) {
	var u users.User
	if err := json.NewDecoder(request.Body).Decode(&u); err != nil {
		return
	}
	err := h.Users.CreateUser(request.Context(), &u)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.NewEncoder(writer).Encode(u); err != nil {
		log.Panicln(err)
	}
}

// GetUserByID extracts the id from the URL parameters and then fetches the user with that id from the database using the GetUserByID method of the UserService interface. If the user is found, it encodes and sends the user as a response.
func (h *Handler) GetUserByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	u, err := h.Users.GetUserByID(request.Context(), id)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.NewEncoder(writer).Encode(u); err != nil {
		log.Panicln(err)
	}
}

// GetByEmail extracts the email from the URL parameters and then fetches the user with that email from the database using the GetByEmail method of the UserService interface. If the user is found, it encodes and sends the user as a response.
func (h *Handler) GetByEmail(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	email := vars["email"]
	u, err := h.Users.GetByEmail(request.Context(), email)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.NewEncoder(writer).Encode(u); err != nil {
		log.Panicln(err)
	}
}

// GetByUsername extracts the username from the URL parameters and then fetches the user with that username from the database using the GetByUsername method of the UserService interface. If the user is found, it encodes and sends the user as a response.
func (h *Handler) GetByUsername(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	username := vars["username"]
	u, err := h.Users.GetByUsername(request.Context(), username)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.NewEncoder(writer).Encode(u); err != nil {
		log.Panicln(err)
	}
}

// UpdateUser decodes a User object from the HTTP request body and then updates the user in the database using the UpdateUser method of the UserService interface. If the user is successfully updated, it encodes and sends the updated user as a response.
func (h *Handler) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	var u users.User
	// Decode request body
	if err := json.NewDecoder(request.Body).Decode(&u); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	// Update user
	err := h.Users.UpdateUser(request.Context(), u)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	// Encode and send response
	if err := json.NewEncoder(writer).Encode(u); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		log.Panicln(err)
	}
}

// DeactivateUserByID extracts the id from the URL parameters and then deletes the user with that id from the database using the DeactivateUserByID method of the UserService interface. If the user is successfully deleted, it sends a No Content status code as a response.
func (h *Handler) DeactivateUserByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.Users.DeactivateUserByID(request.Context(), id)
	if err != nil {
		log.Println(err)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
	if err := json.NewEncoder(writer).Encode(map[string]string{"status": "OK"}); err != nil {
		log.Panicln(err)
	}

}

func (h *Handler) Ping(writer http.ResponseWriter, request *http.Request) {
	err := h.Users.ReadyCheck(request.Context())
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.NewEncoder(writer).Encode(map[string]string{"status": "OK"}); err != nil {
		log.Panicln(err)
	}
}

// ResetPassword decodes a User object from the HTTP request body, then attempts to reset the user's password in the database using the ResetPassword method of the UserService interface. If the password is successfully reset, it sends an OK response.
func (h *Handler) ResetPassword(writer http.ResponseWriter, request *http.Request) {
	var u users.User
	// Decode request body
	if err := json.NewDecoder(request.Body).Decode(&u); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	// Reset password
	err := h.Users.ResetPassword(request.Context(), u)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	// Send response
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(map[string]string{"status": "Password reset successful"}); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		log.Panicln(err)
	}
}
