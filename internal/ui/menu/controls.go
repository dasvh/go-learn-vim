package menu

import "github.com/charmbracelet/bubbles/key"

// Controls holds the controls for menu navigation
type Controls struct {
	Up     key.Binding
	Down   key.Binding
	Select key.Binding
	Back   key.Binding
	Quit   key.Binding
}

// NewControls creates and returns a new Controls instance with default controls
func NewControls() Controls {
	return Controls{
		Up: key.NewBinding(
			key.WithKeys("k", "up"),
			key.WithHelp("k,↑", "move up")),
		Down: key.NewBinding(
			key.WithKeys("j", "down"),
			key.WithHelp("j,↓", "move down")),
		Select: key.NewBinding(
			key.WithKeys("l", "right", "enter"),
			key.WithHelp("l,→,⏎", "select")),
		Back: key.NewBinding(
			key.WithKeys("h", "left", "esc"),
			key.WithHelp("h,←,esc", "go back")),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit")),
	}
}

// Info returns a slice of key bindings for displaying control information
func (k Controls) Info() []key.Binding {
	return []key.Binding{
		k.Up,
		k.Down,
		k.Select,
		k.Back,
		k.Quit,
	}
}
