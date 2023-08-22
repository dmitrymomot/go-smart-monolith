package storage

import (
	"context"

	"github.com/dmitrymomot/go-smart-monolith/internal/user/domain"
)

type (
	// Storage is a storage service adapter.
	// It's just an implementation of the repository pattern.
	Storage struct {
		client storageClient
	}

	// Low-level storage client. Redis, mongo, pg, etc.
	storageClient interface {
		Get(ctx context.Context, key string) (interface{}, error)
		Set(ctx context.Context, key string, value interface{}) error
	}
)

// New is a factory function that creates a new storage service adapter.
func New(client storageClient) *Storage {
	return &Storage{
		client: client,
	}
}

// GetUserByID gets a user by ID.
func (s *Storage) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	v, err := s.client.Get(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return v.(domain.User), nil
}

// GetUserByEmail gets a user by email.
// It's just an example of a different method. Don't pay attention to implementation.
func (s *Storage) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	v, err := s.client.Get(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return v.(domain.User), nil
}

// StoreUser stores a user.
func (s *Storage) StoreUser(ctx context.Context, user domain.User) error {
	return s.client.Set(ctx, user.Email, user)
}
