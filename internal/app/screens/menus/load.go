package menus

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/app/screens/adventure"
	"github.com/dasvh/go-learn-vim/internal/app/state"
	"github.com/dasvh/go-learn-vim/internal/storage"
	"github.com/dasvh/go-learn-vim/internal/views"
)

const (
	ButtonLoadSelection = "Load Selection"
)

// Load represents the load game screen
type Load struct {
	*views.MenuView
	repo storage.AdventureGameRepository
	size tea.WindowSizeMsg
}

// NewLoad creates a new load game screen
func NewLoad(repo storage.AdventureGameRepository) views.Menu {
	base := views.NewBaseMenu("Load Game", []views.ButtonConfig{
		{Label: ButtonLoadSelection},
	})
	return &Load{MenuView: base, repo: repo}
}

// LoadAdventureGame loads the saved app and returns a command to transition to the adventure mode screen
func (l *Load) LoadAdventureGame() tea.Cmd {
	return func() tea.Msg {
		hasIncompleteGame, err := l.repo.HasIncompleteGame()
		if !hasIncompleteGame || err != nil {
			fmt.Println("No save file found")
			return state.ChangeScreen(state.MainMenuScreen)
		}

		loadedGame, err := l.repo.LoadAdventureGame()
		if err != nil {
			fmt.Printf("Failed to load saved app: %v\n", err)
			return state.ChangeScreen(state.MainMenuScreen)
		}

		loadedAdventure, err := adventure.LoadAdventure(l.repo, loadedGame, l.size)
		if err != nil {
			fmt.Printf("Failed to load saved app: %v\n", err)
			return state.ChangeScreen(state.MainMenuScreen)
		}

		return state.ScreenTransitionMsg{
			Screen: state.AdventureModeScreen,
			Model:  loadedAdventure,
		}
	}
}

// Update handles state updates based on incoming messages
func (l *Load) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := l.MenuView.Update(msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if msg.Width > 0 && msg.Height > 0 {
			l.size = msg
		}
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

// HandleSelection handles the selection of a button
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
