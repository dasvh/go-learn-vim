package info

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/components/view"
	"github.com/dasvh/go-learn-vim/internal/game/state"
)

const (
	ButtonVim        = "Vim"
	ButtonMotions    = "Motions"
	ButtonCheatsheet = "Cheatsheet"
)

// Info represents the information screen of the game
type Info struct {
	*view.MenuView
}

// NewInfoMenu creates a new info menu with predefined buttons for Vim navigation
func NewInfoMenu() view.Menu {
	info := view.NewBaseMenu("Info Menu", []view.ButtonConfig{
		{Label: ButtonVim},
		{Label: ButtonMotions},
		{Label: ButtonCheatsheet, Inactive: true},
	})
	return &Info{MenuView: info}
}

// Update handles input messages and updates the info screen state accordingly
func (i *Info) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := i.MenuView.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, i.Controls().Select) {
			return i, i.HandleSelection()
		}
		if key.Matches(msg, i.Controls().Back) {
			return i, state.ChangeScreen(state.MainMenuScreen)
		}
	}

	return i, cmd
}

// HandleSelection implements ButtonHandler interface
func (i *Info) HandleSelection() tea.Cmd {
	selected := i.CurrentButton()
	if selected == nil || selected.Inactive {
		return nil
	}

	switch selected.Label {
	case ButtonVim:
		return state.ChangeScreen(state.VimInfoScreen)
	case ButtonMotions:
		return state.ChangeScreen(state.MotionsInfoScreen)
	case ButtonCheatsheet:
		return state.ChangeScreen(state.CheatsheetInfoScreen)
	default:
		return nil
	}
}
