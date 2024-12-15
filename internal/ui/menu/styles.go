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
	ContentTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("0")).
				Background(lipgloss.Color("200")).
				MarginBottom(1).
				Bold(true).
				Underline(true)
	ContentBodyStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("120")).
				PaddingLeft(1).
				PaddingRight(1)
	ContentSectionStyle = lipgloss.NewStyle().
				BorderForeground(lipgloss.Color("120")).
				Border(lipgloss.NormalBorder()).
				PaddingTop(1).
				PaddingLeft(1).
				PaddingRight(1)
)
