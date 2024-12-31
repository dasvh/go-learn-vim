package storage

// AdventureGameRepository is the interface that wraps
// the basic methods to interact with the storage of an adventure app
type AdventureGameRepository interface {
	SaveAdventureGame(state AdventureGameState) error
	LoadAdventureGame() (AdventureGameState, error)
	HasIncompleteGame() (bool, error)
}
