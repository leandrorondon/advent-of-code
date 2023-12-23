package day10

import (
	"fmt"
	"github.com/leandrorondon/advent-of-code/pkg/stack"
	"slices"
)

type Node struct {
	R           int
	C           int
	Connections map[Direction]*Node
	Value       string
}

func (n *Node) Print() {
	if n == nil {
		return
	}

	fmt.Printf("(%d,%d) %s\n", n.R, n.C, n.Value)
	for k, v := range n.Connections {
		fmt.Printf("  %s: %s\n", k, v.Value)
	}
}

type Direction string

const (
	up    Direction = "u"
	down  Direction = "d"
	left  Direction = "l"
	right Direction = "r"
)

var directions = []Direction{up, down, left, right}

type Position struct {
	R int
	C int
}

const (
	vertical    = "|"
	horizontal  = "-"
	topleft     = "F"
	topright    = "7"
	bottomleft  = "L"
	bottomright = "J"
	start       = "S"
)

var directionsToCheck = map[string][]Direction{
	vertical:    {up, down},
	horizontal:  {left, right},
	topleft:     {down, right},
	topright:    {down, left},
	bottomleft:  {up, right},
	bottomright: {up, left},
	start:       {up, down, left, right},
}

var oppositeConnection = map[Direction][]string{
	up:    {vertical, topleft, topright, start},
	down:  {vertical, bottomleft, bottomright, start},
	left:  {horizontal, topleft, bottomleft, start},
	right: {horizontal, topright, bottomright, start},
}

type Map struct {
	Nodes [][]*Node
	Start Position
	Rows  int
	Cols  int
}

type SearchStep struct {
	Node     *Node
	Previous *Node
	LoopSize int
	Path     []*Node
}

func (m *Map) Farthest(p Position) int {
	return (m.LoopSize(p) - 1) / 2
}

func (m *Map) LoopSize(p Position) int {
	start := m.Nodes[p.R][p.C]
	st := stack.NewStack[SearchStep]()

	st.Push(SearchStep{
		Node:     start,
		LoopSize: 1,
		Path:     []*Node{start},
	})

	for st.Size() > 0 {
		step := st.Pop()
		for _, toCheck := range directionsToCheck[step.Node.Value] {
			next := m.Go(step.Node, toCheck)
			if next == nil || next == step.Previous {
				continue
			}

			oc := oppositeConnection[toCheck]
			if !slices.Contains(oc, next.Value) {
				continue
			}

			if next == start {
				return step.LoopSize + 1
			}

			st.Prepend(SearchStep{
				Node:     next,
				LoopSize: step.LoopSize + 1,
				Path:     append(step.Path, next),
				Previous: step.Node,
			})
		}
	}

	return 0
}

func (m *Map) BuildNeighbourhood() {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			if m.Nodes[i][j] == nil {
				continue
			}

			toCheck := directionsToCheck[m.Nodes[i][j].Value]
			for _, dir := range toCheck {
				next := m.Go(m.Nodes[i][j], dir)
				if next == nil {
					continue
				}

				oc := oppositeConnection[dir]
				if slices.Contains(oc, next.Value) {
					m.Nodes[i][j].Connections[dir] = next
				}
			}
		}
	}
}

func (m *Map) Print() {
	for _, v := range m.Nodes {
		for _, n := range v {
			n.Print()
		}
	}
}

func (m *Map) Go(from *Node, dir Direction) *Node {
	if from == nil {
		return nil
	}

	switch dir {
	case up:
		if from.R == 0 {
			return nil
		}

		return m.Nodes[from.R-1][from.C]

	case down:
		if from.R == len(m.Nodes)-1 {
			return nil
		}

		return m.Nodes[from.R+1][from.C]

	case left:
		if from.C == 0 {
			return nil
		}

		return m.Nodes[from.R][from.C-1]

	case right:
		if from.C == len(m.Nodes[0])-1 {
			return nil
		}

		return m.Nodes[from.R][from.C+1]
	}

	return nil
}
