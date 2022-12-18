package tetris

import (
	"fmt"
	"strconv"
	"strings"
)

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

func (r *Rock) Print() {
	for i := r.Rows - 1; i >= 0; i-- {
		bitsPattern := r.Pattern[i]

		s := strconv.FormatInt(int64(bitsPattern), 2)
		v, _ := strconv.Atoi(s)
		l := strings.Repeat("0", r.Cols-len(s))
		s = fmt.Sprintf("%s%d", l, v)
		s = strings.ReplaceAll(s, "1", "#")
		s = strings.ReplaceAll(s, "0", ".")
		fmt.Println(s)
	}
	fmt.Printf("\n")
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
