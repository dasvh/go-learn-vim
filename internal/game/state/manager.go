package state

import (
	tea "github.com/charmbracelet/bubbletea"
	ui "github.com/dasvh/go-learn-vim/internal/ui/menu"
)

// MenuManager manages the game views
type MenuManager struct {
	// activeView represents the current active view
	activeView View
	// views represents the registered views
	views map[View]ui.Menu
}

// NewManager returns a new MenuManager instance
func NewManager() *MenuManager {
	return &MenuManager{
		activeView: MainView,
		views:      make(map[View]ui.Menu),
	}
}

// Register registers a view with its corresponding model
// with `view` as the key and `model` as the value
func (m *MenuManager) Register(view View, model tea.Model) {
	if menu, ok := model.(ui.Menu); ok {
		m.views[view] = menu
	}
}

// Current returns the current active model
func (m *MenuManager) Current() tea.Model {
	return m.views[m.activeView]
}

// Switch switches the active view to the specified view
func (m *MenuManager) Switch(to View) tea.Cmd {
	if model, exists := m.views[to]; exists {
		m.activeView = to
		return model.Init()
	}
	return nil
}

// ActiveView returns the current active view
func (m *MenuManager) ActiveView() View {
	return m.activeView
}

// Views returns a map of all registered views
func (m *MenuManager) Views() map[View]tea.Model {
	views := make(map[View]tea.Model)
	for k, v := range m.views {
		views[k] = v
	}
	return views
}
