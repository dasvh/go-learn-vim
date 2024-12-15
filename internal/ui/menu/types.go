package menu

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Menu represents a UI menu interface that combines the tea.Model and ButtonHandler interfaces
type Menu interface {
	tea.Model
	ButtonHandler
}

// ButtonHandler handles button-related operations
type ButtonHandler interface {
	// HandleSelection handles the selection of a button
	HandleSelection() tea.Cmd
}

// ButtonConfig represents the configuration for a button in the UI menu
type ButtonConfig struct {
	// Label represents the text displayed on the button
	Label string
	// Inactive determines if the button is rendered
	Inactive bool
}
