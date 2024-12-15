package game

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/game/state"
	"github.com/dasvh/go-learn-vim/internal/game/views"
)

// Game represents the main game structure which holds the state manager
// and the window size message
type Game struct {
	manager *state.MenuManager
	size    tea.WindowSizeMsg
}

// NewGame initializes a new Game instance with a state manager
// and registers the respective views
func NewGame() *Game {
	manager := state.NewManager()

	manager.Register(state.MainView, views.NewMainView())
	manager.Register(state.InfoView, views.NewInfoView())

	return &Game{
		manager: manager,
	}
}

// Init handles the initialization of the game state
func (g *Game) Init() tea.Cmd {
	return g.manager.Current().Init()
}

// Update handles messages and updates the game state accordingly
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
	case state.View:
		return g, g.manager.Switch(msg)
	}

	model, cmd := g.manager.Current().Update(msg)
	g.manager.Register(g.manager.ActiveView(), model)
	return g, cmd
}

// View returns the current view of the game
func (g *Game) View() string {
	return g.manager.Current().View()
}
