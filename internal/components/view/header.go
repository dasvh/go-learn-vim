package view

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

// LevelInfo displays the current level
type LevelInfo struct {
	Text  string
	Style lipgloss.Style
}

// SetLevel sets the level text
func (li *LevelInfo) SetLevel(level int) {
	li.Text = fmt.Sprintf("Level: %d", level)
}

// GameModeInfo displays the current game mode
type GameModeInfo struct {
	Text  string
	Style lipgloss.Style
}

// SetMode sets the game mode text
func (gmi *GameModeInfo) SetMode(mode string) {
	gmi.Text = fmt.Sprintf("Mode: %s", mode)
}

// GameStats displays the current game stats
type GameStats struct {
	Text  string
	Style lipgloss.Style
}

// Instructions displays the instructions for the current level
type Instructions struct {
	Text  string
	Style lipgloss.Style
}

// SetInstructions sets the instructions text
func (i *Instructions) SetInstructions(instructions string) {
	i.Text = fmt.Sprintf("Instructions: %s", instructions)
}

// TopBarWidths contains the width of each section of the top bar
type TopBarWidths struct {
	Level int
	Mode  int
	Stats int
}

// CalculateTopBarWidths calculates the width of each section of the top bar
func CalculateTopBarWidths(totalWidth int, borderSpace int) TopBarWidths {
	widthMinusBorders := totalWidth - borderSpace
	baseWidth := widthMinusBorders / 3
	remainingSpace := widthMinusBorders % 3

	widths := TopBarWidths{
		Level: baseWidth,
		Mode:  baseWidth,
		Stats: baseWidth,
	}

	if remainingSpace > 0 {
		widths.Level++
		remainingSpace--
	}
	if remainingSpace > 0 {
		widths.Mode++
		remainingSpace--
	}
	if remainingSpace > 0 {
		widths.Stats++
	}

	return widths
}

// RenderTopBar renders the top bar
func RenderTopBar(levelInfo LevelInfo, gameMode GameModeInfo, statsInfo GameStats, widths TopBarWidths) string {
	level := levelInfo.Style.Render(levelInfo.Text)
	mode := gameMode.Style.Render(gameMode.Text)
	stats := statsInfo.Style.Render(statsInfo.Text)

	topBar := lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().Width(widths.Level).Align(lipgloss.Left).Render(level),
		lipgloss.NewStyle().Width(widths.Mode).Align(lipgloss.Center).Render(mode),
		lipgloss.NewStyle().Width(widths.Stats).Align(lipgloss.Right).Render(stats),
	)

	return Styles.Adventure.Header.Border.Render(topBar)
}

// RenderLevelInstructions renders the instructions for the current level
func RenderLevelInstructions(instructions Instructions, width int) string {
	return instructions.Style.Align(lipgloss.Center).Width(width).Render(instructions.Text)
}
