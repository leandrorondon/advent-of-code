package physics

import (
	"github.com/leandrorondon/advent-of-code/pkg/2022/day14/grid"
)

type Physics struct{}

func NewPhysics() Physics {
	return Physics{}
}

func (ph Physics) Fall(pos *grid.Coordinate, g *grid.Grid) bool {
	settled := false
	for !settled && ph.inGrid(pos, g) {
		settled = ph.fallStep(pos, g)
	}

	return settled
}

func (ph Physics) inGrid(pos *grid.Coordinate, g *grid.Grid) bool {
	return pos.X >= g.Min.Y &&
		pos.X <= g.Max.X &&
		pos.Y >= g.Min.Y &&
		pos.Y <= g.Max.Y
}

func (ph Physics) fallStep(pos *grid.Coordinate, g *grid.Grid) bool {
	old := *pos
	if ph.fallDown(pos, g) {
		if pos.Y <= g.Max.Y {
			g.Set(old, grid.SandPath)
			g.Set(*pos, grid.Sand)
		}
		return false
	}

	if ph.fallLeft(pos, g) {
		if pos.Y <= g.Max.Y {
			g.Set(old, grid.SandPath)
			g.Set(*pos, grid.Sand)
		}
		return false
	}

	if ph.fallRight(pos, g) {
		if pos.Y <= g.Max.Y {
			g.Set(old, grid.SandPath)
			g.Set(*pos, grid.Sand)
		}
		return false
	}

	return true
}

func (ph Physics) fallDown(pos *grid.Coordinate, g *grid.Grid) bool {
	next := &grid.Coordinate{X: pos.X, Y: pos.Y + 1}
	return ph.fallTo(pos, next, g)
}

func (ph Physics) fallLeft(pos *grid.Coordinate, g *grid.Grid) bool {
	next := &grid.Coordinate{X: pos.X - 1, Y: pos.Y + 1}
	return ph.fallTo(pos, next, g)
}

func (ph Physics) fallRight(pos *grid.Coordinate, g *grid.Grid) bool {
	next := &grid.Coordinate{X: pos.X + 1, Y: pos.Y + 1}
	return ph.fallTo(pos, next, g)
}

func (ph Physics) fallTo(cur, next *grid.Coordinate, g *grid.Grid) bool {
	m := g.Get(*next)
	if m == grid.SandPath || m == grid.Air {
		*cur = *next
		return true
	}

	return false
}
