package screens

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/components/view"
	"github.com/dasvh/go-learn-vim/internal/game/state"
)

const (
	ButtonInfo   = "Info"
	ButtonNew    = "New Game"
	ButtonLoad   = "Load Game"
	ButtonScores = "Scores"
	ButtonStats  = "Stats"
	ButtonQuit   = "Quit"
)

// Main represents the main menu view
type Main struct {
	*view.MenuView
}

// NewMainMenu returns a new Main instance
func NewMainMenu() view.Menu {
	base := view.NewBaseMenu("Main Menu", []view.ButtonConfig{
		{Label: ButtonInfo},
		{Label: ButtonNew},
		{Label: ButtonLoad, Inactive: true},
		{Label: ButtonScores, Inactive: true},
		{Label: ButtonStats, Inactive: true},
		{Label: ButtonQuit},
	})
	return &Main{MenuView: base}
}

// Update handles messages and transitions between menu states
func (m *Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := m.MenuView.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.Controls().Select) {
			return m, m.HandleSelection()
		}
	}

	return m, cmd
}

// HandleSelection implements ButtonHandler interface
func (m *Main) HandleSelection() tea.Cmd {
	selected := m.CurrentButton()
	if selected == nil || selected.Inactive {
		return nil
	}

	switch selected.Label {
	case ButtonInfo:
		return state.ChangeScreen(state.InfoMenuScreen)
	case ButtonNew:
		return state.ChangeScreen(state.NewGameScreen)
	case ButtonQuit:
		return tea.Quit
	default:
		return nil
	}
}
