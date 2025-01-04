package models

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

// TargetBehavior defines the behavior of a target
type TargetBehavior interface {
	DefineTargets(width, height int) []Target
	UpdateGrid(grid [][]rune, targets []Target, current int, chars *Characters)
	GetTargetCount() int
}
