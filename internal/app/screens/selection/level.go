package selection

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/app/controllers"
	cl "github.com/dasvh/go-learn-vim/internal/components/list"
	"github.com/dasvh/go-learn-vim/internal/models"
	"github.com/dasvh/go-learn-vim/internal/style"
	"github.com/dasvh/go-learn-vim/internal/views"
	"strconv"
)

// LevelSelection is a screen that allows the user to select a level
type LevelSelection struct {
	view  *views.SelectionView
	size  tea.WindowSizeMsg
	lc    *controllers.Level
	items []cl.Item
}

// NewLevelSelection creates a new LevelSelection screen
func NewLevelSelection(lc *controllers.Level) *LevelSelection {
	levels := lc.GetLevels()
	items := make([]cl.Item, len(levels))

	for i, level := range levels {
		n := level.Number()
		items[i] = cl.Item{
			Name:    "Level " + strconv.Itoa(n),
			Details: level.Description(),
			Number:  5,
		}
	}
	return &LevelSelection{
		lc:    lc,
		items: items,
	}
}

// setSelectionView sets the view of the level selection screen
func (ls *LevelSelection) setSelectionView() {

	width := ls.size.Width
	height := ls.size.Height / 2
	levelList := cl.NewList(ls.items, width, height,
		cl.WithItemsIdentifiers("Select a level to play", "level", "levels"),
		cl.WithShowDescription(true),
		cl.WithDisableQuitKeybindings(),
		cl.WithTitleStyle(style.LevelSelection.Title),
		cl.WithSelectedTitleStyle(style.LevelSelection.SelectedItem),
		cl.WithNormalTitleStyle(style.LevelSelection.Item),
		cl.WithFilterMatchStyle(style.LevelSelection.FilterMatch),
		cl.WithDimmedTitleStyle(style.LevelSelection.DimmedItem),
	)

	ls.view = views.NewSelectionView(
		"Level Selection",
		levelList,
		nil,
		ls.handleSelect,
		nil,
	)
}

// handleSelect handles the selection of a level
func (ls *LevelSelection) handleSelect(item cl.Item) tea.Cmd {
	if lvl, ok := ls.lc.GetLevels()[item.Number]; ok {
		ls.lc.SetLevel(lvl)
	} else {
		return nil // todo: return batch with an error msg
	}
	return tea.Batch(
		models.ChangeScreen(models.AdventureModeScreen),
		func() tea.Msg {
			return models.SetLevelMsg{
				LevelNumber: ls.lc.GetLevelNumber()}
		},
	)
}

func (ls *LevelSelection) Init() tea.Cmd { return nil }

func (ls *LevelSelection) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if msg.Width != ls.size.Width || msg.Height != ls.size.Height {
			ls.size = msg
			ls.setSelectionView()
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, ls.view.SelectionControls().Back) && !ls.view.List.IsFiltering():
			return ls, models.ChangeScreen(models.NewGameScreen)
		}
	}

	ls.view.List.Model.Update(msg)

	var cmd tea.Cmd
	_, cmd = ls.view.Update(msg)
	return ls, cmd
}

func (ls *LevelSelection) View() string {
	return ls.view.View()
}
