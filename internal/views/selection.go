package views

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dasvh/go-learn-vim/internal/components"
	cl "github.com/dasvh/go-learn-vim/internal/components/list"
	"github.com/dasvh/go-learn-vim/internal/style"
)

// SelectionView is a view that displays a list of items and allows the user to
// select one of them. It also allows the user to insert a new item.
type SelectionView struct {
	title          *components.Title
	size           tea.WindowSizeMsg
	controls       cl.SelectionControls
	InsertControls cl.InsertControls
	Help           help.Model
	TextInput      *textinput.Model
	List           *cl.List
	onSelect       func(item cl.Item) tea.Cmd
	onInsert       func() tea.Cmd
}

// NewSelectionView creates a new SelectionView
func NewSelectionView(
	subtitle string,
	list *cl.List,
	textInput *textinput.Model,
	onSelect func(item cl.Item) tea.Cmd,
	onInsert func() tea.Cmd,
) *SelectionView {
	return &SelectionView{
		title: components.NewTitle(
			style.Styles.Title.Render(mainTitle), style.Styles.Subtitle.Render(subtitle),
		),
		controls:       cl.NewSelectionControls(),
		InsertControls: cl.NewInsertControls(),
		Help:           help.New(),
		TextInput:      textInput,
		List:           list,
		onSelect:       onSelect,
		onInsert:       onInsert,
	}
}

// SelectionControls returns the selection controls
func (sv *SelectionView) SelectionControls() cl.SelectionControls { return sv.controls }

func (sv *SelectionView) Init() tea.Cmd { return nil }

func (sv *SelectionView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		sv.size = msg
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, sv.controls.Select) && !sv.List.IsFiltering():
			selected := sv.List.SelectedItem().(cl.Item)
			return sv, sv.onSelect(selected)
		case key.Matches(msg, sv.InsertControls.Insert) && !sv.List.IsFiltering():
			if sv.onInsert != nil {
				return sv, sv.onInsert()
			}
		case key.Matches(msg, sv.controls.Quit) && !sv.List.IsFiltering():
			return sv, tea.Quit
		}
	}

	// todo: figure out how to fix the view after applying a filter

	var cmd tea.Cmd
	sv.List.Model, cmd = sv.List.Model.Update(msg)
	return sv, cmd
}

func (sv *SelectionView) View() string {
	views := []string{
		sv.renderCentered(sv.title.Main),
		sv.renderCentered(sv.title.Subtitle),
		sv.renderCentered(sv.List.View()),
	}

	helpBindings := []key.Binding{
		sv.InsertControls.Select,
		sv.InsertControls.Insert,
		sv.controls.Quit,
	}

	views = append(views, sv.renderCentered(sv.Help.ShortHelpView(helpBindings)))

	if sv.TextInput != nil {
		views = append(views, "\n"+sv.renderCentered(sv.TextInput.View()))
	}

	return lipgloss.Place(
		sv.size.Width,
		sv.size.Height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center, views...),
	)
}

// renderCentered renders content centered
func (sv *SelectionView) renderCentered(content string) string {
	return lipgloss.NewStyle().
		Width(sv.size.Width).
		Align(lipgloss.Center).
		Render(content)
}
