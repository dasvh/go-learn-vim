package controllers

import (
	"fmt"
	"github.com/dasvh/go-learn-vim/internal/app/screens/adventure/level"
	"github.com/dasvh/go-learn-vim/internal/models"
)

// Level is a controller for level related actions
type Level struct {
	levels  map[int]models.Level
	current models.Level
}

// NewLevel creates a new Level controller with levels
func NewLevel() *Level {
	return &Level{
		levels: map[int]models.Level{
			level.NewLevelZero().Number(): level.NewLevelZero(),
			level.NewLevelOne().Number():  level.NewLevelOne(),
		},
	}
}

// GetLevels returns all levels
func (lc *Level) GetLevels() map[int]models.Level {
	return lc.levels
}

// GetLevelsCount returns the number of levels
func (lc *Level) GetLevelsCount() int {
	return len(lc.levels)
}

// SetLevel sets the current level to the given level
func (lc *Level) SetLevel(lvl models.Level) {
	lc.current = lvl
}

// InitCurrentLevel initializes the current level with the given width and height
func (lc *Level) InitCurrentLevel(width, height int) {
	lc.current.Init(width, height)
}

// GetCurrentLevel returns the current level
func (lc *Level) GetCurrentLevel() models.Level {
	return lc.current
}

// GetLevelNumber returns the number of the current level
func (lc *Level) GetLevelNumber() int {
	return lc.current.Number()
}

// RestoreLevel restores the level state from a saved state
func (lc *Level) RestoreLevel(state models.SavedLevel) error {
	if state.Width <= 0 || state.Height <= 0 {
		return fmt.Errorf("invalid level dimensions in save state")
	}

	lvl, exists := lc.levels[state.Number]
	if !exists {
		return fmt.Errorf("level %d not found", state.Number)
	}

	lc.current = lvl
	return lc.current.Restore(state)
}

// InitOrResizeLevel initializes or resizes the current level with the given width and height
func (lc *Level) InitOrResizeLevel(width, height int) error {
	if lc.current == nil {
		return fmt.Errorf("no level found")
	}

	if !lc.current.InProgress() {
		lc.InitCurrentLevel(width, height)
		return nil
	}

	params := models.SavedLevel{
		Number:         lc.current.Number(),
		Width:          width,
		Height:         height,
		PlayerPosition: lc.current.GetCurrentPosition(),
		Targets:        lc.current.GetTargets(),
		CurrentTarget:  lc.current.GetCurrentTarget(),
		Completed:      lc.current.IsCompleted(),
		InProgress:     lc.current.InProgress(),
	}
	return lc.current.Restore(params)
}
