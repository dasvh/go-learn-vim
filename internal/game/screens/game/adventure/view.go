package adventure

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
	"github.com/dasvh/go-learn-vim/internal/components/view"
	"github.com/dasvh/go-learn-vim/internal/game/screens/game/adventure/character"
)

const (
	StatsFormat = "Stats: Keystrokes: %d, Time: %d s"
)

// GameView contains the components to render the game view
type GameView struct {
	Field      [][]rune
	Border     lipgloss.Style
	Background lipgloss.Style
}

// InitializeComponents initializes the components for the game view
func InitializeComponents() (view.LevelInfo, view.GameModeInfo, view.GameStats, view.Instructions, GameView) {
	levelInfo := view.LevelInfo{
		Style: view.Styles.Adventure.Header.Level,
	}
	levelInfo.SetLevel(0)

	currentMode := view.GameModeInfo{
		Style: view.Styles.Adventure.Header.Mode,
	}
	currentMode.SetMode("Adventure")

	stats := view.GameStats{
		Text:  fmt.Sprintf(StatsFormat, 0, 0),
		Style: view.Styles.Adventure.Header.Stats,
	}

	instructions := view.Instructions{
		Style: view.Styles.Adventure.Instructions.Style,
	}
	instructions.SetInstructions("Use the hjkl keys to move the player")

	gameView := GameView{
		Field:      [][]rune{},
		Border:     view.Styles.Adventure.Map.Border,
		Background: view.Styles.Adventure.Map.Background,
	}

	return levelInfo, currentMode, stats, instructions, gameView
}

// Render the game view
func renderGameView(gameView GameView, totalWidth int, bordersSpace int) string {
	var lines []string

	// Convert grid content to strings with proper spacing
	for _, row := range gameView.Field {
		var lineRunes []string
		for _, r := range row {
			style := character.ToDefaultCharacterStyle(r)
			cell := style.Render(string(r))
			lineRunes = append(lineRunes, cell)
		}
		lines = append(lines, lipgloss.JoinHorizontal(lipgloss.Left, lineRunes...))
	}

	// Join all lines with newlines
	content := lipgloss.JoinVertical(lipgloss.Left, lines...)

	// Apply border style with proper width
	return gameView.Border.Width(totalWidth - bordersSpace).Render(content)
}

// RenderScreen renders the game screen
func RenderScreen(params ScreenParams) string {
	topBorderSpace := view.GetComponentWidth(view.Styles.Adventure.Header.Border)
	widths := view.CalculateTopBarWidths(params.Size.Width, topBorderSpace)
	topBar := view.RenderTopBar(params.LevelInfo, params.GameMode, params.GameStats, widths)

	levelInstructions := view.RenderLevelInstructions(params.Instructions, params.Size.Width)

	gameViewBorderSpace := view.GetComponentWidth(params.GameView.Border)
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
