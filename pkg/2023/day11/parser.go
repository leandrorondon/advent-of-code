package day11

import "strings"

func Parse(s string) *Map {
	lines := strings.Split(s, "\n")
	var galaxies []*Position
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == galaxy {
				galaxies = append(galaxies, &Position{i, j})
			}
		}
	}

	return &Map{
		Galaxies: galaxies,
		Rows:     len(lines),
		Cols:     len(lines[0]),
	}
}
