package volcano

import "github.com/leandrorondon/advent-of-code/pkg/2022/day16/valve"

type Iterator struct {
	valves map[string]*valve.Valve
}

func (it Iterator) Next(n *valve.Valve) []*valve.Valve {
	return n.Tunnels
}

func (it Iterator) Constraint(_, _ *valve.Valve) bool {
	return false
}
