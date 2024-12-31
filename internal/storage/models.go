package storage

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/app/screens/adventure/level"
	"github.com/dasvh/go-learn-vim/internal/app/state"
)

// AdventureGameState represents the state of an adventure app
type AdventureGameState struct {
	Size  tea.WindowSizeMsg `json:"size"`
	Level level.SavedLevel  `json:"level"`
	Stats state.Stats       `json:"stats"`
}
