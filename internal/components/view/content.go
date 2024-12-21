package view

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Content represents an interface that extends the View interface
// it provides a method to set the content of the implementing type
type Content interface {
	View
	SetContent(string)
}

// ContentView represents a view component that embeds
// the BaseView struct and includes a content field to manage the content
type ContentView struct {
	*BaseView
	content string
}

// NewDisplayView returns a new ContentView instance
func NewDisplayView(title string) *ContentView {
	return &ContentView{
		BaseView: NewBaseView(title),
	}
}

// Init initializes the ContentView
func (cv *ContentView) Init() tea.Cmd { return nil }

// Update handles messages and updates the ContentView accordingly
func (cv *ContentView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		cv.size = msg
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, cv.controls.Quit):
			return cv, tea.Quit
		}
	}
	return cv, nil
}

// View renders the ContentView
func (cv *ContentView) View() string {
	views := []string{
		cv.mainTitle,
		cv.subtitle,
		cv.content,
		cv.help.ShortHelpView(cv.controls.ContentHelp()),
	}

	return lipgloss.Place(cv.size.Width, cv.size.Height, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center, views...))
}

// SetContent sets the content of the ContentView
func (cv *ContentView) SetContent(content string) {
	cv.content = content
}
