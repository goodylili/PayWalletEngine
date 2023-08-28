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

func (d *Database) HealthCheck() error {
	sqlDB, err := d.Client.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

func LoadConfig() (string, error) {
	configurations := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("SSL_MODE"))

	return configurations, nil
}

func NewDatabase(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	database := &Database{
		Client: db,
	}

	return database, nil
}
