package day10

import (
	"fmt"
	"slices"

	"github.com/leandrorondon/advent-of-code/pkg/stack"
)

type Position struct {
	R int
	C int
}

type Node struct {
	R           int
	C           int
	Connections map[Direction]*Node
	Value       string
	WayOut      bool
	Loop        bool
}

func (n *Node) Print() {
	if n == nil {
		fmt.Printf(" ")
		return
	}
	fmt.Printf(n.Value)
}

type SearchStep struct {
	Node     *Node
	Previous *Node
	LoopSize int
	Path     []*Node
}

type WayOutStep struct {
	Node    *Node
	Visited map[*Node]bool
}

type Map struct {
	Nodes [][]*Node
	Start Position
	Rows  int
	Cols  int
}

func (m *Map) Farthest(p Position) int {
	return (m.LoopSize(p) - 1) / 2
}

func (m *Map) IsEdge(r, c int) bool {
	return r == 0 || c == 0 || r == m.Rows-1 || c == m.Cols-1
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
				for _, n := range step.Path {
					n.Loop = true
				}
				next.Loop = true
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

func (m *Map) BuildGraph() {
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

func (m *Map) PrintMap() {
	for i, v := range m.Nodes {
		fmt.Printf("%2d ", i)
		for _, n := range v {
			n.Print()
		}
		fmt.Printf("\n")
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

func (m *Map) Expand() *Map {
	newM := &Map{
		Nodes: nil,
		Start: Position{m.Start.R*2 - 1, m.Start.C*2 - 1},
		Rows:  m.Rows*2 - 1,
		Cols:  m.Cols*2 - 1,
	}

	newM.Nodes = make([][]*Node, newM.Rows)

	for r := 0; r < m.Rows; r++ {
		newM.Nodes[2*r] = make([]*Node, newM.Cols)

		if r > 0 {
			newM.Nodes[2*r-1] = make([]*Node, newM.Cols)
		}

		for i := 0; i < m.Cols; i++ {
			newM.Nodes[2*r][2*i] = &Node{
				R:           2 * r,
				C:           2 * i,
				Value:       m.Nodes[r][i].Value,
				Connections: make(map[Direction]*Node),
			}

			if i > 0 {
				newM.Nodes[2*r][2*i-1] = &Node{
					R:           2 * r,
					C:           2*i - 1,
					Value:       ground,
					Connections: make(map[Direction]*Node),
				}
				expandLeftRight(m.Nodes[r][i-1], m.Nodes[r][i], newM.Nodes[2*r][2*i-1], newM.Nodes[2*r][2*i-2], newM.Nodes[2*r][2*i])

				if r > 0 {
					newM.Nodes[2*r-1][2*i-1] = &Node{
						R:           2*r - 1,
						C:           2*i - 1,
						Value:       ground,
						Connections: make(map[Direction]*Node),
					}
				}
			}

			if r > 0 {
				newM.Nodes[2*r-1][2*i] = &Node{
					R:           2*r - 1,
					C:           2 * i,
					Value:       ground,
					Connections: make(map[Direction]*Node),
				}

				expandTopDown(m.Nodes[r-1][i], m.Nodes[r][i], newM.Nodes[2*r-1][2*i], newM.Nodes[2*r-2][2*i], newM.Nodes[2*r][2*i])
			}
		}
	}

	return newM
}

func (m *Map) findWayOut(node *Node, expanded *Map) bool {
	if node.Value != ground && node.Loop {
		return true
	}

	expandedNode := expanded.Nodes[node.R*2][node.C*2]

	st := stack.NewStack[WayOutStep]()

	st.Push(WayOutStep{
		Node:    expandedNode,
		Visited: make(map[*Node]bool),
	})

	for st.Size() > 0 {
		step := st.Pop()
		for _, toCheck := range directions {
			next := expanded.Go(step.Node, toCheck)
			if next == nil || (next.Value != ground && next.Loop) || step.Visited[next] {
				continue
			}

			if next.WayOut || expanded.IsEdge(next.R, next.C) {
				for n := range step.Visited {
					n.WayOut = true
				}
				next.WayOut = true
				return true
			}

			visited := step.Visited
			visited[next] = true
			st.Push(WayOutStep{
				Node:    next,
				Visited: visited,
			})
		}
	}

	return false
}

func (m *Map) Internals() int {
	exp := m.Expand()
	internals := 0

	for r := range m.Nodes {
		for c := range m.Nodes[r] {
			if m.Nodes[r][c].Loop {
				continue
			}

			m.Nodes[r][c].WayOut = m.findWayOut(m.Nodes[r][c], exp)

			if !m.Nodes[r][c].WayOut {
				internals++
			}
		}
	}

	return internals
}

func expandLeftRight(originalLeft, originalRight, prev1, prev2, cur *Node) {
	if originalRight.Connections[left] == originalLeft {
		// expanded node
		prev1.Value = horizontal
		prev1.Connections[left] = prev2
		prev1.Connections[right] = cur
		// current node
		cur.Connections[left] = prev1
		// previous node
		prev2.Connections[right] = prev1

		if originalRight.Loop {
			prev1.Loop = true
			prev2.Loop = true
			cur.Loop = true
		}
	}
}

func expandTopDown(originalTop, originalDown, prev1, prev2, cur *Node) {
	if originalTop.Connections[down] == originalDown {
		// expanded node
		prev1.Value = vertical
		prev1.Connections[up] = prev2
		prev1.Connections[down] = cur
		// current node
		cur.Connections[up] = prev1
		// previous node
		prev2.Connections[down] = prev1

		if originalTop.Loop {
			prev1.Loop = true
			prev2.Loop = true
			cur.Loop = true
		}
	}
}
