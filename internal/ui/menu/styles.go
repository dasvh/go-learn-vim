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
	ContentTitleStyle   = lipgloss.NewStyle().Bold(true).Underline(true).Foreground(lipgloss.Color("33"))
	ContentBodyStyle    = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("231"))
	ContentSectionStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(1).Margin(1).BorderForeground(lipgloss.Color("69"))
)
