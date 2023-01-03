package day23

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
