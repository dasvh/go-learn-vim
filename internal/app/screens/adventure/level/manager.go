package level

import "fmt"

type Manager struct {
	currentLevel Level
	levelNumber  int
}

func NewManager() *Manager {
	level, number := NewLevelZero()
	return &Manager{
		currentLevel: level,
		levelNumber:  number,
	}
}

func (m *Manager) InitCurrentLevel(width, height int) {
	m.currentLevel.Init(width, height)
}

func (m *Manager) GetCurrentLevel() Level {
	return m.currentLevel
}

func (m *Manager) GetLevelNumber() int {
	return m.levelNumber
}

func (m *Manager) RestoreLevel(state SavedLevel) error {
	if len(state.Targets) == 0 {
		return fmt.Errorf("invalid save state: no targets found")
	}

	level, number := NewLevelZero()
	m.currentLevel = level
	m.levelNumber = number

	return m.currentLevel.Restore(state)
}

func (m *Manager) InitOrResizeLevel(width, height int) error {
	if !m.currentLevel.InProgress() {
		m.InitCurrentLevel(width, height)
		return nil
	}

	params := SavedLevel{
		Number:         m.levelNumber,
		Width:          width,
		Height:         height,
		PlayerPosition: m.currentLevel.GetCurrentPosition(),
		Targets:        m.currentLevel.GetTargets(),
		CurrentTarget:  m.currentLevel.GetCurrentTarget(),
		Completed:      m.currentLevel.IsCompleted(),
		InProgress:     m.currentLevel.InProgress(),
	}
	return m.currentLevel.Restore(params)
}
