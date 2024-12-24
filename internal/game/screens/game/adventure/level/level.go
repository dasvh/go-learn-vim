package level

type Level interface {
	Init(width, height int)
	Update(playerX, playerY int) bool
	GetStartPosition() (int, int)
	GetGrid() [][]rune
	GetInstructions() string
	GetTargetCharacter() rune
	GetTargetReachedCharacter() rune
	IsCompleted() bool
}

type Position struct {
	X int
	Y int
}

type LevelZero struct {
	grid      [][]rune
	current   int
	width     int
	height    int
	targets   []Position
	completed bool
	init      bool
}

func NewLevelZero() Level {
	return &LevelZero{}
}

func (level0 *LevelZero) GetTargetCharacter() rune {
	return 'X'
}

func (level0 *LevelZero) GetTargetReachedCharacter() rune {
	return '✓'
}

func (level0 *LevelZero) Init(width, height int) {
	level0.width = width
	level0.height = height

	// define targets
	level0.targets = []Position{
		{0, 0},                  // top left
		{width - 1, 0},          // top right
		{0, height - 1},         // bottom left
		{width - 1, height - 1}, // bottom right
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

	if level0.current < len(level0.targets) {
		target := level0.targets[level0.current]
		level0.grid[target.Y][target.X] = level0.GetTargetCharacter()
	}
}

func (level0 *LevelZero) Update(playerX, playerY int) bool {
	target := level0.targets[level0.current]

	// check if player reached target
	if playerX == target.X && playerY == target.Y {
		level0.grid[target.Y][target.X] = level0.GetTargetReachedCharacter()

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
	return "Reach the X in each corner using hjkl keys"
}

func (level0 *LevelZero) IsCompleted() bool {
	return level0.completed
}
