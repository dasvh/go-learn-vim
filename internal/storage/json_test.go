package storage

import (
	"fmt"
	"github.com/dasvh/go-learn-vim/internal/models"
	"os"
	"testing"
	"time"
)

func Test_JSONRepository_ComputeHighScores(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test_repo_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	initialJSON := `{"players":[],"saves":[]}`
	if _, err := tempFile.Write([]byte(initialJSON)); err != nil {
		t.Fatalf("Failed to write initial JSON: %v", err)
	}
	tempFile.Close()

	repo, err := NewJSONRepository(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}

	player1 := models.Player{ID: "p1", Name: "Player 1"}
	player2 := models.Player{ID: "p2", Name: "Player 2"}
	player3 := models.Player{ID: "p3", Name: "Player 3"}
	player4 := models.Player{ID: "p3", Name: "Player 4"}

	game1 := createTestGameSaveWithID(player1, "g1", 1, 50, 100, true)  // 14450
	game2 := createTestGameSaveWithID(player2, "g2", 1, 80, 150, true)  // 8290
	game3 := createTestGameSaveWithID(player3, "g3", 1, 30, 70, true)   // 18500
	game4 := createTestGameSaveWithID(player4, "g4", 1, 40, 40, false)  // not completed
	game5 := createTestGameSaveWithID(player4, "g5", 0, 150, 100, true) // -3250 -> 5000
	game6 := createTestGameSaveWithID(player1, "g6", 1, 40, 90, true)   // 16390

	repo.SaveGame(game1)
	repo.SaveGame(game2)
	repo.SaveGame(game3)
	repo.SaveGame(game4)
	repo.SaveGame(game5)
	repo.SaveGame(game6)

	scores, err := repo.ComputeHighScores()

	fmt.Println(scores)
	if err != nil {
		t.Fatalf("Failed to compute high scores: %v", err)
	}

	if len(scores) != 5 {
		t.Errorf("Expected 4 high scores, got %d", len(scores))
	}

	if scores[0].PlayerName != "Player 3" {
		t.Errorf("Expected Player 3 to be top score, got %s", scores[0].PlayerName)
	}

	if scores[1].PlayerName != "Player 1" {
		t.Errorf("Expected Player 1 to be second, got %s", scores[1].PlayerName)
	}

	if scores[2].PlayerName != "Player 1" {
		t.Errorf("Expected Player 1 to be third, got %s", scores[2].PlayerName)
	}

	if scores[3].PlayerName != "Player 2" {
		t.Errorf("Expected Player 2 to be fourth, got %s", scores[3].PlayerName)
	}

	if scores[4].PlayerName != "Player 4" {
		t.Errorf("Expected Player 4 to be fourth, got %s", scores[4].PlayerName)
	}

	if len(scores) >= 2 {
		expectedScore := 25000 - (30 * 177) - (70 * 17)
		if scores[0].Score != expectedScore {
			t.Errorf("Expected score %d for %s, got %d", expectedScore, scores[0].PlayerName, scores[0].Score)
		}

		expectedMinScore := 5000
		if scores[4].Score != expectedMinScore {
			t.Errorf("Expected score %d for %s, got %d", expectedMinScore, scores[4].PlayerName, scores[4].Score)
		}
	}
}

func createTestGameSaveWithID(player models.Player, id string, level int,
	timeElapsed int, keystrokes int, completed bool) models.GameSave {
	stats := models.Stats{
		TimeElapsed:     timeElapsed,
		TotalKeystrokes: keystrokes,
		KeyPresses:      make(map[string]int),
	}

	gameState := models.AdventureGameState{
		Level: models.SavedLevel{
			Number:    level,
			Completed: completed,
		},
		Stats: stats,
	}

	return models.GameSave{
		ID:        id,
		Player:    player,
		GameMode:  "Adventure",
		GameState: gameState,
		Timestamp: time.Now(),
	}
}
