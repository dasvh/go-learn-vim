package adventure

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
	"github.com/dasvh/go-learn-vim/internal/app/screens/adventure/character"
	"github.com/dasvh/go-learn-vim/internal/style"
	"github.com/dasvh/go-learn-vim/internal/views"
)

const (
	StatsFormat = "Stats: Keystrokes: %d, Time: %d s"
)

// GameView contains the components to render the app views
type GameView struct {
	Field      [][]rune
	Border     lipgloss.Style
	Background lipgloss.Style
}

// InitializeComponents initializes the components for the app views
func InitializeComponents() (views.LevelInfo, views.GameModeInfo, views.GameStats, views.Instructions, GameView) {
	levelInfo := views.LevelInfo{
		Style: style.Styles.Adventure.Header.Level,
	}
	levelInfo.SetLevel(0)

	currentMode := views.GameModeInfo{
		Style: style.Styles.Adventure.Header.Mode,
	}
	currentMode.SetMode("Adventure")

	stats := views.GameStats{
		Text:  fmt.Sprintf(StatsFormat, 0, 0),
		Style: style.Styles.Adventure.Header.Stats,
	}

	instructions := views.Instructions{
		Style: style.Styles.Adventure.Instructions.Style,
	}
	instructions.SetInstructions("Use the hjkl keys to move the player")

	gameView := GameView{
		Field:      [][]rune{},
		Border:     style.Styles.Adventure.Map.Border,
		Background: style.Styles.Adventure.Map.Background,
	}

	return levelInfo, currentMode, stats, instructions, gameView
}

// Render the app views
func renderGameView(gameView GameView, totalWidth int, bordersSpace int) string {
	var lines []string

	// Convert grid content to strings with proper spacing
	for _, row := range gameView.Field {
		var lineRunes []string
		for _, r := range row {
			characterStyle := character.ToDefaultCharacterStyle(r)
			cell := characterStyle.Render(string(r))
			lineRunes = append(lineRunes, cell)
		}
		lines = append(lines, lipgloss.JoinHorizontal(lipgloss.Left, lineRunes...))
	}

	// Join all lines with newlines
	content := lipgloss.JoinVertical(lipgloss.Left, lines...)

	// Apply border style with proper width
	return gameView.Border.Width(totalWidth - bordersSpace).Render(content)
}

// RenderScreen renders the app screen
func RenderScreen(params ScreenParams) string {
	topBorderSpace := style.GetComponentWidth(style.Styles.Adventure.Header.Border)
	widths := views.CalculateTopBarWidths(params.Size.Width, topBorderSpace)
	topBar := views.RenderTopBar(params.LevelInfo, params.GameMode, params.GameStats, widths)

	levelInstructions := views.RenderLevelInstructions(params.Instructions, params.Size.Width)

	gameViewBorderSpace := style.GetComponentWidth(params.GameView.Border)
	gameViewContent := renderGameView(params.GameView, params.Size.Width, gameViewBorderSpace)

	controlsBar := help.New().ShortHelpView(params.Bindings)

	return lipgloss.JoinVertical(
		lipgloss.Center,
		topBar,
		levelInstructions,
		gameViewContent,
		controlsBar,
	)
}
