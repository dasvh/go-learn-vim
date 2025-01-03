package app

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/app/controllers"
	"github.com/dasvh/go-learn-vim/internal/app/screens"
	"github.com/dasvh/go-learn-vim/internal/app/screens/adventure"
	"github.com/dasvh/go-learn-vim/internal/app/screens/info"
	"github.com/dasvh/go-learn-vim/internal/app/screens/menus"
	"github.com/dasvh/go-learn-vim/internal/app/screens/selection"
	"github.com/dasvh/go-learn-vim/internal/storage"
)

// App represents the main app structure which holds the screen manager
// and the window size message
type App struct {
	manager *screens.ScreenManager
	size    tea.WindowSizeMsg
}

// NewApp initializes a new App instance with a screen manager
// and registers the respective screens
func NewApp(repo storage.GameRepository) *App {
	manager := screens.NewManager()
	gc := controllers.NewGame(repo)

	hasIncompleteGame, err := repo.HasIncompleteGames()
	if err != nil {
		fmt.Println("Failed to check for app saves + ", err)
	}

	manager.Register(screens.MainMenuScreen, menus.NewMainMenu(hasIncompleteGame))
	manager.Register(screens.LoadGameScreen, menus.NewLoad(repo, gc))
	manager.Register(screens.InfoMenuScreen, menus.NewInfoMenu())
	manager.Register(screens.VimInfoScreen, info.NewVimInfo())
	manager.Register(screens.MotionsInfoScreen, info.NewMotionsInfo())
	manager.Register(screens.NewGameScreen, menus.NewGameModes())
	manager.Register(screens.PlayerSelectionScreen, selection.NewPlayerSelection(gc, screens.NewGameScreen))
	manager.Register(screens.AdventureModeScreen, adventure.NewAdventure(repo, gc))

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
	// pass the player from the player selection screen to the adventure mode screen
	case screens.SetPlayerMsg:
		if model, ok := g.manager.Screens()[screens.AdventureModeScreen].(*adventure.Adventure); ok {
			model.Update(msg)
		}
		return g, nil
	// handle screen transitions with model registration
	case screens.ScreenTransitionMsg:
		g.manager.Register(msg.Screen, msg.Model)
		return g, g.manager.SwitchTo(msg.Screen)
	case screens.Screen:
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
