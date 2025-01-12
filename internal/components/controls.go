package components

import "github.com/charmbracelet/bubbles/key"

// Controls represents a set of key bindings for navigating and interacting
// with a user interface
type Controls struct {
	Up     key.Binding
	Down   key.Binding
	Select key.Binding
	Back   key.Binding
	Quit   key.Binding
}

// NewControls creates a new Controls instance with predefined key bindings
func NewControls() Controls {
	return Controls{
		Up: key.NewBinding(
			key.WithKeys("k", "up"),
			key.WithHelp("↑/k", "move up")),
		Down: key.NewBinding(
			key.WithKeys("j", "down"),
			key.WithHelp("↓/j", "move down")),
		Select: key.NewBinding(
			key.WithKeys("l", "right", "enter"),
			key.WithHelp("→/l/⏎", "select")),
		Back: key.NewBinding(
			key.WithKeys("h", "left", "esc"),
			key.WithHelp("←/h/esc", "go back")),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q/ctrl+c", "quit")),
	}
}

// NavigationHelp returns a slice of key bindings for displaying control information
func (c Controls) NavigationHelp() []key.Binding {
	return []key.Binding{
		c.Up,
		c.Down,
		c.Select,
		c.Back,
		c.Quit,
	}
}

// ContentHelp returns a slice of key bindings for displaying control information
func (c Controls) ContentHelp() []key.Binding {
	return []key.Binding{
		c.Up,
		c.Down,
		c.Back,
		c.Quit,
	}
}
