package info

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dasvh/go-learn-vim/internal/models"
	"github.com/dasvh/go-learn-vim/internal/views"
)

// VimInfo represents the screen model for vim information screens
type VimInfo struct {
	*views.ContentView
}

// newScreenModel creates a new VimInfo screen model
func newScreenModel(title string, sections []models.Section, extraContent string) *VimInfo {
	display := views.NewContentView(title)
	content := lipgloss.JoinVertical(lipgloss.Center,
		display.RenderSections(sections), extraContent)
	display.SetContent(content)
	return &VimInfo{ContentView: display}
}

// NewVimInfo creates a new VimInfo screen model for the Vim Information screen
func NewVimInfo() *VimInfo {
	return newScreenModel("Vim Information", models.VimInfoSections, models.BramTribute)
}

// NewVimCheatsheet creates a new VimInfo screen model for the Vim Cheatsheet
func NewVimCheatsheet() *VimInfo {
	return newScreenModel("Vim Cheatsheet", models.CheatsheetSection, "")
}

// Update handles messages and updates the VimInfo screen model accordingly
func (vi *VimInfo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := vi.ContentView.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, vi.Controls().Back) {
			return vi, models.ChangeScreen(models.InfoMenuScreen)
		}
	}

	return vi, cmd
}
