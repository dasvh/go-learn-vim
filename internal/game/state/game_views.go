package state

import tea "github.com/charmbracelet/bubbletea"

// GameScreen represents the different screens/views available in the game
type GameScreen uint8

const (
	// MainMenuScreen represents the main menu view
	MainMenuScreen GameScreen = iota
	// InfoMenuScreen represents the information menu view
	InfoMenuScreen
	// VimInfoScreen represents the vim information view
	VimInfoScreen
	// MotionsInfoScreen represents the motions information view
	MotionsInfoScreen
	// CheatsheetInfoScreen represents the cheatsheet information view
	CheatsheetInfoScreen
	// NewGameScreen represents the new game view
	NewGameScreen
	// AdventureModeScreen represents the adventure game mode view
	AdventureModeScreen
)

// ChangeScreen returns a command to change the current view to the specified view
func ChangeScreen(to GameScreen) tea.Cmd {
	return func() tea.Msg {
		return to
	}
}
