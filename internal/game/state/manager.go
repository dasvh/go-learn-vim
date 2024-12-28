package state

import (
	tea "github.com/charmbracelet/bubbletea"
)

// ViewManager handles the management and switching of game models
type ViewManager struct {
	activeView GameScreen
	views      map[GameScreen]tea.Model
}

// NewManager creates and returns a new ViewManager instance with MainMenuScreen as the default active view
func NewManager() *ViewManager {
	return &ViewManager{
		activeView: AdventureModeScreen,
		views:      make(map[GameScreen]tea.Model),
	}
}

// Register adds a new view to the ViewManager
func (vm *ViewManager) Register(view GameScreen, model tea.Model) {
	vm.views[view] = model
}

// CurrentView returns the currently active view model from the ViewManager
func (vm *ViewManager) CurrentView() tea.Model {
	return vm.views[vm.activeView]
}

// SwitchTo changes the active view to the specified GameScreen if it exists
func (vm *ViewManager) SwitchTo(to GameScreen) tea.Cmd {
	if model, exists := vm.views[to]; exists {
		vm.activeView = to
		return model.Init()
	}
	return nil
}

// ActiveView returns the currently active game screen view
func (vm *ViewManager) ActiveView() GameScreen {
	return vm.activeView
}

// Views returns a map containing all registered views
func (vm *ViewManager) Views() map[GameScreen]tea.Model {
	return vm.views
}
