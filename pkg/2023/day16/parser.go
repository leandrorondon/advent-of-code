package day16

import (
	"github.com/leandrorondon/advent-of-code/pkg/bits"
	"strings"
)

func Parse(s string) *Map {
	lines := strings.Split(s, "\n")
	m := make([][]*Mirror, len(lines))
	energised := make([]*bits.Bits, len(lines))

	closestInRow := make([]*Mirror, len(lines))
	closestInCol := make([]*Mirror, len(lines[0]))

	for r := range lines {
		energised[r] = bits.New(len(lines[r]))
		m[r] = make([]*Mirror, len(lines[r]))
		for c := range lines[r] {
			if lines[r][c] != '.' {
				mirror := &Mirror{
					R:    r,
					C:    c,
					Type: lines[r][c],
				}

				if closestInRow[r] != nil {
					if mirror.Type == '-' || mirror.Type == '\\' || mirror.Type == '/' {
						mirror.left = closestInRow[r]
					}
					if closestInRow[r].Type == '-' || closestInRow[r].Type == '\\' || closestInRow[r].Type == '/' {
						closestInRow[r].right = mirror
					}
				}

				if closestInCol[c] != nil {
					if mirror.Type == '|' || mirror.Type == '\\' || mirror.Type == '/' {
						mirror.up = closestInCol[c]
					}
					if closestInCol[c].Type == '|' || closestInCol[c].Type == '\\' || closestInCol[c].Type == '/' {
						closestInCol[c].down = mirror
					}
				}

				closestInRow[r] = mirror
				closestInCol[c] = mirror
				m[r][c] = mirror
			}
		}
	}

	return &Map{
		mirrors:   m,
		energised: energised,
	}
}
