package level

import (
	"fmt"
	"github.com/dasvh/go-learn-vim/internal/models"
	"time"
)

const levelNumberOne = 1

// One represents level one of the adventure mode
type One struct {
	width          int
	height         int
	currentMaze    int
	totalMazes     int
	seeds          []int64
	mazes          []*Maze
	completed      bool
	restore        bool
	inProgress     bool
	movementBlock  bool
	grid           [][]rune
	chars          *models.Characters
	player         models.Position
	targets        []models.Target
	targetBehavior []*MazeTargets
	blockEnds      time.Time
}

// NewLevelOne returns a new instance of models.Level one
func NewLevelOne() models.Level {
	chars := &models.DefaultCharacters
	return &One{
		chars:          chars,
		totalMazes:     2,
		seeds:          []int64{42, 69},
		targetBehavior: []*MazeTargets{},
	}
}

// Number returns the number of the level
func (level1 *One) Number() int {
	return levelNumberOne
}

// Description returns the description of the level
func (level1 *One) Description() string {
	return "Navigate two mazes using hjkl keys"
}

// Init initializes the level with the given dimensions
func (level1 *One) Init(width, height int) {
	level1.inProgress = true
	level1.currentMaze = 0

	maxMazeSize := min(width/2, height)

	if maxMazeSize < 5 {
		panic(fmt.Sprintf("Grid too small for mazes: Width=%d, Height=%d", width, height))
	}

	mazeSize := maxMazeSize

	// offsets
	paddingBetweenMazes := 3
	maze1OffsetX := (width - 2*mazeSize - paddingBetweenMazes) / 2
	maze2OffsetX := maze1OffsetX + mazeSize + paddingBetweenMazes
	centerY := (height - mazeSize) / 2

	level1.mazes = []*Maze{
		NewMaze(mazeSize, level1.seeds[0], maze1OffsetX, centerY, 3),
		NewMaze(mazeSize, level1.seeds[1], maze2OffsetX, centerY, 2),
	}
	level1.targetBehavior = []*MazeTargets{
		NewMazeTargets(level1.chars, level1.mazes[0]),
		NewMazeTargets(level1.chars, level1.mazes[1]),
	}

	level1.width = width
	level1.height = height
	level1.setDimensions(width, height)
	level1.resetTargets()
	level1.initializeGrid()
}

// PlayerMove handles models.PlayerMovement and transitions between mazes
func (level1 *One) PlayerMove(delta models.Position) models.PlayerMovement {
	newPos := models.Position{
		X: level1.player.X + delta.X,
		Y: level1.player.Y + delta.Y,
	}

	// check if the player collides with a wall
	for _, wall := range level1.targetBehavior[level1.currentMaze].maze.GetWalls() {
		if wall == newPos {
			return models.PlayerMovement{
				UpdatedPosition:    level1.player,
				Completed:          false,
				ValidMove:          false,
				InstructionMessage: level1.GetInstructions(),
			}
		}
	}

	// check if the player has reached the target
	if level1.targets[0].Position == newPos {
		level1.targets[0].Reached = true
		if level1.currentMaze < level1.totalMazes-1 {
			level1.currentMaze++
			level1.resetTargets()
			level1.initializeGrid()
			return models.PlayerMovement{
				UpdatedPosition:    level1.GetStartPosition(),
				Completed:          false,
				ValidMove:          true,
				InstructionMessage: "Maze completed! Moving to the next maze...",
			}
		} else {
			level1.completed = true
			level1.inProgress = false
			return models.PlayerMovement{
				UpdatedPosition:    level1.GetStartPosition(),
				Completed:          true,
				ValidMove:          true,
				InstructionMessage: "All mazes completed! Level finished!",
			}
		}
	}

	// update the player's position
	level1.grid[level1.player.Y][level1.player.X] = level1.chars.Player.Trail.Rune
	level1.player = newPos
	level1.grid[newPos.Y][newPos.X] = level1.chars.Player.Cursor.Rune

	return models.PlayerMovement{
		UpdatedPosition:    newPos,
		Completed:          false,
		ValidMove:          true,
		InstructionMessage: level1.GetInstructions(),
	}
}

// PlacePlayer places the player at the given models.Position
func (level1 *One) PlacePlayer(position models.Position) {
	if !level1.restore {
		level1.player = position
	}
	level1.restore = false
	level1.grid[level1.player.Y][level1.player.X] = level1.chars.Player.Cursor.Rune
}

// Render returns the grid to render
func (level1 *One) Render() [][]rune { return level1.grid }

// GetStartPosition returns the starting player models.Position
func (level1 *One) GetStartPosition() models.Position {
	// todo: maybe StartPosition is a concern of the targetBehavior
	return level1.mazes[level1.currentMaze].StartPosition
}

// GetCurrentPosition returns the current player models.Position
func (level1 *One) GetCurrentPosition() models.Position {
	return level1.player
}

// GetTargets returns the models.Target's for the level
func (level1 *One) GetTargets() []models.Target {
	return level1.targets
}

// GetCurrentTarget returns the currentTarget
func (level1 *One) GetCurrentTarget() int {
	return level1.currentMaze
}

// GetInstructions returns the instructions for the level
func (level1 *One) GetInstructions() string {
	return fmt.Sprintf("Instructions: Maze %d/%d: Reach the X using hjkl keys", level1.currentMaze+1, level1.totalMazes)
}

// InProgress returns whether the level is in progress
func (level1 *One) InProgress() bool {
	return level1.inProgress
}

// IsCompleted returns whether the level is completed
func (level1 *One) IsCompleted() bool {
	return level1.completed
}

// Restore a models.SavedLevel from a game state
func (level1 *One) Restore(state models.SavedLevel) error {
	if state.Width <= 0 || state.Height <= 0 {
		return fmt.Errorf("invalid dimensions in save state")
	}

	level1.restore = true
	return nil
}

// initializeGrid updates the grid for the current maze
func (level1 *One) initializeGrid() {
	level1.clearGrid()
	for _, maze := range level1.mazes {
		for _, wall := range maze.GetWalls() {
			if wall.X >= 0 && wall.X < level1.width && wall.Y >= 0 && wall.Y < level1.height {
				level1.grid[wall.Y][wall.X] = models.DefaultCharacters.Wall.Rune
			}
		}
	}

	level1.PlacePlayer(level1.GetStartPosition())
	level1.targetBehavior[level1.currentMaze].UpdateGrid(level1.grid, level1.targets, 0, level1.chars)
}

// clearGrid clears the grid
func (level1 *One) clearGrid() {
	for y := range level1.grid {
		for x := range level1.grid[y] {
			level1.grid[y][x] = ' '
		}
	}
}

// setDimensions sets the level dimensions and initializes the grid
func (level1 *One) setDimensions(width, height int) {
	level1.width = width
	level1.height = height
	level1.grid = make([][]rune, height)
	for y := range level1.grid {
		level1.grid[y] = make([]rune, width)
	}
}

// resetTargets sets up the targets for the current maze
func (level1 *One) resetTargets() {
	level1.targets = level1.targetBehavior[level1.currentMaze].DefineTargets()
}
