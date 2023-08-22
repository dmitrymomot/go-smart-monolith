package domain

type Player struct {
	UserID     string `json:"user_id"`
	PlayerName string `json:"player_name"`
}

// NewPlayer creates a new player.
func NewPlayer(userID, playerName string) Player {
	return Player{
		UserID:     userID,
		PlayerName: playerName,
	}
}
