package day11

import (
	"fmt"
	"github.com/leandrorondon/advent-of-code/pkg/maputils"
	mymath "github.com/leandrorondon/advent-of-code/pkg/math"
	"slices"
)

const (
	empty  = '.'
	galaxy = '#'
)

type Position struct {
	R int
	C int
}

func (p *Position) Equals(r, c int) bool {
	return r == p.R && c == p.C
}

func (p *Position) Distance(p2 *Position) int {
	return mymath.Abs(p.R-p2.R) + mymath.Abs(p.C-p2.C)
}

type Map struct {
	Rows     int
	Cols     int
	Galaxies []*Position
}

func (m *Map) Print() {
	gid := 0

	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			if (gid < len(m.Galaxies)) && m.Galaxies[gid].Equals(i, j) {
				gid++
				fmt.Printf(string(galaxy))
				continue
			}
			fmt.Printf(string(empty))
		}
		fmt.Printf("\n")
	}
}

func (m *Map) Expand(n int) *Map {
	n = n - 1

	rowGalaxyMap := make(map[int]bool)
	colGalaxyMap := make(map[int]bool)

	galaxies := make([]*Position, len(m.Galaxies))
	for i, g := range m.Galaxies {
		rowGalaxyMap[g.R] = true
		colGalaxyMap[g.C] = true
		galaxies[i] = &Position{g.R, g.C}
	}

	rowGalaxy := maputils.KeysToSlice(rowGalaxyMap)
	colGalaxy := maputils.KeysToSlice(colGalaxyMap)

	emptyRows := inverseSlice(rowGalaxy, m.Rows)
	emptyCols := inverseSlice(colGalaxy, m.Cols)

	for i := range emptyRows {
		for j, g := range galaxies {
			if emptyRows[i] < m.Galaxies[j].R {
				g.R += n
			}
		}
	}

	for i := range emptyCols {
		for j, g := range galaxies {
			if emptyCols[i] < m.Galaxies[j].C {
				g.C += n
			}
		}
	}

	return &Map{
		Rows:     m.Rows + len(emptyRows)*n,
		Cols:     m.Cols + len(emptyCols)*n,
		Galaxies: galaxies,
	}
}

func (m *Map) SumDistances() int {
	sum := 0
	for i := 0; i < len(m.Galaxies)-1; i++ {
		for j := i + 1; j < len(m.Galaxies); j++ {
			dist := m.Galaxies[i].Distance(m.Galaxies[j])
			//fmt.Printf("Between galaxy %d (%d,%d) and galaxy %d (%d,%d) : %d\n", i+1, m.Galaxies[i].R, m.Galaxies[i].C, j+1, m.Galaxies[j].R, m.Galaxies[j].C, dist)
			sum += dist

		}
	}

	return sum
}

func inverseSlice(in []int, size int) []int {
	var inverse []int
	for i := 0; i < size; i++ {
		if slices.Contains(in, i) {
			continue
		}
		inverse = append(inverse, i)
	}

	return inverse
}
