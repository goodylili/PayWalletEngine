package db

import (
	"PayWalletEngine/internal/users"
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID         int     `gorm:"unique;not null"` // username for the user
	Username       string  `gorm:"unique;not null"` // username for the user
	Email          string  `gorm:"unique;not null"` // email address for the user
	HashedPassword string  `gorm:"not null"`        // hashed password for the user
	Balance        float64 `gorm:"default:0"`       // current balance for the user's wallet
}

func (d *Database) UpdateUser(ctx context.Context, user users.User) error {
	var dbUser User
	// Check if user exists
	if err := d.Client.WithContext(ctx).Where("user_id = ?", user.UserID).First(&dbUser).Error; err != nil {
		return err
	}
	dbUser = User{
		UserID:         user.UserID,
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Balance:        user.Balance,
	}
	if err := d.Client.WithContext(ctx).Save(&dbUser).Error; err != nil {
		return err
	}
	return nil
}

func (d *Database) DeleteUser(ctx context.Context, s string) error {
	user := User{}
	if err := d.Client.WithContext(ctx).Where("username = ?", s).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (d *Database) GetUser(ctx context.Context, s string) (users.User, error) {
	dbUser := User{}
	if err := d.Client.WithContext(ctx).Where("username = ?", s).First(&dbUser).Error; err != nil {
		return users.User{}, err
	}
	return users.User{
		UserID:         dbUser.UserID,
		Username:       dbUser.Username,
		Email:          dbUser.Email,
		HashedPassword: dbUser.HashedPassword,
		Balance:        dbUser.Balance,
	}, nil
}

func (d *Database) CreateUser(ctx context.Context, user *users.User) error {
	dbUser := User{
		UserID:         user.UserID,
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Balance:        user.Balance,
	}
	if err := d.Client.WithContext(ctx).Create(&dbUser).Error; err != nil {
		return err
	}
	return nil
}

func (d *Database) GetByEmail(ctx context.Context, email string) (*users.User, error) {
	var dbUser User
	err := d.Client.WithContext(ctx).Where("email = ?", email).First(&dbUser).Error
	if err != nil {
		return nil, err
	}
	return &users.User{
		UserID:         dbUser.UserID,
		Username:       dbUser.Username,
		Email:          dbUser.Email,
		HashedPassword: dbUser.HashedPassword,
		Balance:        dbUser.Balance,
	}, nil
}

func (d *Database) GetByUsername(ctx context.Context, username string) (*users.User, error) {
	var dbUser User
	err := d.Client.WithContext(ctx).Where("username = ?", username).First(&dbUser).Error
	if err != nil {
		return nil, err
	}
	return &users.User{
		UserID:         dbUser.UserID,
		Username:       dbUser.Username,
		Email:          dbUser.Email,
		HashedPassword: dbUser.HashedPassword,
		Balance:        dbUser.Balance,
	}, nil
}
