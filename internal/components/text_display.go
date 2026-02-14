package components

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

// TextDisplay displays styled text
type TextDisplay struct {
	Text  string
	Style lipgloss.Style
}

// SetText sets the display text for the component
func (dc *TextDisplay) SetText(format string, args ...any) {
	dc.Text = fmt.Sprintf(format, args...)
}

// Render generates the styled text
func (dc *TextDisplay) Render() string {
	return dc.Style.Render(dc.Text)
}
