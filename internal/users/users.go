package users

import (
	"context"
	"log"
	"time"
)

// User -  a representation of the users of the wallet engine
type User struct {
	ID        int64     // unique identifier for the user
	Username  string    // username for the user
	Email     string    // email address for the user
	Password  string    // hashed password for the user
	Balance   float64   // current balance for the user's wallet
	CreatedAt time.Time // timestamp for when the user's account was created
	UpdatedAt time.Time // timestamp for when the user's account was last updated
}

type Store interface {
	CreateUser(user *User) error
	UpdateUser(userID string, user User) error
	DeleteUser(userID string) error
	GetUser(context.Context, string) (User, error)
	GetAllUsers() ([]*User, error)
}

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
		log.Println(err)
		return user, err
	}

	return user, nil
}
