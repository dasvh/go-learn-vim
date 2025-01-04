package controllers

import (
	"fmt"
	"github.com/dasvh/go-learn-vim/internal/models"
	"github.com/dasvh/go-learn-vim/internal/storage"
	"github.com/google/uuid"
	"time"
)

// Game is a controller for game related actions
type Game struct {
	repo          storage.GameRepository
	currentPlayer *models.Player
}

// NewGame creates a new Game controller
func NewGame(repo storage.GameRepository) *Game {
	return &Game{repo: repo}
}

// CreatePlayer creates a new player with the given name
func (gc *Game) CreatePlayer(name string) (models.Player, error) {
	players, err := gc.repo.Players()
	if err != nil {
		return models.Player{}, err
	}

	for _, player := range players {
		if player.Name == name {
			return models.Player{}, fmt.Errorf("player with name %q already exists", name)
		}
	}

	player := models.Player{
		ID:   uuid.NewString(),
		Name: name,
	}

	err = gc.repo.AddPlayer(player)
	return player, err
}

// Players returns all players
func (gc *Game) Players() ([]models.Player, error) {
	return gc.repo.Players()
}

// SetPlayer sets the current player
func (gc *Game) SetPlayer(player models.Player) {
	gc.currentPlayer = &player
}

// SaveGame saves the current game state
func (gc *Game) SaveGame(mode string, gameState models.GameState) error {
	if gc.currentPlayer == nil {
		return fmt.Errorf("no player selected")
	}

	gameSave := models.GameSave{
		ID:        uuid.NewString(),
		Player:    *gc.currentPlayer,
		Timestamp: time.Now(),
		GameMode:  mode,
		GameState: gameState,
	}

	return gc.repo.SaveGame(gameSave)
}
