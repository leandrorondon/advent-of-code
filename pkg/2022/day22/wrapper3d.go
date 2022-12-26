package day22

type Wrapper3D struct {
	grid     [][]byte
	facesMin [6]Coordinate
	facesMax [6]Coordinate
}

func NewWrapper3D(grid [][]byte) Wrapper3D {
	w := Wrapper3D{
		grid: grid,
	}

	w.facesMin[0] = Coordinate{0, 50}
	w.facesMax[0] = Coordinate{49, 99}
	w.facesMin[1] = Coordinate{0, 100}
	w.facesMax[1] = Coordinate{49, 149}
	w.facesMin[2] = Coordinate{50, 50}
	w.facesMax[2] = Coordinate{99, 99}
	w.facesMin[3] = Coordinate{100, 0}
	w.facesMax[3] = Coordinate{149, 49}
	w.facesMin[4] = Coordinate{100, 50}
	w.facesMax[4] = Coordinate{149, 99}
	w.facesMin[5] = Coordinate{150, 0}
	w.facesMax[5] = Coordinate{199, 49}

	return w
}

func (w Wrapper3D) Password(pos Position) int {
	return 1000*(pos.Pos.Row+1) + 4*(pos.Pos.Col+1) + dirValue[pos.Face]
}

func (w Wrapper3D) Next(pos *Position) Position {
	next := pos.Pos.Next(pos.Face)

	if !w.isVoid(next) {
		return Position{Pos: next, Face: pos.Face}
	}

	return w.wrap(pos)
}

func (w Wrapper3D) wrap(pos *Position) Position {
	switch w.face(pos.Pos) {
	case 1:
		switch pos.Face {
		case '^':
			return Position{
				Pos:  Coordinate{Row: pos.Pos.Col + 100, Col: 0},
				Face: '>',
			}
		case '<':
			return Position{
				Pos:  Coordinate{Row: 149 - pos.Pos.Row, Col: 0},
				Face: '>',
			}
		}
	case 2:
		switch pos.Face {
		case '^':
			return Position{
				Pos:  Coordinate{Row: 199, Col: pos.Pos.Col - 100},
				Face: '^',
			}
		case 'v':
			return Position{
				Pos:  Coordinate{Row: pos.Pos.Col - 50, Col: 99},
				Face: '<',
			}
		case '>':
			return Position{
				Pos:  Coordinate{Row: 149 - pos.Pos.Row, Col: 99},
				Face: '<',
			}
		}
	case 3:
		switch pos.Face {
		case '<':
			return Position{
				Pos:  Coordinate{Row: 100, Col: pos.Pos.Row - 50},
				Face: 'v',
			}
		case '>':
			return Position{
				Pos:  Coordinate{Row: 49, Col: pos.Pos.Row + 50},
				Face: '^',
			}
		}
	case 4:
		switch pos.Face {
		case '^':
			return Position{
				Pos:  Coordinate{Row: pos.Pos.Col + 50, Col: 50},
				Face: '>',
			}
		case '<':
			return Position{
				Pos:  Coordinate{Row: 149 - pos.Pos.Row, Col: 50},
				Face: '>',
			}
		}
	case 5:
		switch pos.Face {
		case '>':
			return Position{
				Pos:  Coordinate{Row: 149 - pos.Pos.Row, Col: 149},
				Face: '<',
			}
		case 'v':
			return Position{
				Pos:  Coordinate{Row: pos.Pos.Col + 100, Col: 49},
				Face: '<',
			}
		}
	case 6:
		switch pos.Face {
		case '<':
			return Position{
				Pos:  Coordinate{Row: 0, Col: pos.Pos.Row - 100},
				Face: 'v',
			}
		case 'v':
			return Position{
				Pos:  Coordinate{Row: 0, Col: pos.Pos.Col + 100},
				Face: 'v',
			}
		case '>':
			return Position{
				Pos:  Coordinate{Row: 149, Col: pos.Pos.Row - 100},
				Face: '^',
			}
		}
	}

	panic("shouldn't be here")

	return Position{}
}

func (w Wrapper3D) next(pos Coordinate, dir byte) Coordinate {
	next := pos.Next(dir)

	return next
}

func (w Wrapper3D) isVoid(pos Coordinate) bool {
	return pos.Row < 0 || pos.Col < 0 || pos.Row > len(w.grid)-1 || pos.Col > len(w.grid[0])-1 ||
		(w.grid[pos.Row][pos.Col] == Void ||
			w.grid[pos.Row][pos.Col] == Space)
}

func (w Wrapper3D) face(pos Coordinate) int {
	for i := 0; i < 6; i++ {
		if pos.Row >= w.facesMin[i].Row && pos.Col >= w.facesMin[i].Col &&
			pos.Row <= w.facesMax[i].Row && pos.Col <= w.facesMax[i].Col {
			return i + 1
		}
	}

	return -1
}
