package state

import tea "github.com/charmbracelet/bubbletea"

// View represents the different views in the game
type View uint8

const (
	// MainView represents the main menu view
	MainView View = iota
	// InfoView represents the information menu view
	InfoView
	// InfoVimView represents the vim information view
	InfoVimView
	// GameView represents the game menu view
	GameView
)

// ChangeView returns a command to change the current view to the specified view
func ChangeView(to View) tea.Cmd {
	return func() tea.Msg {
		return to
	}
}
