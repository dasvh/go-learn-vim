package info

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/game/state"
	ui "github.com/dasvh/go-learn-vim/internal/ui/menu"
)

const (
	ButtonVim        = "Vim"
	ButtonMotions    = "Motions"
	ButtonCheatsheet = "Cheatsheet"
)

// Menu represents the information menu view
type Menu struct {
	*ui.BaseMenu
}

// NewInfoMenu returns a new Menu instance
func NewInfoMenu() ui.Menu {
	base := ui.NewBaseMenu("Info Menu", []ui.ButtonConfig{
		{Label: ButtonVim},
		{Label: ButtonMotions},
		{Label: ButtonCheatsheet, Inactive: true},
	})
	return &Menu{BaseMenu: base}
}

// Update handles messages and transitions between menu states
func (m *Menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := m.BaseMenu.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.Controls().Select) {
			return m, m.HandleSelection()
		}
		if key.Matches(msg, m.Controls().Back) {
			return m, state.ChangeView(state.MainMenu)
		}
	}

	return m, cmd
}

// HandleSelection implements ButtonHandler interface
func (m *Menu) HandleSelection() tea.Cmd {
	selected := m.CurrentButton()
	if selected == nil || selected.Inactive {
		return nil
	}

	switch selected.Label {
	case ButtonVim:
		return state.ChangeView(state.InfoVim)
	case ButtonMotions:
		return tea.Println("Motions")
	default:
		return nil
	}
}
