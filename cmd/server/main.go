package main

import (
	"PayWalletEngine/internal/accounts"
	"PayWalletEngine/internal/db"
	"PayWalletEngine/internal/transactions"
	transportHTTP "PayWalletEngine/transport/http"

	//transportHTTP "PayWalletEngine/internal/transport/http"
	"PayWalletEngine/internal/users"

	"fmt"
	"log"
)

// Run - is going to be responsible for / the instantiation and startup of our / go application
func Run() error {
	fmt.Println("starting up the application...")

	store, err := db.NewDatabase()
	if err != nil {
		log.Println("Database Connection Failure")
		return err
	}

	//if err := store.Ping(); err != nil {
	//	return err
	//}
	//log.Println("Successfully connected to the store")

	if err := store.MigrateDB(); err != nil {
		log.Println("failed to setup store migrations")
		return err
	}

	userService := users.NewService(store)
	transactionService := transactions.NewTransactionService(store)
	accountService := accounts.NewAccountService(store)
	handler := transportHTTP.NewHandler(userService, transactionService accountService)

	if err := handler.Serve(); err != nil {
		log.Println("failed to gracefully serve our application")
		return err
	}

	return nil


}
func main() {
	fmt.Println("GO REST API Course")
	if err := Run(); err != nil {
		log.Println(err)
	}

}
