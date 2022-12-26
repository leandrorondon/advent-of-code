package day22

type Coordinate struct {
	Row int
	Col int
}

func (c Coordinate) Add(other Coordinate) Coordinate {
	return Coordinate{
		Row: c.Row + other.Row,
		Col: c.Col + other.Col,
	}
}

func (c Coordinate) Next(dir byte) Coordinate {
	return c.Add(increment[dir])
}

var increment = map[byte]Coordinate{
	Right: {0, 1},
	Down:  {1, 0},
	Left:  {0, -1},
	Up:    {-1, 0},
}
