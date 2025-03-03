package components

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/exp/teatest"
	"github.com/muesli/termenv"
	"io"
	"testing"
	"time"
)

func Test_NewButtons(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *Buttons
		layout   lipgloss.Position
		expected string
	}{
		{
			name: "single_button",
			setup: func() *Buttons {
				return NewButtons("Start")
			},
		},
		{
			name: "multiple_buttons",
			setup: func() *Buttons {
				return NewButtons("Start", "Options", "Quit")
			},
		},
		{
			name: "with_inactive_button",
			setup: func() *Buttons {
				buttons := NewButtons("Start", "Options", "Quit")
				buttons.UpdateButtonState("Options", true)
				return buttons
			},
		},
		{
			name: "all_buttons_inactive",
			setup: func() *Buttons {
				buttons := NewButtons("Start", "Options", "Quit")
				buttons.UpdateButtonState("Start", true)
				buttons.UpdateButtonState("Options", true)
				buttons.UpdateButtonState("Quit", true)
				return buttons
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buttons := tt.setup()
			tm := teatest.NewTestModel(t, buttons, teatest.WithInitialTermSize(100, 50))

			time.Sleep(50 * time.Millisecond)

			out := readOutput(t, tm.Output())
			teatest.RequireEqualOutput(t, out)
		})
	}
}

func Test_ButtonsStyles(t *testing.T) {
	tests := []struct {
		name         string
		setup        func() *Buttons
		colorProfile termenv.Profile
	}{
		{
			name: "custom_style",
			setup: func() *Buttons {
				buttons := NewButtons("Custom", "Style")
				buttons.styles.Selected = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#FF0000")).
					Bold(true)
				buttons.styles.Default = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#0000FF")).
					Italic(true)
				return buttons
			},
			colorProfile: termenv.TrueColor,
		},
		{
			name: "custom_layout",
			setup: func() *Buttons {
				buttons := NewButtons("Left", "Aligned")
				buttons.layout = lipgloss.Left
				return buttons
			},
			colorProfile: termenv.Ascii,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buttons := tt.setup()
			lipgloss.SetColorProfile(tt.colorProfile)
			tm := teatest.NewTestModel(t, buttons, teatest.WithInitialTermSize(100, 50))

			time.Sleep(50 * time.Millisecond)

			out := readOutput(t, tm.Output())
			teatest.RequireEqualOutput(t, out)
		})
	}
}

func Test_ButtonsNavigation(t *testing.T) {
	tests := []struct {
		name           string
		setup          func() *Buttons
		actions        []Move
		expectedCursor int
	}{
		{
			name: "move_down_once",
			setup: func() *Buttons {
				return NewButtons("First", "Second", "Third")
			},
			actions:        []Move{MoveDown},
			expectedCursor: 1,
		},
		{
			name: "move_down_and_up",
			setup: func() *Buttons {
				return NewButtons("First", "Second", "Third")
			},
			actions:        []Move{MoveDown, MoveDown, MoveUp},
			expectedCursor: 1,
		},
		{
			name: "top_boundary_respected",
			setup: func() *Buttons {
				return NewButtons("First", "Second", "Third")
			},
			actions:        []Move{MoveUp, MoveUp, MoveUp, MoveUp},
			expectedCursor: 0,
		},
		{
			name: "bottom_boundary_respected",
			setup: func() *Buttons {
				return NewButtons("First", "Second", "Third")
			},
			actions:        []Move{MoveDown, MoveDown, MoveDown, MoveDown},
			expectedCursor: 2,
		},
		{
			name: "skip_inactive_buttons_down",
			setup: func() *Buttons {
				buttons := NewButtons("First", "Second", "Third")
				buttons.UpdateButtonState("Second", true)
				return buttons
			},
			actions:        []Move{MoveDown},
			expectedCursor: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buttons := tt.setup()

			for _, action := range tt.actions {
				buttons.Update(action)
			}

			if buttons.cursor != tt.expectedCursor {
				t.Errorf("Expected cursor position %d, got %d", tt.expectedCursor, buttons.cursor)
			}
		})
	}
}

func Test_CurrentButton(t *testing.T) {
	tests := []struct {
		name           string
		setup          func() *Buttons
		actions        []Move
		expectedLabel  string
		expectedExists bool
	}{
		{
			name: "get_first_button",
			setup: func() *Buttons {
				return NewButtons("First", "Second", "Third")
			},
			actions:        []Move{},
			expectedLabel:  "First",
			expectedExists: true,
		},
		{
			name: "get_button_after_navigation",
			setup: func() *Buttons {
				return NewButtons("First", "Second", "Third")
			},
			actions:        []Move{MoveDown, MoveDown},
			expectedLabel:  "Third",
			expectedExists: true,
		},
		{
			name: "empty_buttons",
			setup: func() *Buttons {
				return NewButtons()
			},
			actions:        []Move{},
			expectedExists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buttons := tt.setup()

			for _, action := range tt.actions {
				buttons.Update(action)
			}

			current := buttons.CurrentButton()

			if tt.expectedExists && current == nil {
				t.Fatalf("Expected a current button but got nil")
			}

			if !tt.expectedExists && current != nil {
				t.Fatalf("Expected no current button but got %v", current)
			}

			if tt.expectedExists && current.Label != tt.expectedLabel {
				t.Errorf("Expected button label %q, got %q", tt.expectedLabel, current.Label)
			}
		})
	}
}

func Test_UpdateButtonState(t *testing.T) {
	tests := []struct {
		name           string
		setup          func() *Buttons
		updateLabel    string
		updateInactive bool
		checkLabel     string
		expectedState  bool
	}{
		{
			name: "make_button_inactive",
			setup: func() *Buttons {
				return NewButtons("First", "Second", "Third")
			},
			updateLabel:    "Second",
			updateInactive: true,
			checkLabel:     "Second",
			expectedState:  true,
		},
		{
			name: "make_inactive_button_active",
			setup: func() *Buttons {
				b := NewButtons("First", "Second", "Third")
				b.UpdateButtonState("Second", true)
				return b
			},
			updateLabel:    "Second",
			updateInactive: false,
			checkLabel:     "Second",
			expectedState:  false,
		},
		{
			name: "update_nonexistent_button",
			setup: func() *Buttons {
				return NewButtons("First", "Second", "Third")
			},
			updateLabel:    "Fourth",
			updateInactive: true,
			checkLabel:     "First",
			expectedState:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buttons := tt.setup()

			buttons.UpdateButtonState(tt.updateLabel, tt.updateInactive)

			var found bool
			var actualState bool

			for _, b := range buttons.Items() {
				if b.Label == tt.checkLabel {
					found = true
					actualState = b.Inactive
					break
				}
			}

			if !found {
				t.Fatalf("Button with label %q not found", tt.checkLabel)
			}

			if actualState != tt.expectedState {
				t.Errorf("Expected button %q inactive state to be %v, got %v",
					tt.checkLabel, tt.expectedState, actualState)
			}
		})
	}
}

func readOutput(t *testing.T, r io.Reader) []byte {
	t.Helper()
	data, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	return data
}
