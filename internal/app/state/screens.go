package state

import tea "github.com/charmbracelet/bubbletea"

// Screen represents the different screens available in the app
type Screen uint8

// ScreenTransitionMsg represents a message to transition to a different screen
type ScreenTransitionMsg struct {
	Screen Screen
	Model  tea.Model
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
	// LoadGameScreen represents the load app screen
	LoadGameScreen
	// AdventureModeScreen represents the adventure app mode screen
	AdventureModeScreen
)

// ChangeScreen returns a command to change the current screen to the specified screen
func ChangeScreen(to Screen) tea.Cmd {
	return func() tea.Msg {
		return to
	}
}
