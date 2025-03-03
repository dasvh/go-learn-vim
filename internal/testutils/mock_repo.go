package testutils

import (
	"fmt"
	"github.com/dasvh/go-learn-vim/internal/models"
)

type MockGameRepository struct {
	PlayersData   []models.Player
	GameSavesData []models.GameSave
}

func NewMockGameRepository() *MockGameRepository {
	return &MockGameRepository{
		PlayersData:   []models.Player{},
		GameSavesData: []models.GameSave{},
	}
}

func NewMockGameRepositoryWithData(players []models.Player, saves []models.GameSave) *MockGameRepository {
	return &MockGameRepository{
		PlayersData:   players,
		GameSavesData: saves,
	}
}

// AddPlayer adds a player to the mock repository
func (m *MockGameRepository) AddPlayer(player models.Player) error {
	m.PlayersData = append(m.PlayersData, player)
	return nil
}

// Players returns all players in the mock repository
func (m *MockGameRepository) Players() ([]models.Player, error) {
	return m.PlayersData, nil
}

// SaveGame saves a game to the mock repository
func (m *MockGameRepository) SaveGame(save models.GameSave) error {
	for i, s := range m.GameSavesData {
		if s.ID == save.ID {
			m.GameSavesData[i] = save
			return nil
		}
	}
	m.GameSavesData = append(m.GameSavesData, save)
	return nil
}

// LoadGame loads a game from the mock repository by its ID
func (m *MockGameRepository) LoadGame(gameID string) (models.GameSave, error) {
	for _, save := range m.GameSavesData {
		if save.ID == gameID {
			return save, nil
		}
	}
	return models.GameSave{}, fmt.Errorf("game with ID %q not found", gameID)
}

// HasIncompleteGames checks if there are incomplete games
func (m *MockGameRepository) HasIncompleteGames() bool {
	for _, save := range m.GameSavesData {
		if !save.GameState.IsCompleted() {
			return true
		}
	}
	return false
}

// IncompleteGames returns all incomplete games
func (m *MockGameRepository) IncompleteGames() ([]models.GameSave, error) {
	var incompleteGames []models.GameSave
	for _, save := range m.GameSavesData {
		if !save.GameState.IsCompleted() {
			incompleteGames = append(incompleteGames, save)
		}
	}
	return incompleteGames, nil
}

// LoadGameState loads the game state by its ID
func (m *MockGameRepository) LoadGameState(gameID string) (models.GameState, error) {
	for _, save := range m.GameSavesData {
		if save.ID == gameID {
			switch save.GameState.(type) {
			case models.AdventureGameState:
				return save.GameState.(models.AdventureGameState), nil
			default:
				return nil, fmt.Errorf("unsupported game mode: %s", save.GameMode)
			}
		}
	}
	return nil, fmt.Errorf("game with ID %q not found", gameID)
}

// LifetimeStats returns the lifetime stats
func (m *MockGameRepository) LifetimeStats() (*models.LifetimeStats, error) {
	return &models.LifetimeStats{}, nil
}

// PlayerLifetimeStats returns the lifetime stats for a player
func (m *MockGameRepository) PlayerLifetimeStats(playerID string) (*models.LifetimeStats, error) {
	return &models.LifetimeStats{}, nil
}

// ComputeHighScores computes the high scores
func (m *MockGameRepository) ComputeHighScores() ([]models.HighScore, error) {
	return []models.HighScore{}, nil
}
