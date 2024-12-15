package menu

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dasvh/go-learn-vim/internal/ui/buttons"
)

const (
	// Title represents the subtitle of the menu
	Title = `

 ██████╗  ██████╗     ██╗     ███████╗ █████╗ ██████╗ ███╗   ██╗    ██╗   ██╗██╗███╗   ███╗
██╔════╝ ██╔═══██╗    ██║     ██╔════╝██╔══██╗██╔══██╗████╗  ██║    ██║   ██║██║████╗ ████║
██║  ███╗██║   ██║    ██║     █████╗  ███████║██████╔╝██╔██╗ ██║    ██║   ██║██║██╔████╔██║
██║   ██║██║   ██║    ██║     ██╔══╝  ██╔══██║██╔══██╗██║╚██╗██║    ╚██╗ ██╔╝██║██║╚██╔╝██║
╚██████╔╝╚██████╔╝    ███████╗███████╗██║  ██║██║  ██║██║ ╚████║     ╚████╔╝ ██║██║ ╚═╝ ██║
 ╚═════╝  ╚═════╝     ╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝      ╚═══╝  ╚═╝╚═╝     ╚═╝
`
)

// BaseMenu represents the base menu structure
type BaseMenu struct {
	size     tea.WindowSizeMsg
	controls Controls
	help     help.Model
	buttons  *buttons.Buttons
	title    string
	content  string
}

// NewBaseMenu returns a new BaseMenu instance
// with the provided subtitle and button configurations
func NewBaseMenu(subtitle string, buttonConfigs []ButtonConfig) *BaseMenu {
	labels := make([]string, len(buttonConfigs))
	for i, config := range buttonConfigs {
		labels[i] = config.Label
	}

	m := &BaseMenu{
		controls: NewControls(),
		help:     help.New(),
		buttons:  buttons.New(labels...),
		title:    subtitle,
	}

	for i, config := range buttonConfigs {
		m.buttons.Items()[i].Inactive = config.Inactive
	}

	return m
}

// Init initializes the Menu
func (m *BaseMenu) Init() tea.Cmd {
	return nil
}

// Update handles messages and transitions between menu states
func (m *BaseMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

// View renders the Menu
func (m *BaseMenu) View() string {
	views := []string{
		DefaultTitleStyle.Render(Title),
		DefaultSubtitleStyle.Align(lipgloss.Center).Render(m.title),
	}

	if m.content != "" {
		views = append(views,
			m.content,
			m.help.ShortHelpView(m.controls.ContentHelp()),
		)
	} else {
		views = append(views,
			m.buttons.View()+"\n",
			m.help.ShortHelpView(m.controls.NavigationHelp()),
		)
	}
	return lipgloss.Place(m.size.Width, m.size.Height, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center, views...))
}

// Controls returns the key bindings for the menu
func (m *BaseMenu) Controls() Controls {
	return m.controls
}

// CurrentButton returns the currently selected button
func (m *BaseMenu) CurrentButton() *buttons.Button {
	return m.buttons.CurrentButton()
}

func (m *BaseMenu) SetContent(content string) {
	m.content = content
}
