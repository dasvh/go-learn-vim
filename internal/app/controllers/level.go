package controllers

import (
	"fmt"
	"github.com/dasvh/go-learn-vim/internal/app/screens/adventure/level"
	"github.com/dasvh/go-learn-vim/internal/models"
)

type Level struct {
	currentLevel models.Level
	levelNumber  int
}

func NewLevel() *Level {
	zero, number := level.NewLevelZero()
	return &Level{
		currentLevel: zero,
		levelNumber:  number,
	}
}

func (lc *Level) InitCurrentLevel(width, height int) {
	lc.currentLevel.Init(width, height)
}

func (lc *Level) GetCurrentLevel() models.Level {
	return lc.currentLevel
}

func (lc *Level) GetLevelNumber() int {
	return lc.levelNumber
}

func (lc *Level) RestoreLevel(state models.SavedLevel) error {
	if len(state.Targets) == 0 {
		return fmt.Errorf("invalid save state: no targets found")
	}

	zero, number := level.NewLevelZero()
	lc.currentLevel = zero
	lc.levelNumber = number

	return lc.currentLevel.Restore(state)
}

func (lc *Level) InitOrResizeLevel(width, height int) error {
	if !lc.currentLevel.InProgress() {
		lc.InitCurrentLevel(width, height)
		return nil
	}

	params := models.SavedLevel{
		Number:         lc.levelNumber,
		Width:          width,
		Height:         height,
		PlayerPosition: lc.currentLevel.GetCurrentPosition(),
		Targets:        lc.currentLevel.GetTargets(),
		CurrentTarget:  lc.currentLevel.GetCurrentTarget(),
		Completed:      lc.currentLevel.IsCompleted(),
		InProgress:     lc.currentLevel.InProgress(),
	}
	return lc.currentLevel.Restore(params)
}
