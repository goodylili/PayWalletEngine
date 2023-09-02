package http

import (
	"PayWalletEngine/internal/accounts"
	"PayWalletEngine/internal/transactions"
	"PayWalletEngine/internal/users"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router      *mux.Router
	Transaction transactions.TransactionService
	Users       users.UserService
	Accounts    accounts.AccountService
	Server      *http.Server
}

// Response object
type Response struct {
	Message string `json:"message"`
}

// NewHandler - returns a pointer to a Handler
func NewHandler(users users.UserService, transactions transactions.TransactionService, accounts accounts.AccountService) *Handler {
	log.Info("setting up our handler")
	h := &Handler{
		Users:       users,
		Transaction: transactions,
		Accounts:    accounts,
	}

	h.Router = mux.NewRouter()

	h.mapRoutes()

	// set the content type to application/json
	h.Router.Use(JSONMiddleware)
	//  log every incoming request
	h.Router.Use(LoggingMiddleware)
	// timeout all requests that take longer than 15 seconds
	h.Router.Use(TimeoutMiddleware)

	h.Server = &http.Server{
		Addr:         "0.0.0.0:8080", // Good practice to set timeouts to avoid Slow-loris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h.Router,
	}

	return h
}

// mapRoutes - sets up all the routes for our application
func (h *Handler) mapRoutes() {

	// Server Health and Ready Check Routes
	h.Router.HandleFunc("/alive", h.AliveCheck).Methods("GET")
	h.Router.HandleFunc("/ready", h.ReadyCheck).Methods("GET")

	// Users Routes
	h.Router.HandleFunc("/api/v1/user", h.CreateUser).Methods("POST")
	h.Router.HandleFunc("/api/v1/users/{id}", JWTAuth(h.GetUser)).Methods("GET")
	h.Router.HandleFunc("/api/v1/users/email/{email}", JWTAuth(h.GetByEmail)).Methods("GET")
	h.Router.HandleFunc("/api/v1/users/username/{username}", JWTAuth(h.GetByUsername)).Methods("GET")
	h.Router.HandleFunc("/api/v1/users/{id}", JWTAuth(h.UpdateUser)).Methods("PUT")
	h.Router.HandleFunc("/api/v1/users/{id}", JWTAuth(h.DeleteUser)).Methods("DELETE")
	h.Router.HandleFunc("/api/v1/users/ping", h.Ping).Methods("GET")

	// Account Routes
	h.Router.HandleFunc("/api/v1/account", JWTAuth(h.CreateAccount)).Methods("POST")
	h.Router.HandleFunc("/api/v1/accounts/{id}", JWTAuth(h.GetAccountByID)).Methods("GET")
	h.Router.HandleFunc("/api/v1/accounts/number/{number}", JWTAuth(h.GetAccountByNumber)).Methods("GET")
	h.Router.HandleFunc("/api/v1/accounts/{id}", JWTAuth(h.UpdateAccountDetails)).Methods("PUT")
	h.Router.HandleFunc("/api/v1/accounts/balance/{account_number}/{amount}", JWTAuth(h.UpdateAccountBalance)).Methods("PUT")
	h.Router.HandleFunc("/api/v1/accounts/{id}", JWTAuth(h.DeleteAccountDetails)).Methods("DELETE")

	// Transactions Routes
	h.Router.HandleFunc("/api/v1/transaction", JWTAuth(h.CreateTransaction)).Methods("POST")
	h.Router.HandleFunc("/api/v1/transactions/{transaction_id}", JWTAuth(h.GetTransactionByTransactionID)).Methods("GET")
	h.Router.HandleFunc("/api/v1/transactions/sender/{accountNumber}", JWTAuth(h.GetTransactionsBySender)).Methods("GET")
	h.Router.HandleFunc("/api/v1/transactions/receiver/{accountNumber}", JWTAuth(h.GetTransactionsByReceiver)).Methods("GET")
	h.Router.HandleFunc("/api/v1/transaction/", JWTAuth(h.UpdateTransaction)).Methods("PUT")
	h.Router.HandleFunc("/api/v1/transaction/{transaction_id}", JWTAuth(h.DeleteTransactionByID)).Methods("DELETE")
	h.Router.HandleFunc("/api/v1/transactions/{transaction_reference}", JWTAuth(h.GetTransactionByReference)).Methods("GET")

}

func (h *Handler) AliveCheck(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(Response{Message: "I am Alive!"}); err != nil {
		panic(err)
	}
}

func (h *Handler) ReadyCheck(writer http.ResponseWriter, request *http.Request) {
	if err := h.Users.ReadyCheck(request.Context()); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(Response{Message: "I am Ready!"}); err != nil {
		panic(err)
	}

}

// Serve - gracefully serves our newly set up handler function
func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	<-c

	// CreateAccount a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := h.Server.Shutdown(ctx)
	if err != nil {
		return err
	}

	log.Println("shutting down gracefully")
	return nil
}
