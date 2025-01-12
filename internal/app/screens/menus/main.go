package menus

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/models"
	"github.com/dasvh/go-learn-vim/internal/views"
)

const (
	ButtonInfo   = "Info"
	ButtonLoad   = "Load Game"
	ButtonNew    = "New Game"
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
		{Label: ButtonScores},
		{Label: ButtonStats},
		{Label: ButtonQuit},
	})
	return &Main{MenuView: base}
}

// UpdateLoadButton updates the load button based on the canLoadGame flag
func (m *Main) UpdateLoadButton(canLoadGame bool) {
	m.UpdateButtonState(ButtonLoad, !canLoadGame)
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
		return models.ChangeScreen(models.InfoMenuScreen)
	case ButtonLoad:
		return models.ChangeScreen(models.LoadSaveSelectionScreen)
	case ButtonNew:
		return models.ChangeScreen(models.PlayerSelectionScreen)
	case ButtonScores:
		return models.ChangeScreen(models.ScoresScreen)
	case ButtonStats:
		return models.ChangeScreen(models.StatsScreen)
	case ButtonQuit:
		return tea.Quit
	default:
		return nil
	}
}
