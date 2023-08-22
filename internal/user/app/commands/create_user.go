package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/dmitrymomot/go-smart-monolith/internal/user/domain"
)

// ErrUserAlreadyExists is returned when the user already exists.
var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrFailedToCreateUser = errors.New("failed to create user")
)

type (
	// CreateUserCommand represents the request body for CreateUser.
	CreateUserCommand struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// UserCreatedEvent represents the event body for UserCreated.
	UserCreatedEvent struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}

	// createUserRepository represents the repository interface for CreateUser.
	createUserRepository interface {
		GetUserByEmail(ctx context.Context, email string) (domain.User, error)
		StoreUser(ctx context.Context, user domain.User) error
	}
)

// CreateUser creates a new user.
func CreateUser(
	repo createUserRepository,
	flag bool,
) func(ctx context.Context, cmd CreateUserCommand) ([]interface{}, error) {
	return func(ctx context.Context, cmd CreateUserCommand) ([]interface{}, error) {
		// Check if the email is already taken.
		if _, err := repo.GetUserByEmail(ctx, cmd.Email); err == nil {
			return nil, ErrUserAlreadyExists
		}

		// Create the user.
		user := domain.NewUser(cmd.Email, cmd.Password)
		if err := repo.StoreUser(ctx, user); err != nil {
			return nil, fmt.Errorf("%w: %v", ErrFailedToCreateUser, err)
		}

		// Return the event.
		return []interface{}{
			UserCreatedEvent{
				ID:    user.ID,
				Email: user.Email,
			},
		}, nil
	}
}
