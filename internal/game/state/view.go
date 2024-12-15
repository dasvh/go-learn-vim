package state

import tea "github.com/charmbracelet/bubbletea"

// View represents the different views in the game
type View uint8

const (
	// MainMenu represents the main menu view
	MainMenu View = iota
	// InfoMenu represents the information menu view
	InfoMenu
	// InfoVim represents the vim information view
	InfoVim
	// InfoMotions represents the motions information view
	InfoMotions
	// InfoCheatsheet represents the cheatsheet information view
	InfoCheatsheet
	// GameMenu represents the game menu view
	GameMenu
)

// ChangeView returns a command to change the current view to the specified view
func ChangeView(to View) tea.Cmd {
	return func() tea.Msg {
		return to
	}
}
