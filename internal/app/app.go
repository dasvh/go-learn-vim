package app

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/app/controllers"
	"github.com/dasvh/go-learn-vim/internal/app/screens/adventure"
	"github.com/dasvh/go-learn-vim/internal/app/screens/info"
	"github.com/dasvh/go-learn-vim/internal/app/screens/leaderboards"
	"github.com/dasvh/go-learn-vim/internal/app/screens/menus"
	"github.com/dasvh/go-learn-vim/internal/app/screens/selection"
	"github.com/dasvh/go-learn-vim/internal/models"
	"github.com/dasvh/go-learn-vim/internal/storage"
)

// App represents the main app structure which holds the screen controller
// and the window size message
type App struct {
	sc   *controllers.Screen
	size tea.WindowSizeMsg
}

// NewApp initializes a new App instance with a screen controller
// and registers the respective screens
func NewApp(repo storage.GameRepository) *App {
	screen := controllers.NewScreen()
	game := controllers.NewGame(repo)

	hasIncompleteGame, err := repo.HasIncompleteGames()
	if err != nil {
		fmt.Println("Failed to check for app saves + ", err)
	}

	screen.Register(models.MainMenuScreen, menus.NewMainMenu(hasIncompleteGame))
	screen.Register(models.LoadGameScreen, menus.NewLoad(repo, game))
	screen.Register(models.InfoMenuScreen, menus.NewInfoMenu())
	screen.Register(models.VimInfoScreen, info.NewVimInfo())
	screen.Register(models.MotionsInfoScreen, info.NewMotionsInfo())
	screen.Register(models.NewGameScreen, menus.NewGameModes())
	screen.Register(models.PlayerSelectionScreen, selection.NewPlayerSelection(game, models.NewGameScreen))
	screen.Register(models.AdventureModeScreen, adventure.NewAdventure(game))
	screen.Register(models.StatsScreen, leaderboards.NewStatsScreen(repo))
	screen.Register(models.ScoresScreen, leaderboards.NewScoresScreen(repo))

	return &App{
		sc: screen,
	}
}

// Init initializes the current app views and returns any initial commands
func (a *App) Init() tea.Cmd {
	return a.sc.CurrentScreen().Init()
}

// Update handles state updates based on incoming messages
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if msg.Width != 0 && msg.Height != 0 {
			a.size = msg
			for view, model := range a.sc.Screens() {
				if updatedModel, _ := model.Update(msg); updatedModel != nil {
					a.sc.Register(view, updatedModel)
				}
			}
		}
	// pass the player from the player selection screen to the adventure mode screen
	case models.SetPlayerMsg:
		if model, ok := a.sc.Screens()[models.AdventureModeScreen].(*adventure.Adventure); ok {
			model.Update(msg)
		}
		return a, nil
	// handle screen transitions with model registration
	case models.ScreenTransitionMsg:
		a.sc.Register(msg.Screen, msg.Model)
		return a, a.sc.SwitchTo(msg.Screen)
	case models.Screen:
		return a, a.sc.SwitchTo(msg)
	}

	model, cmd := a.sc.CurrentScreen().Update(msg)
	a.sc.Register(a.sc.ActiveScreen(), model)
	return a, cmd
}

// View returns the string representation of the current views managed by the app
func (a *App) View() string {
	return a.sc.CurrentScreen().View()
}
