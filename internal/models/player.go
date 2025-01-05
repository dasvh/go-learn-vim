package models

import "time"

const PlayerNameMaxLength = 20

// Player represents a player with an ID and a name
type Player struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// PlayerMovement represents the result of a player action
type PlayerMovement struct {
	UpdatedPosition    Position
	Completed          bool
	ValidMove          bool
	InstructionMessage string
}

// HighScore represents a player's high score entry
type HighScore struct {
	PlayerName string    `json:"player_name"`
	Level      int       `json:"level"`
	Score      int       `json:"score"`
	Timestamp  time.Time `json:"timestamp"`
}
