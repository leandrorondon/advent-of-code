package day22

type Wrapper2D struct {
	lastRow int
	lastCol int
	grid    [][]byte
}

func NewWrapper2D(grid [][]byte) Wrapper2D {
	return Wrapper2D{
		grid:    grid,
		lastRow: len(grid) - 1,
		lastCol: len(grid[0]) - 1,
	}
}

func (w Wrapper2D) Password(pos Position) int {
	return 1000*(pos.Pos.Row+1) + 4*(pos.Pos.Col+1) + dirValue[pos.Face]
}

func (w Wrapper2D) Next(pos *Position) Position {
	next := w.next(pos.Pos, pos.Face)

	for w.isVoid(next) {
		next = w.next(next, pos.Face)
	}

	return Position{Pos: next, Face: pos.Face}
}

func (w Wrapper2D) next(pos Coordinate, dir byte) Coordinate {
	next := pos.Next(dir)

	if next.Row < 0 {
		next.Row = w.lastRow
	} else if next.Row > w.lastRow {
		next.Row = 0
	}

	if next.Col < 0 {
		next.Col = w.lastCol
	} else if next.Col > w.lastCol {
		next.Col = 0
	}

	return next
}

func (w Wrapper2D) isVoid(pos Coordinate) bool {
	return w.grid[pos.Row][pos.Col] == Void ||
		w.grid[pos.Row][pos.Col] == Space
}

var dirValue = map[byte]int{
	Right: 0,
	Down:  1,
	Left:  2,
	Up:    3,
}
