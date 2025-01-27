package controllers

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dasvh/go-learn-vim/internal/models"
	"github.com/dasvh/go-learn-vim/internal/storage"
	"github.com/dasvh/go-learn-vim/internal/testutils"
	"testing"
)

func Test_CreatePlayer(t *testing.T) {
	type fields struct {
		repo          *testutils.MockGameRepository
		currentPlayer *models.Player
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Create new player successfully",
			fields: fields{
				repo: testutils.NewMockGameRepository(),
			},
			args:    args{name: "Alice"},
			wantErr: false,
		},
		{
			name: "Duplicate player name",
			fields: fields{
				repo: testutils.NewMockGameRepositoryWithData(
					[]models.Player{{ID: "1", Name: "Alice"}}, nil,
				),
			},
			args:    args{name: "Alice"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gc := &Game{
				repo: tt.fields.repo,
			}
			got, err := gc.CreatePlayer(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePlayer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if got.Name != tt.args.name {
					t.Errorf("expected player name %q, got %q", tt.args.name, got.Name)
				}

				players, _ := tt.fields.repo.Players()
				if len(players) != 1 {
					t.Errorf("expected 1 player in the repository, got %d", len(players))
				}
			}
		})
	}
}

func Test_Players(t *testing.T) {
	tests := []struct {
		name      string
		setupRepo func() storage.GameRepository
		wantCount int
		wantErr   bool
	}{
		{
			name: "No players in repository",
			setupRepo: func() storage.GameRepository {
				return testutils.NewMockGameRepository()
			},
			wantCount: 0,
			wantErr:   false,
		},
		{
			name: "Multiple players in repository",
			setupRepo: func() storage.GameRepository {
				return testutils.NewMockGameRepositoryWithData(
					[]models.Player{
						{ID: "1", Name: "Alice"},
						{ID: "2", Name: "Bob"},
					}, nil,
				)
			},
			wantCount: 2,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame(tt.setupRepo())

			players, err := game.Players()
			if (err != nil) != tt.wantErr {
				t.Errorf("Players() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(players) != tt.wantCount {
				t.Errorf("expected %d players, got %d", tt.wantCount, len(players))
			}
		})
	}
}

func Test_SetPlayer(t *testing.T) {
	player := models.Player{ID: "1", Name: "Alice"}
	game := NewGame(testutils.NewMockGameRepository())

	game.SetPlayer(player)

	if game.currentPlayer == nil || game.currentPlayer.Name != "Alice" {
		t.Errorf("SetPlayer() did not set the current player correctly, got %v", game.currentPlayer)
	}
}

func Test_SaveGame(t *testing.T) {
	tests := []struct {
		name      string
		setup     func() *Game
		mode      string
		gameState models.GameState
		saveID    string
		wantErr   bool
		validate  func(t *testing.T, repo *testutils.MockGameRepository)
	}{
		{
			name: "Save game without selecting a player",
			setup: func() *Game {
				return NewGame(testutils.NewMockGameRepository())
			},
			mode:      "adventure",
			gameState: testGameState,
			saveID:    "",
			wantErr:   true,
		},
		{
			name: "Save game with auto-generated saveID",
			setup: func() *Game {
				game := NewGame(testutils.NewMockGameRepository())
				game.SetPlayer(models.Player{ID: "1", Name: "Alice"})
				return game
			},
			mode:      "adventure",
			gameState: testGameState,
			saveID:    "",
			wantErr:   false,
			validate: func(t *testing.T, repo *testutils.MockGameRepository) {
				saves := repo.GameSavesData
				if len(saves) != 1 {
					t.Fatalf("expected 1 save, got %d", len(saves))
				}
				if saves[0].ID == "" {
					t.Errorf("expected auto-generated save ID, got empty ID")
				}
			},
		},
		{
			name: "Save game with existing saveID",
			setup: func() *Game {
				game := NewGame(testutils.NewMockGameRepositoryWithData(nil, []models.GameSave{
					{ID: "existing-id", Player: models.Player{ID: "1", Name: "Alice"}},
				}))
				game.SetPlayer(models.Player{ID: "1", Name: "Alice"})
				return game
			},
			mode:      "adventure",
			gameState: testGameState,
			saveID:    "existing-id",
			wantErr:   false,
			validate: func(t *testing.T, repo *testutils.MockGameRepository) {
				saves := repo.GameSavesData
				if len(saves) != 1 {
					t.Fatalf("expected 1 save, got %d", len(saves))
				}
				if saves[0].ID != "existing-id" {
					t.Errorf("expected save ID 'existing-id', got %q", saves[0].ID)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := tt.setup()
			err := game.SaveGame(tt.mode, tt.gameState, tt.saveID)

			if (err != nil) != tt.wantErr {
				t.Errorf("SaveGame() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

var testGameState = models.AdventureGameState{
	WindowSize: tea.WindowSizeMsg{
		Width:  140,
		Height: 75,
	},
	Level: models.SavedLevel{
		Number: 1,
		Width:  138,
		Height: 68,
		PlayerPosition: models.Position{
			X: 4,
			Y: 10,
		},
		Targets:       make([]models.Target, 0),
		CurrentTarget: 2,
		Completed:     false,
		InProgress:    true,
	},
}
