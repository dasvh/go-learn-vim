package selection

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dasvh/go-learn-vim/internal/app/controllers"
	cl "github.com/dasvh/go-learn-vim/internal/components/list"
	"github.com/dasvh/go-learn-vim/internal/models"
	"github.com/dasvh/go-learn-vim/internal/style"
	"github.com/dasvh/go-learn-vim/internal/views"
)

// PlayerSelection is a screen that allows the user to select a player or create a new one
type PlayerSelection struct {
	view        *views.SelectionView
	size        tea.WindowSizeMsg
	gc          *controllers.Game
	gameScreen  models.Screen
	inInputMode bool
	textInput   textinput.Model
	players     []models.Player
	items       []cl.Item
}

// NewPlayerSelection creates a new PlayerSelection screen.
// It takes a game controller and a game screen
func NewPlayerSelection(gc *controllers.Game, gameScreen models.Screen) *PlayerSelection {
	players, _ := gc.Players()
	playerItems := make([]cl.Item, len(players))

	for i, player := range players {
		playerItems[i] = cl.Item{Name: player.Name}
	}

	ps := &PlayerSelection{
		gc:         gc,
		gameScreen: gameScreen,
		textInput:  textinput.New(),
		players:    players,
		items:      playerItems,
	}

	ps.textInput.Prompt = "Enter new player name: "
	ps.textInput.CharLimit = models.PlayerNameMaxLength
	return ps
}

// setSelectionView sets the view of the player selection screen
func (ps *PlayerSelection) setSelectionView() {
	width := ps.size.Width
	height := ps.size.Height / 2
	playerList := cl.NewList(ps.items, width, height,
		cl.WithItemsIdentifiers("Select a player or create a new one", "player", "players"),
		cl.WithDisableQuitKeybindings(),
		cl.WithShowDescription(false),
		cl.WithTitleStyle(style.PlayerSelection.Title),
		cl.WithSelectedTitleStyle(style.PlayerSelection.SelectedItem),
		cl.WithNormalTitleStyle(style.PlayerSelection.Item),
		cl.WithFilterMatchStyle(style.PlayerSelection.FilterMatch),
		cl.WithDimmedTitleStyle(style.PlayerSelection.DimmedItem),
	)

	ps.view = views.NewSelectionView(
		"Player Selection",
		playerList,
		&ps.textInput,
		ps.handleSelect,
		ps.handleInsert,
	)
}

// handleInsert toggles the input mode
func (ps *PlayerSelection) handleInsert() tea.Cmd {
	ps.inInputMode = !ps.inInputMode
	if ps.inInputMode {
		ps.focusInput()
	} else {
		ps.resetInput()
	}

	return nil
}

// handleSelect sets the selected player in the game controller and changes the screen to the game screen
func (ps *PlayerSelection) handleSelect(item cl.Item) tea.Cmd {
	for _, player := range ps.players {
		if player.Name == item.Name {
			ps.gc.SetPlayer(player)
			// return a batch command to change the screen and set the player
			return tea.Batch(
				models.ChangeScreen(ps.gameScreen),
				func() tea.Msg { return models.SetPlayerMsg{Player: player} },
			)
		}
	}
	return nil
}

// focusInput focuses the text input
func (ps *PlayerSelection) focusInput() {
	ps.textInput.Focus()
	ps.textInput.Reset()
}

// resetInput resets the input field, clearing its content and placeholder,
// and exits input mode.
func (ps *PlayerSelection) resetInput() {
	ps.textInput.Placeholder = ""
	ps.textInput.Blur()
	ps.textInput.Reset()
	ps.inInputMode = false
}

// handlePlayerCreationUpdate handles the player creation update
func (ps *PlayerSelection) handlePlayerCreationUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, ps.view.InsertControls.Confirm):
			return ps.createPlayer()
		case key.Matches(msg, ps.view.InsertControls.Cancel):
			ps.resetInput()
		case key.Matches(msg, ps.view.InsertControls.Quit):
			return ps, tea.Quit
		}
	}

	var cmd tea.Cmd
	ps.textInput, cmd = ps.textInput.Update(msg)
	return ps, cmd
}

// setPlaceholder sets the placeholder of the text input
func (ps *PlayerSelection) setPlaceholder(message, color string) {
	ps.textInput.SetValue("")
	ps.textInput.Placeholder = message
	ps.textInput.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(color))
}

// createPlayer creates a new player using the text input and adds it to the list
func (ps *PlayerSelection) createPlayer() (tea.Model, tea.Cmd) {
	newPlayerName := ps.textInput.Value()
	if newPlayerName == "" {
		ps.setPlaceholder("Name cannot be empty!", "#FF0000")
		return ps, nil
	}

	createdPlayer, err := ps.gc.CreatePlayer(newPlayerName)
	if err != nil {
		ps.setPlaceholder(fmt.Sprintf("Error creating player: %s", err), "#FF0000")
		return ps, nil
	}

	ps.addPlayerToList(createdPlayer)
	ps.resetInput()
	ps.view.List.CursorToLastItem()
	return ps, nil
}

// addPlayerToList adds a player to the list
func (ps *PlayerSelection) addPlayerToList(player models.Player) {
	ps.players = append(ps.players, player)
	newItem := cl.Item{Name: player.Name}
	ps.view.List.AddItem(newItem)
}

func (ps *PlayerSelection) Init() tea.Cmd { return nil }

func (ps *PlayerSelection) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if ps.inInputMode {
		return ps.handlePlayerCreationUpdate(msg)
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if msg.Width != ps.size.Width || msg.Height != ps.size.Height {
			ps.size = msg
			ps.setSelectionView()
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, ps.view.SelectionControls().Back) && !ps.view.List.IsFiltering():
			return ps, models.ChangeScreen(models.MainMenuScreen)
		}
	}

	// update the SelectionView and return its command
	var cmd tea.Cmd
	_, cmd = ps.view.Update(msg)
	return ps, cmd
}

func (ps *PlayerSelection) View() string {
	content := []string{
		ps.view.View(),
		ps.view.Help.ShortHelpView(ps.view.InsertControls.InputHelp()),
	}

	if ps.inInputMode {
		return lipgloss.Place(
			ps.size.Width,
			ps.size.Height,
			lipgloss.Center,
			lipgloss.Center,
			lipgloss.JoinVertical(lipgloss.Center, content...),
		)
	}

	// extra padding for the list if not in input mode
	return ps.view.View() + "\n"
}
