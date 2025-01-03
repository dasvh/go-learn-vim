package stats

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/storage"
	"github.com/dasvh/go-learn-vim/internal/views"
	"sort"
	"strconv"
)

// StatsScreen represents the screen model for the Stats screen
type StatsScreen struct {
	view          *views.TableView
	lifetimeStats *storage.LifetimeStats
	playerStats   map[string]*storage.LifetimeStats
	error         error
	repo          storage.GameRepository
}

// NewStatsScreen creates a new StatsScreen screen model
func NewStatsScreen(repo storage.GameRepository) *StatsScreen {
	tv := views.NewTableView("Stats")

	tv.SetColumns([]table.Column{
		{Title: "Player", Width: 20},
		{Title: "Games Played", Width: 15},
		{Title: "Keystrokes", Width: 12},
		{Title: "Playtime (s)", Width: 12},
		{Title: "Key Presses", Width: 30},
	})

	return &StatsScreen{
		view:        tv,
		playerStats: make(map[string]*storage.LifetimeStats),
		repo:        repo,
	}
}

// Init initializes the StatsScreen screen model and populates it with data
func (sv *StatsScreen) Init() tea.Cmd {
	return func() tea.Msg {
		lifetimeStats, err := sv.repo.LifetimeStats()
		if err != nil {
			return statsScreenError{err}
		}

		playerStats := make(map[string]*storage.LifetimeStats)
		players, err := sv.repo.Players()
		if err == nil {
			for _, p := range players {
				stats, _ := sv.repo.PlayerLifetimeStats(p.ID)
				playerStats[p.Name] = stats
			}
		}

		return statsScreenData{
			LifetimeStats: lifetimeStats,
			PlayerStats:   playerStats,
		}
	}
}

// Update processes input and updates the table
func (sv *StatsScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := sv.view.Update(msg)

	switch msg := msg.(type) {
	case statsScreenData:
		sv.lifetimeStats = msg.LifetimeStats
		sv.playerStats = msg.PlayerStats
		sv.populateTable()
		return sv, nil
	case statsScreenError:
		sv.error = msg.Err
		return sv, nil
	}

	return sv, cmd
}

// View renders the TableView
func (sv *StatsScreen) View() string {
	if sv.error != nil {
		return fmt.Sprintf("Error: %v\nPress 'q' to quit.", sv.error)
	}
	return sv.view.View()
}

// populateTable fills the TableView with stats data
func (sv *StatsScreen) populateTable() {
	rows := []table.Row{
		{
			"Global",
			strconv.Itoa(sv.lifetimeStats.TotalGames),
			strconv.Itoa(sv.lifetimeStats.TotalKeystrokes),
			strconv.Itoa(sv.lifetimeStats.TotalPlaytime),
			formatKeyPresses(sv.lifetimeStats.KeyPresses),
		},
	}
	for player, stats := range sv.playerStats {
		rows = append(rows, table.Row{
			player,
			strconv.Itoa(stats.TotalGames),
			strconv.Itoa(stats.TotalKeystrokes),
			strconv.Itoa(stats.TotalPlaytime),
			formatKeyPresses(stats.KeyPresses),
		})
	}

	sv.view.SetRows(rows)
}

// formatKeyPresses formats key presses as a sorted string
func formatKeyPresses(keyPresses map[string]int) string {
	if len(keyPresses) == 0 {
		return "None"
	}
	keys := make([]string, 0, len(keyPresses))
	for k := range keyPresses {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	result := ""
	for _, k := range keys {
		result += fmt.Sprintf("%s: %d, ", k, keyPresses[k])
	}
	return result[:len(result)-2] // Remove trailing ", "
}

// statsScreenData holds the stats data for updating the screen
type statsScreenData struct {
	LifetimeStats *storage.LifetimeStats
	PlayerStats   map[string]*storage.LifetimeStats
}

// statsScreenError represents an error encountered in StatsScreen
type statsScreenError struct {
	Err error
}
