package main

import (
	"PayWalletEngine/internal/db"
	//transportHTTP "PayWalletEngine/internal/transport/http"
	"PayWalletEngine/internal/users"

	"fmt"
	"log"
)

// Run - is going to be responsible for / the instantiation and startup of our / go application
func Run() error {
	fmt.Println("starting up the application...")

	store, err := db.NewDatabase(dsn)
	if err != nil {
		log.Println("Database Connection Failure")
		return err
	}
	if err := store.Ping(); err != nil {
		return err
	}
	log.Println("Successfully connected to the store")

	if err := store.MigrateDB(); err != nil {
		log.Println("failed to setup store migrations")
		return err
	}

	userService := users.NewService(store)

	return nil

}
func main() {
	fmt.Println("GO REST API Course")
	if err := Run(); err != nil {
		log.Println(err)
	}

}
