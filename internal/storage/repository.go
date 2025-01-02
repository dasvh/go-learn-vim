package storage

type GameRepository interface {
	AddPlayer(Player) error
	Players() ([]Player, error)

	SaveGame(save GameSave) error
	LoadGame(gameID string) (GameSave, error)
	HasIncompleteGames() (bool, error)
	IncompleteGames() ([]GameSave, error)

	LoadGameState(gameID string) (GameState, error)
}
