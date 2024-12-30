package adventure

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/components/view"
	"github.com/dasvh/go-learn-vim/internal/game/screens/game"
	"github.com/dasvh/go-learn-vim/internal/game/screens/game/adventure/level"
	"github.com/dasvh/go-learn-vim/internal/storage"
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
	controls         Controls
	stats            *game.Stats
	currentTime      int
	levelManager     *level.Manager
	repo             storage.AdventureGameRepository
}

// NewAdventure creates a new Adventure instance
func NewAdventure(repo storage.AdventureGameRepository) *Adventure {
	controls := NewBasicControls()
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
		currentTime:      0,
		levelManager:     levelManager,
		repo:             repo,
	}
}

// LoadAdventure creates a new Adventure instance from a saved game state
func LoadAdventure(repo storage.AdventureGameRepository, state storage.AdventureGameState, size tea.WindowSizeMsg) (*Adventure, error) {
	controls := NewBasicControls()
	levelInfo, gameMode, statsInfo, levelInstructions, gameView := InitializeComponents()
	levelManager := level.NewManager()

	adventure := &Adventure{
		levelInfo:        levelInfo,
		gameModeInfo:     gameMode,
		gameStats:        statsInfo,
		gameInstructions: levelInstructions,
		gameView:         gameView,
		controls:         controls,
		stats:            &state.Stats,
		levelManager:     levelManager,
		repo:             repo,
		size:             size,
	}

	// Calculate scaling factors
	xScale := float64(size.Width) / float64(state.Size.Width)
	yScale := float64(size.Height) / float64(state.Size.Height)

	// debug info
	debugString := fmt.Sprintf("size: %+v, xScale: %+v yScale: %+v, state.Size: %+v", size, xScale, yScale, state.Size)
	adventure.gameInstructions.SetInstructions(debugString)

	// resize the grid based on the (new) window size
	adventure.gridWidth = size.Width - view.GetComponentWidth(view.Styles.Adventure.Map.Border)
	adventure.gridHeight = size.Height - (view.GetComponentHeight(view.Styles.Adventure.Header.Border) +
		view.GetComponentHeight(view.Styles.Adventure.Instructions.Style) +
		view.GetComponentHeight(view.Styles.Adventure.Map.Border))

	// scale the player position based on the new grid dimensions
	state.Level.PlayerPosition = scalePosition(state.Level.PlayerPosition, xScale, yScale)

	// update the Level.Width and Level.Height
	state.Level.Width = adventure.gridWidth
	state.Level.Height = adventure.gridHeight

	if err := adventure.levelManager.RestoreLevel(state.Level); err != nil {
		return nil, fmt.Errorf("failed to load level state: %w", err)
	}

	adventure.levelInfo.SetLevel(state.Level.Number)
	adventure.gameStats.Text = fmt.Sprintf(StatsFormat, state.Stats.TotalKeystrokes, state.Stats.TimeElapsed)
	//adventure.gameInstructions.SetInstructions(adventure.levelManager.GetCurrentLevel().GetInstructions())

	return adventure, nil
}

// scalePosition scales a position based on the x and y scaling factors
func scalePosition(pos level.Position, xScale, yScale float64) level.Position {
	return level.Position{
		X: int(float64(pos.X) * xScale),
		Y: int(float64(pos.Y) * yScale),
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

func (a *Adventure) initializeLevel() {
	err := a.levelManager.InitOrResizeLevel(a.gridWidth, a.gridHeight)
	if err != nil {
		return
	}
	a.levelInfo.SetLevel(a.levelManager.GetLevelNumber())
	a.gameInstructions.SetInstructions(a.levelManager.GetCurrentLevel().GetInstructions())
}

// Save saves the game state
func (a *Adventure) Save(repo storage.AdventureGameRepository) error {
	state := storage.AdventureGameState{
		Size: a.size,
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
		Stats: game.Stats{
			KeyPresses:      a.stats.KeyPresses,
			TotalKeystrokes: a.stats.TotalKeystrokes,
			TimeElapsed:     a.stats.TimeElapsed,
		},
	}
	return repo.SaveAdventureGame(state)
}

// TickMsg represents a tick message
type TickMsg time.Time

// Init initializes the game
func (a *Adventure) Init() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (a *Adventure) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TickMsg:
		a.stats.IncrementTime()
		a.gameStats.Text = fmt.Sprintf(StatsFormat,
			a.stats.TotalKeystrokes,
			a.stats.TimeElapsed,
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
			// save the game state before quitting
			err := a.Save(a.repo)
			if err != nil {
				fmt.Println("Failed to save game state:", err)
			}
			return a, tea.Quit
		}

		// update player action
		result := a.levelManager.GetCurrentLevel().UpdatePlayerAction(delta)

		// update game instructions
		if result.InstructionMessage != "" {
			a.gameInstructions.SetInstructions(result.InstructionMessage)
		}
		if result.Completed {
			// save the game state after completing the level
			err := a.Save(a.repo)
			if err != nil {
				fmt.Println("Failed to save game state:", err)
			}

			return a, tea.Quit
		}

		// only register keystrokes if it's a motion key and the move is valid
		if isMotionKey && result.ValidMove {
			a.stats.RegisterKey(keyString, true)
		}

		a.gameStats.Text = fmt.Sprintf(StatsFormat, a.stats.TotalKeystrokes, a.stats.TimeElapsed)
		return a, nil
	}
	return a, nil
}

// View renders the entire game screen
func (a *Adventure) View() string {
	// update game view content
	currentView := a.levelManager.GetCurrentLevel().Render()
	a.gameView.Field = currentView

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
