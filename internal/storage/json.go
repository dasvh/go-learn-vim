package storage

import (
	"encoding/json"
	"fmt"
	"github.com/dasvh/go-learn-vim/internal/models"
	"os"
	"sort"
)

// JSONRepository stores the data in a JSON file
type JSONRepository struct {
	filePath string
	data     struct {
		Players []models.Player   `json:"players"`
		Saves   []models.GameSave `json:"saves"`
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
func (repo *JSONRepository) AddPlayer(player models.Player) error {
	repo.data.Players = append(repo.data.Players, player)
	return repo.save()
}

// Players returns all players in the repository
func (repo *JSONRepository) Players() ([]models.Player, error) {
	return repo.data.Players, nil
}

// SaveGame saves a game to the repository
func (repo *JSONRepository) SaveGame(save models.GameSave) error {
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
func (repo *JSONRepository) LoadGame(gameID string) (models.GameSave, error) {
	for _, save := range repo.data.Saves {
		if save.ID == gameID {
			return save, nil
		}
	}

	return models.GameSave{}, fmt.Errorf("game with ID %s not found", gameID)
}

// HasIncompleteGames returns whether there are any incomplete games
func (repo *JSONRepository) HasIncompleteGames() bool {
	file, err := os.Open(repo.filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	for _, save := range repo.data.Saves {
		if !save.GameState.IsCompleted() {
			return true
		}
	}

	return false
}

// IncompleteGames returns all incomplete games
func (repo *JSONRepository) IncompleteGames() ([]models.GameSave, error) {
	var incomplete []models.GameSave
	for _, save := range repo.data.Saves {
		if !save.GameState.IsCompleted() {
			incomplete = append(incomplete, save)
		}
	}

	return incomplete, nil
}

// LoadGameState loads a specific GameState from the repository
func (repo *JSONRepository) LoadGameState(gameID string) (models.GameState, error) {
	for _, save := range repo.data.Saves {
		if save.ID == gameID {
			// Deserialize the specific GameState type
			switch save.GameState.(type) {
			case models.AdventureGameState:
				return save.GameState.(models.AdventureGameState), nil
			default:
				return nil, fmt.Errorf("unsupported game controllers mode")
			}
		}
	}
	return nil, fmt.Errorf("game with ID %s not found", gameID)
}

// LifetimeStats computes aggregated stats across all game saves
func (repo *JSONRepository) LifetimeStats() (*models.LifetimeStats, error) {
	lifetimeStats := models.NewLifetimeStats()

	uniqueGames := make(map[string]struct{})

	for _, save := range repo.data.Saves {
		if adventureState, ok := save.GameState.(models.AdventureGameState); ok {
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
func (repo *JSONRepository) PlayerLifetimeStats(playerID string) (*models.LifetimeStats, error) {
	lifetimeStats := models.NewLifetimeStats()

	for _, save := range repo.data.Saves {
		if save.Player.ID == playerID {
			if adventureState, ok := save.GameState.(models.AdventureGameState); ok {
				lifetimeStats.Merge(adventureState.Stats)
				lifetimeStats.TotalGames++
			}
		}
	}

	return lifetimeStats, nil
}

// ComputeHighScores computes high scores for the repository
func (repo *JSONRepository) ComputeHighScores() ([]models.HighScore, error) {
	const (
		BaseScore       = 25000
		TimeWeight      = 177
		KeystrokeWeight = 17
		MinScore        = 5000
	)
	var highScores []models.HighScore

	for _, save := range repo.data.Saves {
		if save.GameState.IsCompleted() {
			gameStats := save.GameState.(models.AdventureGameState).Stats
			if gameStats.TimeElapsed > 0 && gameStats.TotalKeystrokes > 0 {
				timePenalty := gameStats.TimeElapsed * TimeWeight
				keystrokePenalty := gameStats.TotalKeystrokes * KeystrokeWeight
				score := max(BaseScore-timePenalty-keystrokePenalty, MinScore)
				highScores = append(highScores, models.HighScore{
					PlayerName: save.Player.Name,
					Level:      save.GameState.(models.AdventureGameState).Level.Number,
					Score:      score,
					Timestamp:  save.Timestamp,
				})
			}
		}
	}

	sort.Slice(highScores, func(i, j int) bool {
		return highScores[i].Score > highScores[j].Score
	})

	return highScores, nil
}
