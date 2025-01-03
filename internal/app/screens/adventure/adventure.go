package adventure

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/app/controllers"
	"github.com/dasvh/go-learn-vim/internal/app/screens"
	"github.com/dasvh/go-learn-vim/internal/app/screens/adventure/level"
	"github.com/dasvh/go-learn-vim/internal/storage"
	"github.com/dasvh/go-learn-vim/internal/views"
	"time"
)

const gameMode = "Adventure"

// Adventure represents the adventure mode
type Adventure struct {
	controls     Controls
	stats        *storage.Stats
	levelManager *level.Manager
	gc           *controllers.Game
	view         views.AdventureView
	gridWidth    int
	gridHeight   int
}

// NewAdventure creates a new Adventure instance
func NewAdventure(gc *controllers.Game) *Adventure {
	controls := NewBasicControls()
	levelManager := level.NewManager()
	view := views.InitializeAdventureView()
	view.SetMode(gameMode)
	view.SetLevel(levelManager.GetLevelNumber())
	view.SetStats(0, 0)
	view.SetInfo(levelManager.GetCurrentLevel().GetInstructions())
	view.Help = controls.BasicHelp()

	return &Adventure{
		controls:     controls,
		stats:        storage.NewStats(),
		levelManager: levelManager,
		gc:           gc,
		view:         view,
	}
}

// scalePosition scales a position based on the x and y scaling factors
func scalePosition(pos level.Position, xScale, yScale float64) level.Position {
	return level.Position{
		X: int(float64(pos.X) * xScale),
		Y: int(float64(pos.Y) * yScale),
	}
}

func (a *Adventure) initializeLevel() {
	err := a.levelManager.InitOrResizeLevel(a.gridWidth, a.gridHeight)
	if err != nil {
		return
	}
	a.view.Level.SetText(fmt.Sprintf("Level: %d", a.levelManager.GetLevelNumber()))
	a.view.Info.SetText(a.levelManager.GetCurrentLevel().GetInstructions())
}

// Save saves the app gameState
func (a *Adventure) Save() error {
	gameState := storage.AdventureGameState{
		WindowSize: a.view.Size,
		Level: level.SavedLevel{
			Number:         a.levelManager.GetLevelNumber(),
			Width:          a.gridWidth,
			Height:         a.gridHeight,
			PlayerPosition: a.levelManager.GetCurrentLevel().GetCurrentPosition(),
			Targets:        a.levelManager.GetCurrentLevel().GetTargets(),
			CurrentTarget:  a.levelManager.GetCurrentLevel().GetCurrentTarget(),
			Completed:      a.levelManager.GetCurrentLevel().IsCompleted(),
			InProgress:     a.levelManager.GetCurrentLevel().InProgress(),
		},
		Stats: storage.Stats{
			KeyPresses:      a.stats.KeyPresses,
			TotalKeystrokes: a.stats.TotalKeystrokes,
			TimeElapsed:     a.stats.TimeElapsed,
		},
	}
	return a.gc.SaveGame(gameMode, gameState)
}

