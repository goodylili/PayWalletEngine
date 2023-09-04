package users

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockUserStore is a mocked object that implements the UserStore interface
type MockUserStore struct {
	MockCreateUser      func(ctx context.Context, user *User) error
	MockGetUserByID     func(ctx context.Context, id string) (User, error)
	MockGetByEmail      func(ctx context.Context, email string) (*User, error)
	MockGetByUsername   func(ctx context.Context, username string) (*User, error)
	MockUpdateUser      func(ctx context.Context, user User) error
	MockDeactivateUsers func(ctx context.Context, id string) error
	MockPing            func(ctx context.Context) error
	MockResetPassword   func(ctx context.Context, user User) error
}

func (m *MockUserStore) CreateUser(ctx context.Context, user *User) error {
	return m.MockCreateUser(ctx, user)
}

func (m *MockUserStore) GetUserByID(ctx context.Context, id string) (User, error) {
	return m.MockGetUserByID(ctx, id)
}

func (m *MockUserStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	return m.MockGetByEmail(ctx, email)
}

func (m *MockUserStore) GetByUsername(ctx context.Context, username string) (*User, error) {
	return m.MockGetByUsername(ctx, username)
}

func (m *MockUserStore) UpdateUser(ctx context.Context, user User) error {
	return m.MockUpdateUser(ctx, user)
}

func (m *MockUserStore) DeactivateUsers(ctx context.Context, id string) error {
	return m.MockDeactivateUsers(ctx, id)
}

func (m *MockUserStore) Ping(ctx context.Context) error {
	return m.MockPing(ctx)
}

func (m *MockUserStore) ResetPassword(ctx context.Context, user User) error {
	return m.MockResetPassword(ctx, user)
}

func TestUserService_CreateUser(t *testing.T) {
	store := &MockUserStore{
		MockCreateUser: func(ctx context.Context, user *User) error {
			return nil // Simulating successful user creation
		},
	}

	service := NewService(store)

	user := &User{
		Username: "testUser",
		Email:    "test@email.com",
		Password: "testPassword",
	}

	err := service.CreateUser(context.Background(), user)
	assert.NoError(t, err)
}

func TestUserService_GetUserByID(t *testing.T) {
	store := &MockUserStore{
		MockGetUserByID: func(ctx context.Context, id string) (User, error) {
			if id == "notfound" {
				return User{}, errors.New("user not found")
			}

			return User{Username: "testUser", Email: "test@email.com"}, nil
		},
	}

	service := NewService(store)

	_, err := service.GetUserByID(context.Background(), "notfound")
	assert.Error(t, err)

	user, err := service.GetUserByID(context.Background(), "1")
	assert.NoError(t, err)
	assert.Equal(t, "testUser", user.Username)
}

func TestUserService_GetByEmail(t *testing.T) {
	store := &MockUserStore{
		MockGetByEmail: func(ctx context.Context, email string) (*User, error) {
			if email == "notfound@email.com" {
				return nil, errors.New("user not found")
			}

			return &User{Username: "testUser", Email: "test@email.com"}, nil
		},
	}

	service := NewService(store)

	_, err := service.GetByEmail(context.Background(), "notfound@email.com")
	assert.Error(t, err)

	user, err := service.GetByEmail(context.Background(), "test@email.com")
	assert.NoError(t, err)
	assert.Equal(t, "testUser", user.Username)
}

func TestUserService_GetByUsername(t *testing.T) {
	store := &MockUserStore{
		MockGetByUsername: func(ctx context.Context, username string) (*User, error) {
			if username == "notfoundUser" {
				return nil, errors.New("user not found")
			}

			return &User{Username: "testUser", Email: "test@email.com"}, nil
		},
	}

	service := NewService(store)

	_, err := service.GetByUsername(context.Background(), "notfoundUser")
	assert.Error(t, err)

	user, err := service.GetByUsername(context.Background(), "testUser")
	assert.NoError(t, err)
	assert.Equal(t, "testUser", user.Username)
}

func TestUserService_UpdateUser(t *testing.T) {
	store := &MockUserStore{
		MockUpdateUser: func(ctx context.Context, user User) error {
			if user.Username == "errorUser" {
				return errors.New("error updating user")
			}
			return nil
		},
	}

	service := NewService(store)

	user := User{
		Username: "errorUser",
		Email:    "error@email.com",
	}

	err := service.UpdateUser(context.Background(), user)
	assert.Error(t, err)

	user.Username = "testUser"
	err = service.UpdateUser(context.Background(), user)
	assert.NoError(t, err)
}

func TestUserService_DeactivateUser(t *testing.T) {
	store := &MockUserStore{
		MockDeactivateUsers: func(ctx context.Context, id string) error {
			if id == "invalidID" {
				return errors.New("error deactivating user")
			}
			return nil
		},
	}

	service := NewService(store)

	err := service.DeactivateUser(context.Background(), "invalidID")
	assert.Error(t, err)

	err = service.DeactivateUser(context.Background(), "validID")
	assert.NoError(t, err)
}

func TestUserService_ReadyCheck(t *testing.T) {
	store := &MockUserStore{
		MockPing: func(ctx context.Context) error {
			return nil // Assuming always ready
		},
	}

	service := NewService(store)

	err := service.ReadyCheck(context.Background())
	assert.NoError(t, err)
}

func TestUserService_ResetPassword(t *testing.T) {
	store := &MockUserStore{
		MockResetPassword: func(ctx context.Context, user User) error {
			if user.Username == "errorUser" {
				return errors.New("error resetting password")
			}
			return nil
		},
	}

	service := NewService(store)

	user := User{
		Username: "errorUser",
		Email:    "error@email.com",
		Password: "newPassword",
	}

	err := service.ResetPassword(context.Background(), user)
	assert.Error(t, err)

	user.Username = "testUser"
	err = service.ResetPassword(context.Background(), user)
	assert.NoError(t, err)
}
