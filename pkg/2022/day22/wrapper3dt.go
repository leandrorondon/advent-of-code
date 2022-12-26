package day22

type Wrapper3DT struct {
	grid     [][]byte
	facesMin [6]Coordinate
	facesMax [6]Coordinate
}

func NewWrapper3DT(grid [][]byte) Wrapper3DT {
	w := Wrapper3DT{
		grid: grid,
	}

	w.facesMin[0] = Coordinate{0, 8}
	w.facesMax[0] = Coordinate{3, 11}
	w.facesMin[1] = Coordinate{4, 0}
	w.facesMax[1] = Coordinate{7, 3}
	w.facesMin[2] = Coordinate{0, 4}
	w.facesMax[2] = Coordinate{7, 7}
	w.facesMin[3] = Coordinate{0, 8}
	w.facesMax[3] = Coordinate{7, 11}
	w.facesMin[4] = Coordinate{8, 8}
	w.facesMax[4] = Coordinate{11, 11}
	w.facesMin[5] = Coordinate{8, 12}
	w.facesMax[5] = Coordinate{11, 15}

	return w
}

func (w Wrapper3DT) Password(pos Position) int {
	return 1000*(pos.Pos.Row+1) + 4*(pos.Pos.Col+1) + dirValue[pos.Face]
}

func (w Wrapper3DT) Next(pos *Position) Position {
	next := pos.Pos.Next(pos.Face)

	if !w.isVoid(next) {
		return Position{Pos: next, Face: pos.Face}
	}

	return w.wrap(pos)
}

func (w Wrapper3DT) wrap(pos *Position) Position {
	switch w.face(pos.Pos) {
	case 1:
		switch pos.Face {
		case '^':
			return Position{
				Pos:  Coordinate{Row: 4, Col: 11 - pos.Pos.Col},
				Face: 'v',
			}
		case '>':
			return Position{
				Pos:  Coordinate{Row: 11 - pos.Pos.Row, Col: 15},
				Face: '<',
			}
		case '<':
			return Position{
				Pos:  Coordinate{Row: 4, Col: pos.Pos.Row + 4},
				Face: 'v',
			}
		}
	case 2:
		switch pos.Face {
		case '^':
			return Position{
				Pos:  Coordinate{Row: 0, Col: 11 - pos.Pos.Col},
				Face: 'v',
			}
		case 'v':
			return Position{
				Pos:  Coordinate{Row: 11, Col: 11 - pos.Pos.Col},
				Face: '^',
			}
		case '<':
			return Position{
				Pos:  Coordinate{Row: 11, Col: 19 - pos.Pos.Row},
				Face: '^',
			}
		}
	case 3:
		switch pos.Face {
		case '^':
			return Position{
				Pos:  Coordinate{Row: pos.Pos.Col - 4, Col: 7},
				Face: '>',
			}
		case 'v':
			return Position{
				Pos:  Coordinate{Row: 15 - pos.Pos.Col, Col: 7},
				Face: '<',
			}
		}
	case 4:
		switch pos.Face {
		case '>':
			return Position{
				Pos:  Coordinate{Row: 8, Col: 19 - pos.Pos.Row},
				Face: 'v',
			}
		}
	case 5:
		switch pos.Face {
		case '<':
			return Position{
				Pos:  Coordinate{Row: 15 - pos.Pos.Row, Col: 7},
				Face: '^',
			}
		case 'v':
			return Position{
				Pos:  Coordinate{Row: 7, Col: 11 - pos.Pos.Col},
				Face: '^',
			}
		}
	case 6:
		switch pos.Face {
		case '>':
			return Position{
				Pos:  Coordinate{Row: 11 - pos.Pos.Row, Col: 11},
				Face: '<',
			}
		case 'v':
			return Position{
				Pos:  Coordinate{Row: 19 - pos.Pos.Col, Col: 0},
				Face: '>',
			}
		case '^':
			return Position{
				Pos:  Coordinate{Row: 19 - pos.Pos.Col, Col: 11},
				Face: '<',
			}
		}
	}

	panic("shouldn't be here")

	return Position{}
}

func (w Wrapper3DT) next(pos Coordinate, dir byte) Coordinate {
	next := pos.Next(dir)

	return next
}

func (w Wrapper3DT) isVoid(pos Coordinate) bool {
	return pos.Row < 0 || pos.Col < 0 || pos.Row > len(w.grid)-1 || pos.Col > len(w.grid[0])-1 ||
		(w.grid[pos.Row][pos.Col] == Void ||
			w.grid[pos.Row][pos.Col] == Space)
}

func (w Wrapper3DT) face(pos Coordinate) int {
	for i := 0; i < 6; i++ {
		if pos.Row >= w.facesMin[i].Row && pos.Col >= w.facesMin[i].Col &&
			pos.Row <= w.facesMax[i].Row && pos.Col <= w.facesMax[i].Col {
			return i + 1
		}
	}

	return -1
}
