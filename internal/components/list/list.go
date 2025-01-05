package list

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// List is a wrapper around the bubbles list.Model
type List struct {
	Model list.Model
}

// NewList creates a new bubbles list.Model with the given items, size and options
func NewList(items []Item, width, height int, opts ...Option) *List {
	delegate := list.NewDefaultDelegate()
	listItems := make([]list.Item, len(items))
	// need to convert the items to list.Item to be able to use the bubbles list.Model
	for i, item := range items {
		listItems[i] = item
	}
	l := list.New(listItems, delegate, width, height)
	// apply the options
	for _, opt := range opts {
		opt(&l, &delegate)
	}
	// need to set the delegate again after applying the options
	l.SetDelegate(delegate)
	return &List{Model: l}
}

// SelectedItem returns the current selected item in the list
func (l *List) SelectedItem() list.Item {
	return l.Model.SelectedItem()
}

// AddItem adds an item to the list
func (l *List) AddItem(item list.Item) {
	items := l.Model.Items()
	l.Model.SetItems(append(items, item))
}

// CursorToLastItem moves the cursor to the last item in the list
func (l *List) CursorToLastItem() {
	last := len(l.Model.Items()) - 1
	l.Model.Select(last)
}

// IsFiltering returns true if the list is currently in the filtering state
func (l *List) IsFiltering() bool {
	return l.Model.FilterState() == list.Filtering
}

func (l *List) Init() tea.Cmd { return nil }

func (l *List) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	l.Model, cmd = l.Model.Update(msg)
	return l, cmd
}

func (l *List) View() string {
	return l.Model.View()
}

// Item is a wrapper around the bubbles list.Item
type Item struct {
	Name    string
	Details string
	Number  int
}

func (i Item) Title() string       { return i.Name }
func (i Item) Description() string { return i.Details }
func (i Item) FilterValue() string { return i.Name }
