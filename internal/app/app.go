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
	gc   *controllers.Game
	lc   *controllers.Level
	size tea.WindowSizeMsg
}

// NewApp initializes a new App instance with a screen controller
// and registers the respective screens
func NewApp(repo storage.GameRepository) *App {
	screen := controllers.NewScreen()
	game := controllers.NewGame(repo)
	level := controllers.NewLevel()

	app := &App{
		sc: screen,
		gc: game,
		lc: level,
	}

	screen.Register(models.MainMenuScreen, menus.NewMainMenu(repo.HasIncompleteGames()))
	screen.Register(models.InfoMenuScreen, menus.NewInfoMenu())
	screen.Register(models.VimInfoScreen, info.NewVimInfo())
	screen.Register(models.CheatsheetInfoScreen, info.NewVimCheatsheet())
	screen.Register(models.LoadSaveSelectionScreen, selection.NewSaveSelection(repo, app.handleSaveSelection))
	screen.Register(models.NewGameScreen, menus.NewGameModes())
	screen.Register(models.PlayerSelectionScreen, selection.NewPlayerSelection(game, models.NewGameScreen))
	screen.Register(models.AdventureModeScreen, adventure.NewAdventure(game, level))
	screen.Register(models.LevelSelectionScreen, selection.NewLevelSelection(level))
	screen.Register(models.ScoresScreen, leaderboards.NewScoresScreen(repo))
	screen.Register(models.StatsScreen, leaderboards.NewStatsScreen(repo))

	return app
}

// handleSaveSelection handles the selection of a save and loads the adventure game
func (a *App) handleSaveSelection(save models.GameSave) tea.Cmd {
	adventureState, ok := save.GameState.(models.AdventureGameState)
	if !ok {
		fmt.Println("Invalid game state type in save:", save.GameState)
		return nil
	}

	// pass SaveID to adventure game
	adventureState.SaveID = save.ID

	loadedAdventure, err := adventure.Load(a.gc, a.lc, adventureState, a.size)
	if err != nil {
		fmt.Printf("Failed to load adventure: %v\n", err)
		return nil
	}

	return tea.Batch(
		func() tea.Msg { return models.SetPlayerMsg{Player: save.Player} },
		func() tea.Msg {
			return models.ScreenTransitionMsg{Screen: models.AdventureModeScreen, Model: loadedAdventure}
		},
	)
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
	case models.SetPlayerMsg:
		// set the player for the game controller
		a.gc.SetPlayer(msg.Player)
		// pass the player to the adventure mode screen
		if model, ok := a.sc.Screens()[models.AdventureModeScreen].(*adventure.Adventure); ok {
			model.Update(msg)
		}
		return a, nil
	// pass the level from the level selection screen to the adventure mode screen to init the level
	case models.SetLevelMsg:
		if model, ok := a.sc.Screens()[models.AdventureModeScreen].(*adventure.Adventure); ok {
			model.Update(msg)
		}
		return a, nil
	// update the load button in the main menu screen
	case models.UpdateLoadButtonMsg:
		if model, ok := a.sc.Screens()[models.MainMenuScreen].(*menus.Main); ok {
			model.UpdateLoadButton(msg.CanLoadGame)
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
