package views

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/game/state"
	ui "github.com/dasvh/go-learn-vim/internal/ui/menu"
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
	*ui.BaseMenu
}

// NewMainView returns a new Main instance
func NewMainView() ui.Menu {
	base := ui.NewBaseMenu("Main Menu", []ui.ButtonConfig{
		{Label: ButtonInfo},
		{Label: ButtonNew},
		{Label: ButtonLoad, Inactive: true},
		{Label: ButtonScores},
		{Label: ButtonStats},
		{Label: ButtonQuit},
	})
	return &Main{BaseMenu: base}
}

// Update handles messages and transitions between menu states
func (m *Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := m.BaseMenu.Update(msg)

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
		return state.ChangeView(state.InfoView)
	case ButtonQuit:
		return tea.Quit
	default:
		return nil
	}
}
