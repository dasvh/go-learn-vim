package info

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	ui "github.com/dasvh/go-learn-vim/internal/components/view"
	"github.com/dasvh/go-learn-vim/internal/game/state"
)

// MotionsInfo represents Vim Motions Information view
type MotionsInfo struct {
	*ui.ContentView
}

// NewMotionsInfo creates and returns a new MotionsInfo instance with initialized display view
func NewMotionsInfo() *MotionsInfo {
	display := ui.NewDisplayView("Motions Information")
	content := "Hello World!"
	display.SetContent(content)
	return &MotionsInfo{ContentView: display}
}

// Update handles updates to the MotionsInfo screen model based on incoming messages
func (m *MotionsInfo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := m.ContentView.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.Controls().Back) {
			return m, state.ChangeScreen(state.InfoMenuScreen)
		}
	}

	return m, cmd
}
