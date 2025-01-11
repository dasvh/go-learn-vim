package adventure

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/app/controllers"
	"github.com/dasvh/go-learn-vim/internal/models"
	"github.com/dasvh/go-learn-vim/internal/views"
	"time"
)

// TODO: there is a bug with the player position when the window is resized before starting a game
// 		 need to send a Msg to AdventureModel before starting a game to init the size related state

const gameMode = "Adventure"

// Adventure represents the adventure mode
type Adventure struct {
	controls   Controls
	stats      *models.Stats
	lc         *controllers.Level
	gc         *controllers.Game
	view       views.AdventureView
	gridWidth  int
	gridHeight int
	saveID     string
}

// NewAdventure creates a new Adventure instance
func NewAdventure(gc *controllers.Game, lc *controllers.Level) *Adventure {
	controls := NewBasicControls()
	view := views.InitializeAdventureView()
	view.SetMode(gameMode)
	view.SetStats(0, 0)
	view.Help = controls.BasicHelp()
	return &Adventure{
		controls: controls,
		stats:    models.NewStats(),
		lc:       lc,
		gc:       gc,
		view:     view,
	}
}

// scalePosition scales a models.Position based on the x and y scaling factors
func scalePosition(pos models.Position, xScale, yScale float64) models.Position {
	return models.Position{
		X: int(float64(pos.X) * xScale),
		Y: int(float64(pos.Y) * yScale),
	}
}

func (a *Adventure) initializeLevel() {
	err := a.lc.InitOrResizeLevel(a.gridWidth, a.gridHeight)
	if err != nil {
		fmt.Println("Failed to initialize level:", err)
		return
	}
	a.view.Level.SetText(fmt.Sprintf("Level: %d", a.lc.GetLevelNumber()))
	a.view.Info.SetText(a.lc.GetCurrentLevel().GetInstructions())
}

// Save saves the app models.AdventureGameState with models.Stats
func (a *Adventure) Save() error {
	gameState := models.AdventureGameState{
		WindowSize: a.view.Size,
		Level: models.SavedLevel{
			Number:         a.lc.GetLevelNumber(),
			Width:          a.gridWidth,
			Height:         a.gridHeight,
			PlayerPosition: a.lc.GetCurrentLevel().GetCurrentPosition(),
			Targets:        a.lc.GetCurrentLevel().GetTargets(),
			CurrentTarget:  a.lc.GetCurrentLevel().GetCurrentTarget(),
			Completed:      a.lc.GetCurrentLevel().IsCompleted(),
			InProgress:     a.lc.GetCurrentLevel().InProgress(),
		},
		Stats:  *a.stats,
		SaveID: a.saveID,
	}
	return a.gc.SaveGame(gameMode, gameState, a.saveID)
}

// Load creates a new Adventure instance from a saved models.GameState
func Load(gc *controllers.Game, lc *controllers.Level, gameState models.GameState, size tea.WindowSizeMsg) (*Adventure, error) {
	// Ensure the GameState is of the correct type
	ags, ok := gameState.(models.AdventureGameState)
	if !ok {
		return nil, fmt.Errorf("invalid game state type: expected AdventureGameState")
	}

	controls := NewBasicControls()
	gameView := views.InitializeAdventureView()
	gameView.Size = size

	adventure := &Adventure{
		controls: controls,
		stats:    &ags.Stats,
		lc:       lc,
		gc:       gc,
		view:     gameView,
		saveID:   ags.SaveID,
	}
	// update the grid dimensions
	adventure.gridWidth, adventure.gridHeight = adventure.view.UpdateGridDimensions()

	// calculate scaling factors
	xScale := float64(size.Width) / float64(ags.WindowSize.Width)
	yScale := float64(size.Height) / float64(ags.WindowSize.Height)

	// scale the player position based on the new grid dimensions
	ags.Level.PlayerPosition = scalePosition(ags.Level.PlayerPosition, xScale, yScale)

	// update the Level.Width and Level.Height
	ags.Level.Width = adventure.gridWidth
	ags.Level.Height = adventure.gridHeight

	if err := adventure.lc.RestoreLevel(ags.Level); err != nil {
		return nil, fmt.Errorf("failed to load level gameSave: %w", err)
	}

	adventure.view.SetMode(gameMode)
	adventure.view.SetLevel(adventure.lc.GetLevelNumber())
	adventure.view.SetStats(adventure.stats.TotalKeystrokes, adventure.stats.TimeElapsed)
	adventure.view.SetInfo(adventure.lc.GetCurrentLevel().GetInstructions())

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
	case models.SetPlayerMsg:
		a.view.SetPlayer(msg.Player.Name)
		return a, nil
	case models.SetLevelMsg:
		a.initializeLevel()
	case TickMsg:
		a.stats.IncrementTime()
		a.view.SetStats(a.stats.TotalKeystrokes, a.stats.TimeElapsed)
		return a, tea.Tick(time.Second, func(t time.Time) tea.Msg {
			return TickMsg(t)
		})
	case tea.WindowSizeMsg:
		if msg.Width != a.view.Size.Width || msg.Height != a.view.Size.Height {
			a.view.Size = msg
			a.gridWidth, a.gridHeight = a.view.UpdateGridDimensions()
			if a.lc.GetCurrentLevel() != nil {
				a.initializeLevel()
			}
		}
	case tea.KeyMsg:
		keyString := msg.String()
		isMotionKey := false
		var delta models.Position
		switch {
		case key.Matches(msg, a.controls.MoveLeft):
			isMotionKey = true
			delta = models.Position{X: -1, Y: 0}
		case key.Matches(msg, a.controls.MoveRight):
			isMotionKey = true
			delta = models.Position{X: 1, Y: 0}
		case key.Matches(msg, a.controls.MoveUp):
			isMotionKey = true
			delta = models.Position{X: 0, Y: -1}
		case key.Matches(msg, a.controls.MoveDown):
			isMotionKey = true
			delta = models.Position{X: 0, Y: 1}
		case key.Matches(msg, a.controls.Quit):
			// save the app controllers before quitting
			err := a.Save()
			if err != nil {
				fmt.Println("Failed to save app controllers:", err)
			}
			return a, tea.Quit
		}

		// update player action
		result := a.lc.GetCurrentLevel().PlayerMove(delta)

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
	game := a.lc.GetCurrentLevel().Render()
	a.view.GameMap.Field = game
	return a.view.RenderScreen()
}
