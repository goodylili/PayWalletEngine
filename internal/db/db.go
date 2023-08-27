package db

import (
	"fmt"
	"gorm.io/gorm"
)

type Database struct {
	Client *gorm.DB
}

func NewDatabase() (*Database, error) {
	connectionString := fmt.Sprintf()
}
