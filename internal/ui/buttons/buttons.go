package buttons

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Button represents a button with a label and an inactive state
type Button struct {
	Label    string
	Inactive bool
}

// Buttons represents a collection of buttons with styles and current selection
type Buttons struct {
	items  []*Button
	cursor int
	layout lipgloss.Position
	styles buttonStyles
}

// buttonStyles holds the buttonStyles for active and inactive buttons
type buttonStyles struct {
	Selected lipgloss.Style
	Default  lipgloss.Style
}

// New creates a new Buttons instance with the given labels
func New(labels ...string) *Buttons {
	buttons := make([]*Button, 0, len(labels))
	for _, label := range labels {
		buttons = append(buttons, &Button{Label: label})
	}

	selectedStyle := lipgloss.NewStyle().
		Width(30).
		Height(3).
		Border(lipgloss.NormalBorder()).
		Padding(1, 1).
		Align(lipgloss.Center).
		BorderForeground(activeBgColor).
		Background(activeBgColor).
		Foreground(foregroundColor)

	defaultStyle := lipgloss.NewStyle().
		Width(30).
		Height(3).
		Border(lipgloss.NormalBorder()).
		Padding(1, 1).
		Align(lipgloss.Center).
		BorderForeground(inactiveBgColor).
		Background(inactiveBgColor)

	return &Buttons{
		items:  buttons,
		layout: lipgloss.Center,
		styles: buttonStyles{
			Selected: selectedStyle,
			Default:  defaultStyle,
		},
	}
}

// Init initializes the Buttons model
func (b *Buttons) Init() tea.Cmd {
	return nil
}

// Update handles the update messages for the Buttons model
func (b *Buttons) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case Move:
		switch msg {
		case MoveUp:
			b.MoveUp()
		case MoveDown:
			b.MoveDown()
		}
	}
	return b, nil
}

// View renders the buttons
func (b *Buttons) View() string {
	var views []string
	for i, item := range b.items {
		if item.Inactive {
			continue // skip rendering inactive buttons
		}
		style := b.styles.Default
		if i == b.cursor {
			style = b.styles.Selected
		}
		views = append(views, style.Render(item.Label))
	}
	return lipgloss.JoinVertical(b.layout, views...)
}

// Items returns the buttons in the collection
func (b *Buttons) Items() []*Button {
	return b.items
}

// CurrentButton returns the currently selected button
func (b *Buttons) CurrentButton() *Button {
	if b.cursor < 0 || b.cursor >= len(b.items) {
		return nil
	}
	return b.items[b.cursor]
}

// MoveUp moves the selection up, skipping inactive buttons
func (b *Buttons) MoveUp() {
	if b.cursor > 0 {
		b.cursor--
		if b.items[b.cursor].Inactive {
			b.MoveUp() // skip inactive buttons
		}
	}
}

// MoveDown moves the selection down, skipping inactive buttons
func (b *Buttons) MoveDown() {
	lastActive := b.lastActiveIndex()
	if b.cursor < lastActive {
		b.cursor++
		if b.items[b.cursor].Inactive {
			b.MoveDown() // skip inactive buttons
		}
	}
}

// lastActiveIndex returns the index of the last active button in the list
func (b *Buttons) lastActiveIndex() int {
	for i := len(b.items) - 1; i >= 0; i-- {
		if !b.items[i].Inactive {
			return i
		}
	}
	return 0
}

var (
	activeBgColor   = lipgloss.AdaptiveColor{Light: "60", Dark: "120"}
	inactiveBgColor = lipgloss.AdaptiveColor{Light: "120", Dark: "200"}
	foregroundColor = lipgloss.CompleteColor{TrueColor: "#00000F"}
)

// Move represents a movement direction for button selection
type Move uint8

const (
	MoveUp   Move = iota // Move up in the list
	MoveDown             // Move down in the list
)
