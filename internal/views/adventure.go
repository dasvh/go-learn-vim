package views

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dasvh/go-learn-vim/internal/components"
	"github.com/dasvh/go-learn-vim/internal/style"
)

const (
	StatsFormat = "Stats: Keystrokes: %d, Time: %d s"
)

// GameMap represents the game map of the adventure mode
type GameMap struct {
	Field      [][]rune
	Border     lipgloss.Style
	Background lipgloss.Style
}

// AdventureView represents the adventure mode view
type AdventureView struct {
	Size    tea.WindowSizeMsg
	Level   components.TextDisplay
	Player  components.TextDisplay
	Mode    components.TextDisplay
	Stats   components.TextDisplay
	Info    components.TextDisplay
	GameMap GameMap
	Help    []key.Binding
}

// InitializeAdventureView creates a new instance of AdventureView
func InitializeAdventureView() AdventureView {
	return AdventureView{
		Level: components.TextDisplay{
			Style: style.Styles.Adventure.Header.Level,
		},
		Player: components.TextDisplay{
			Style: style.Styles.Adventure.Header.Level,
		},
		Mode: components.TextDisplay{
			Style: style.Styles.Adventure.Header.Mode,
		},
		Stats: components.TextDisplay{
			Style: style.Styles.Adventure.Header.Stats,
		},
		Info: components.TextDisplay{
			Style: style.Styles.Adventure.Instructions.Style,
		},
		GameMap: GameMap{
			Field:      make([][]rune, 0),
			Border:     style.Styles.Adventure.Map.Border,
			Background: style.Styles.Adventure.Map.Background,
		},
	}
}

// UpdateGridDimensions updates the grid dimensions
func (av *AdventureView) UpdateGridDimensions() (int, int) {
	topHeight := style.GetComponentHeight(style.Styles.Adventure.Header.Border) +
		style.GetComponentHeight(style.Styles.Adventure.Instructions.Style)
	mapBorderHeight := style.GetComponentHeight(av.GameMap.Border)

	gridHeight := av.Size.Height - (topHeight + mapBorderHeight)
	gridWidth := av.Size.Width - style.GetComponentWidth(av.GameMap.Border)

	return gridWidth, gridHeight
}

// RenderScreen renders the adventure mode screen
func (av *AdventureView) RenderScreen() string {
	topBorderWidth := style.GetComponentWidth(style.Styles.Adventure.Header.Border)
	gameMapBorderWidth := style.GetComponentWidth(av.GameMap.Border)

	sections := []components.TextDisplay{av.Level, av.Player, av.Mode, av.Stats}
	widths := calculateSectionWidths(av.Size.Width, topBorderWidth, 4)
	positions := []lipgloss.Position{lipgloss.Left, lipgloss.Top, lipgloss.Center, lipgloss.Right}

	topBar := renderTopBar(sections, widths, positions, style.Styles.Adventure.Header.Border)
	levelInstructions := av.Info.Render()
	gameMap := renderGameMap(av.GameMap, av.Size.Width, gameMapBorderWidth)
	controlsBar := help.New().ShortHelpView(av.Help)

	return lipgloss.JoinVertical(
		lipgloss.Center,
		topBar,
		levelInstructions,
		gameMap,
		controlsBar,
	)
}

// SetLevel sets the level text
func (av *AdventureView) SetLevel(level int) {
	av.Level.SetText("Level: %d", level)
}

// SetPlayer sets the player text
func (av *AdventureView) SetPlayer(player string) {
	av.Player.SetText("Player: %s", player)
}

// SetMode sets the mode text
func (av *AdventureView) SetMode(mode string) {
	av.Mode.SetText("Mode: %s", mode)
}

// SetStats sets the stats text with keystrokes and time
func (av *AdventureView) SetStats(keystrokes int, time int) {
	av.Stats.SetText(StatsFormat, keystrokes, time)
}

// SetInfo sets the info text
func (av *AdventureView) SetInfo(info string) {
	av.Info.SetText(info)
}

// renderGameMap renders the game map
func renderGameMap(gameMap GameMap, totalWidth int, bordersSpace int) string {
	var lines []string

	// Convert grid content to strings with proper spacing
	for _, row := range gameMap.Field {
		var lineRunes []string
		for _, r := range row {
			characterStyle := components.ToDefaultCharacterStyle(r)
			cell := characterStyle.Render(string(r))
			lineRunes = append(lineRunes, cell)
		}
		lines = append(lines, lipgloss.JoinHorizontal(lipgloss.Left, lineRunes...))
	}

	// Join all lines with newlines
	content := lipgloss.JoinVertical(lipgloss.Left, lines...)

	// Apply border style with proper width
	return gameMap.Border.Width(totalWidth - bordersSpace).Render(content)
}

// renderTopBar renders the top bar
func renderTopBar(sections []components.TextDisplay, widths []int, positions []lipgloss.Position, border lipgloss.Style) string {
	if len(sections) != len(widths) || len(sections) != len(positions) {
		panic("sections, widths, and positions must have the same length")
	}

	var renderedSections []string
	for i, section := range sections {
		renderedSections = append(renderedSections,
			lipgloss.NewStyle().
				Width(widths[i]).
				Align(positions[i]).
				Render(section.Render()),
		)
	}

	topBar := lipgloss.JoinHorizontal(lipgloss.Top, renderedSections...)
	return border.Render(topBar)
}

// calculateSectionWidths calculates the width of each section in a top bar
func calculateSectionWidths(totalWidth int, borderSpace int, numSections int) []int {
	if numSections <= 0 {
		return []int{}
	}

	widthMinusBorders := totalWidth - borderSpace
	baseWidth := widthMinusBorders / numSections
	remainingSpace := widthMinusBorders % numSections

	// Initialize widths with the base width
	widths := make([]int, numSections)
	for i := range widths {
		widths[i] = baseWidth
	}

	// Distribute the remaining space among sections
	for i := 0; remainingSpace > 0; i++ {
		widths[i%numSections]++
		remainingSpace--
	}

	return widths
}
