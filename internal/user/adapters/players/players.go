package players

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dmitrymomot/go-smart-monolith/internal/user/domain"
)

type (
	// Player is a player service adapter. It is used to communicate with the
	// players service. It is used to decouple the user service from the players
	// service. You can use to map the players service responses to the internal
	// domain models.
	Player struct {
		httpClient httpClient
		config     Config
	}

	// httpClient is a low-level abstraction for the HTTP client.
	httpClient interface {
		Get(url string) (resp *http.Response, err error)
	}

	// Config is a configuration for the players service adapter.
	Config struct {
		Endpoint string
	}

	// PlayerResponse represents the response body for GetPlayer.
	PlayerResponse struct {
		UserID     string `json:"user_id"`
		PlayerName string `json:"player_name"`
	}
)

// New is a factory function that creates a new player service adapter.
func New(cnf Config, httpc httpClient) *Player {
	return &Player{
		httpClient: httpc,
		config:     cnf,
	}
}

// GetPlayer gets a player by userID from the players HTTP service.
func (p *Player) GetPlayer(ctx context.Context, userID string) (domain.Player, error) {
	res, err := p.httpClient.Get(p.config.Endpoint + "/players/" + userID)
	if err != nil {
		return domain.Player{}, fmt.Errorf("failed to get player: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return domain.Player{}, fmt.Errorf("failed to get player: %w", err)
	}

	var player PlayerResponse
	if err := json.NewDecoder(res.Body).Decode(&player); err != nil {
		return domain.Player{}, fmt.Errorf("failed to get player: %w", err)
	}

	// Map the player response to the domain model.
	// Note: you can't pass the player directly to the domain model,
	// follow single responsibility principle and create a separate struct for the domain model.
	return domain.Player{
		UserID:     player.UserID,
		PlayerName: player.PlayerName,
	}, nil
}
