package models

// Level represents a level in the adventure mode
type Level interface {
	Init(width, height int)
	PlayerMove(position Position) PlayerMovement
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

// SavedLevel represents a saved level
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
