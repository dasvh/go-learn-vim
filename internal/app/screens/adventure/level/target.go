package level

import (
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
		{models.Position{X: offsetX, Y: offsetY}, false},
		{models.Position{X: width - offsetX - 1, Y: offsetY}, false},
		{models.Position{X: offsetX, Y: height - offsetY - 1}, false},
		{models.Position{X: width - offsetX - 1, Y: height - offsetY - 1}, false},
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
