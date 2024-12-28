package level

import (
	"fmt"
	"time"

	"github.com/dasvh/go-learn-vim/internal/game/screens/game/adventure/character"
)

// Level represents a game level
type Level interface {
	Init(width, height int)
	UpdatePlayerAction(position Position) PlayerActionResult
	PlacePlayer(position Position)
	Render() [][]rune
	GetStartPosition() Position
	GetCurrentPosition() Position
	GetTargets() []Target
	GetCurrentTarget() int
	GetInstructions() string
	IsCompleted() bool
	Restore(state SavedLevel) error
}

const levelNumberZero = 0

// SavedLevel represents the state of a saved level
type SavedLevel struct {
	Number         int      `json:"number"`
	Width          int      `json:"width"`
	Height         int      `json:"height"`
	PlayerPosition Position `json:"player_position"`
	Targets        []Target `json:"targets"`
	CurrentTarget  int      `json:"current_target"`
	Completed      bool     `json:"completed"`
}

// Position represents a 2D Position
type Position struct {
	X int
	Y int
}

// Target represents a Position and whether it has been Reached
type Target struct {
	Position Position
	Reached  bool
}

// PlayerActionResult represents the result of a player action
type PlayerActionResult struct {
	UpdatedPosition    Position
	Completed          bool
	ValidMove          bool
	InstructionMessage string
}

type LevelZero struct {
	grid          [][]rune
	current       int
	width         int
	height        int
	player        Position
	targets       []Target
	completed     bool
	init          bool
	chars         *character.Characters
	movementBlock bool
	blockEnds     time.Time
}

func NewLevelZero() (Level, int) {
	return &LevelZero{
		chars: &character.DefaultCharacters,
	}, levelNumberZero
}

func (level0 *LevelZero) initializeGrid() {
	// clear grid on init
	if !level0.init {
		level0.grid = make([][]rune, level0.height)
		for y := 0; y < level0.height; y++ {
			level0.grid[y] = make([]rune, level0.width)
			for x := 0; x < level0.width; x++ {
				level0.grid[y][x] = ' '
			}
		}
		// restore player position from save state
		if level0.player != (Position{}) {
			level0.PlacePlayer(level0.player)
		} else {
			level0.PlacePlayer(level0.GetStartPosition())
		}
		level0.init = true
		level0.updateTargets()
		return
	}

	level0.PlacePlayer(level0.GetStartPosition())
	level0.updateTargets()
}

// updateTargets updates the Target positions on the grid
func (level0 *LevelZero) updateTargets() {
	for i, target := range level0.targets {
		switch {
		case i == level0.current:
			level0.grid[target.Position.Y][target.Position.X] = level0.chars.Target.Active.Rune
		case target.Reached:
			level0.grid[target.Position.Y][target.Position.X] = level0.chars.Target.Reached.Rune
		default:
			level0.grid[target.Position.Y][target.Position.X] = level0.chars.Target.Inactive.Rune
		}
	}
}

// Init initializes the level with the given dimensions
func (level0 *LevelZero) Init(width, height int) {
	level0.width = width
	level0.height = height

	// offset targets from the border by 20%
	offsetX := int(float64(width) * 0.2)
	offsetY := int(float64(height) * 0.2)

	// define targets
	level0.targets = []Target{
		{Position{offsetX, offsetY}, false},                          // top left
		{Position{width - offsetX - 1, offsetY}, false},              // top right
		{Position{offsetX, height - offsetY - 1}, false},             // bottom left
		{Position{width - offsetX - 1, height - offsetY - 1}, false}, // bottom right
	}

	level0.current = 0
	level0.initializeGrid()
}

// UpdatePlayerAction handles player movement and Target completion
func (level0 *LevelZero) UpdatePlayerAction(delta Position) PlayerActionResult {
	// block movement if cooldown is active
	if level0.movementBlock && time.Now().Before(level0.blockEnds) {
		return PlayerActionResult{
			UpdatedPosition:    level0.player,
			Completed:          level0.completed,
			ValidMove:          false,
			InstructionMessage: level0.GetInstructions(),
		}
	}

	// unblock movement if cooldown is over
	if level0.movementBlock {
		level0.movementBlock = false
	}

	newPos := Position{
		X: level0.player.X + delta.X,
		Y: level0.player.Y + delta.Y,
	}

	// check if player is within bounds
	if newPos.X < 0 || newPos.X >= level0.width || newPos.Y < 0 || newPos.Y >= level0.height {
		return PlayerActionResult{
			UpdatedPosition:    level0.player,
			Completed:          level0.completed,
			ValidMove:          false,
			InstructionMessage: level0.GetInstructions(),
		}
	}

	currentTarget := level0.targets[level0.current]
	// check if player has Reached the Target
	if currentTarget.Position == newPos {
		level0.targets[level0.current].Reached = true
		if level0.current == len(level0.targets)-1 {
			level0.completed = true
			return PlayerActionResult{
				UpdatedPosition:    level0.GetStartPosition(),
				Completed:          true,
				ValidMove:          true,
				InstructionMessage: "Level completed!",
			}
		}

		level0.current++
		level0.initializeGrid()

		// block movement for 500ms after reaching a Target
		level0.movementBlock = true
		level0.blockEnds = time.Now().Add(500 * time.Millisecond)

		return PlayerActionResult{
			UpdatedPosition:    level0.GetStartPosition(),
			Completed:          false,
			ValidMove:          true,
			InstructionMessage: level0.GetInstructions(),
		}
	}

	// update player Position
	level0.grid[level0.player.Y][level0.player.X] = level0.chars.Player.Trail.Rune
	level0.grid[newPos.Y][newPos.X] = level0.chars.Player.Cursor.Rune
	level0.player = newPos

	return PlayerActionResult{
		UpdatedPosition:    newPos,
		Completed:          false,
		ValidMove:          true,
		InstructionMessage: level0.GetInstructions(),
	}
}

// PlacePlayer places the player at the given Position
func (level0 *LevelZero) PlacePlayer(position Position) {
	level0.player.X = position.X
	level0.player.Y = position.Y
	level0.grid[position.Y][position.X] = level0.chars.Player.Cursor.Rune
}

// Render provides the visual representation of the level
func (level0 *LevelZero) Render() [][]rune {
	return level0.grid
}

// GetStartPosition returns the starting player Position
func (level0 *LevelZero) GetStartPosition() Position {
	return Position{level0.width / 2, level0.height / 2}
}

// GetCurrentPosition returns the current player Position
func (level0 *LevelZero) GetCurrentPosition() Position {
	return level0.player
}

// GetTargets returns the Target positions
func (level0 *LevelZero) GetTargets() []Target {
	return level0.targets
}

// GetCurrentTarget returns the current Target index
func (level0 *LevelZero) GetCurrentTarget() int {
	return level0.current
}

// GetInstructions returns the instructions for the current level
func (level0 *LevelZero) GetInstructions() string {
	return fmt.Sprintf("Target %d/%d: Reach the X using hjkl keys", level0.current+1, len(level0.targets))
}

// IsCompleted returns whether the level is completed
func (level0 *LevelZero) IsCompleted() bool {
	return level0.completed
}

// Restore restores the level state from a saved state
func (level0 *LevelZero) Restore(state SavedLevel) error {
	if state.Width <= 0 || state.Height <= 0 {
		return fmt.Errorf("invalid dimensions in save state")
	}

	level0.width = state.Width
	level0.height = state.Height
	level0.player = state.PlayerPosition
	level0.targets = state.Targets
	level0.current = state.CurrentTarget
	level0.completed = state.Completed
	level0.init = false

	level0.initializeGrid()

	return nil
}
