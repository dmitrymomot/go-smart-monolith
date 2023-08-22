package commands_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dmitrymomot/go-smart-monolith/internal/user/app/commands"
	"github.com/dmitrymomot/go-smart-monolith/internal/user/domain"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// createUserRepository is a mock implementation of the createUserRepository
// interface.
type createUserRepository struct {
	mock.Mock
}

// GetUserByEmail is a mock implementation of the GetUser method.
func (m *createUserRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(domain.User), args.Error(1)
}

// StoreUser is a mock implementation of the StoreUser method.
func (m *createUserRepository) StoreUser(ctx context.Context, user domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	// Test data.
	var (
		email    = "test@mail.dev"
		password = "password"
	)

	// Success case.
	t.Run("success", func(t *testing.T) {
		// Create the repository mock and set the expectations.
		repo := &createUserRepository{}
		repo.On("GetUserByEmail", mock.Anything, email).Return(domain.User{}, errors.New("not found"))
		repo.On("StoreUser", mock.Anything, mock.Anything).Return(nil)

		// Create the command.
		cmd := commands.CreateUser(repo, true)

		// Call the method under test.
		events, err := cmd(context.Background(), commands.CreateUserCommand{
			Email:    email,
			Password: password,
		})
		require.NoError(t, err)
		require.Len(t, events, 1)
		require.IsType(t, commands.UserCreatedEvent{}, events[0])

		// Assert the expectations.
		repo.AssertExpectations(t)
	})

	// Email is already taken.
	t.Run("email_taken", func(t *testing.T) {
		// Create the repository mock and set the expectations.
		repo := &createUserRepository{}
		user := domain.NewUser(email, password)
		repo.On("GetUserByEmail", mock.Anything, email).Return(user, nil)

		// Create the command.
		cmd := commands.CreateUser(repo, true)

		// Call the method under test.
		events, err := cmd(context.Background(), commands.CreateUserCommand{
			Email:    email,
			Password: password,
		})
		require.Error(t, err)
		require.ErrorIs(t, err, commands.ErrUserAlreadyExists)
		require.Len(t, events, 0)

		// Assert the expectations.
		repo.AssertExpectations(t)
	})

	// Failed to create user.
	t.Run("failed_to_create_user", func(t *testing.T) {
		// Create the repository mock and set the expectations.
		repo := &createUserRepository{}
		repo.On("GetUserByEmail", mock.Anything, email).Return(domain.User{}, errors.New("not found"))
		repo.On("StoreUser", mock.Anything, mock.Anything).Return(errors.New("failed to create user"))

		// Create the command.
		cmd := commands.CreateUser(repo, true)

		// Call the method under test.
		events, err := cmd(context.Background(), commands.CreateUserCommand{
			Email:    email,
			Password: password,
		})
		require.Error(t, err)
		require.ErrorIs(t, err, commands.ErrFailedToCreateUser)
		require.Len(t, events, 0)

		// Assert the expectations.
		repo.AssertExpectations(t)
	})
}
