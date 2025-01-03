package controllers

import (
	"fmt"
	"github.com/dasvh/go-learn-vim/internal/storage"
	"github.com/google/uuid"
	"time"
)

// Game is a controller for game related actions
type Game struct {
	repo          storage.GameRepository
	currentPlayer *storage.Player
}

// NewGame creates a new Game controller
func NewGame(repo storage.GameRepository) *Game {
	return &Game{repo: repo}
}

// CreatePlayer creates a new player with the given name
func (g *Game) CreatePlayer(name string) (storage.Player, error) {
	players, err := g.repo.Players()
	if err != nil {
		return storage.Player{}, err
	}

	for _, player := range players {
		if player.Name == name {
			return storage.Player{}, fmt.Errorf("player with name %q already exists", name)
		}
	}

	player := storage.Player{
		ID:   uuid.NewString(),
		Name: name,
	}

	err = g.repo.AddPlayer(player)
	return player, err
}

// Players returns all players
func (g *Game) Players() ([]storage.Player, error) {
	return g.repo.Players()
}

// SetPlayer sets the current player
func (g *Game) SetPlayer(player storage.Player) {
	g.currentPlayer = &player
}

// SaveGame saves the current game state
func (g *Game) SaveGame(mode string, gameState storage.GameState) error {
	if g.currentPlayer == nil {
		return fmt.Errorf("no player selected")
	}

	gameSave := storage.GameSave{
		ID:        uuid.NewString(),
		Player:    *g.currentPlayer,
		Timestamp: time.Now(),
		GameMode:  mode,
		GameState: gameState,
	}

	return g.repo.SaveGame(gameSave)
}
