package menus

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/app/state"
	"github.com/dasvh/go-learn-vim/internal/views"
)

const (
	ButtonAdventure = "Adventure Mode"
	ButtonChallenge = "Challenge Mode"
)

// Mode represents the mode selection screen
type Mode struct {
	*views.MenuView
}

// NewGameModes creates a new game mode selection screen
func NewGameModes() views.Menu {
	mode := views.NewBaseMenu("New Game Menu", []views.ButtonConfig{
		{Label: ButtonAdventure},
		{Label: ButtonChallenge, Inactive: true},
	})
	return &Mode{MenuView: mode}
}

// Update handles state updates based on incoming messages
func (m *Mode) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := m.MenuView.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.Controls().Select) {
			return m, m.HandleSelection()
		}
		if key.Matches(msg, m.Controls().Back) {
			return m, state.ChangeScreen(state.MainMenuScreen)
		}
	}

	return m, cmd
}

// HandleSelection handles the selection of a button
func (m *Mode) HandleSelection() tea.Cmd {
	selected := m.CurrentButton()
	if selected == nil || selected.Inactive {
		return nil
	}

	switch selected.Label {
	case ButtonAdventure:
		return state.ChangeScreen(state.AdventureModeScreen)
	default:
		return nil
	}
}
