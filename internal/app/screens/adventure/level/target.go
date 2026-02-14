package level

import (
	"slices"

	"github.com/dasvh/go-learn-vim/internal/models"
)

// CornerTargets defines the behavior of a target that is placed in the corners of the screen
type CornerTargets struct {
	chars *models.Characters
}

// NewCornerTargets creates a new CornerTargets
func NewCornerTargets(chars *models.Characters) models.TargetBehavior {
	return &CornerTargets{chars: chars}
}

// DefineTargets defines the targets in the corners of the screen
func (ct *CornerTargets) DefineTargets(width, height int) []models.Target {
	offsetX := int(float64(width) * 0.2)
	offsetY := int(float64(height) * 0.2)

	return []models.Target{
		{Position: models.Position{X: offsetX, Y: offsetY}, Reached: false},
		{Position: models.Position{X: width - offsetX - 1, Y: offsetY}, Reached: false},
		{Position: models.Position{X: offsetX, Y: height - offsetY - 1}, Reached: false},
		{Position: models.Position{X: width - offsetX - 1, Y: height - offsetY - 1}, Reached: false},
	}
}

// UpdateGrid updates the grid with the targets
func (ct *CornerTargets) UpdateGrid(grid [][]rune, targets []models.Target, current int, chars *models.Characters) {
	for i, target := range targets {
		switch {
		case i == current:
			grid[target.Position.Y][target.Position.X] = chars.Target.Active.Rune
		case target.Reached:
			grid[target.Position.Y][target.Position.X] = chars.Target.Reached.Rune
		default:
			grid[target.Position.Y][target.Position.X] = chars.Target.Inactive.Rune
		}
	}
}

// GetTargetCount returns the number of targets
func (ct *CornerTargets) GetTargetCount() int {
	return 4
}

// MazeTargets integrates Maze for walls and manages targets
type MazeTargets struct {
	chars *models.Characters
	maze  *Maze
}

// NewMazeTargets creates a new MazeTargets instance
func NewMazeTargets(chars *models.Characters, maze *Maze) *MazeTargets {
	return &MazeTargets{
		chars: chars,
		maze:  maze,
	}
}

// DefineTargets dynamically places a single target in an open cell
func (mt *MazeTargets) DefineTargets() []models.Target {
	var openCells []models.Position

	// get all open cells
	for y := 1; y < mt.maze.height-1; y++ {
		for x := 1; x < mt.maze.width-1; x++ {
			// account for the maze offset
			pos := models.Position{
				X: x + mt.maze.offsetX,
				Y: y + mt.maze.offsetY,
			}
			if !mt.isPositionCollidingWithWall(pos) {
				openCells = append(openCells, pos)
			}
		}
	}

	// need a valid target
	if len(openCells) == 0 {
		panic("no open cells available for target placement")
	}

	// select a "random" open cell, not truly random since the maze is deterministic
	targetPos := openCells[mt.maze.rand.Intn(len(openCells))]
	return []models.Target{
		{Position: targetPos, Reached: false},
	}
}

// GetTargetCount returns the number of targets
func (mt *MazeTargets) GetTargetCount() int {
	return 1
}

// UpdateGrid updates the grid with targets and walls
func (mt *MazeTargets) UpdateGrid(grid [][]rune, targets []models.Target, current int, chars *models.Characters) {
	target := targets[0]
	grid[target.Position.Y][target.Position.X] = chars.Target.Active.Rune

	for _, wall := range mt.maze.GetWalls() {
		grid[wall.Y][wall.X] = chars.Wall.Rune
	}
}

// Helper function to check if a position collides with walls
func (mt *MazeTargets) isPositionCollidingWithWall(pos models.Position) bool {
	return slices.Contains(mt.maze.GetWalls(), pos)
}
