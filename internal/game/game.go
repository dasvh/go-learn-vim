package game

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/game/screens"
	"github.com/dasvh/go-learn-vim/internal/game/screens/info"
	"github.com/dasvh/go-learn-vim/internal/game/state"
)

// Game represents the main game structure which holds the state manager
// and the window size message
type Game struct {
	manager *state.ViewManager
	size    tea.WindowSizeMsg
}

// NewGame initializes a new Game instance with a state manager
// and registers the respective screens
func NewGame() *Game {
	manager := state.NewManager()

	manager.Register(state.MainMenuScreen, screens.NewMainMenu())
	manager.Register(state.InfoMenuScreen, info.NewInfoMenu())
	manager.Register(state.VimInfoScreen, info.NewVimInfo())
	manager.Register(state.MotionsInfoScreen, info.NewMotionsInfo())

	return &Game{
		manager: manager,
	}
}

// Init initializes the current game view and returns any initial commands
func (g *Game) Init() tea.Cmd {
	return g.manager.CurrentView().Init()
}

// Update handles state updates based on incoming messages
func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if msg.Width != 0 && msg.Height != 0 {
			g.size = msg
			for view, model := range g.manager.Views() {
				if updatedModel, _ := model.Update(msg); updatedModel != nil {
					g.manager.Register(view, updatedModel)
				}
			}
		}
	case state.GameScreen:
		return g, g.manager.SwitchTo(msg)
	}

	model, cmd := g.manager.CurrentView().Update(msg)
	g.manager.Register(g.manager.ActiveView(), model)
	return g, cmd
}

// View returns the string representation of the current view managed by the game
func (g *Game) View() string {
	return g.manager.CurrentView().View()
}
