package queries

import (
	"context"

	"github.com/dmitrymomot/go-smart-monolith/internal/user/domain"
)

type (
	// GetUserQuery represents the request body for GetUser.
	GetUserQuery struct {
		ID string
	}

	// User represents the response body for GetUser.
	User struct {
		ID         string
		Email      string
		PlayerName string
	}

	// getUserRepository represents the repository for GetUser.
	getUserRepository interface {
		GetUserByID(ctx context.Context, id string) (domain.User, error)
	}

	// playersSvcClient represents the client for GetUser.
	playersSvcClient interface {
		GetPlayer(ctx context.Context, key string) (domain.Player, error)
	}
)

// GetUser gets a user.
func GetUser(
	repo getUserRepository,
	playersClient playersSvcClient,
) func(ctx context.Context, query GetUserQuery) (User, error) {
	return func(ctx context.Context, query GetUserQuery) (User, error) {
		u, err := repo.GetUserByID(ctx, query.ID)
		if err != nil {
			return User{}, err
		}

		// Get additional data from the players service.
		p, err := playersClient.GetPlayer(ctx, u.ID)
		if err != nil {
			return User{}, err
		}

		return User{
			ID:         u.ID,
			Email:      u.Email,
			PlayerName: p.PlayerName,
		}, nil
	}
}
