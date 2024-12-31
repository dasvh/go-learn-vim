package storage

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/game/screens/game"
	"github.com/dasvh/go-learn-vim/internal/game/screens/game/adventure/level"
)

// AdventureGameState represents the state of an adventure game
type AdventureGameState struct {
	Size  tea.WindowSizeMsg `json:"size"`
	Level level.SavedLevel  `json:"level"`
	Stats game.Stats        `json:"stats"`
}
