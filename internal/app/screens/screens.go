package screens

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/storage"
)

// Screen represents the different screens available in the app
type Screen uint8

// ScreenTransitionMsg represents a message to transition to a different screen
type ScreenTransitionMsg struct {
	Screen Screen
	Model  tea.Model
}

// SetPlayerMsg represents a message to set the player for the game modes
type SetPlayerMsg struct {
	Player storage.Player
}

const (
	// MainMenuScreen represents the main menu screen
	MainMenuScreen Screen = iota
	// InfoMenuScreen represents the information menu screen
	InfoMenuScreen
	// VimInfoScreen represents the vim information screen
	VimInfoScreen
	// MotionsInfoScreen represents the motions information screen
	MotionsInfoScreen
	// CheatsheetInfoScreen represents the cheatsheet information screen
	CheatsheetInfoScreen
	// NewGameScreen represents the new app screen
	NewGameScreen
	// PlayerSelectionScreen represents the player selection screen
	PlayerSelectionScreen
	// LoadGameScreen represents the load app screen
	LoadGameScreen
	// AdventureModeScreen represents the adventure app mode screen
	AdventureModeScreen
	// StatsScreen represents the stats screen
	StatsScreen
)

// ChangeScreen returns a command to change the current screen to the specified screen
func ChangeScreen(to Screen) tea.Cmd {
	return func() tea.Msg {
		return to
	}
}
