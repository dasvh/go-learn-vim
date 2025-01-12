package leaderboards

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/models"
	"github.com/dasvh/go-learn-vim/internal/storage"
	"github.com/dasvh/go-learn-vim/internal/views"
	"strconv"
)

// ScoresScreen represents the screen model for the Scores screen
type ScoresScreen struct {
	view       *views.TableView
	highScores []models.HighScore
	error      error
	repo       storage.GameRepository
}

// NewScoresScreen creates a new ScoresScreen screen model
func NewScoresScreen(repo storage.GameRepository) *ScoresScreen {
	tv := views.NewTableView("High Scores")

	tv.SetColumns([]table.Column{
		{Title: "", Width: 3},
		{Title: "Player", Width: models.PlayerNameMaxLength},
		{Title: "Level", Width: 5},
		{Title: "Score", Width: 15},
		{Title: "Date", Width: 20},
	})

	return &ScoresScreen{
		view:       tv,
		repo:       repo,
		highScores: make([]models.HighScore, 0),
	}
}

// Init initializes the ScoresScreen screen model and populates it with data
func (ss *ScoresScreen) Init() tea.Cmd {
	return func() tea.Msg {
		highScores, err := ss.repo.ComputeHighScores()
		if err != nil {
			return scoresScreenError{err}
		}

		return scoresScreenData{HighScores: highScores}
	}
}

// Update updates the ScoresScreen screen model
func (ss *ScoresScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := ss.view.Update(msg)

	switch msg := msg.(type) {
	case scoresScreenData:
		ss.highScores = msg.HighScores
		ss.populateTable()
		return ss, nil
	case scoresScreenError:
		ss.error = msg.error
		return ss, nil
	}

	return ss, cmd
}

// View returns the view for the ScoresScreen screen model
func (ss *ScoresScreen) View() string {
	if ss.error != nil {
		return fmt.Sprintf("Error: %v\nPress 'q' to quit.", ss.error)
	}
	return ss.view.View()
}

// populateTable populates the table with high scores
func (ss *ScoresScreen) populateTable() {
	rows := make([]table.Row, 0)

	if len(ss.highScores) == 0 {
		rows = append(rows, table.Row{"", "No scores yet", "", "", ""})
	} else {
		for place, hs := range ss.highScores {
			rows = append(rows, table.Row{
				strconv.Itoa(place + 1),
				hs.PlayerName,
				strconv.Itoa(hs.Level),
				strconv.Itoa(hs.Score),
				hs.Timestamp.Format("2006-01-02 15:04"),
			})
		}
	}

	ss.view.SetRows(rows)
}

// scoresScreenData is a message that contains the data for the ScoresScreen screen model
type scoresScreenData struct {
	HighScores []models.HighScore
}

// scoresScreenError is a message that contains an error for the ScoresScreen screen model
type scoresScreenError struct {
	error
}
