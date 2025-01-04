package storage

import "github.com/dasvh/go-learn-vim/internal/models"

type GameRepository interface {
	AddPlayer(models.Player) error
	Players() ([]models.Player, error)

	SaveGame(save models.GameSave) error
	LoadGame(gameID string) (models.GameSave, error)
	HasIncompleteGames() (bool, error)
	IncompleteGames() ([]models.GameSave, error)

	LoadGameState(gameID string) (models.GameState, error)

	LifetimeStats() (*models.LifetimeStats, error)
	PlayerLifetimeStats(playerID string) (*models.LifetimeStats, error)

	ComputeHighScores() ([]models.HighScore, error)
}
