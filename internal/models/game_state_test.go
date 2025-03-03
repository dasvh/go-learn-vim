package models

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

var (
	testDataPath            = filepath.Join("..", "testutils", "adventure.json")
	firstSavePlayerPosition = Position{X: 50, Y: 16}
)

const (
	firstUUID = "106ee258-21a1-4704-8bad-107c36263e82"
	saveCount = 6
)

func TestGameSave_UnmarshalJSON_Adventure(t *testing.T) {
	data, err := os.ReadFile(testDataPath)
	if err != nil {
		t.Fatalf("failed to read JSON file: %v", err)
	}

	var saveData struct {
		Saves []GameSave `json:"saves"`
	}
	if err := json.Unmarshal(data, &saveData); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	if len(saveData.Saves) != saveCount {
		t.Errorf("expected %d saves, got %d", saveCount, len(saveData.Saves))
	}

	firstSave := saveData.Saves[0]
	if firstSave.ID != firstUUID {
		t.Errorf("expected ID to be '%s', got '%s'", firstUUID, firstSave.ID)
	}
	if firstSave.Player.Name != "Alice" {
		t.Errorf("expected Player.Name to be 'Alice', got '%s'", firstSave.Player.Name)
	}
	if firstSave.GameMode != "Adventure" {
		t.Errorf("expected GameMode to be 'Adventure', got '%s'", firstSave.GameMode)
	}

	ags, ok := firstSave.GameState.(AdventureGameState)
	if !ok {
		t.Errorf("expected GameState to be AdventureGameState, got %T", firstSave.GameState)
	} else {
		if ags.Level.Number != 0 {
			t.Errorf("expected level number to be 0, got %d", ags.Level.Number)
		}
		if ags.Level.PlayerPosition != firstSavePlayerPosition {
			t.Errorf("expected player position to be %+v, got %+v", firstSavePlayerPosition, ags.Level.PlayerPosition)
		}
	}
}

func TestGameSave_UnmarshalJSON_EmptyFile(t *testing.T) {
	data := ""
	var saveData struct {
		Saves []GameSave `json:"saves"`
	}
	if err := json.Unmarshal([]byte(data), &saveData); err == nil {
		t.Errorf("expected error for empty JSON, got nil")
	}
}

func TestGameSave_UnmarshalJSON_EdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		wantErr bool
	}{
		{
			name:    "Missing GameState",
			json:    `{"id": "123", "game_mode": "Adventure"}`,
			wantErr: true,
		},
		{
			name:    "Invalid GameMode",
			json:    `{"id": "123", "game_mode": "InvalidMode", "game_state": {}}`,
			wantErr: true,
		},
		{
			name:    "Empty JSON",
			json:    `{}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gameSave GameSave
			err := json.Unmarshal([]byte(tt.json), &gameSave)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGameSave_MarshalJSON_Adventure(t *testing.T) {
	data, err := os.ReadFile(testDataPath)
	if err != nil {
		t.Fatalf("failed to read JSON file: %v", err)
	}

	var originalData struct {
		Saves []GameSave `json:"saves"`
	}
	if err := json.Unmarshal(data, &originalData); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	// Marshal the data back to JSON
	marshaledData, err := json.MarshalIndent(originalData, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal JSON: %v", err)
	}

	// Validate key fields manually
	var roundTripData struct {
		Saves []GameSave `json:"saves"`
	}
	if err := json.Unmarshal(marshaledData, &roundTripData); err != nil {
		t.Fatalf("failed to unmarshal marshaled JSON: %v", err)
	}

	if len(originalData.Saves) != len(roundTripData.Saves) {
		t.Fatalf("expected %d saves, got %d", len(originalData.Saves), len(roundTripData.Saves))
	}

	for i, save := range originalData.Saves {
		if !reflect.DeepEqual(save, roundTripData.Saves[i]) {
			t.Errorf("save %d: round-trip data mismatch\nOriginal:\n%+v\nGot:\n%+v", i, save, roundTripData.Saves[i])
		}
	}
}
