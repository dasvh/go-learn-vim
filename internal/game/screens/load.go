package screens

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/components/view"
	"github.com/dasvh/go-learn-vim/internal/game/screens/game/adventure"
	"github.com/dasvh/go-learn-vim/internal/game/state"
	"github.com/dasvh/go-learn-vim/internal/storage"
)

const (
	ButtonLoadSelection = "Load Selection"
)

// Load represents the load game view
type Load struct {
	*view.MenuView
	repo storage.AdventureGameRepository
}

// NewLoad returns a new Load instance
func NewLoad(repo storage.AdventureGameRepository) view.Menu {
	base := view.NewBaseMenu("Load Game", []view.ButtonConfig{
		{Label: ButtonLoadSelection},
	})
	return &Load{MenuView: base, repo: repo}
}

// LoadAdventureGame loads the saved game and returns a command to transition to the adventure mode screen
func (l *Load) LoadAdventureGame() tea.Cmd {
	return func() tea.Msg {
		hasIncompleteGame, err := l.repo.HasIncompleteGame()
		if !hasIncompleteGame || err != nil {
			fmt.Println("No save file found")
			return state.ChangeScreen(state.MainMenuScreen)
		}

		// create a new adventure instance
		newAdventure := adventure.NewAdventure(l.repo)

		// load the saved game into the new adventure instance
		if err := newAdventure.Load(l.repo); err != nil {
			fmt.Printf("Failed to load saved game: %v\n", err)
			return state.ChangeScreen(state.MainMenuScreen)
		}

		return state.ScreenTransitionMsg{
			Screen: state.AdventureModeScreen,
			Model:  newAdventure,
		}
	}
}

// Update handles messages and transitions between menu states
func (l *Load) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := l.MenuView.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, l.Controls().Select) {
			return l, l.HandleSelection()
		}
		if key.Matches(msg, l.Controls().Back) {
			return l, state.ChangeScreen(state.MainMenuScreen)
		}
	}
	return l, cmd
}

// HandleSelection implements ButtonHandler interface
func (l *Load) HandleSelection() tea.Cmd {
	selected := l.CurrentButton()
	if selected == nil || selected.Inactive {
		return nil
	}

	switch selected.Label {
	case ButtonLoadSelection:
		return l.LoadAdventureGame()
	default:
		return nil
	}
}
