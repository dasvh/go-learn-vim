package menus

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/app/screens"
	"github.com/dasvh/go-learn-vim/internal/views"
)

const (
	ButtonInfo   = "Info"
	ButtonNew    = "New Game"
	ButtonLoad   = "Load Game"
	ButtonScores = "Scores"
	ButtonStats  = "Stats"
	ButtonQuit   = "Quit"
)

// Main represents the main menu screen
type Main struct {
	*views.MenuView
}

// NewMainMenu creates a new main menu screen
func NewMainMenu(canLoadGame bool) views.Menu {
	base := views.NewBaseMenu("Main Menu", []views.ButtonConfig{
		{Label: ButtonInfo},
		{Label: ButtonLoad, Inactive: !canLoadGame},
		{Label: ButtonNew},
		{Label: ButtonScores, Inactive: true},
		{Label: ButtonStats, Inactive: true},
		{Label: ButtonQuit},
	})
	return &Main{MenuView: base}
}

// Update handles state updates based on incoming messages
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

// HandleSelection handles the selection of a button
func (m *Main) HandleSelection() tea.Cmd {
	selected := m.CurrentButton()
	if selected == nil || selected.Inactive {
		return nil
	}

	switch selected.Label {
	case ButtonInfo:
		return screens.ChangeScreen(screens.InfoMenuScreen)
	case ButtonLoad:
		return screens.ChangeScreen(screens.LoadGameScreen)
	case ButtonNew:
		return screens.ChangeScreen(screens.PlayerSelectionScreen)
	case ButtonQuit:
		return tea.Quit
	default:
		return nil
	}
}
