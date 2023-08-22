package domain

import "github.com/google/uuid"

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
}

// NewUser creates a new user.
func NewUser(email, password string) User {
	return User{
		ID:           uuid.New().String(),
		Email:        email,
		PasswordHash: password,
	}
}
