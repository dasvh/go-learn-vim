package adventure

import "github.com/charmbracelet/bubbles/key"

// Controls represents Controls for any level in the adventure mode
type Controls struct {
	MoveLeft    key.Binding
	MoveRight   key.Binding
	MoveUp      key.Binding
	MoveDown    key.Binding
	WordForward key.Binding
	WordBack    key.Binding
	EndOfWord   key.Binding
	StartOfLine key.Binding
	EndOfLine   key.Binding
	FirstLine   key.Binding
	LastLine    key.Binding
	Escape      key.Binding
	Quit        key.Binding
}

// NewBasicControls creates a new BasicControls instance with predefined key bindings
func NewBasicControls() Controls {
	return Controls{
		MoveLeft: key.NewBinding(
			key.WithKeys("h"),
			key.WithHelp("h", "move left")),
		MoveRight: key.NewBinding(
			key.WithKeys("l"),
			key.WithHelp("l", "move right")),
		MoveUp: key.NewBinding(
			key.WithKeys("k"),
			key.WithHelp("k", "move up")),
		MoveDown: key.NewBinding(
			key.WithKeys("j"),
			key.WithHelp("j", "move down")),
		Escape: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "back to selection")),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit")),
	}
}

// BasicHelp returns a slice of key bindings for displaying basic adventure control information
func (c Controls) BasicHelp() []key.Binding {
	return []key.Binding{
		c.MoveLeft,
		c.MoveDown,
		c.MoveUp,
		c.MoveRight,
		c.Escape,
		c.Quit,
	}
}
