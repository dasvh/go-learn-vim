package views

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dasvh/go-learn-vim/internal/components"
	ct "github.com/dasvh/go-learn-vim/internal/components/table"
	"github.com/dasvh/go-learn-vim/internal/models"
	"github.com/dasvh/go-learn-vim/internal/style"
	"strconv"
)

// TableView represents a view that displays a table using the bubbles table component
type TableView struct {
	title         *components.Title
	size          tea.WindowSizeMsg
	tableControls ct.Controls
	help          help.Model
	table         table.Model
	onSelect      func(row int) tea.Cmd
}

// NewTableView creates a new TableView
func NewTableView(subtitle string) *TableView {
	return &TableView{
		title: components.NewTitle(
			style.Styles.Title.Render(mainTitle), style.Styles.Subtitle.Render(subtitle),
		),
		help:          help.New(),
		tableControls: ct.NewTableControls(),
	}
}

// SetOnSelect sets the onSelect function
func (tv *TableView) SetOnSelect(onSelect func(index int) tea.Cmd) {
	tv.onSelect = onSelect
}

// SetColumns sets the columns of the TableView
func (tv *TableView) SetColumns(columns []table.Column) {
	tv.table = table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithStyles(style.Table),
	)
}

// SetRows sets the rows of the TableView
func (tv *TableView) SetRows(rows []table.Row) {
	tv.table.SetRows(rows)
}

func (tv *TableView) Init() tea.Cmd { return nil }

func (tv *TableView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if msg.Width != tv.size.Height || msg.Height != tv.size.Height {
			tv.size = msg
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, tv.tableControls.Up):
			tv.table.MoveUp(1)
		case key.Matches(msg, tv.tableControls.Down):
			tv.table.MoveDown(1)
		case key.Matches(msg, tv.tableControls.GotoTop):
			tv.table.GotoTop()
		case key.Matches(msg, tv.tableControls.GotoBottom):
			tv.table.GotoBottom()
		case key.Matches(msg, tv.tableControls.Select) && tv.onSelect != nil:
			row := tv.table.SelectedRow()
			index, err := strconv.Atoi(row[0])
			if err != nil {
				fmt.Println("Error converting row to int:", err)
				return tv, nil
			}
			return tv, tv.onSelect(index)
		case key.Matches(msg, tv.tableControls.Back):
			return tv, models.ChangeScreen(models.MainMenuScreen)
		case key.Matches(msg, tv.tableControls.Quit):
			return tv, tea.Quit
		}
	}

	return tv, nil
}

func (tv *TableView) View() string {
	views := []string{
		tv.title.Main,
		tv.title.Subtitle,
		tv.table.View(),
		tv.help.ShortHelpView(tv.tableControls.ShortHelp()),
	}

	return lipgloss.Place(tv.size.Width, tv.size.Height, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center, views...))
}
