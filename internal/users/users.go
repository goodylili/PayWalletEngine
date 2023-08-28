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
}

var (
	ErrFetchingUser  = errors.New("could not fetch comment by ID")
	ErrUpdatingUser  = errors.New("could not update comment")
	ErrNoCommentUser = errors.New("no comment found")
	ErrDeletingUser  = errors.New("could not delete comment")
)

// Service is the blueprint for the user logic
type Service struct {
	Store Store
}

// NewService creates a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetUser(ctx context.Context, userID string) (User, error) {
	user, err := s.Store.GetUser(ctx, userID)
	if err != nil {
		log.Printf("Error fetching user with ID %s: %v", userID, err)
		return user, err
	}

	return user, nil
}

func (s *Service) CreateUser(ctx context.Context, user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return err
	}
	user.Password = string(hashedPassword)

	if err := s.Store.CreateUser(ctx, user); err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}

	return nil
}

func (s *Service) UpdateUser(ctx context.Context, user User) error {
	if err := s.Store.UpdateUser(ctx, user); err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}

	return nil
}

func (s *Service) DeleteUser(ctx context.Context, userID string) error {
	if err := s.Store.DeleteUser(ctx, userID); err != nil {
		log.Printf("Error deleting user with ID %s: %v", userID, err)
		return err
	}

	return nil
}

func (s *Service) GetAllUsers(ctx context.Context) ([]*User, error) {
	users, err := s.Store.GetAllUsers(ctx)
	if err != nil {
		log.Printf("Error fetching all users: %v", err)
		return nil, err
	}

	return users, nil
}

func (s *Service) UpdatePassword(ctx context.Context, userID string, newPassword string) error {
	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return err
	}

	// Get the user by userID
	user, err := s.Store.GetUser(ctx, userID)
	if err != nil {
		log.Printf("Error fetching user with ID %s: %v", userID, err)
		return err
	}

	// Update the password
	user.Password = string(hashedPassword)

	// Update the user in the store
	if err := s.Store.UpdateUser(ctx, user); err != nil {
		log.Printf("Error updating password for user with ID %s: %v", userID, err)
		return err
	}

	return nil
}
