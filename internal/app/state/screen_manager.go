package state

import (
	tea "github.com/charmbracelet/bubbletea"
)

// ScreenManager handles the management and switching of app screens in the app
type ScreenManager struct {
	startScreen Screen
	screens     map[Screen]tea.Model
}

// NewManager creates and returns a new ScreenManager instance with MainScreen as the start screen
func NewManager() *ScreenManager {
	return &ScreenManager{
		startScreen: MainMenuScreen,
		screens:     make(map[Screen]tea.Model),
	}
}

// Register adds a new screen to the ScreenManager
func (sm *ScreenManager) Register(view Screen, model tea.Model) {
	sm.screens[view] = model
}

// CurrentScreen returns the currently active screen model from the ScreenManager
func (sm *ScreenManager) CurrentScreen() tea.Model {
	return sm.screens[sm.startScreen]
}

// SwitchTo changes the current screen to the specified screen
func (sm *ScreenManager) SwitchTo(to Screen) tea.Cmd {
	if model, exists := sm.screens[to]; exists {
		sm.startScreen = to
		return model.Init()
	}
	return nil
}

// ActiveScreen returns the currently active screen
func (sm *ScreenManager) ActiveScreen() Screen {
	return sm.startScreen
}

// Screens returns all the screens in the ScreenManager
func (sm *ScreenManager) Screens() map[Screen]tea.Model {
	return sm.screens
}
