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
	ValveID          int
	Time             int
	PressureReleased int
	Visited          []bool
}

func (v Volcano) HighestPossiblePressureReleased(from *valve.Valve, maxTime int) (int, []*Node) {
	var queue []*Node
	queue = append(queue, &Node{
		ValveID: from.ID,
		Visited: make([]bool, len(v.valves)),
	})

	highest := 0
	var all []*Node

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		highest = math.Max(highest, current.PressureReleased)
		all = append(all, current)

		for _, next := range v.valves {
			dist := v.distances[current.ValveID][next.ID]

			if next.Rate > 0 && current.Time+dist <= maxTime && !current.Visited[next.ID] {
				visited := make([]bool, len(current.Visited))
				copy(visited, current.Visited)
				visited[next.ID] = true
				moveAndOpen := dist + 1
				node := &Node{
					ValveID:          next.ID,
					Time:             current.Time + moveAndOpen, // +1 to open the valve
					Visited:          visited,
					PressureReleased: current.PressureReleased + (maxTime-(current.Time+moveAndOpen))*next.Rate,
				}
				queue = append(queue, node)
			}
		}
	}

	return highest, all
}

func (v Volcano) FindBestCombination(paths []*Node) int {
	highest := 0
	for i := 0; i < len(paths)-1; i++ {
	outer:
		for j := i + 1; j < len(paths); j++ {
			// ignore crossed paths
			for id, visited := range paths[i].Visited {
				if visited && paths[j].Visited[id] {
					continue outer
				}
			}

			highest = math.Max(highest, paths[i].PressureReleased+paths[j].PressureReleased)
		}
	}

	return highest
}
