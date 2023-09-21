package users

import (
	"PayWalletEngine/internal/accounts"
	"context"
	"gorm.io/gorm"
	"log"
)

// User -  a representation of the users of the wallet engine
type User struct {
	gorm.Model `json:"-"`
	Username   string             `json:"username"`  // username for the user
	Email      string             `json:"email"`     // email address for the user
	Password   string             `json:"password"`  // hashed password for the user
	IsActive   bool               `json:"is_active"` // status of the user, true means active
	Account    []accounts.Account `json:"accounts"`
}

type UserStore interface {
	CreateUser(context.Context, *User) error
	GetUserByID(context.Context, int64) (User, error)
	GetByEmail(context.Context, string) (*User, error)
	GetByUsername(context.Context, string) (*User, error)
	UpdateUser(context.Context, User, uint) error
	ResetPassword(context.Context, User) error
	ChangeUserStatus(context.Context, User, uint) error
	PingDatabase(ctx context.Context) error
}

// UserService is the blueprint for the user logic
type UserService struct {
	Store UserStore
}

// NewService creates a new service
func NewService(store UserStore) UserService {
	return UserService{
		Store: store,
	}
}

func (u *UserService) CreateUser(ctx context.Context, user *User) error {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return err
	}
	user.Password = hashedPassword

	if err := u.Store.CreateUser(ctx, user); err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return nil
}

func (u *UserService) GetUserByID(ctx context.Context, id int64) (User, error) {
	user, err := u.Store.GetUserByID(ctx, id)
	if err != nil {
		log.Printf("Error fetching user with ID %v: %v", id, err)
		return user, err
	}
	return user, nil
}

func (u *UserService) UpdateUser(ctx context.Context, user User, id uint) error {
	if err := u.Store.UpdateUser(ctx, user, id); err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}

	return nil
}

func (u *UserService) ChangeUserStatus(ctx context.Context, user User, id uint) error {
	if err := u.Store.ChangeUserStatus(ctx, user, id); err != nil {
		log.Printf("Error deactivating user with ID %v: %v", id, err)
		return err
	}

	return nil
}

func (u *UserService) GetByEmail(ctx context.Context, email string) (*User, error) {
	user, err := u.Store.GetByEmail(ctx, email)
	if err != nil {
		log.Printf("Error fetching user with email %v: %v", email, err)
		return nil, err
	}

	return user, nil
}

func (u *UserService) GetByUsername(ctx context.Context, username string) (*User, error) {
	user, err := u.Store.GetByUsername(ctx, username)
	if err != nil {
		log.Printf("Error fetching user with username %v: %v", username, err)
		return nil, err
	}

	return user, nil
}

// ReadyCheck - a function that tests we are functionally ready to serve requests
func (u *UserService) ReadyCheck(ctx context.Context) error {
	log.Println("Checking readiness")
	return u.Store.PingDatabase(ctx)
}

func (u *UserService) ResetPassword(ctx context.Context, user User) error {

	// Next, we'll call the store's ResetPassword method
	if err := u.Store.ResetPassword(ctx, user); err != nil {
		log.Printf("Error resetting password for user with username %v: %v", user.Username, err)
		return err
	}

	return nil
}
