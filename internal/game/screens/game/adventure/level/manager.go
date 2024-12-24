package level

type Manager struct {
	currentLevel Level
	levelNumber  int
}

func NewManager() *Manager {
	return &Manager{
		currentLevel: NewLevelZero(),
		levelNumber:  1,
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
