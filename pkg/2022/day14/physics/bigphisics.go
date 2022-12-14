package physics

import (
	"github.com/leandrorondon/advent-of-code/pkg/2022/day14/grid"
)

type BigPhysics struct {
	floor int
}

func NewBigPhysics(floor int) BigPhysics {
	return BigPhysics{
		floor: floor,
	}
}

func (ph BigPhysics) Fall(pos *grid.Coordinate, g *grid.Grid) bool {
	settled := false
	for !settled && ph.inGrid(pos, g) {
		settled = ph.fallStep(pos, g)
	}

	return settled
}

func (ph BigPhysics) inGrid(pos *grid.Coordinate, g *grid.Grid) bool {
	return pos.X >= g.Min.Y &&
		pos.X <= g.Max.X &&
		pos.Y >= g.Min.Y &&
		pos.Y <= g.Max.Y
}

func (ph BigPhysics) fallStep(pos *grid.Coordinate, g *grid.Grid) bool {
	if ph.fallDown(pos, g) {
		return false
	}

	if ph.fallLeft(pos, g) {
		return false
	}

	if ph.fallRight(pos, g) {
		return false
	}

	return true
}

func (ph BigPhysics) fallDown(pos *grid.Coordinate, g *grid.Grid) bool {
	next := &grid.Coordinate{X: pos.X, Y: pos.Y + 1}
	return ph.fallTo(pos, next, g)
}

func (ph BigPhysics) fallLeft(pos *grid.Coordinate, g *grid.Grid) bool {
	next := &grid.Coordinate{X: pos.X - 1, Y: pos.Y + 1}
	if next.X == g.Min.X {
		g.Min.X--
		g.SetXY(g.Min.X, g.Max.Y, grid.Rock)
	}

	return ph.fallTo(pos, next, g)
}

func (ph BigPhysics) fallRight(pos *grid.Coordinate, g *grid.Grid) bool {
	next := &grid.Coordinate{X: pos.X + 1, Y: pos.Y + 1}
	if next.X == g.Max.X {
		g.Max.X++
		g.Set(g.Max, grid.Rock)
	}

	return ph.fallTo(pos, next, g)
}

func (ph BigPhysics) fallTo(cur, next *grid.Coordinate, g *grid.Grid) bool {
	m := g.Get(*next)
	if m == grid.SandPath || m == grid.Air {
		g.Set(*cur, grid.SandPath)
		g.Set(*next, grid.Sand)
		*cur = *next
		return true
	}

	return false
}
