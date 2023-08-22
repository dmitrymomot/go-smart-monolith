package storage

import (
	"context"
	"errors"
	"sync"
)

// Storage represents a kv storage example.
// It's an example of low-level storage implementation.
// It can be a database, a cache, a file, etc.
// And this package cannot know anything about the domain models or the
// business logic.
type Storage struct {
	sync.RWMutex
	kv map[string]interface{}
}

// New creates a new storage.
func New() *Storage {
	return &Storage{
		kv: make(map[string]interface{}),
	}
}

// Get gets a value from the storage.
func (s *Storage) Get(ctx context.Context, key string) (interface{}, error) {
	s.RLock()
	defer s.RUnlock()
	if v, ok := s.kv[key]; ok {
		return v, nil
	}
	return nil, errors.New("not found")
}

// Set sets a value to the storage.
func (s *Storage) Set(ctx context.Context, key string, value interface{}) error {
	s.Lock()
	defer s.Unlock()
	s.kv[key] = value
	return nil
}