// Load creates a new Adventure instance from a saved app controllers
func Load(gc *controllers.Game, gameState storage.GameState, size tea.WindowSizeMsg) (*Adventure, error) {
	// Ensure the GameState is of the correct type
	ags, ok := gameState.(storage.AdventureGameState)
	if !ok {
		return nil, fmt.Errorf("invalid game controllers type: expected AdventureGameState")
	}

	controls := NewBasicControls()
	levelManager := level.NewManager()

	gameView := views.InitializeAdventureView()
	gameView.Size = size
	gameView.SetMode(gameMode)
	gameView.SetLevel(ags.Level.Number)
	gameView.SetPlayer("Player")
	gameView.SetStats(ags.Stats.TotalKeystrokes, ags.Stats.TimeElapsed)
	gameView.SetInfo(levelManager.GetCurrentLevel().GetInstructions())
	gameView.Help = controls.BasicHelp()

	adventure := &Adventure{
		controls:     controls,
		stats:        &ags.Stats,
		levelManager: levelManager,
		gc:           gc,
		view:         gameView,
	}

	// update the grid dimensions
	adventure.gridWidth, adventure.gridHeight = adventure.view.UpdateGridDimensions()

	// calculate scaling factors
	xScale := float64(size.Width) / float64(ags.WindowSize.Width)
	yScale := float64(size.Height) / float64(ags.WindowSize.Height)

	// debug info
	debugString := fmt.Sprintf("size: %+v, xScale: %+v yScale: %+v, gameSave.size: %+v", size, xScale, yScale, ags.WindowSize)
	adventure.view.SetInfo(debugString)

	// scale the player position based on the new grid dimensions
	ags.Level.PlayerPosition = scalePosition(ags.Level.PlayerPosition, xScale, yScale)

	// update the Level.Width and Level.Height
	ags.Level.Width = adventure.gridWidth
	ags.Level.Height = adventure.gridHeight

	if err := adventure.levelManager.RestoreLevel(ags.Level); err != nil {
		return nil, fmt.Errorf("failed to load level gameSave: %w", err)
	}

	adventure.view.SetLevel(adventure.levelManager.GetLevelNumber())
	adventure.view.SetStats(adventure.stats.TotalKeystrokes, adventure.stats.TimeElapsed)
	//adventure.view.SetInfo(adventure.levelManager.GetCurrentLevel().GetInstructions())

	return adventure, nil
}

// TickMsg represents a tick message
type TickMsg time.Time

// Init initializes the app
func (a *Adventure) Init() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (a *Adventure) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case screens.SetPlayerMsg:
		a.view.SetPlayer(msg.Player.Name)
		return a, nil
	case TickMsg:
		a.stats.IncrementTime()
		a.view.SetStats(a.stats.TotalKeystrokes, a.stats.TimeElapsed)
		return a, tea.Tick(time.Second, func(t time.Time) tea.Msg {
			return TickMsg(t)
		})
	case tea.WindowSizeMsg:
		if msg.Width > 0 && msg.Height > 0 {
			a.view.Size = msg
			a.gridWidth, a.gridHeight = a.view.UpdateGridDimensions()
			a.initializeLevel()
		}
	case tea.KeyMsg:
		keyString := msg.String()
		isMotionKey := false
		var delta level.Position
		switch {
		case key.Matches(msg, a.controls.MoveLeft):
			isMotionKey = true
			delta = level.Position{X: -1, Y: 0}
		case key.Matches(msg, a.controls.MoveRight):
			isMotionKey = true
			delta = level.Position{X: 1, Y: 0}
		case key.Matches(msg, a.controls.MoveUp):
			isMotionKey = true
			delta = level.Position{X: 0, Y: -1}
		case key.Matches(msg, a.controls.MoveDown):
			isMotionKey = true
			delta = level.Position{X: 0, Y: 1}
		case key.Matches(msg, a.controls.Quit):
			// save the app controllers before quitting
			err := a.Save()
			if err != nil {
				fmt.Println("Failed to save app controllers:", err)
			}
			return a, tea.Quit
		}

		// update player action
		result := a.levelManager.GetCurrentLevel().UpdatePlayerAction(delta)

		// update app instructions
		if result.InstructionMessage != "" {
			a.view.SetInfo(result.InstructionMessage)
		}
		if result.Completed {
			// save the app controllers after completing the level
			err := a.Save()
			if err != nil {
				fmt.Println("Failed to save app controllers:", err)
			}

			return a, tea.Quit
		}

		// only register keystrokes if it's a motion key and the move is valid
		if isMotionKey && result.ValidMove {
			a.stats.RegisterKey(keyString, true)
		}

		a.view.SetStats(a.stats.TotalKeystrokes, a.stats.TimeElapsed)
		return a, nil
	}
	return a, nil
}

// View renders the entire app screen
func (a *Adventure) View() string {
	// render the game
	game := a.levelManager.GetCurrentLevel().Render()
	a.view.GameMap.Field = game
	return a.view.RenderScreen()
}
