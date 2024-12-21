package info

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/components/view"
	"github.com/dasvh/go-learn-vim/internal/game/state"
)

// MotionsInfo represents Vim Motions Information view
type MotionsInfo struct {
	*view.ContentView
}

// NewMotionsInfo creates and returns a new MotionsInfo instance with initialized display view
func NewMotionsInfo() *MotionsInfo {
	display := view.NewDisplayView("Motions Information")
	display.SetContent("Hello world!")
	return &MotionsInfo{ContentView: display}
}

// Update handles updates to the MotionsInfo screen model based on incoming messages
func (mi *MotionsInfo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := mi.ContentView.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, mi.Controls().Back) {
			return mi, state.ChangeScreen(state.InfoMenuScreen)
		}
	}

	return mi, cmd
}
