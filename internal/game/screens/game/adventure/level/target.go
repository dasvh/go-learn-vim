package level

import "github.com/dasvh/go-learn-vim/internal/game/screens/game/adventure/character"

// Target represents a Position and whether it has been Reached
type Target struct {
	Position Position
	Reached  bool
}

// TargetBehavior defines the behavior of a target
type TargetBehavior interface {
	DefineTargets(width, height int) []Target
	UpdateGrid(grid [][]rune, targets []Target, current int, chars *character.Characters)
	GetTargetCount() int
}

// CornerTargets defines the behavior of a target that is placed in the corners of the screen
type CornerTargets struct {
	chars *character.Characters
}

// NewCornerTargets creates a new CornerTargets
func NewCornerTargets(chars *character.Characters) TargetBehavior {
	return &CornerTargets{chars: chars}
}

// DefineTargets defines the targets in the corners of the screen
func (ct *CornerTargets) DefineTargets(width, height int) []Target {
	offsetX := int(float64(width) * 0.2)
	offsetY := int(float64(height) * 0.2)

	return []Target{
		{Position{offsetX, offsetY}, false},
		{Position{width - offsetX - 1, offsetY}, false},
		{Position{offsetX, height - offsetY - 1}, false},
		{Position{width - offsetX - 1, height - offsetY - 1}, false},
	}
}

// UpdateGrid updates the grid with the targets
func (ct *CornerTargets) UpdateGrid(grid [][]rune, targets []Target, current int, chars *character.Characters) {
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
