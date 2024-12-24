package adventure

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/components/view"
	"github.com/dasvh/go-learn-vim/internal/game/screens/game"
	"github.com/dasvh/go-learn-vim/internal/game/screens/game/adventure/level"
	"time"
)

// Adventure represents the state of the adventure game
type Adventure struct {
	levelInfo        view.LevelInfo
	gameModeInfo     view.GameModeInfo
	gameStats        view.GameStats
	gameInstructions view.Instructions
	gameView         GameView
	size             tea.WindowSizeMsg
	gridWidth        int
	gridHeight       int
	playerX          int
	playerY          int
	grid             [][]rune
	controls         Controls
	stats            *game.Stats
	currentTime      int
	levelManager     *level.Manager
}

// NewAdventure creates a new Adventure instance
func NewAdventure() *Adventure {
	controls := NewBasicControls() // Using basic Controls for now
	levelInfo, gameMode, statsInfo, levelInstructions, gameView := InitializeComponents()
	levelManager := level.NewManager()
	return &Adventure{
		levelInfo:        levelInfo,
		gameModeInfo:     gameMode,
		gameStats:        statsInfo,
		gameInstructions: levelInstructions,
		gameView:         gameView,
		controls:         controls,
		stats:            game.NewStats(),
		gridWidth:        60,
		gridHeight:       30,
		currentTime:      0,
		levelManager:     levelManager,
	}
}

// ScreenParams represents the parameters to render the screen
type ScreenParams struct {
	Size         tea.WindowSizeMsg
	LevelInfo    view.LevelInfo
	GameMode     view.GameModeInfo
	GameStats    view.GameStats
	Instructions view.Instructions
	GameView     GameView
	Bindings     []key.Binding
}

// initializeGrid sets up the grid and places the player in the center
// func (a *Adventure) initializeGrid() {
// 	newGrid := make([][]rune, a.gridHeight)
// 	for y := 0; y < a.gridHeight; y++ {
// 		newGrid[y] = make([]rune, a.gridWidth)
// 	}

//		a.playerX = a.gridWidth / 2
//		a.playerY = a.gridHeight / 2
//		newGrid[a.playerY][a.playerX] = a.gameView.PlayerCharacter.rune
//		a.grid = newGrid
//	}
func (a *Adventure) initializeGrid() {
	// Calculate grid dimensions based on window size
	if a.size.Width > 0 && a.size.Height > 0 {
		a.gridWidth = a.size.Width - view.GetComponentWidth(view.Styles.Adventure.Map.Border)
		a.gridHeight = a.size.Height - (view.GetComponentHeight(view.Styles.Adventure.Header.Border) +
			view.GetComponentHeight(view.Styles.Adventure.Instructions.Style) +
			view.GetComponentHeight(view.Styles.Adventure.Map.Border))

		// Initialize level with current dimensions
		a.levelManager.InitCurrentLevel(a.gridWidth, a.gridHeight)

		// Get fresh grid from level
		a.grid = a.levelManager.GetCurrentLevel().GetGrid()

		// Get starting position and place player
		a.playerX, a.playerY = a.levelManager.GetCurrentLevel().GetStartPosition()
		a.grid[a.playerY][a.playerX] = a.gameView.PlayerCharacter.rune

		// Update game view with initial grid
		a.gameView.Field = a.grid

		// Update target character for rendering
		a.gameView.TargetCharacter = Character{
			rune:   a.levelManager.GetCurrentLevel().GetTargetCharacter(),
			string: string(a.levelManager.GetCurrentLevel().GetTargetCharacter()),
		}

		// Update level info display
		a.levelInfo.SetLevel(a.levelManager.GetLevelNumber())

		// Update instructions
		a.gameInstructions.SetInstructions(a.levelManager.GetCurrentLevel().GetInstructions())
	}
}

// movePlayer moves the player and leaves a trail
func (a *Adventure) movePlayer(dx, dy int) tea.Cmd {
	newX := a.playerX + dx
	newY := a.playerY + dy

	if newX >= 0 && newX < a.gridWidth && newY >= 0 && newY < a.gridHeight {
		// create a trail
		a.grid[a.playerY][a.playerX] = a.gameView.TrailCharacter.rune

		// update player
		a.playerX = newX
		a.playerY = newY

		// check if target reached before updating player position in grid
		if a.levelManager.GetCurrentLevel().Update(a.playerX, a.playerY) {
			if a.levelManager.GetCurrentLevel().IsCompleted() {
				return tea.Quit
			}
			// get grid with target
			a.grid = a.levelManager.GetCurrentLevel().GetGrid()
			// set player at new position
			a.playerX, a.playerY = a.levelManager.GetCurrentLevel().GetStartPosition()
		}

		// position player
		a.grid[a.playerY][a.playerX] = a.gameView.PlayerCharacter.rune

		// update gameView with current grid
		a.gameView.Field = a.grid
	}

	return nil
}

// TickMsg represents a tick message
type TickMsg time.Time

// Init initializes the game
func (a *Adventure) Init() tea.Cmd {
	a.initializeGrid()
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

// Update updates the game state
func (a *Adventure) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TickMsg:
		a.stats.IncrementTime()
		a.gameStats.Text = fmt.Sprintf(StatsFormat,
			a.stats.TotalKeystrokes,
			a.stats.Time,
		)
		return a, tea.Tick(time.Second, func(t time.Time) tea.Msg {
			return TickMsg(t)
		})
	case tea.WindowSizeMsg:
		if msg.Width > 0 && msg.Height > 0 {
			a.size = msg
			a.gridWidth = msg.Width - view.GetComponentWidth(view.Styles.Adventure.Map.Border)
			a.gridHeight = msg.Height - (view.GetComponentHeight(view.Styles.Adventure.Header.Border) +
				view.GetComponentHeight(view.Styles.Adventure.Instructions.Style) +
				view.GetComponentHeight(view.Styles.Adventure.Map.Border))

			a.initializeGrid()
		}
	case tea.KeyMsg:
		keyString := msg.String()
		isMotionKey := false
		var cmd tea.Cmd
		switch {
		case key.Matches(msg, a.controls.MoveLeft):
			isMotionKey = true
			cmd = a.movePlayer(-1, 0)
		case key.Matches(msg, a.controls.MoveRight):
			isMotionKey = true
			cmd = a.movePlayer(1, 0)
		case key.Matches(msg, a.controls.MoveUp):
			isMotionKey = true
			cmd = a.movePlayer(0, -1)
		case key.Matches(msg, a.controls.MoveDown):
			isMotionKey = true
			cmd = a.movePlayer(0, 1)
		case key.Matches(msg, a.controls.Quit):
			return a, tea.Quit
		}

		// only register motion keys
		a.stats.RegisterKey(keyString, isMotionKey)
		a.gameStats.Text = fmt.Sprintf(StatsFormat, a.stats.TotalKeystrokes, a.stats.Time)

		return a, cmd
	}
	return a, nil
}

// View renders the entire game screen
func (a *Adventure) View() string {
	// update game view content
	a.gameView.Field = a.grid

	return RenderScreen(ScreenParams{
		Size:         a.size,
		LevelInfo:    a.levelInfo,
		GameMode:     a.gameModeInfo,
		GameStats:    a.gameStats,
		Instructions: a.gameInstructions,
		GameView:     a.gameView,
		Bindings:     a.controls.BasicHelp(),
	})
}
