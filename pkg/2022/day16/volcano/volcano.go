package volcano

import (
	"github.com/leandrorondon/advent-of-code/pkg/2022/day16/valve"
	"github.com/leandrorondon/advent-of-code/pkg/graph"
	"github.com/leandrorondon/advent-of-code/pkg/maputils"
	"github.com/leandrorondon/advent-of-code/pkg/math"
)

type Volcano struct {
	valves    map[string]*valve.Valve
	distances map[*valve.Valve]map[*valve.Valve]int
}

func New(valvesMap map[string]*valve.Valve) Volcano {
	var it graph.Iterator[*valve.Valve] = Iterator{
		valves: valvesMap,
	}

	m := make(map[*valve.Valve]map[*valve.Valve]int)
	for _, v1 := range valvesMap {
		m[v1] = make(map[*valve.Valve]int)
		for _, v2 := range valvesMap {
			if v2.Rate > 0 && v1 != v2 {
				m[v1][v2] = graph.ShortestPath(v1, v2, it)
			}
		}
	}

	return Volcano{
		valves:    valvesMap,
		distances: m,
	}
}

type Node struct {
	Valve            *valve.Valve
	Time             int
	PressureReleased int
	Visited          map[*valve.Valve]bool
}

func (v Volcano) HighestPossiblePressureReleased(from *valve.Valve) int {
	visited := make(map[*valve.Valve]bool)
	visited[from] = true

	var queue []Node
	queue = append(queue, Node{
		Valve:   from,
		Visited: make(map[*valve.Valve]bool),
	})

	highest := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		highest = math.Max(highest, current.PressureReleased)

		for _, next := range v.valves {
			dist := v.distances[current.Valve][next]

			if next.Rate > 0 && current.Time+dist <= 30 && !current.Visited[next] {
				visited[next] = true
				m := maputils.Copy(current.Visited)
				m[next] = true
				moveAndOpen := dist + 1
				node := Node{
					Valve:            next,
					Time:             current.Time + moveAndOpen, // +1 to open the valve
					Visited:          m,
					PressureReleased: current.PressureReleased + (30-(current.Time+moveAndOpen))*next.Rate,
				}
				queue = append(queue, node)
			}
		}
	}

	return highest
}
