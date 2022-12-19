package tetris

func NewRock(t Rock) *Rock {
	return &Rock{
		Pattern: t.Pattern,
		Cols:    t.Cols,
		Rows:    t.Rows,
	}
}

type Rock struct {
	Position Coordinate
	Pattern  []int
	Cols     int
	Rows     int
	Rest     bool
}

var (
	/*
		####
	*/
	rock1 = Rock{
		Pattern: []int{15},
		Cols:    4, Rows: 1,
	}

	/*
		.#.
		###
		.#.
	*/
	rock2 = Rock{
		Pattern: []int{2, 7, 2},
		Cols:    3, Rows: 3,
	}

	/*
		..#
		..#
		###
	*/
	rock3 = Rock{
		Pattern: []int{7, 1, 1},
		Cols:    3, Rows: 3,
	}

	/*
		#
		#
		#
		#
	*/
	rock4 = Rock{
		Pattern: []int{1, 1, 1, 1},
		Cols:    1, Rows: 4,
	}

	/*
		##
		##
	*/
	rock5 = Rock{
		Pattern: []int{3, 3},
		Cols:    2, Rows: 2,
	}
)
