package level

import (
	"github.com/dasvh/go-learn-vim/internal/models"
	"testing"
)

func Test_MazeGeneration(t *testing.T) {
	t.Run("basic maze properties", func(t *testing.T) {
		size, seed := 21, int64(42)
		offsetX, offsetY := 5, 5
		pathWidth := 1

		maze := NewMaze(size, seed, offsetX, offsetY, pathWidth)

		if len(maze.GetWalls()) == 0 {
			t.Error("expected walls to be generated")
		}

		startPos := maze.StartPosition
		for _, wall := range maze.GetWalls() {
			if wall == startPos {
				t.Error("expected start position to be a path cell")
			}
		}

		if startPos.X < offsetX || startPos.X >= size+offsetX ||
			startPos.Y < offsetY || startPos.Y >= size+offsetY {
			t.Error("expected start position to be within maze bounds")
		}
	})

	t.Run("deterministic maze generation", func(t *testing.T) {
		maze1 := NewMaze(21, 42, 5, 5, 1)
		maze2 := NewMaze(21, 42, 5, 5, 1)

		walls1 := maze1.GetWalls()
		walls2 := maze2.GetWalls()

		if len(walls1) != len(walls2) {
			t.Errorf("expected same number of walls, got %d and %d", len(walls1), len(walls2))
		}

		// Compare wall positions
		wallMap1 := makeWallMap(walls1)
		wallMap2 := makeWallMap(walls2)

		for pos := range wallMap1 {
			if !wallMap2[pos] {
				t.Error("expected maze with same seed to have identical walls")
				break
			}
		}
	})

	t.Run("path width validation", func(t *testing.T) {
		pathWidth := 2
		maze := NewMaze(21, 42, 5, 5, pathWidth)
		walls := maze.GetWalls()
		wallMap := makeWallMap(walls)

		for y := maze.offsetY; y < maze.height+maze.offsetY; y += 2 * pathWidth {
			for x := maze.offsetX; x < maze.width+maze.offsetX; x += 2 * pathWidth {
				pos := models.Position{X: x, Y: y}
				if !wallMap[pos] {
					if !hasPathWidthSpace(pos, pathWidth, wallMap) {
						t.Errorf("Path cell at (%d,%d) doesn't have %dx%d clear space",
							x, y, pathWidth, pathWidth)
					}

					if !hasValidConnection(pos, pathWidth, wallMap, maze.width, maze.height) {
						t.Errorf("Path cell at (%d,%d) has no valid connections", x, y)
					}
				}
			}
		}
	})

	t.Run("path existence", func(t *testing.T) {
		maze := NewMaze(21, 42, 5, 5, 1)
		walls := maze.GetWalls()
		wallMap := makeWallMap(walls)

		endPos := findAccessiblePosition(maze)
		path := findPath(maze.StartPosition, endPos, wallMap, maze.offsetX, maze.offsetY, maze.width, maze.height)

		if path == nil {
			t.Error("expected path to exist between start and end positions")
		}
	})
}

func makeWallMap(walls []models.Position) map[models.Position]bool {
	wallMap := make(map[models.Position]bool)
	for _, w := range walls {
		wallMap[w] = true
	}
	return wallMap
}

func findPath(start, end models.Position, walls map[models.Position]bool,
	offsetX, offsetY, width, height int) []models.Position {
	visited := make(map[models.Position]bool)
	queue := [][]models.Position{{start}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		current := path[len(path)-1]

		if current == end {
			return path
		}

		dirs := []models.Position{{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: 0, Y: 1}}
		for _, dir := range dirs {
			next := models.Position{X: current.X + dir.X, Y: current.Y + dir.Y}

			if next.X < offsetX || next.X >= width+offsetX ||
				next.Y < offsetY || next.Y >= height+offsetY ||
				walls[next] || visited[next] {
				continue
			}

			visited[next] = true
			newPath := make([]models.Position, len(path))
			copy(newPath, path)
			queue = append(queue, append(newPath, next))
		}
	}

	return nil
}

func findAccessiblePosition(maze *Maze) models.Position {
	walls := makeWallMap(maze.GetWalls())
	for y := maze.offsetY; y < maze.height+maze.offsetY; y++ {
		for x := maze.offsetX; x < maze.width+maze.offsetX; x++ {
			pos := models.Position{X: x, Y: y}
			if !walls[pos] {
				return pos
			}
		}
	}
	return maze.StartPosition
}

func hasPathWidthSpace(pos models.Position, pathWidth int, walls map[models.Position]bool) bool {
	for dy := range pathWidth {
		for dx := range pathWidth {
			checkPos := models.Position{X: pos.X + dx, Y: pos.Y + dy}
			if walls[checkPos] {
				return false
			}
		}
	}
	return true
}

func hasValidConnection(pos models.Position, pathWidth int, walls map[models.Position]bool, width, height int) bool {
	directions := []models.Position{
		{X: 0, Y: -pathWidth}, // up
		{X: 0, Y: pathWidth},  // down
		{X: -pathWidth, Y: 0}, // left
		{X: pathWidth, Y: 0},  // right
	}

	for _, dir := range directions {
		nextPos := models.Position{X: pos.X + dir.X, Y: pos.Y + dir.Y}
		if isValidCorridor(pos, nextPos, walls) {
			return true
		}
	}
	return false
}

func isValidCorridor(from, to models.Position, walls map[models.Position]bool) bool {
	minX := min(from.X, to.X)
	maxX := max(from.X, to.X)
	minY := min(from.Y, to.Y)
	maxY := max(from.Y, to.Y)

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if walls[models.Position{X: x, Y: y}] {
				return false
			}
		}
	}
	return true
}
