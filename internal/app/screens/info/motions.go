package info

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/app/screens"
	"github.com/dasvh/go-learn-vim/internal/views"
)

// MotionsInfo represents the screen model for the Motions Information screen
type MotionsInfo struct {
	*views.ContentView
}

// NewMotionsInfo creates a new MotionsInfo screen model
func NewMotionsInfo() *MotionsInfo {
	display := views.NewContentView("Motions Information")
	display.SetContent("Hello world!")
	return &MotionsInfo{ContentView: display}
}

// Update handles messages and updates the MotionsInfo screen model accordingly
func (mi *MotionsInfo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := mi.ContentView.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, mi.Controls().Back) {
			return mi, screens.ChangeScreen(screens.InfoMenuScreen)
		}
	}

	return mi, cmd
}
