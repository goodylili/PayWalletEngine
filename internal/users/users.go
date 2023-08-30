package users

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

// User -  a representation of the users of the wallet engine
type User struct {
	gorm.Model
	UserID   int     `json:"userID"`
	Username string  `json:"username"` // username for the user
	Email    string  `json:"email"`    // email address for the user
	Password string  `json:"password"` // hashed password for the user
	Balance  float64 `json:"balance"`  // current balance for the user's wallet
}

type Store interface {
	UpdateUser(context.Context, User) error
	DeleteUser(context.Context, string) error
	GetUser(context.Context, string) (User, error)
	GetAllUsers(context.Context) ([]*User, error)
	CreateUser(ctx context.Context, user *User) error
	Ping(context.Context) error
}

var (
	ErrFetchingUser  = errors.New("could not fetch comment by ID")
	ErrUpdatingUser  = errors.New("could not update comment")
	ErrNoCommentUser = errors.New("no comment found")
	ErrDeletingUser  = errors.New("could not delete comment")
)

// UserService is the blueprint for the user logic
type UserService struct {
	Store Store
}

// NewService creates a new service
func NewService(store Store) *UserService {
	return &UserService{
		Store: store,
	}
}

func (u *UserService) GetUser(ctx context.Context, userID string) (User, error) {
	user, err := u.Store.GetUser(ctx, userID)
	if err != nil {
		log.Printf("Error fetching user with ID %s: %v", userID, err)
		return user, err
	}

	return user, nil
}

func (u *UserService) CreateUser(ctx context.Context, user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return err
	}
	user.Password = string(hashedPassword)

	if err := u.Store.CreateUser(ctx, user); err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}

	return nil
}

func (u *UserService) UpdateUser(ctx context.Context, user User) error {
	if err := u.Store.UpdateUser(ctx, user); err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}

	return nil
}

func (u *UserService) DeleteUser(ctx context.Context, userID string) error {
	if err := u.Store.DeleteUser(ctx, userID); err != nil {
		log.Printf("Error deleting user with ID %s: %v", userID, err)
		return err
	}

	return nil
}

func (u *UserService) GetAllUsers(ctx context.Context) ([]*User, error) {
	users, err := u.Store.GetAllUsers(ctx)
	if err != nil {
		log.Printf("Error fetching all users: %v", err)
		return nil, err
	}

	return users, nil
}

func (u *UserService) UpdatePassword(ctx context.Context, userID string, newPassword string) error {
	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return err
	}

	// Get the user by userID
	user, err := u.Store.GetUser(ctx, userID)
	if err != nil {
		log.Printf("Error fetching user with ID %s: %v", userID, err)
		return err
	}

	// Update the password
	user.Password = string(hashedPassword)

	// Update the user in the store
	if err := u.Store.UpdateUser(ctx, user); err != nil {
		log.Printf("Error updating password for user with ID %s: %v", userID, err)
		return err
	}

	return nil
}

// ReadyCheck - a function that tests we are functionally ready to serve requests
func (u *UserService) ReadyCheck(ctx context.Context) error {
	log.Println("Checking readiness")
	return u.Store.Ping(ctx)
}
