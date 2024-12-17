package view

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dasvh/go-learn-vim/internal/components/buttons"
)

// Menu represents an interface that extends the View interface
// it provides a method to handle button selection
type Menu interface {
	View
	ButtonHandler
}

// MenuView represents a view component that embeds
// the BaseView struct and includes a Buttons field to manage the buttons
type MenuView struct {
	*BaseView
	buttons *buttons.Buttons
}

// ButtonHandler defines an interface for handling button selections
type ButtonHandler interface {
	// HandleSelection handles the selection of a button
	HandleSelection() tea.Cmd
}

// ButtonConfig represents the configuration for a button component
type ButtonConfig struct {
	// Label represents the text displayed on the button.
	Label string
	// Inactive determines if the button is rendered.
	Inactive bool
}

// NewBaseMenu returns a new MenuView instance
// with the provided subtitle and button configurations
func NewBaseMenu(subtitle string, buttonConfigs []ButtonConfig) *MenuView {
	labels := make([]string, len(buttonConfigs))
	for i, config := range buttonConfigs {
		labels[i] = config.Label
	}

	m := &MenuView{
		BaseView: NewBaseView(subtitle),
		buttons:  buttons.New(labels...),
	}

	for i, config := range buttonConfigs {
		m.buttons.Items()[i].Inactive = config.Inactive
	}

	return m
}

// Init initializes the MenuView
func (m *MenuView) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the MenuView accordingly
func (m *MenuView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.size = msg
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.controls.Up):
			m.buttons.Update(buttons.MoveUp)
		case key.Matches(msg, m.controls.Down):
			m.buttons.Update(buttons.MoveDown)
		case key.Matches(msg, m.controls.Select):
			// Leave `Select` handling to derived menus
			return m, nil
		case key.Matches(msg, m.controls.Back):
			// Leave `Back` handling to derived menus
			return m, nil
		case key.Matches(msg, m.controls.Quit):
			return m, tea.Quit
		}
	}
	return m, nil
}

// View renders the MenuView
func (m *MenuView) View() string {
	views := []string{
		m.mainTitle,
		m.subtitle,
	}

	views = append(views,
		m.buttons.View()+"\n",
		m.help.ShortHelpView(m.controls.NavigationHelp()),
	)

	return lipgloss.Place(m.size.Width, m.size.Height, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center, views...))
}

// CurrentButton returns the currently selected button
func (m *MenuView) CurrentButton() *buttons.Button {
	return m.buttons.CurrentButton()
}
