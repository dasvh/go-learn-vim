package game

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/components/view"
	"github.com/dasvh/go-learn-vim/internal/game/state"
)

const (
	ButtonAdventure = "Adventure Mode"
	ButtonChallenge = "Challenge Mode"
)

// Mode represents the game mode selection screen
type Mode struct {
	*view.MenuView
}

// NewModes creates a new game mode selection screen with predefined buttons for Vim navigation
func NewModes() view.Menu {
	mode := view.NewBaseMenu("New Game Menu", []view.ButtonConfig{
		{Label: ButtonAdventure},
		{Label: ButtonChallenge, Inactive: true},
	})
	return &Mode{MenuView: mode}
}

// Update handles input messages and updates the game mode selection screen state accordingly
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

// HandleSelection implements ButtonHandler interface
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
