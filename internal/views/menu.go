package views

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dasvh/go-learn-vim/internal/components"
)

// Menu represents an interface that extends the View interface
// it provides a method to handle button selection
type Menu interface {
	View
	ButtonHandler
}

// MenuView represents a views component that embeds
// the BaseView struct and includes a Buttons field to manage the buttons
type MenuView struct {
	*BaseView
	buttons *components.Buttons
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
		buttons:  components.NewButtons(labels...),
	}

	for i, config := range buttonConfigs {
		m.buttons.Items()[i].Inactive = config.Inactive
	}

	return m
}

// Init initializes the MenuView
func (mv *MenuView) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the MenuView accordingly
func (mv *MenuView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		mv.size = msg
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, mv.controls.Up):
			mv.buttons.Update(components.MoveUp)
		case key.Matches(msg, mv.controls.Down):
			mv.buttons.Update(components.MoveDown)
		case key.Matches(msg, mv.controls.Select):
			// Leave `Select` handling to derived menus
			return mv, nil
		case key.Matches(msg, mv.controls.Back):
			// Leave `Back` handling to derived menus
			return mv, nil
		case key.Matches(msg, mv.controls.Quit):
			return mv, tea.Quit
		}
	}
	return mv, nil
}

// View renders the MenuView
func (mv *MenuView) View() string {
	views := []string{
		mv.title.Main,
		mv.title.Subtitle,
	}

	views = append(views,
		mv.buttons.View()+"\n",
		mv.help.ShortHelpView(mv.controls.NavigationHelp()),
	)

	return lipgloss.Place(mv.size.Width, mv.size.Height, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center, views...))
}

// CurrentButton returns the currently selected button
func (mv *MenuView) CurrentButton() *components.Button {
	return mv.buttons.CurrentButton()
}
