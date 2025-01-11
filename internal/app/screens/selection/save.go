package selection

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/models"
	"github.com/dasvh/go-learn-vim/internal/storage"
	"github.com/dasvh/go-learn-vim/internal/views"
	"strconv"
)

// SaveSelection represents the screen model for the Saves screen
type SaveSelection struct {
	view         *views.TableView
	saves        []models.GameSave
	error        error
	repo         storage.GameRepository
	onSaveSelect func(save models.GameSave) tea.Cmd
}

// NewSaveSelection creates a new SaveSelection screen model
func NewSaveSelection(repo storage.GameRepository, onSaveSelect func(save models.GameSave) tea.Cmd) *SaveSelection {
	ss := views.NewTableView("Load Game Saves")

	ss.SetColumns([]table.Column{
		{Title: "", Width: 1},
		{Title: "Player", Width: models.PlayerNameMaxLength},
		{Title: "Mode", Width: 9},
		{Title: "Level", Width: 5},
		{Title: "Date", Width: 20},
	})

	// call handleSelection on onSelect
	saveSelection := &SaveSelection{
		view:         ss,
		repo:         repo,
		saves:        make([]models.GameSave, 0),
		onSaveSelect: onSaveSelect,
	}

	ss.SetOnSelect(saveSelection.handleSelection)
	return saveSelection
}

// handleSelection handles the selection of a save
func (ss *SaveSelection) handleSelection(index int) tea.Cmd {
	if index < 0 || index >= len(ss.saves) {
		fmt.Println("Invalid row index:", index)
		return nil
	}

	if ss.onSaveSelect != nil {
		return ss.onSaveSelect(ss.saves[index])
	}
	return nil
}

// Init initializes the SaveSelection screen model and populates it with data
func (ss *SaveSelection) Init() tea.Cmd {
	return func() tea.Msg {
		saves, err := ss.repo.IncompleteGames()
		if err != nil {
			return saveSelectionError{err}
		}
		return saveSelectionData{Saves: saves}
	}
}

// Update updates the SaveSelection screen model
func (ss *SaveSelection) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := ss.view.Update(msg)
	switch msg := msg.(type) {
	case saveSelectionData:
		ss.saves = msg.Saves
		ss.populateTable()
		return ss, nil
	case saveSelectionError:
		ss.error = msg.error
		return ss, nil
	}

	return ss, cmd
}

// View returns the view for the SaveSelection screen model
func (ss *SaveSelection) View() string {
	if ss.error != nil {
		return fmt.Sprintf("Error: %v\nPress 'q' to quit.", ss.error)
	}
	return ss.view.View()
}

// populateTable populates the table with data
func (ss *SaveSelection) populateTable() {
	rows := make([]table.Row, len(ss.saves))
	for i, save := range ss.saves {
		rows[i] = table.Row{
			strconv.Itoa(i),
			save.Player.Name,
			save.GameMode,
			strconv.Itoa(save.GameState.(models.AdventureGameState).Level.Number),
			save.Timestamp.Format("2006-01-02 15:04:05"),
		}
	}
	ss.view.SetRows(rows)
}

// saveSelectionData is a message that contains the saves data
type saveSelectionData struct {
	Saves []models.GameSave
}

// saveSelectionError is a message that contains an error
type saveSelectionError struct {
	error
}
