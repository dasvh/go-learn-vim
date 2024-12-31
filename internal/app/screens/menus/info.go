package menus

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/app/state"
	"github.com/dasvh/go-learn-vim/internal/views"
)

const (
	ButtonVim        = "Vim"
	ButtonMotions    = "Motions"
	ButtonCheatsheet = "Cheatsheet"
)

// Info represents the info menu screen
type Info struct {
	*views.MenuView
}

// NewInfoMenu creates a new info menu screen
func NewInfoMenu() views.Menu {
	info := views.NewBaseMenu("Info Menu", []views.ButtonConfig{
		{Label: ButtonVim},
		{Label: ButtonMotions},
		{Label: ButtonCheatsheet, Inactive: true},
	})
	return &Info{MenuView: info}
}

// Update handles state updates based on incoming messages
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

// HandleSelection handles the selection of a button
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
