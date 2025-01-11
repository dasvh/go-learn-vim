package level

import (
	"github.com/dasvh/go-learn-vim/internal/models"
	"math/rand"
)

// Maze represents a maze with walls
type Maze struct {
	width         int
	height        int
	walls         []models.Position
	offsetX       int
	offsetY       int
	rand          *rand.Rand
	StartPosition models.Position
}

// NewMaze initializes a new square Maze with an offset
func NewMaze(size int, seed int64, offsetX, offsetY int, pathWidth int) *Maze {
	r := rand.New(rand.NewSource(seed))
	m := &Maze{
		width:   size,
		height:  size,
		offsetX: offsetX,
		offsetY: offsetY,
		rand:    r,
	}
	m.generateGridWalls()
	m.generateMazeDFS(pathWidth)
	m.StartPosition = models.Position{X: offsetX + pathWidth, Y: offsetY + pathWidth}
	return m
}

// GetWalls returns the walls with offsets applied
func (m *Maze) GetWalls() []models.Position {
	var offsetWalls []models.Position
	for _, wall := range m.walls {
		offsetWalls = append(offsetWalls, models.Position{X: wall.X + m.offsetX, Y: wall.Y + m.offsetY})
	}
	return offsetWalls
}

// generateGridWalls generates the walls for the Maze
func (m *Maze) generateGridWalls() {
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			m.walls = append(m.walls, models.Position{X: x, Y: y})
		}
	}
}

// generateMazeDFS generates a maze using the depth-first search algorithm
// with a given path width to create corridors
func (m *Maze) generateMazeDFS(pathWidth int) {
	visited := make([][]bool, m.height)
	for y := range visited {
		visited[y] = make([]bool, m.width)
	}

	// start from inside the maze in the top-left corner
	start := models.Position{X: pathWidth, Y: pathWidth}
	m.dfs(start, visited, pathWidth)
}

func (m *Maze) dfs(cell models.Position, visited [][]bool, pathWidth int) {
	// mark the visited cells
	for dy := 0; dy < pathWidth; dy++ {
		for dx := 0; dx < pathWidth; dx++ {
			nx, ny := cell.X+dx, cell.Y+dy
			if ny < m.height && nx < m.width {
				visited[ny][nx] = true
				m.removeWall(nx, ny)
			}
		}
	}

	directions := []models.Position{
		{X: 0, Y: -2 * pathWidth}, // up
		{X: 0, Y: 2 * pathWidth},  // down
		{X: -2 * pathWidth, Y: 0}, // left
		{X: 2 * pathWidth, Y: 0},  // right
	}

	// shuffle the directions to randomize the path
	m.rand.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] })

	for _, dir := range directions {
		next := models.Position{X: cell.X + dir.X, Y: cell.Y + dir.Y}

		// check if the next cell is within the maze bounds and not visited
		if next.X >= pathWidth && next.X+pathWidth < m.width &&
			next.Y >= pathWidth && next.Y+pathWidth < m.height &&
			!isVisited(next, visited, pathWidth) {
			// remove walls between the cells
			removeCorridorWalls(m, cell, next, pathWidth)

			// recursively visit the next cell
			m.dfs(next, visited, pathWidth)
		}
	}
}

// removeWall removes a wall from the Maze
func (m *Maze) removeWall(x, y int) {
	m.walls = filterWalls(m.walls, func(w models.Position) bool {
		return !(w.X == x && w.Y == y)
	})
}

// isVisited checks if the entire "pathWidth x pathWidth" area is visited
func isVisited(cell models.Position, visited [][]bool, pathWidth int) bool {
	for dy := 0; dy < pathWidth; dy++ {
		for dx := 0; dx < pathWidth; dx++ {
			nx, ny := cell.X+dx, cell.Y+dy
			if ny >= len(visited) || nx >= len(visited[0]) || !visited[ny][nx] {
				return false
			}
		}
	}
	return true
}

// removeCorridorWalls removes walls between two "pathWidth x pathWidth" cells
func removeCorridorWalls(m *Maze, current, next models.Position, pathWidth int) {
	// Calculate the midpoint between the cells
	midX := (current.X + next.X) / 2
	midY := (current.Y + next.Y) / 2

	// Remove walls in the corridor
	for dy := 0; dy < pathWidth; dy++ {
		for dx := 0; dx < pathWidth; dx++ {
			m.removeWall(midX+dx, midY+dy)
		}
	}
}

// filterWalls filters the walls based on a predicate that returns true
// for which walls to keep
func filterWalls(walls []models.Position, predicate func(models.Position) bool) []models.Position {
	var result []models.Position
	for _, w := range walls {
		if predicate(w) {
			result = append(result, w)
		}
	}
	return result
}

// isWall checks if a position is a wall in the Maze
func (m *Maze) isWall(pos models.Position) bool {
	for _, wall := range m.walls {
		if wall.X == pos.X && wall.Y == pos.Y {
			return true
		}
	}
	return false
}
