package menus

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/app/screens"
	"github.com/dasvh/go-learn-vim/internal/app/screens/adventure"
	"github.com/dasvh/go-learn-vim/internal/storage"
	"github.com/dasvh/go-learn-vim/internal/views"
)

const (
	ButtonLoadSelection = "Load Selection"
)

// Load represents the load game screen
type Load struct {
	*views.MenuView
	repo storage.GameRepository
	size tea.WindowSizeMsg
}

// NewLoad creates a new load game screen
func NewLoad(repo storage.GameRepository) views.Menu {
	base := views.NewBaseMenu("Load Game", []views.ButtonConfig{
		{Label: ButtonLoadSelection},
	})
	return &Load{MenuView: base, repo: repo}
}

// LoadAdventureGame loads the saved app and returns a command to transition to the adventure mode screen
func (l *Load) LoadAdventureGame() tea.Cmd {
	return func() tea.Msg {
		hasIncompleteGame, err := l.repo.HasIncompleteGames()
		if !hasIncompleteGame || err != nil {
			fmt.Println("No save file found")
			return screens.ChangeScreen(screens.MainMenuScreen)
		}

		// load game state
		gameState, err := l.repo.LoadGameState("0") // Replace "0" with actual game ID
		if err != nil {
			fmt.Printf("Failed to load saved game controllers: %v\n", err)
			return screens.ChangeScreen(screens.MainMenuScreen)
		}

		loadedAdventure, err := adventure.Load(l.repo, gameState, l.size)
		if err != nil {
			fmt.Printf("Failed to initialize adventure: %v\n", err)
			return screens.ChangeScreen(screens.MainMenuScreen)
		}

		return screens.ScreenTransitionMsg{
			Screen: screens.AdventureModeScreen,
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
			return l, screens.ChangeScreen(screens.MainMenuScreen)
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
