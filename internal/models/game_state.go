package models

import (
	"encoding/json"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

// GameState represents the controllers of a game
type GameState interface {
	IsCompleted() bool
}

// AdventureGameState represents the controllers of an adventure game
type AdventureGameState struct {
	WindowSize tea.WindowSizeMsg `json:"window_size"`
	Level      SavedLevel        `json:"level"`
	Stats      Stats             `json:"stats"`
	SaveID     string            `json:"save_id"`
}

// IsCompleted returns true if the level is completed
func (ags AdventureGameState) IsCompleted() bool { return ags.Level.Completed }

// GameSave represents a saved game
type GameSave struct {
	ID        string    `json:"id"`
	Player    Player    `json:"player"`
	Timestamp time.Time `json:"timestamp"`
	GameMode  string    `json:"game_mode"`
	GameState GameState `json:"game_state"`
	Score     int       `json:"score"`
}

// UnmarshalJSON decodes a GameSave from JSON
func (gs *GameSave) UnmarshalJSON(data []byte) error {
	type Alias GameSave
	aux := &struct {
		GameState json.RawMessage `json:"game_state"`
		*Alias
	}{
		Alias: (*Alias)(gs),
	}

	// decode GameSave
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	// decode GameState based on GameMode
	switch aux.GameMode {
	case "Adventure":
		var ags AdventureGameState
		if err := json.Unmarshal(aux.GameState, &ags); err != nil {
			return fmt.Errorf("failed to decode AdventureGameState: %w", err)
		}
		ags.SaveID = aux.ID
		gs.GameState = ags
	default:
		return fmt.Errorf("unsupported game mode: %s", aux.GameMode)
	}

	return nil
}

// MarshalJSON encodes a GameSave to JSON
func (gs *GameSave) MarshalJSON() ([]byte, error) {
	type Alias GameSave
	aux := &struct {
		GameState any `json:"game_state"`
		*Alias
	}{
		Alias:     (*Alias)(gs),
		GameState: gs.GameState,
	}
	return json.Marshal(aux)
}
