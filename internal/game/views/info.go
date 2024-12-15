package views

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

// Info represents the information menu view
type Info struct {
	*ui.BaseMenu
}

// NewInfoView returns a new Info instance
func NewInfoView() ui.Menu {
	base := ui.NewBaseMenu("Info Menu", []ui.ButtonConfig{
		{Label: ButtonVim},
		{Label: ButtonMotions},
		{Label: ButtonCheatsheet, Inactive: true},
	})
	return &Info{BaseMenu: base}
}

// Update handles messages and transitions between menu states
func (m *Info) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := m.BaseMenu.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.Controls().Select) {
			//return m, m.handleSelection()
			return m, m.HandleSelection()
		}
		if key.Matches(msg, m.Controls().Back) {
			return m, state.ChangeView(state.MainView)
		}
	}

	return m, cmd
}

// HandleSelection implements ButtonHandler interface
func (m *Info) HandleSelection() tea.Cmd {
	selected := m.CurrentButton()
	if selected == nil || selected.Inactive {
		return nil
	}

	switch selected.Label {
	case ButtonVim:
		return tea.Println("Vim")
	case ButtonMotions:
		return tea.Println("Motions")
	default:
		return nil
	}
}
