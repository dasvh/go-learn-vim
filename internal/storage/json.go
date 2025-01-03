package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

// JSONRepository stores the data in a JSON file
type JSONRepository struct {
	filePath string
	data     struct {
		Players []Player   `json:"players"`
		Saves   []GameSave `json:"saves"`
	}
}

// NewJSONRepository creates a new JSONRepository
func NewJSONRepository(filePath string) (*JSONRepository, error) {
	repo := &JSONRepository{filePath: filePath}

	if _, err := os.Stat(filePath); err == nil {
		file, err := os.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()

		if err := json.NewDecoder(file).Decode(&repo.data); err != nil {
			return nil, fmt.Errorf("failed to decode JSON: %w", err)
		}
	}

	return repo, nil
}

// save writes the repository data to the JSON file
func (repo *JSONRepository) save() error {
	file, err := os.Create(repo.filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(repo.data)
}

// AddPlayer adds a new player to the repository
func (repo *JSONRepository) AddPlayer(player Player) error {
	repo.data.Players = append(repo.data.Players, player)
	return repo.save()
}

// Players returns all players in the repository
func (repo *JSONRepository) Players() ([]Player, error) {
	return repo.data.Players, nil
}

// SaveGame saves a game to the repository
func (repo *JSONRepository) SaveGame(save GameSave) error {
	for i, s := range repo.data.Saves {
		if s.ID == save.ID {
			repo.data.Saves[i] = save
			return repo.save()
		}
	}
	repo.data.Saves = append(repo.data.Saves, save)
	return repo.save()
}

// LoadGame loads a game from the repository
func (repo *JSONRepository) LoadGame(gameID string) (GameSave, error) {
	for _, save := range repo.data.Saves {
		if save.ID == gameID {
			return save, nil
		}
	}

	return GameSave{}, fmt.Errorf("game with ID %s not found", gameID)
}

// HasIncompleteGames returns whether there are any incomplete games
func (repo *JSONRepository) HasIncompleteGames() (bool, error) {
	file, err := os.Open(repo.filePath)
	if err != nil {
		return false, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	for _, save := range repo.data.Saves {
		if !save.GameState.IsCompleted() {
			return true, nil
		}
	}

	return false, nil
}

// IncompleteGames returns all incomplete games
func (repo *JSONRepository) IncompleteGames() ([]GameSave, error) {
	var incomplete []GameSave
	for _, save := range repo.data.Saves {
		if !save.GameState.IsCompleted() {
			incomplete = append(incomplete, save)
		}
	}

	return incomplete, nil
}

// LoadGameState loads a specific GameState from the repository
func (repo *JSONRepository) LoadGameState(gameID string) (GameState, error) {
	for _, save := range repo.data.Saves {
		if save.ID == gameID {
			// Deserialize the specific GameState type
			switch save.GameState.(type) {
			case AdventureGameState:
				return save.GameState.(AdventureGameState), nil
			default:
				return nil, fmt.Errorf("unsupported game controllers mode")
			}
		}
	}
	return nil, fmt.Errorf("game with ID %s not found", gameID)
}

// LifetimeStats computes aggregated stats across all game saves
func (repo *JSONRepository) LifetimeStats() (*LifetimeStats, error) {
	lifetimeStats := NewLifetimeStats()

	uniqueGames := make(map[string]struct{}) // Track unique game IDs

	for _, save := range repo.data.Saves {
		if adventureState, ok := save.GameState.(AdventureGameState); ok {
			lifetimeStats.Merge(adventureState.Stats)

			if _, exists := uniqueGames[save.ID]; !exists {
				uniqueGames[save.ID] = struct{}{}
				lifetimeStats.TotalGames++
			}
		}
	}

	return lifetimeStats, nil
}

// PlayerLifetimeStats computes stats for a specific player
func (repo *JSONRepository) PlayerLifetimeStats(playerID string) (*LifetimeStats, error) {
	lifetimeStats := NewLifetimeStats()

	for _, save := range repo.data.Saves {
		if save.Player.ID == playerID {
			if adventureState, ok := save.GameState.(AdventureGameState); ok {
				lifetimeStats.Merge(adventureState.Stats)
				lifetimeStats.TotalGames++
			}
		}
	}

	return lifetimeStats, nil
}
