package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dasvh/go-learn-vim/internal/style"
)

// Button represents a component with a label
type Button struct {
	Label    string
	Inactive bool
}

// Buttons represents a collection of interactive buttons that can be navigated through
type Buttons struct {
	items  []*Button
	cursor int
	layout lipgloss.Position
	styles buttonStyles
}

// buttonStyles defines the visual styles for buttons in different states
type buttonStyles struct {
	Selected lipgloss.Style
	Default  lipgloss.Style
}

// NewButtons creates a new Buttons component with the provided labels
func NewButtons(labels ...string) *Buttons {
	buttons := make([]*Button, 0, len(labels))
	for _, label := range labels {
		buttons = append(buttons, &Button{Label: label})
	}

	return &Buttons{
		items:  buttons,
		layout: style.Buttons.Layout,
		styles: buttonStyles{
			Selected: style.Buttons.Selected,
			Default:  style.Buttons.Default,
		},
	}
}

// Init initializes the Buttons component
func (b *Buttons) Init() tea.Cmd {
	return nil
}

// Update handles the state updates for the Buttons component
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

// View renders the button list component, displaying active buttons vertically
func (b *Buttons) View() string {
	var views []string
	for i, item := range b.items {
		if item.Inactive {
			continue // skip rendering inactive buttons
		}
		s := b.styles.Default
		if i == b.cursor {
			s = b.styles.Selected
		}
		views = append(views, s.Align(b.layout).Render(item.Label))
	}
	return lipgloss.JoinVertical(b.layout, views...)
}

// Items returns a slice of all buttons stored in the Buttons component
func (b *Buttons) Items() []*Button {
	return b.items
}

// CurrentButton returns the currently selected Button based on the cursor position
func (b *Buttons) CurrentButton() *Button {
	if b.cursor < 0 || b.cursor >= len(b.items) {
		return nil
	}
	return b.items[b.cursor]
}

// UpdateButtonState updates the state of a button
func (b *Buttons) UpdateButtonState(label string, inactive bool) {
	for _, button := range b.items {
		if button.Label == label {
			button.Inactive = inactive
			break
		}
	}
}

// MoveUp moves the cursor one position up in the buttons list
func (b *Buttons) MoveUp() {
	if b.cursor > 0 {
		b.cursor--
		if b.items[b.cursor].Inactive {
			b.MoveUp() // skip inactive buttons
		}
	}
}

// MoveDown moves the cursor to the next active button in the list
func (b *Buttons) MoveDown() {
	lastActive := b.lastActiveIndex()
	if b.cursor < lastActive {
		b.cursor++
		if b.items[b.cursor].Inactive {
			b.MoveDown() // skip inactive buttons
		}
	}
}

// lastActiveIndex iterates through the button items in reverse order and returns
// the index of the last active button
func (b *Buttons) lastActiveIndex() int {
	for i := len(b.items) - 1; i >= 0; i-- {
		if !b.items[i].Inactive {
			return i
		}
	}
	return 0
}

// Move represents a directional movement command
type Move uint8

const (
	MoveUp   Move = iota // Move up in the list
	MoveDown             // Move down in the list
)
