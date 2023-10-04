package db

import (
	"PayWalletEngine/internal/accounts"
	"PayWalletEngine/internal/users"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string  `gorm:"unique;not null"`
	Email    string  `gorm:"unique;not null"`
	Password string  `gorm:"not null"`
	IsActive bool    `gorm:"not null"`
	Account  Account `gorm:"foreignKey:UserID;references:ID"`
}

func (d *Database) CreateUser(ctx context.Context, user *users.User) error {
	dbUser := &User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		IsActive: false,
		Account: Account{
			Balance: 0,
			UserID:  user.ID,
		},
	}

	if err := d.Client.WithContext(ctx).Create(dbUser).Error; err != nil {
		return err
	}

	return nil
}

// GetUserByID returns the user with a specified id
func (d *Database) GetUserByID(ctx context.Context, id int64) (users.User, error) {
	dbUser := User{}
	if err := d.Client.WithContext(ctx).Where("id = ?", id).First(&dbUser).Error; err != nil {
		return users.User{}, err
	}
	return users.User{
		Username: dbUser.Username,
		Email:    dbUser.Email,
		IsActive: dbUser.IsActive,
		Password: dbUser.Password,
	}, nil
}

func (d *Database) GetByEmail(ctx context.Context, email string) (*users.User, error) {
	var dbUser User
	err := d.Client.WithContext(ctx).Where("email = ?", email).First(&dbUser).Error
	if err != nil {
		return nil, err
	}
	return &users.User{
		Username: dbUser.Username,
		Email:    dbUser.Email,
		IsActive: dbUser.IsActive,
		Password: dbUser.Password,
	}, nil
}

func (d *Database) GetAccountByUserID(ctx context.Context, userID uint) (*Account, error) {
	var acct Account
	err := d.Client.WithContext(ctx).Where("user_id = ?", userID).First(&acct).Error
	if err != nil {
		return nil, err
	}
	return &acct, nil
}

func (d *Database) GetByUsername(ctx context.Context, username string) (*users.User, error) {
	var dbUser User
	err := d.Client.WithContext(ctx).Where("username = ?", username).First(&dbUser).Error
	if err != nil {
		return nil, err
	}
	return &users.User{
		Username: dbUser.Username,
		Email:    dbUser.Email,
		IsActive: dbUser.IsActive,
		Password: dbUser.Password,
	}, nil
}

func (d *Database) UpdateUser(ctx context.Context, user users.User, id uint) error {
	// Check if the user exists based on the provided ID
	var existingUser users.User
	if err := d.Client.WithContext(ctx).Where("id = ?", id).First(&existingUser).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return fmt.Errorf("user with ID %d not found", id)
		}
		log.Println("Error querying user:", err)
		return err
	}

	// Create a map of columns and their values that you want to update
	updateColumns := map[string]interface{}{
		"username": user.Username,
		"email":    user.Email,
	}

	// Update only the specified columns in the database
	if err := d.Client.WithContext(ctx).Model(&existingUser).Updates(updateColumns).Error; err != nil {
		log.Println("Error updating user:", err)
		return err
	}

	return nil
}

func (d *Database) ChangeUserStatus(ctx context.Context, user users.User, id uint) error {
	// Check if the user exists based on the provided ID
	var existingUser users.User
	if err := d.Client.WithContext(ctx).Where("id = ?", id).First(&existingUser).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return fmt.Errorf("user with ID %d not found", id)
		}
		log.Println("Error querying user:", err)
		return err
	}

	// Create a map of columns and their values that you want to update
	updateColumns := map[string]interface{}{
		"IsActive": user.IsActive,
	}

	// Update only the specified columns in the database
	if err := d.Client.WithContext(ctx).Model(&existingUser).Updates(updateColumns).Error; err != nil {
		log.Println("Error updating user:", err)
		return err
	}

	return nil
}

func (d *Database) PingDatabase(ctx context.Context) error {
	db, err := d.Client.DB()
	if err != nil {
		return err
	}

	if err := db.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

func (d *Database) ResetPassword(ctx context.Context, newUser users.User) error {
	// Hash the new password
	hashedPassword, err := users.HashPassword(newUser.Password)
	if err != nil {
		return err
	}

	// Log the provided username and email
	log.Printf("Username: %s, Email: %s\n", newUser.Username, newUser.Email)

	// Update user password where username, email match and the user is active
	result := d.Client.WithContext(ctx).Model(&User{}).
		Where("username = ? AND email = ? AND is_active = ?", newUser.Username, newUser.Email, newUser.IsActive).
		Updates(map[string]interface{}{"password": hashedPassword})

	// Log the result of the query
	log.Printf("RowsAffected: %d, Error: %v\n", result.RowsAffected, result.Error)

	// Check if any rows were affected
	if result.RowsAffected == 0 {
		return errors.New("no matching active user found with the provided username and email")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetAccountsByUserID retrieves all accounts associated with a user
func (d *Database) GetAccountsByUserID(ctx context.Context, userID uint) ([]*accounts.Account, error) {
	var userAccounts []*accounts.Account

	// Retrieve all accounts associated with the provided userID
	err := d.Client.WithContext(ctx).Where("user_id = ?", userID).Find(&userAccounts).Error
	if err != nil {
		return nil, err
	}

	return userAccounts, nil
}
