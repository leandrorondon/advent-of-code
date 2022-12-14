package grid

import (
	"fmt"
	mymath "github.com/leandrorondon/advent-of-code/pkg/math"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type Path []Coordinate

type Material string

const (
	Air      Material = "."
	Rock              = "#"
	Sand              = "o"
	Source            = "+"
	SandPath          = "~"
)

func New(source Coordinate) *Grid {
	return &Grid{
		m:      make(map[Coordinate]Material),
		Min:    source,
		Max:    source,
		source: source,
	}
}

type Grid struct {
	m      map[Coordinate]Material
	Min    Coordinate
	Max    Coordinate
	source Coordinate
}

func (g *Grid) Set(pos Coordinate, mat Material) {
	g.m[pos] = mat
	g.checkMaxMin(pos.X, pos.Y)
}

func (g *Grid) SetXY(x, y int, mat Material) {
	pos := Coordinate{x, y}
	g.m[pos] = mat
	g.checkMaxMin(pos.X, pos.Y)
}

func (g *Grid) Get(pos Coordinate) Material {
	mat, ok := g.m[Coordinate{pos.X, pos.Y}]
	if ok {
		return mat
	}

	return Air
}

func (g *Grid) checkMaxMin(x, y int) {
	g.Min.X = mymath.Min(x, g.Min.X)
	g.Max.X = mymath.Max(x, g.Max.X)
	g.Min.Y = mymath.Min(y, g.Min.Y)
	g.Max.Y = mymath.Max(y, g.Max.Y)
}

func (g *Grid) Print() {
	h1 := fmt.Sprintf("   %3d", g.Min.X)
	h2 := "    |"
	prev := g.Min.X
	for x := g.Min.X; x <= g.Max.X; x++ {
		if x%10 == 0 && x > prev+5 {
			fmt.Println(x, prev)
			h1 += strings.Repeat(" ", x-prev-3) + fmt.Sprintf("%3d", x)
			h2 += strings.Repeat(" ", x-prev-1) + "|"
			prev = x
		}
	}
	fmt.Println(h1)
	fmt.Println(h2)
	for y := g.Min.Y; y <= g.Max.Y; y++ {
		s := fmt.Sprintf("%3d ", y)
		for x := g.Min.X; x <= g.Max.X; x++ {
			coord := Coordinate{X: x, Y: y}
			mat, ok := g.m[coord]
			if ok {
				if coord == g.source && mat != Sand {
					s += Source
				} else {
					s += string(mat)
				}

			} else {
				s += string(Air)
			}
		}
		fmt.Println(s)
	}
}
