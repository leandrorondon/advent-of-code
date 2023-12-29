package day13

import (
	"strings"
)

func Parse(s string) Patterns {
	lines := strings.Split(s, "\n")

	var ps Patterns
	var p Pattern
	p.Columns = make([]string, len(lines[0]))

	for i := range lines {
		if lines[i] == "" {
			ps = append(ps, p)

			p = Pattern{}
			p.Columns = make([]string, len(lines[i+1]))
			continue
		}

		p.Rows = append(p.Rows, lines[i])
		for c := 0; c < len(lines[i]); c++ {
			p.Columns[c] = p.Columns[c] + string(lines[i][c])
		}
	}
	ps = append(ps, p)

	return ps
}
