package volcano

import (
	"github.com/leandrorondon/advent-of-code/pkg/2022/day16/valve"
	"github.com/leandrorondon/advent-of-code/pkg/graph"
	"github.com/leandrorondon/advent-of-code/pkg/math"
)

type Volcano struct {
	valves    []*valve.Valve
	distances [][]int
}

func New(valvesMap map[string]*valve.Valve) Volcano {
	var it graph.Iterator[*valve.Valve] = Iterator{
		valves: valvesMap,
	}

	valves := make([]*valve.Valve, len(valvesMap))
	for _, v := range valvesMap {
		valves[v.ID] = v
	}

	m := make([][]int, len(valvesMap))
	for _, v1 := range valvesMap {
		m[v1.ID] = make([]int, len(valvesMap))
		for _, v2 := range valvesMap {
			if v2.Rate > 0 && v1 != v2 {
				m[v1.ID][v2.ID] = graph.ShortestPath(v1, v2, it)
			}
		}
	}

	return Volcano{
		valves:    valves,
		distances: m,
	}
}

type Node struct {
	Valve            *valve.Valve
	Time             int
	PressureReleased int
	Visited          []bool
}

func (v Volcano) HighestPossiblePressureReleased(from *valve.Valve, maxTime int) int {
	var queue []*Node
	queue = append(queue, &Node{
		Valve:   from,
		Visited: make([]bool, len(v.valves)),
	})

	highest := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		highest = math.Max(highest, current.PressureReleased)

		for _, next := range v.valves {
			dist := v.distances[current.Valve.ID][next.ID]

			if next.Rate > 0 && current.Time+dist <= maxTime && !current.Visited[next.ID] {
				visited := make([]bool, len(current.Visited))
				copy(visited, current.Visited)
				visited[next.ID] = true
				moveAndOpen := dist + 1
				node := &Node{
					Valve:            next,
					Time:             current.Time + moveAndOpen, // +1 to open the valve
					Visited:          visited,
					PressureReleased: current.PressureReleased + (maxTime-(current.Time+moveAndOpen))*next.Rate,
				}
				queue = append(queue, node)
			}
		}
	}

	return highest
}
