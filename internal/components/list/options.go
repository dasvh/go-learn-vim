package list

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

// Option is a function that modifies the list.Model and list.DefaultDelegate.
// Can be used to set various options on the list such as title, styles, etc.
type Option func(*list.Model, *list.DefaultDelegate)

// WithItemsIdentifiers sets the title and status bar item names of the list
func WithItemsIdentifiers(title, singular, plural string) Option {
	return func(l *list.Model, _ *list.DefaultDelegate) {
		l.Title = title
		l.SetStatusBarItemName(singular, plural)
	}
}

// WithTitleStyle sets the title style of the list.
func WithTitleStyle(style lipgloss.Style) Option {
	return func(l *list.Model, _ *list.DefaultDelegate) {
		l.Styles.Title = style
	}
}

// WithDisableQuitKeybindings disables the quit keybindings in the list
func WithDisableQuitKeybindings() Option {
	return func(l *list.Model, _ *list.DefaultDelegate) {
		l.DisableQuitKeybindings()
	}
}

// WithShowDescription sets whether the list should show the description of the items.
func WithShowDescription(show bool) Option {
	return func(l *list.Model, d *list.DefaultDelegate) {
		d.ShowDescription = show
	}
}

// WithNormalTitleStyle sets the normal title style of the list delegate.
func WithNormalTitleStyle(style lipgloss.Style) Option {
	return func(l *list.Model, d *list.DefaultDelegate) {
		d.Styles.NormalTitle = style
	}
}

// WithNormalDescriptionStyle sets the normal description style of the list delegate.
func WithNormalDescriptionStyle(style lipgloss.Style) Option {
	return func(l *list.Model, d *list.DefaultDelegate) {
		d.Styles.NormalDesc = style
	}
}

// WithSelectedTitleStyle sets the selected title style of the list delegate.
func WithSelectedTitleStyle(style lipgloss.Style) Option {
	return func(l *list.Model, d *list.DefaultDelegate) {
		d.Styles.SelectedTitle = style
	}
}

// WithSelectedDescriptionStyle sets the selected description style of the list delegate.
func WithSelectedDescriptionStyle(style lipgloss.Style) Option {
	return func(l *list.Model, d *list.DefaultDelegate) {
		d.Styles.SelectedDesc = style
	}
}

// WithDimmedTitleStyle sets the dimmed title style of the list delegate.
func WithDimmedTitleStyle(style lipgloss.Style) Option {
	return func(l *list.Model, d *list.DefaultDelegate) {
		d.Styles.DimmedTitle = style
	}
}

// WithFilterMatchStyle sets the filter match style of the list delegate.
func WithFilterMatchStyle(style lipgloss.Style) Option {
	return func(l *list.Model, d *list.DefaultDelegate) {
		d.Styles.FilterMatch = style
	}
}
