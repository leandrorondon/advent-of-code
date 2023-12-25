package day10

import (
	"strings"
)

func ParseMap(s string) Map {
	lines := strings.Split(s, "\n")
	var m Map
	m.Nodes = make([][]*Node, len(lines))
	m.Rows = len(lines)
	m.Cols = len(lines[0])
	for r := range lines {
		m.Nodes[r] = make([]*Node, len(lines[r]))
		if c := strings.Index(lines[r], "S"); c >= 0 {
			m.Start = Position{r, c}
		}

		for i := 0; i < len(lines[r]); i++ {
			m.Nodes[r][i] = &Node{
				Value:       string(lines[r][i]),
				R:           r,
				C:           i,
				Connections: make(map[Direction]*Node),
			}
		}
	}

	return m
}
