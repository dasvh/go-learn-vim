package controllers

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/models"
)

// Screen handles the management and switching of app screens in the app
type Screen struct {
	startScreen models.Screen
	screens     map[models.Screen]tea.Model
}

// NewScreen creates and returns a new Screen instance with MainScreen as the start screen
func NewScreen() *Screen {
	return &Screen{
		startScreen: models.MainMenuScreen,
		screens:     make(map[models.Screen]tea.Model),
	}
}

// Register adds a new screen to the ScreenController
func (sc *Screen) Register(view models.Screen, model tea.Model) {
	sc.screens[view] = model
}

// CurrentScreen returns the currently active screen model from the ScreenController
func (sc *Screen) CurrentScreen() tea.Model {
	return sc.screens[sc.startScreen]
}

// SwitchTo changes the current screen to the specified screen
func (sc *Screen) SwitchTo(to models.Screen) tea.Cmd {
	if model, exists := sc.screens[to]; exists {
		sc.startScreen = to
		return model.Init()
	}
	return nil
}

// ActiveScreen returns the currently active screen
func (sc *Screen) ActiveScreen() models.Screen {
	return sc.startScreen
}

// Screens returns all the screens in the ScreenController
func (sc *Screen) Screens() map[models.Screen]tea.Model {
	return sc.screens
}
