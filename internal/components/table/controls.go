package table

import "github.com/charmbracelet/bubbles/key"

// Controls defines key bindings for interacting with a table
type Controls struct {
	Up         key.Binding
	Down       key.Binding
	GotoTop    key.Binding
	GotoBottom key.Binding
	Select     key.Binding
	Back       key.Binding
	Quit       key.Binding
}

// NewTableControls initializes a new set of TableControls
func NewTableControls() Controls {
	return Controls{
		Up: key.NewBinding(
			key.WithKeys("k", "up"),
			key.WithHelp("↑/k", "move up")),
		Down: key.NewBinding(
			key.WithKeys("j", "down"),
			key.WithHelp("↓/j", "move down")),
		GotoTop: key.NewBinding(
			key.WithKeys("g", "home"),
			key.WithHelp("g/home", "go to top")),
		GotoBottom: key.NewBinding(
			key.WithKeys("G", "end"),
			key.WithHelp("G/end", "go to bottom")),
		Select: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("⏎", "select row")),
		Back: key.NewBinding(
			key.WithKeys("h", "left", "esc"),
			key.WithHelp("←/h/esc", "go back")),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q/ctrl+c", "quit")),
	}
}

// ShortHelp returns a slice of TableControls bindings for displaying help information
func (tc Controls) ShortHelp() []key.Binding {
	return []key.Binding{
		tc.Up,
		tc.Down,
		tc.GotoTop,
		tc.GotoBottom,
		tc.Select,
		tc.Back,
		tc.Quit,
	}
}
