package level

import (
	"fmt"
	"github.com/dasvh/go-learn-vim/internal/models"
	"time"
)

const levelNumberZero = 0

type Zero struct {
	grid           [][]rune
	currentTarget  int
	width          int
	height         int
	player         models.Position
	targets        []models.Target
	completed      bool
	restore        bool
	inProgress     bool
	chars          *models.Characters
	movementBlock  bool
	blockEnds      time.Time
	targetBehavior models.TargetBehavior
}

func NewLevelZero() (models.Level, int) {
	chars := &models.DefaultCharacters
	return &Zero{
		chars:          chars,
		targetBehavior: NewCornerTargets(chars),
	}, levelNumberZero
}

// Init initializes the level with the given dimensions
func (level0 *Zero) Init(width, height int) {
	level0.inProgress = true
	level0.setDimensions(width, height)
	level0.resetTargets()
	level0.initializeGrid()
}

// UpdatePlayerAction handles player movement and Target completion
func (level0 *Zero) UpdatePlayerAction(delta models.Position) models.PlayerMovement {
	// block movement if cooldown is active
	if level0.movementBlock && time.Now().Before(level0.blockEnds) {
		return models.PlayerMovement{
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

	newPos := models.Position{
		X: level0.player.X + delta.X,
		Y: level0.player.Y + delta.Y,
	}

	// check if player is within bounds
	if newPos.X < 0 || newPos.X >= level0.width || newPos.Y < 0 || newPos.Y >= level0.height {
		return models.PlayerMovement{
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
		if level0.currentTarget == level0.targetBehavior.GetTargetCount()-1 {
			level0.completed = true
			level0.inProgress = false
			return models.PlayerMovement{
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

		return models.PlayerMovement{
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

	return models.PlayerMovement{
		UpdatedPosition:    newPos,
		Completed:          false,
		ValidMove:          true,
		InstructionMessage: level0.GetInstructions(),
	}
}

// PlacePlayer places the player at the given Position
func (level0 *Zero) PlacePlayer(position models.Position) {
	if !level0.restore {
		level0.player = position
	}
	level0.restore = false
	level0.grid[level0.player.Y][level0.player.X] = level0.chars.Player.Cursor.Rune
}

// Render provides the visual representation of the level
func (level0 *Zero) Render() [][]rune {
	return level0.grid
}

// GetStartPosition returns the starting player Position
func (level0 *Zero) GetStartPosition() models.Position {
	return models.Position{X: level0.width / 2, Y: level0.height / 2}
}

// GetCurrentPosition returns the currentTarget player Position
func (level0 *Zero) GetCurrentPosition() models.Position {
	return level0.player
}

// GetTargets returns the Target positions
func (level0 *Zero) GetTargets() []models.Target {
	return level0.targets
}

// GetCurrentTarget returns the currentTarget Target index
func (level0 *Zero) GetCurrentTarget() int {
	return level0.currentTarget
}

// GetInstructions returns the instructions for the currentTarget level
func (level0 *Zero) GetInstructions() string {
	return fmt.Sprintf("Instructions: Target %d/%d: Reach the X using hjkl keys", level0.currentTarget+1, level0.targetBehavior.GetTargetCount())
}

// InProgress returns whether the level is in progress
func (level0 *Zero) InProgress() bool {
	return level0.inProgress
}

// IsCompleted returns whether the level is completed
func (level0 *Zero) IsCompleted() bool {
	return level0.completed
}

// Restore a SavedLevel from a game state
func (level0 *Zero) Restore(state models.SavedLevel) error {
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

func (level0 *Zero) initializeGrid() {
	level0.clearGrid()
	level0.PlacePlayer(level0.GetStartPosition())
	level0.targetBehavior.UpdateGrid(level0.grid, level0.targets, level0.currentTarget, level0.chars)
}

// clearGrid clears the grid
func (level0 *Zero) clearGrid() {
	for y := range level0.grid {
		for x := range level0.grid[y] {
			level0.grid[y][x] = ' '
		}
	}
}

// setDimensions sets the level dimensions and initializes the grid
func (level0 *Zero) setDimensions(width, height int) {
	level0.width = width
	level0.height = height
	level0.grid = make([][]rune, height)
	for y := range level0.grid {
		level0.grid[y] = make([]rune, width)
	}
}

// resetTargets resets the Target positions
func (level0 *Zero) resetTargets() {
	level0.targets = level0.targetBehavior.DefineTargets(level0.width, level0.height)
	level0.currentTarget = 0
}

// replaceTargets replaces the currentTarget targets with the saved targets
func (level0 *Zero) replaceTargets(savedTargets []models.Target) {
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
