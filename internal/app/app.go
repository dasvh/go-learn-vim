package app

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/app/screens/adventure"
	"github.com/dasvh/go-learn-vim/internal/app/screens/info"
	"github.com/dasvh/go-learn-vim/internal/app/screens/menus"
	"github.com/dasvh/go-learn-vim/internal/app/state"
	"github.com/dasvh/go-learn-vim/internal/storage"
)

// App represents the main app structure which holds the state manager
// and the window size message
type App struct {
	manager *state.ScreenManager
	size    tea.WindowSizeMsg
}

// NewApp initializes a new App instance with a state manager
// and registers the respective screens
func NewApp(repo storage.AdventureGameRepository) *App {
	manager := state.NewManager()

	hasIncompleteGame, err := repo.HasIncompleteGame()
	if err != nil {
		fmt.Println("Failed to check for app saves + ", err)
	}

	manager.Register(state.MainMenuScreen, menus.NewMainMenu(hasIncompleteGame))
	manager.Register(state.LoadGameScreen, menus.NewLoad(repo))
	manager.Register(state.InfoMenuScreen, menus.NewInfoMenu())
	manager.Register(state.VimInfoScreen, info.NewVimInfo())
	manager.Register(state.MotionsInfoScreen, info.NewMotionsInfo())
	manager.Register(state.NewGameScreen, menus.NewGameModes())
	manager.Register(state.AdventureModeScreen, adventure.NewAdventure(repo))

	return &App{
		manager: manager,
	}
}

// Init initializes the current app views and returns any initial commands
func (g *App) Init() tea.Cmd {
	return g.manager.CurrentScreen().Init()
}

// Update handles state updates based on incoming messages
func (g *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if msg.Width != 0 && msg.Height != 0 {
			g.size = msg
			for view, model := range g.manager.Screens() {
				if updatedModel, _ := model.Update(msg); updatedModel != nil {
					g.manager.Register(view, updatedModel)
				}
			}
		}
	// handle screen transitions with model registration
	case state.ScreenTransitionMsg:
		g.manager.Register(msg.Screen, msg.Model)
		return g, g.manager.SwitchTo(msg.Screen)
	case state.Screen:
		return g, g.manager.SwitchTo(msg)
	}

	model, cmd := g.manager.CurrentScreen().Update(msg)
	g.manager.Register(g.manager.ActiveScreen(), model)
	return g, cmd
}

// View returns the string representation of the current views managed by the app
func (g *App) View() string {
	return g.manager.CurrentScreen().View()
}
