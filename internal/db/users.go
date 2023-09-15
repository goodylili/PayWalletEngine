package db

import (
	"PayWalletEngine/internal/users"
	"PayWalletEngine/utils"
	"context"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account  []Account `gorm:"foreignkey:ID"`
	Username string    `gorm:"unique;not null"`
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Balance  float64   `gorm:"default:0"`
	IsActive bool      `gorm:"default:true"`
}

func (d *Database) CreateUser(ctx context.Context, user *User) error {
	dbUser := &User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Balance:  user.Balance,
		IsActive: user.IsActive,
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
		Balance:  dbUser.Balance,
		IsActive: dbUser.IsActive,
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
		Balance:  dbUser.Balance,
		IsActive: dbUser.IsActive,
	}, nil
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
		Balance:  dbUser.Balance,
		IsActive: dbUser.IsActive,
	}, nil
}

func (d *Database) UpdateUser(ctx context.Context, user users.User) error {
	var dbUser User
	// Check if user exists
	if err := d.Client.WithContext(ctx).Where("id = ?", user.ID).First(&dbUser).Error; err != nil {
		return err
	}

	// Check if the passwords match using the comparePasswords function
	if !utils.ComparePasswords(dbUser.Password, user.Password) {
		return errors.New("password does not match")
	}

	dbUser = User{
		Username: user.Username,
		Email:    user.Email,
		Balance:  user.Balance,
		IsActive: dbUser.IsActive,
	}

	// if the user exists and passwords match, update the database with the user's new details
	if err := d.Client.WithContext(ctx).Save(&dbUser).Error; err != nil {
		return err
	}
	return nil
}

// DeactivateUserByID sets the user's IsActive status to false based on the provided ID.
func (d *Database) DeactivateUserByID(ctx context.Context, id int64) error {
	user := User{}
	if err := d.Client.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		// Return an error if the user is not found
		return err
	}
	// Set the user's IsActive status to false
	user.IsActive = false
	if err := d.Client.WithContext(ctx).Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (d *Database) ResetPassword(ctx context.Context, newUser User) error {
	// Assuming that the newUser.Password is already hashed

	// Update user password where username, email match and the user is active
	result := d.Client.WithContext(ctx).Model(&User{}).Where("username = ? AND email = ? AND is_active = ?", newUser.Username, newUser.Email, true).Update("password", newUser.Password)

	// Check if any rows were affected
	if result.RowsAffected == 0 {
		return errors.New("no matching active user found with the provided username and email")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
