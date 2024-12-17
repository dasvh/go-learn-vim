package view

import "github.com/charmbracelet/lipgloss"

// theme defines the color palette for the view components
var theme = struct {
	Primary   lipgloss.Color
	Secondary lipgloss.Color
	Content   lipgloss.Color
	BgLight   lipgloss.Color
	BgDark    lipgloss.Color
}{
	Primary:   lipgloss.Color("#00ADD8"),
	Secondary: lipgloss.Color("#113344"),
	Content:   lipgloss.Color("120"),
	BgLight:   lipgloss.Color("#f0f0f0"),
	BgDark:    lipgloss.Color("200"),
}

// Styles defines the styling configuration for view components
var Styles = struct {
	Title    lipgloss.Style
	Subtitle lipgloss.Style
	Display  struct {
		Section lipgloss.Style
		Title   lipgloss.Style
		Text    lipgloss.Style
	}
}{
	Title: lipgloss.NewStyle().
		Foreground(theme.Primary).
		Bold(true),
	Subtitle: lipgloss.NewStyle().
		Foreground(theme.Secondary).
		Background(theme.BgLight).
		Height(1).
		Width(45).
		MarginBottom(2).
		Bold(true).
		Blink(true).
		Align(lipgloss.Center),
	Display: struct {
		Section lipgloss.Style
		Title   lipgloss.Style
		Text    lipgloss.Style
	}{
		Section: lipgloss.NewStyle().
			BorderForeground(theme.Content).
			Border(lipgloss.NormalBorder()).
			Padding(1, 1),
		Title: lipgloss.NewStyle().
			Foreground(theme.Secondary).
			Background(theme.BgDark).
			MarginBottom(1).
			Bold(true).
			Underline(true),
		Text: lipgloss.NewStyle().
			Foreground(theme.Content).
			Padding(0, 1),
	},
}
