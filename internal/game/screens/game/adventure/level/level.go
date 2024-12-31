package level

import (
	"fmt"
	"github.com/dasvh/go-learn-vim/internal/game/screens/game/adventure/character"
	"time"
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
	InProgress() bool
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
	InProgress     bool     `json:"in_progress"`
}

// Position represents a 2D Position
type Position struct {
	X int
	Y int
}

// PlayerActionResult represents the result of a player action
type PlayerActionResult struct {
	UpdatedPosition    Position
	Completed          bool
	ValidMove          bool
	InstructionMessage string
}

type LevelZero struct {
	grid           [][]rune
	currentTarget  int
	width          int
	height         int
	player         Position
	targets        []Target
	completed      bool
	restore        bool
	inProgress     bool
	chars          *character.Characters
	movementBlock  bool
	blockEnds      time.Time
	targetBehavior TargetBehavior
}

func NewLevelZero() (Level, int) {
	chars := &character.DefaultCharacters
	return &LevelZero{
		chars:          chars,
		targetBehavior: NewCornerTargets(chars),
	}, levelNumberZero
}

// Init initializes the level with the given dimensions
func (level0 *LevelZero) Init(width, height int) {
	level0.inProgress = true
	level0.setDimensions(width, height)
	level0.resetTargets()
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

	currentTarget := level0.targets[level0.currentTarget]
	// check if player has Reached the Target
	if currentTarget.Position == newPos {
		level0.targets[level0.currentTarget].Reached = true
		if level0.currentTarget == len(level0.targets)-1 {
			level0.completed = true
			level0.inProgress = false
			return PlayerActionResult{
				UpdatedPosition:    level0.GetStartPosition(),
				Completed:          true,
				ValidMove:          true,
				InstructionMessage: "Level completed!",
			}
		}

		level0.currentTarget++
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
	if !level0.restore {
		level0.player = position
	}
	level0.restore = false
	level0.grid[level0.player.Y][level0.player.X] = level0.chars.Player.Cursor.Rune
}

// Render provides the visual representation of the level
func (level0 *LevelZero) Render() [][]rune {
	return level0.grid
}

// GetStartPosition returns the starting player Position
func (level0 *LevelZero) GetStartPosition() Position {
	return Position{level0.width / 2, level0.height / 2}
}

// GetCurrentPosition returns the currentTarget player Position
func (level0 *LevelZero) GetCurrentPosition() Position {
	return level0.player
}

// GetTargets returns the Target positions
func (level0 *LevelZero) GetTargets() []Target {
	return level0.targets
}

// GetCurrentTarget returns the currentTarget Target index
func (level0 *LevelZero) GetCurrentTarget() int {
	return level0.currentTarget
}

// GetInstructions returns the instructions for the currentTarget level
func (level0 *LevelZero) GetInstructions() string {
	return fmt.Sprintf("Target %d/%d: Reach the X using hjkl keys", level0.currentTarget+1, level0.targetBehavior.GetTargetCount())
}

// InProgress returns whether the level is in progress
func (level0 *LevelZero) InProgress() bool {
	return level0.inProgress
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

	level0.restore = true
	level0.setDimensions(state.Width, state.Height)
	level0.player = state.PlayerPosition
	level0.targets = state.Targets
	level0.currentTarget = state.CurrentTarget
	level0.completed = state.Completed
	level0.inProgress = state.InProgress

	level0.replaceTargets(state.Targets)
	level0.initializeGrid()
	return nil
}

func (level0 *LevelZero) initializeGrid() {
	level0.clearGrid()
	level0.PlacePlayer(level0.GetStartPosition())
	level0.targetBehavior.UpdateGrid(level0.grid, level0.targets, level0.currentTarget, level0.chars)
}

// clearGrid clears the grid
func (level0 *LevelZero) clearGrid() {
	for y := range level0.grid {
		for x := range level0.grid[y] {
			level0.grid[y][x] = ' '
		}
	}
}

// setDimensions sets the level dimensions and initializes the grid
func (level0 *LevelZero) setDimensions(width, height int) {
	level0.width = width
	level0.height = height
	level0.grid = make([][]rune, height)
	for y := range level0.grid {
		level0.grid[y] = make([]rune, width)
	}
}

// resetTargets resets the Target positions
func (level0 *LevelZero) resetTargets() {
	level0.targets = level0.targetBehavior.DefineTargets(level0.width, level0.height)
	level0.currentTarget = 0
}

// replaceTargets replaces the currentTarget targets with the saved targets
func (level0 *LevelZero) replaceTargets(savedTargets []Target) {
	// define targets based on currentTarget dimensions
	targets := level0.targetBehavior.DefineTargets(level0.width, level0.height)

	// transfer the Reached state from the saved targets
	for i := range targets {
		if i < len(savedTargets) {
			targets[i].Reached = savedTargets[i].Reached
		}
	}
	level0.targets = targets
	level0.targetBehavior.UpdateGrid(level0.grid, level0.targets, level0.currentTarget, level0.chars)
}
