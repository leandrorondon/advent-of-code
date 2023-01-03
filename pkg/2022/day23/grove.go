package day23

import (
	"fmt"
	"math"
)

const (
	empty = '.'
	elf   = '#'

	N  = "N"
	S  = "S"
	E  = "E"
	W  = "W"
	NE = "NE"
	NW = "NW"
	SE = "SE"
	SW = "SW"
)

var increment = map[string]Coordinate{
	N:  {-1, 0},
	S:  {1, 0},
	E:  {0, 1},
	W:  {0, -1},
	NW: {-1, -1},
	NE: {-1, 1},
	SW: {1, -1},
	SE: {1, 1},
}

type Grove struct {
	grid map[Coordinate]bool
	elfs []Elf
}

func NewGrove(elfs []Elf) *Grove {
	g := &Grove{
		grid: make(map[Coordinate]bool),
		elfs: elfs,
	}

	for _, e := range elfs {
		g.grid[e.Position] = true
	}

	return g
}

func (g *Grove) RunRounds(rounds int) {
	for i := 0; i < rounds; i++ {
		g.round()
	}
}

func (g *Grove) RunAll() int {
	rounds := 1
	for ; g.round(); rounds++ {
	}
	return rounds
}

func (g *Grove) round() bool {
	moved := false
	proposalsByElf := make(map[int]Coordinate)
	proposalsByPos := make(map[Coordinate]int)

	// 1st half
	for i := range g.elfs {
		proposed, pos := g.elfs[i].Propose(g)
		if !proposed {
			continue
		}

		proposalsByElf[i] = pos
		proposalsByPos[pos] = proposalsByPos[pos] + 1
	}

	// 2ns half
	for i, pos := range proposalsByElf {
		if proposalsByPos[pos] > 1 {
			continue
		}

		moved = true
		g.moveElf(i)
	}

	for i := range g.elfs {
		g.elfs[i].EndRound()
	}

	return moved
}

func (g *Grove) moveElf(id int) {
	g.grid[g.elfs[id].Position] = false
	g.elfs[id].AcceptProposal()
	g.grid[g.elfs[id].Position] = true
}

func (g *Grove) Occupied(pos Coordinate) bool {
	return g.grid[pos]
}

func (g *Grove) EmptySpaces() int {
	min := Coordinate{math.MaxInt, math.MaxInt}
	max := Coordinate{math.MinInt, math.MinInt}

	for i := range g.elfs {
		p := g.elfs[i].Position
		if p.Row < min.Row {
			min.Row = p.Row
		} else if p.Row > max.Row {
			max.Row = p.Row
		}
		if p.Col < min.Col {
			min.Col = p.Col
		} else if p.Col > max.Col {
			max.Col = p.Col
		}
	}

	total := (max.Row - min.Row + 1) * (max.Col - min.Col + 1)
	return total - len(g.elfs)
}

func (g *Grove) Print() {
	min := Coordinate{math.MaxInt, math.MaxInt}
	max := Coordinate{math.MinInt, math.MinInt}
	for i := range g.elfs {
		p := g.elfs[i].Position
		if p.Row < min.Row {
			min.Row = p.Row
		} else if p.Row > max.Row {
			max.Row = p.Row
		}
		if p.Col < min.Col {
			min.Col = p.Col
		} else if p.Col > max.Col {
			max.Col = p.Col
		}
	}

	grid := make([][]byte, max.Row-min.Row+1)
	for row := range grid {
		grid[row] = make([]byte, max.Col-min.Col+1)

		for col := range grid[row] {
			if g.Occupied(Coordinate{row + min.Row, col + min.Col}) {
				grid[row][col] = elf
			} else {
				grid[row][col] = empty
			}
		}

		fmt.Println(string(grid[row]))
	}
	fmt.Println("--- ", len(grid), len(grid[0]))
}
