package storage

import (
	"encoding/json"
	"os"
)

// JSONRepository stores the AdventureGameState in a JSON file
type JSONRepository struct {
	filePath string
}

// NewJSONRepository creates a new JSONRepository
func NewJSONRepository(filePath string) *JSONRepository {
	return &JSONRepository{filePath: filePath}
}

// SaveAdventureGame saves the AdventureGameState to a JSON file
func (repo *JSONRepository) SaveAdventureGame(state AdventureGameState) error {
	file, err := os.Create(repo.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(state)
}

// LoadAdventureGame loads the AdventureGameState from a JSON file
func (repo *JSONRepository) LoadAdventureGame() (AdventureGameState, error) {
	file, err := os.Open(repo.filePath)
	if err != nil {
		return AdventureGameState{}, err
	}
	defer file.Close()

	var state AdventureGameState
	err = json.NewDecoder(file).Decode(&state)
	return state, err
}

// HasIncompleteGame checks if there is an incomplete game
func (repo *JSONRepository) HasIncompleteGame() (bool, error) {
	file, err := os.Open(repo.filePath)
	if err != nil {
		// if the file doesn't exist, return false without an error
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	defer file.Close()

	var state AdventureGameState
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&state)
	if err != nil {
		return false, err
	}

	return !state.Level.Completed, nil
}
