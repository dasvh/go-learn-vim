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

func (a *Adventure) initializeLevel() {
	a.levelManager.InitCurrentLevel(a.gridWidth, a.gridHeight)
	a.gameInstructions.SetInstructions(a.levelManager.GetCurrentLevel().GetInstructions())
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
			return a, tea.Quit
		}

		// update player action
		result := a.levelManager.GetCurrentLevel().UpdatePlayerAction(delta)

		// update game instructions
		if result.InstructionMessage != "" {
			a.gameInstructions.SetInstructions(result.InstructionMessage)
		}
		if result.Completed {
			return a, tea.Quit
		}

		// only register keystrokes if it's a motion key and the move is valid
		if isMotionKey && result.ValidMove {
			a.stats.RegisterKey(keyString, true)
		}

		a.gameStats.Text = fmt.Sprintf(StatsFormat, a.stats.TotalKeystrokes, a.stats.Time)
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
