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

func (m *Manager) LevelFromSave(state SavedLevel) error {
	if len(state.Targets) == 0 {
		return fmt.Errorf("invalid save state: no targets found")
	}

	m.currentLevel, m.levelNumber = NewLevelZero()

	if state.Width <= 0 || state.Height <= 0 {
		return fmt.Errorf("invalid dimensions: width=%d height=%d", state.Width, state.Height)
	}

	m.currentLevel.Init(state.Width, state.Height)
	err := m.currentLevel.Restore(state)
	if err != nil {
		return fmt.Errorf("failed to restore level state: %w", err)
	}

	return nil
}
