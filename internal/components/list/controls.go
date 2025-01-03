package list

import "github.com/charmbracelet/bubbles/key"

// InsertControls is a set of key bindings for inserting items into a list
type InsertControls struct {
	Insert  key.Binding
	Select  key.Binding
	Confirm key.Binding
	Cancel  key.Binding
	Quit    key.Binding
}

// NewInsertControls creates a new InsertControls instance with predefined key bindings
func NewInsertControls() InsertControls {
	return InsertControls{
		Insert: key.NewBinding(
			key.WithKeys("i"),
			key.WithHelp("i", "insert")),
		Select: key.NewBinding(
			key.WithKeys("l", "right", "enter"),
			key.WithHelp("→/l/⏎", "select")),
		Confirm: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("⏎", "confirm")),
		Cancel: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "cancel")),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "quit")),
	}
}

// InputHelp returns a slice of InsertControls bindings for displaying help information
func (ic InsertControls) InputHelp() []key.Binding {
	return []key.Binding{
		ic.Confirm,
		ic.Cancel,
		ic.Quit,
	}
}

// SelectionControls is a set of key bindings for selecting items in a list
type SelectionControls struct {
	Up     key.Binding
	Down   key.Binding
	Select key.Binding
	Back   key.Binding
	Quit   key.Binding
}

// NewSelectionControls creates a new SelectionControls instance with predefined key bindings
func NewSelectionControls() SelectionControls {
	return SelectionControls{
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
			key.WithKeys("h", "left"),
			key.WithHelp("←/h", "go back")),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q/ctrl+c", "quit")),
	}
}
