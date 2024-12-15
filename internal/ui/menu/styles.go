package menu

import "github.com/charmbracelet/lipgloss"

var (
	DefaultTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#00ADD8")).
				Bold(true)
	DefaultSubtitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#113344")).
				Background(lipgloss.Color("#f0f0f0")).
				Height(1).
				Width(45).
				MarginBottom(2).
				Bold(true).
				Blink(true)
)
