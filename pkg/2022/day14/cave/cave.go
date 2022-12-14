package cave

import (
	"github.com/leandrorondon/advent-of-code/pkg/2022/day14/grid"
	mymath "github.com/leandrorondon/advent-of-code/pkg/math"
)

type Cave struct {
	Grid        *grid.Grid
	source      grid.Coordinate
	settledSand int
	physics     Faller
}

type Faller interface {
	Fall(pos *grid.Coordinate, g *grid.Grid) bool
}

func New(sandSource grid.Coordinate, rockPaths []grid.Path, physics Faller) *Cave {
	cave := &Cave{
		Grid:    grid.New(sandSource),
		source:  sandSource,
		physics: physics,
	}
	cave.Grid.Set(sandSource, grid.Source)
	cave.setMaterial(rockPaths, grid.Rock)
	return cave
}

func (c *Cave) SetFloor(level int) {
	for x := c.Grid.Min.X; x <= c.Grid.Max.X; x++ {
		c.Grid.SetXY(x, level, grid.Rock)
	}
}

func (c *Cave) FillSand() int {
	var i int

	for {
		s, settled := c.ProduceSand()
		if s == nil || !settled {
			break
		}
		i++
	}

	return i
}

func (c *Cave) ProduceSand() (*grid.Coordinate, bool) {
	if c.Grid.Get(c.source) == grid.Sand {
		return nil, false
	}

	c.Grid.Set(c.source, grid.Sand)
	grain := c.source
	settled := c.physics.Fall(&grain, c.Grid)

	return &grain, settled
}

func (c *Cave) setMaterial(paths []grid.Path, mat grid.Material) {
	for _, path := range paths {
		c.setMaterialPath(path, mat)
	}
}

func (c *Cave) setMaterialPath(path grid.Path, mat grid.Material) {
	for i := 0; i < len(path)-1; i++ {
		if mymath.Abs(path[i+1].X-path[i].X) > 0 {
			c.setMaterialVertical(path[i], path[i+1], mat)
		} else if mymath.Abs(path[i+1].Y-path[i].Y) > 0 {
			c.setMaterialHorizontal(path[i], path[i+1], mat)
		}
	}
}

func (c *Cave) setMaterialHorizontal(c1, c2 grid.Coordinate, mat grid.Material) {
	from, to := mymath.MinMax(c1.Y, c2.Y)
	for y := from; y <= to; y++ {
		c.Grid.Set(grid.Coordinate{X: c1.X, Y: y}, mat)
	}
}

func (c *Cave) setMaterialVertical(c1, c2 grid.Coordinate, mat grid.Material) {
	from, to := mymath.MinMax(c1.X, c2.X)
	for x := from; x <= to; x++ {
		c.Grid.Set(grid.Coordinate{X: x, Y: c1.Y}, mat)
	}
}
