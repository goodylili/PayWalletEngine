package db

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID   int     `gorm:"unique;not null"` // username for the user
	Username string  `gorm:"unique;not null"` // username for the user
	Email    string  `gorm:"unique;not null"` // email address for the user
	Password string  `gorm:"not null"`        // hashed password for the user
	Balance  float64 `gorm:"default:0"`       // current balance for the user's wallet
}

// GetByID fetches a user by its ID from the database.
func (d *Database) GetByID(ctx context.Context, ID uint) (*User, error) {
	var u User
	err := d.Client.WithContext(ctx).First(&u, ID).Error
	return &u, err
}

// GetByEmail fetches a user by its email from the database.
func (d *Database) GetByEmail(ctx context.Context, email string) (*User, error) {
	var u User
	err := d.Client.WithContext(ctx).Where("email = ?", email).First(&u).Error
	return &u, err
}

// GetByUsername fetches a user by its username from the database.
func (d *Database) GetByUsername(ctx context.Context, username string) (*User, error) {
	var u User
	err := d.Client.WithContext(ctx).Where("username = ?", username).First(&u).Error
	return &u, err
}

// UpdateUser updates a user in the database.
func (d *Database) UpdateUser(ctx context.Context, user *User) error {
	return d.Client.WithContext(ctx).Save(user).Error
}

// DeleteUser deletes a user by its ID from the database.
func (d *Database) DeleteUser(ctx context.Context, ID uint) error {
	return d.Client.WithContext(ctx).Delete(&User{}, ID).Error
}

// CreateUser creates a new user in the database.
func (d *Database) CreateUser(ctx context.Context, user *User) error {
	return d.Client.WithContext(ctx).Create(user).Error
}
