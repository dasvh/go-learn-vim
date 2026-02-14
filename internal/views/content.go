package views

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dasvh/go-learn-vim/internal/models"
	"github.com/dasvh/go-learn-vim/internal/style"
)

// Content represents an interface that extends the View interface
// it provides a method to set the content of the implementing type
type Content interface {
	View
	SetContent(string)
}

// ContentView represents a views component that embeds
// the BaseView struct and includes a content field to manage the content
type ContentView struct {
	*BaseView
	viewport viewport.Model
	content  string
}

// NewContentView returns a new ContentView instance
func NewContentView(title string) *ContentView {
	return &ContentView{
		BaseView: NewBaseView(title),
		content:  "",
		viewport: viewport.New(0, 0),
	}
}

// RenderSections takes a slice of Section structs and returns a formatted string representation
func (cv *ContentView) RenderSections(sections []models.Section) string {
	var rendered strings.Builder
	for i, section := range sections {
		rendered.WriteString(lipgloss.JoinVertical(0,
			style.Styles.Display.Title.Render(section.Title),
			style.Styles.Display.Text.Render(section.Content),
		))
		if i < len(sections)-1 {
			rendered.WriteString("\n\n")
		}
	}
	return style.Styles.Display.Section.Render(rendered.String())
}

// Init initializes the ContentView
func (cv *ContentView) Init() tea.Cmd { return nil }

// Update handles messages and updates the ContentView accordingly
func (cv *ContentView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		cv.size = msg
		headerHeight := 10
		helpHeight := 1
		cv.viewport.Height = msg.Height - headerHeight - helpHeight
		cv.viewport.SetContent(cv.content)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, cv.controls.Up):
			cv.viewport.ScrollUp(1)
		case key.Matches(msg, cv.controls.Down):
			cv.viewport.ScrollDown(1)
		case key.Matches(msg, cv.controls.Quit):
			return cv, tea.Quit
		}
	}
	return cv, nil
}

// View renders the ContentView with a fixed title and a scrollable content area
func (cv *ContentView) View() string {
	views := []string{
		cv.title.Main,
		cv.title.Subtitle,
		cv.viewport.View(),
		cv.help.ShortHelpView(cv.controls.ContentHelp()),
	}

	return lipgloss.Place(cv.size.Width, cv.size.Height, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center, views...))
}

// SetContent sets the content of the ContentView
func (cv *ContentView) SetContent(content string) {
	cv.content = content
}
