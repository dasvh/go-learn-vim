package storage

import (
	"encoding/json"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/app/screens/adventure/level"
	"time"
)

// Player represents a player with an ID and a name
type Player struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GameState represents the controllers of a game
type GameState interface {
	IsCompleted() bool
}

// AdventureGameState represents the controllers of an adventure game
type AdventureGameState struct {
	WindowSize tea.WindowSizeMsg `json:"window_size"`
	Level      level.SavedLevel  `json:"level"`
	Stats      Stats             `json:"stats"`
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
		GameState interface{} `json:"game_state"`
		*Alias
	}{
		Alias:     (*Alias)(gs),
		GameState: gs.GameState,
	}
	return json.Marshal(aux)
}

// LifetimeStats represents aggregated statistics for all games
type LifetimeStats struct {
	TotalKeystrokes int            `json:"total_keystrokes"`
	TotalPlaytime   int            `json:"total_playtime"`
	TotalGames      int            `json:"total_games"`
	KeyPresses      map[string]int `json:"key_presses"`
}

// NewLifetimeStats initializes an empty LifetimeStats
func NewLifetimeStats() *LifetimeStats {
	return &LifetimeStats{
		TotalKeystrokes: 0,
		TotalPlaytime:   0,
		TotalGames:      0,
		KeyPresses:      make(map[string]int),
	}
}

// Merge aggregates stats from another Stats instance
func (ls *LifetimeStats) Merge(stats Stats) {
	ls.TotalKeystrokes += stats.TotalKeystrokes
	ls.TotalPlaytime += stats.TimeElapsed

	for key, count := range stats.KeyPresses {
		ls.KeyPresses[key] += count
	}
}
