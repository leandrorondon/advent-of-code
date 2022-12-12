package heightmap

import "github.com/leandrorondon/advent-of-code/pkg/iterator"

const (
	START = 'S'
	END   = 'E'
)

type HeightMap struct {
	grid  [][]byte
	Start Position
	End   Position
}

type Position struct {
	X int
	Y int
}

type Node struct {
	Pos  Position
	Dist int
}

func (n Node) Next(dp Position) Position {
	return Position{
		n.Pos.X + dp.X,
		n.Pos.Y + dp.Y,
	}
}

var directions = []Position{
	{0, -1}, // left
	{0, 1},  // right
	{-1, 0}, // up
	{1, 0},  // down
}

func (hm HeightMap) ShortestPath(start, dest Position) int {
	visited := make(map[Position]bool)
	visited[start] = true

	var queue []Node
	queue = append(queue, Node{start, 0})

	for len(queue) > 0 {
		current := queue[0]

		if current.Pos == dest {
			return current.Dist
		}

		queue = queue[1:]

		for _, dp := range directions {
			next := current.Next(dp)
			if hm.canGo(visited, current.Pos, next) {
				visited[next] = true
				queue = append(queue, Node{next, current.Dist + 1})
			}
		}
	}

	return -1
}

func (hm HeightMap) ShortestPathFromHeight(startHeight byte, dest Position) int {
	it := iterator.New2D(hm.grid)
	shortest := len(hm.grid) * len(hm.grid[0])
	for it.Next() {
		h := it.Get()
		if h == startHeight {
			n := hm.ShortestPath(Position{it.X(), it.Y()}, dest)
			if n > 0 && n < shortest {
				shortest = n
			}
		}
	}
	return shortest
}

func (hm HeightMap) canGo(visited map[Position]bool, current Position, next Position) bool {
	return next.X >= 0 && next.X < len(hm.grid) && next.Y >= 0 && next.Y < len(hm.grid[0]) &&
		!visited[next] &&
		(hm.grid[current.X][current.Y] >= hm.grid[next.X][next.Y]-1)
}
