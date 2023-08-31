package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	Client *gorm.DB
}

func NewDatabase(dsn string) (*Database, error) {
	configurations := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("SSL_MODE"))

	db, err := gorm.Open(postgres.Open(configurations), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	database := &Database{
		Client: db,
	}

	return database, nil
}

func (d *Database) Ping() error {
	sqlDB, err := d.Client.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
