package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	configurations := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("SSL_MODE"),
	)

	db, err := gorm.Open(postgres.Open(configurations), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{
		Client: db,
	}, nil
}

func (d *Database) Ping(ctx context.Context) error {
	sqlDB, err := d.Client.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.PingContext(ctx); err != nil {
		return err
	}
	return nil
}
