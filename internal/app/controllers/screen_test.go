package controllers

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/testutils"
	"testing"

	"github.com/dasvh/go-learn-vim/internal/models"
)

func Test_NewScreen(t *testing.T) {
	sc := NewScreen()

	if sc.ActiveScreen() != models.MainMenuScreen {
		t.Errorf("expected initial screen to be MainMenuScreen, got %v", sc.ActiveScreen())
	}

	if len(sc.Screens()) != 0 {
		t.Errorf("expected no screens to be registered, got %d", len(sc.Screens()))
	}
}

func Test_Register(t *testing.T) {
	sc := NewScreen()
	mockScreen := &testutils.MockTeaModel{}

	sc.Register(models.MainMenuScreen, mockScreen)

	if len(sc.Screens()) != 1 {
		t.Errorf("expected 1 screen to be registered, got %d", len(sc.Screens()))
	}

	if sc.Screens()[models.MainMenuScreen] != mockScreen {
		t.Errorf("expected registered screen to be MainMenuScreen")
	}
}

func Test_SwitchTo(t *testing.T) {
	tests := []struct {
		name         string
		setup        func() (*Screen, *testutils.MockTeaModel)
		switchTo     models.Screen
		wantCmd      tea.Cmd
		wantedScreen models.Screen
		wantInitCall bool
	}{
		{
			name: "Switch to an existing screen",
			setup: func() (*Screen, *testutils.MockTeaModel) {
				sc := NewScreen()
				mockScreen := &testutils.MockTeaModel{}
				sc.Register(models.MainMenuScreen, mockScreen)
				return sc, mockScreen
			},
			switchTo:     models.MainMenuScreen,
			wantCmd:      nil,
			wantedScreen: models.MainMenuScreen,
			wantInitCall: true,
		},
		{
			name: "Switch to a non-existent screen",
			setup: func() (*Screen, *testutils.MockTeaModel) {
				return NewScreen(), nil
			},
			switchTo:     models.AdventureModeScreen,
			wantCmd:      nil,
			wantedScreen: models.MainMenuScreen,
			wantInitCall: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc, mockScreen := tt.setup()
			cmd := sc.SwitchTo(tt.switchTo)

			if (cmd == nil) != (tt.wantCmd == nil) {
				t.Errorf("SwitchTo() = %v, want %v", cmd, tt.wantCmd)
			}

			if sc.ActiveScreen() != tt.wantedScreen {
				t.Errorf("expected active screen to be %v, got %v", tt.wantedScreen, sc.ActiveScreen())
			}

			if mockScreen != nil && mockScreen.InitCalled != tt.wantInitCall {
				t.Errorf("expected InitCalled to be %v, got %v", tt.wantInitCall, mockScreen.InitCalled)
			}
		})
	}
}

func Test_CurrentScreen(t *testing.T) {
	sc := NewScreen()
	mockScreen := &testutils.MockTeaModel{}

	sc.Register(models.MainMenuScreen, mockScreen)
	sc.SwitchTo(models.MainMenuScreen)

	if sc.CurrentScreen() != mockScreen {
		t.Errorf("expected CurrentScreen to return the main menu screen, got %v", sc.CurrentScreen())
	}
}

func Test_Screens(t *testing.T) {
	sc := NewScreen()
	mockScreen := &testutils.MockTeaModel{}
	sc.Register(models.MainMenuScreen, mockScreen)

	screens := sc.Screens()
	if len(screens) != 1 || screens[models.MainMenuScreen] != mockScreen {
		t.Errorf("Screens() returned incorrect map: %v", screens)
	}
}
