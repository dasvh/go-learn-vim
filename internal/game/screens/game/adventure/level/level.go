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
	GetInstructions() string
	IsCompleted() bool
}

// Position represents a 2D position
type Position struct {
	X int
	Y int
}

// target represents a position and whether it has been reached
type target struct {
	position Position
	reached  bool
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
	targets       []target
	completed     bool
	init          bool
	chars         *character.Characters
	movementBlock bool
	blockEnds     time.Time
}

func NewLevelZero() Level {
	return &LevelZero{
		chars: &character.DefaultCharacters,
	}
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
		level0.init = true
	}

	level0.PlacePlayer(level0.GetStartPosition())

	// todo: figure out how to set the characters even after the game is completed
	for i, target := range level0.targets {
		if i == level0.current {
			level0.grid[target.position.Y][target.position.X] = level0.chars.Target.Active.Rune
		} else {
			if target.reached {
				level0.grid[target.position.Y][target.position.X] = level0.chars.Target.Reached.Rune
			} else {
				level0.grid[target.position.Y][target.position.X] = level0.chars.Target.Inactive.Rune
			}
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
	level0.targets = []target{
		{Position{offsetX, offsetY}, false},                          // top left
		{Position{width - offsetX - 1, offsetY}, false},              // top right
		{Position{offsetX, height - offsetY - 1}, false},             // bottom left
		{Position{width - offsetX - 1, height - offsetY - 1}, false}, // bottom right
	}

	level0.current = 0
	level0.initializeGrid()
}

// UpdatePlayerAction handles player movement and target completion
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
	// check if player has reached the target
	if currentTarget.position == newPos {
		level0.targets[level0.current].reached = true
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

		// block movement for 500ms after reaching a target
		level0.movementBlock = true
		level0.blockEnds = time.Now().Add(500 * time.Millisecond)

		return PlayerActionResult{
			UpdatedPosition:    level0.GetStartPosition(),
			Completed:          false,
			ValidMove:          true,
			InstructionMessage: level0.GetInstructions(),
		}
	}

	// update player position
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

// PlacePlayer places the player at the given position
func (level0 *LevelZero) PlacePlayer(position Position) {
	level0.player.X = position.X
	level0.player.Y = position.Y
	level0.grid[position.Y][position.X] = level0.chars.Player.Cursor.Rune
}

// Render provides the visual representation of the level
func (level0 *LevelZero) Render() [][]rune {
	return level0.grid
}

// GetStartPosition returns the starting player position
func (level0 *LevelZero) GetStartPosition() Position {
	return Position{level0.width / 2, level0.height / 2}
}

// GetInstructions returns the instructions for the current level
func (level0 *LevelZero) GetInstructions() string {
	return fmt.Sprintf("Target %d/%d: Reach the X using hjkl keys", level0.current+1, len(level0.targets))
}

// IsCompleted returns whether the level is completed
func (level0 *LevelZero) IsCompleted() bool {
	return level0.completed
}
