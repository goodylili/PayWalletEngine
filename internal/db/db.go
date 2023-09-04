package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

type Database struct {
	Client *gorm.DB
}

func NewDatabase() (*Database, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	configurations := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("SSL_MODE"))

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  configurations,
		PreferSimpleProtocol: true,
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		return nil, err
	}

	// Enable connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return &Database{
		Client: db,
	}, nil
}
