package level

import (
	"fmt"
	"github.com/dasvh/go-learn-vim/internal/game/screens/game/adventure/character"
)

type Level interface {
	Init(width, height int)
	Update(playerX, playerY int) bool
	GetStartPosition() (int, int)
	GetGrid() [][]rune
	GetInstructions() string
	GetCharacters() *character.Characters
	IsCompleted() bool
}

type Position struct {
	X int
	Y int
}

type target struct {
	position Position
	reached  bool
}

type LevelZero struct {
	grid      [][]rune
	current   int
	width     int
	height    int
	targets   []target
	completed bool
	init      bool
	chars     *character.Characters
}

func (level0 *LevelZero) GetCharacters() *character.Characters {
	return level0.chars
}

func NewLevelZero() Level {
	return &LevelZero{
		chars: &character.DefaultCharacters,
	}
}

func (level0 *LevelZero) Init(width, height int) {
	level0.width = width
	level0.height = height

	// offset targets from the border by 20%
	offsetX := int(float64(width) * 0.2)
	offsetY := int(float64(height) * 0.2)

	// define targets
	level0.targets = []target{
		{Position{offsetX, offsetY}, false},                          // top left
		{Position{width - offsetX - 1, offsetY}, false},              // top right
		{Position{offsetX, height - offsetY - 1}, false},             // bottom left
		{Position{width - offsetX - 1, height - offsetY - 1}, false}, // bottom right
	}

	level0.current = 0
	level0.initializeGrid()
}

func (level0 *LevelZero) initializeGrid() {
	// clear grid on init
	if !level0.init {
		level0.grid = make([][]rune, level0.height)
		for y := 0; y < level0.height; y++ {
			level0.grid[y] = make([]rune, level0.width)
			for x := 0; x < level0.width; x++ {
				level0.grid[y][x] = ' '
			}
		}
		level0.init = true
	}

	// set player
	level0.grid[level0.height/2][level0.width/2] = level0.chars.Player.Cursor.Rune

	// todo: figure out how to set the characters even after the game is completed
	for i, target := range level0.targets {
		if i == level0.current {
			level0.grid[target.position.Y][target.position.X] = level0.chars.Target.Active.Rune
		} else {
			if target.reached {
				level0.grid[target.position.Y][target.position.X] = level0.chars.Target.Reached.Rune
			} else {
				level0.grid[target.position.Y][target.position.X] = level0.chars.Target.Inactive.Rune
			}
		}
	}
}

func (level0 *LevelZero) Update(playerX, playerY int) bool {
	target := level0.targets[level0.current]

	// check if player reached target
	if playerX == target.position.X && playerY == target.position.Y {
		level0.targets[level0.current].reached = true

		// check if this was the last target
		if level0.current == len(level0.targets)-1 {
			level0.completed = true
			return true
		}

		level0.current++
		level0.initializeGrid()
		return true // reset player
	}
	return false
}

func (level0 *LevelZero) GetStartPosition() (int, int) {
	return level0.width / 2, level0.height / 2
}

func (level0 *LevelZero) GetGrid() [][]rune {
	return level0.grid
}

func (level0 *LevelZero) GetInstructions() string {
	return fmt.Sprintf("Target %d/%d: Reach the X using hjkl keys", level0.current+1, len(level0.targets))
}

func (level0 *LevelZero) IsCompleted() bool {
	return level0.completed
}
